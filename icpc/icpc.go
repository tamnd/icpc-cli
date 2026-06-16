// Package icpc is the library behind the icpc command line:
// the HTTP client, request shaping, and the typed data models for
// the ICPC (International Collegiate Programming Contest).
//
// Data sources:
//   - Embedded World Finals dataset (1977-2024): compiled-in Go slice,
//     answers wf/contests/results/teams/search without any network.
//   - ICPC Live Archive (icpcarchive.ecs.baylor.edu): HTML scraping for
//     historical problems from regional and world final contests.
//   - CLICS 2020 scoreboards: when a public scoreboard URL is supplied via
//     --scoreboard-url, the client fetches and parses its JSON.
//
// icpc.global itself requires a session cookie for its REST API; the CLI
// does not attempt to authenticate there.
package icpc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// ArchiveHost is the ICPC Live Archive hostname.
const ArchiveHost = "icpcarchive.ecs.baylor.edu"

// ArchiveBaseURL is the root of the ICPC Live Archive.
const ArchiveBaseURL = "https://" + ArchiveHost

// ICPCHost is the main ICPC website.
const ICPCHost = "icpc.global"

// DefaultUserAgent identifies this client.
const DefaultUserAgent = "icpc/dev (+https://github.com/tamnd/icpc-cli)"

// Config holds constructor parameters for Client.
type Config struct {
	// ArchiveBaseURL is the ICPC Live Archive root. Override in tests.
	ArchiveBaseURL string
	UserAgent      string
	// Rate is the minimum spacing between requests. Zero means no pacing.
	Rate    time.Duration
	Retries int
	Timeout time.Duration
}

// DefaultConfig returns sensible defaults.
func DefaultConfig() Config {
	return Config{
		ArchiveBaseURL: ArchiveBaseURL,
		UserAgent:      DefaultUserAgent,
		Rate:           300 * time.Millisecond,
		Retries:        3,
		Timeout:        30 * time.Second,
	}
}

// Client is a rate-limited HTTP client for the ICPC Live Archive and
// CLICS scoreboards.
type Client struct {
	cfg  Config
	http *http.Client
	mu   sync.Mutex
	last time.Time
}

// NewClient returns a Client configured with cfg.
func NewClient(cfg Config) *Client {
	return &Client{
		cfg:  cfg,
		http: &http.Client{Timeout: cfg.Timeout},
	}
}

// Get fetches a URL with pacing and retries.
func (c *Client) Get(ctx context.Context, rawURL string) ([]byte, error) {
	var lastErr error
	for attempt := 0; attempt <= c.cfg.Retries; attempt++ {
		if attempt > 0 {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(backoff(attempt)):
			}
		}
		body, retry, err := c.do(ctx, rawURL)
		if err == nil {
			return body, nil
		}
		lastErr = err
		if !retry {
			return nil, err
		}
	}
	return nil, fmt.Errorf("get %s: %w", rawURL, lastErr)
}

func (c *Client) do(ctx context.Context, rawURL string) ([]byte, bool, error) {
	c.pace()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("User-Agent", c.cfg.UserAgent)
	req.Header.Set("Accept", "text/html,application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, true, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
		return nil, true, fmt.Errorf("http %d", resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, false, fmt.Errorf("http %d", resp.StatusCode)
	}

	b, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return nil, true, err
	}
	return b, false, nil
}

// pace blocks until at least Rate has elapsed since the last request.
func (c *Client) pace() {
	if c.cfg.Rate <= 0 {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if wait := c.cfg.Rate - time.Since(c.last); wait > 0 {
		time.Sleep(wait)
	}
	c.last = time.Now()
}

func backoff(attempt int) time.Duration {
	d := time.Duration(attempt) * 500 * time.Millisecond
	if d > 5*time.Second {
		d = 5 * time.Second
	}
	return d
}
