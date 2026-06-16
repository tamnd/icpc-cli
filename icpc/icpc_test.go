package icpc

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") == "" {
			t.Error("request carried no User-Agent")
		}
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	cfg := DefaultConfig()
	cfg.ArchiveBaseURL = srv.URL
	cfg.Rate = 0
	c := NewClient(cfg)

	body, err := c.Get(context.Background(), srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "ok" {
		t.Errorf("body = %q, want ok", body)
	}
}

func TestGetRetriesOn503(t *testing.T) {
	var hits int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits++
		if hits < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		_, _ = w.Write([]byte("recovered"))
	}))
	defer srv.Close()

	cfg := DefaultConfig()
	cfg.Rate = 0
	cfg.Retries = 5
	cfg.ArchiveBaseURL = srv.URL
	c := NewClient(cfg)

	start := time.Now()
	body, err := c.Get(context.Background(), srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != "recovered" {
		t.Errorf("body = %q after retries", body)
	}
	if hits != 3 {
		t.Errorf("server saw %d hits, want 3", hits)
	}
	if time.Since(start) < 500*time.Millisecond {
		t.Error("retries did not back off")
	}
}

func TestGetPaceRespected(t *testing.T) {
	var count int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	cfg := DefaultConfig()
	cfg.Rate = 100 * time.Millisecond
	cfg.ArchiveBaseURL = srv.URL
	c := NewClient(cfg)

	start := time.Now()
	for i := 0; i < 3; i++ {
		_, err := c.Get(context.Background(), srv.URL)
		if err != nil {
			t.Fatal(err)
		}
	}
	elapsed := time.Since(start)
	if elapsed < 200*time.Millisecond {
		t.Errorf("3 requests in %v, expected >= 200ms pacing", elapsed)
	}
}

func TestListVolumes(t *testing.T) {
	// Mock the Live Archive volume index HTML
	html := `<!DOCTYPE html><html><body>
<table>
<tr><td><a href="index.php?option=com_onlinejudge&Itemid=8&page=show_category&categoryid=2">Volume 1</a></td></tr>
<tr><td><a href="index.php?option=com_onlinejudge&Itemid=8&page=show_category&categoryid=3">Volume 2</a></td></tr>
</table>
</body></html>`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(html))
	}))
	defer srv.Close()

	cfg := DefaultConfig()
	cfg.ArchiveBaseURL = srv.URL
	cfg.Rate = 0
	c := NewClient(cfg)

	vols, err := c.ListVolumes(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(vols) < 2 {
		t.Fatalf("got %d volumes, want >= 2", len(vols))
	}
	if vols[0].Volume == "" {
		t.Error("volume ID should not be empty")
	}
}

func TestListProblems(t *testing.T) {
	html := `<!DOCTYPE html><html><body>
<a href="index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=100">100 - The 3n+1 Problem</a>
<a href="index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=101">101 - The Blocks Problem</a>
</body></html>`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(html))
	}))
	defer srv.Close()

	cfg := DefaultConfig()
	cfg.ArchiveBaseURL = srv.URL
	cfg.Rate = 0
	c := NewClient(cfg)

	probs, err := c.ListProblems(context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}
	if len(probs) < 2 {
		t.Fatalf("got %d problems, want >= 2", len(probs))
	}
	if probs[0].ID != "100" {
		t.Errorf("first problem ID = %q, want 100", probs[0].ID)
	}
}

func TestGetProblem(t *testing.T) {
	html := `<!DOCTYPE html><html>
<head><title>Problem 100 - The 3n+1 Problem</title></head>
<body>
<h2>The 3n+1 Problem</h2>
<a href="/external/1/100.pdf">PDF</a>
</body></html>`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(html))
	}))
	defer srv.Close()

	cfg := DefaultConfig()
	cfg.ArchiveBaseURL = srv.URL
	cfg.Rate = 0
	c := NewClient(cfg)

	p, err := c.GetProblem(context.Background(), "100")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(p.Title, "3n+1") {
		t.Errorf("title = %q, want to contain '3n+1'", p.Title)
	}
	if p.PDFURL == "" {
		t.Error("PDFURL should not be empty")
	}
}

func TestWorldFinals(t *testing.T) {
	all := WorldFinals(0)
	if len(all) == 0 {
		t.Fatal("WorldFinals(0) returned empty")
	}
	year2024 := WorldFinals(2024)
	if len(year2024) != 1 {
		t.Fatalf("WorldFinals(2024) = %d entries, want 1", len(year2024))
	}
	if year2024[0].WinnerCountry != "South Korea" {
		t.Errorf("2024 winner country = %q, want South Korea", year2024[0].WinnerCountry)
	}
}

func TestWorldFinalsNotFound(t *testing.T) {
	got := WorldFinals(1900)
	if len(got) != 0 {
		t.Errorf("WorldFinals(1900) = %d, want 0", len(got))
	}
}

func TestSearchContests(t *testing.T) {
	// Search by "Massachusetts" to match "Massachusetts Institute of Technology"
	results := SearchContests("Massachusetts")
	if len(results) == 0 {
		t.Fatal("SearchContests(Massachusetts) returned empty")
	}
	// MIT won in 2021, 2022, 2023
	found2021 := false
	for _, r := range results {
		if r.Year == 2021 {
			found2021 = true
		}
	}
	if !found2021 {
		t.Error("SearchContests(Massachusetts) should include 2021")
	}
}

func TestSearchContestsNoMatch(t *testing.T) {
	results := SearchContests("zzzzNotARealTeamName99999")
	if len(results) != 0 {
		t.Errorf("expected no results, got %d", len(results))
	}
}

func TestCLICSStandings(t *testing.T) {
	teamsJSON := `[{"id":"1","name":"Alpha Team","organization_id":"u1","country":"USA"},
{"id":"2","name":"Beta Squad","organization_id":"u2","country":"China"}]`
	sbJSON := `{"rows":[
{"rank":1,"team_id":"1","score":{"num_solved":12,"total_time":900}},
{"rank":2,"team_id":"2","score":{"num_solved":11,"total_time":850}}
]}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case strings.HasSuffix(r.URL.Path, "/teams"):
			_, _ = w.Write([]byte(teamsJSON))
		case strings.HasSuffix(r.URL.Path, "/organizations"):
			_, _ = w.Write([]byte("[]"))
		case strings.HasSuffix(r.URL.Path, "/scoreboard"):
			_, _ = w.Write([]byte(sbJSON))
		default:
			http.NotFound(w, r)
		}
	}))
	defer srv.Close()

	cfg := DefaultConfig()
	cfg.Rate = 0
	cfg.ArchiveBaseURL = srv.URL
	c := NewClient(cfg)

	standings, err := c.FetchCLICSStandings(context.Background(), srv.URL, 2024)
	if err != nil {
		t.Fatal(err)
	}
	if len(standings) != 2 {
		t.Fatalf("got %d standings, want 2", len(standings))
	}
	if standings[0].TeamName != "Alpha Team" {
		t.Errorf("rank 1 team = %q, want Alpha Team", standings[0].TeamName)
	}
	if standings[0].Medal != "gold" {
		t.Errorf("rank 1 medal = %q, want gold", standings[0].Medal)
	}
	if standings[1].Medal != "gold" {
		t.Errorf("rank 2 medal = %q, want gold (ICPC gold is top 4)", standings[1].Medal)
	}
}
