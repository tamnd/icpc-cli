package icpc

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// archiveCategoryBase is the URL template for the ICPC Live Archive volume index.
const archiveCategoryBase = "/index.php?option=com_onlinejudge&Itemid=8&page=show_categories&categoryid=1"

// archiveVolumeBase is the URL template for a single volume's problem list.
const archiveVolumeBase = "/index.php?option=com_onlinejudge&Itemid=8&page=show_category&categoryid=%s"

// archiveProblemBase is the URL template for a single problem page.
const archiveProblemBase = "/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=%s"

var (
	// tdRE matches a table cell content.
	tdRE = regexp.MustCompile(`(?i)<td[^>]*>(.*?)</td>`)
	// hrefRE matches href values.
	hrefRE = regexp.MustCompile(`(?i)href="([^"]+)"`)
	// tagRE strips HTML tags.
	tagRE = regexp.MustCompile(`<[^>]+>`)
	// categoryidRE extracts categoryid parameter.
	categoryidRE = regexp.MustCompile(`categoryid=(\d+)`)
	// problemRE extracts problem parameter.
	problemRE = regexp.MustCompile(`problem=(\d+)`)
	// titleRE extracts an HTML title.
	titleRE = regexp.MustCompile(`(?i)<title[^>]*>(.*?)</title>`)
	// h2RE extracts an h2 heading.
	h2RE = regexp.MustCompile(`(?i)<h2[^>]*>(.*?)</h2>`)
)

// stripTags removes HTML tags and trims whitespace from a string.
func stripTags(s string) string {
	return strings.TrimSpace(tagRE.ReplaceAllString(s, ""))
}

// ListVolumes scrapes the Live Archive category index and returns volume stubs.
func (c *Client) ListVolumes(ctx context.Context) ([]*Problem, error) {
	url := c.cfg.ArchiveBaseURL + archiveCategoryBase
	body, err := c.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("list volumes: %w", err)
	}
	return parseVolumeIndex(string(body), c.cfg.ArchiveBaseURL), nil
}

// parseVolumeIndex extracts volume stubs from the category index HTML.
func parseVolumeIndex(html, baseURL string) []*Problem {
	var out []*Problem
	for _, href := range hrefRE.FindAllStringSubmatch(html, -1) {
		link := href[1]
		m := categoryidRE.FindStringSubmatch(link)
		if m == nil {
			continue
		}
		volID := m[1]
		// Skip categoryid=1 (the root index itself)
		if volID == "1" {
			continue
		}
		out = append(out, &Problem{
			ID:     "vol-" + volID,
			Volume: volID,
			Title:  "Volume " + volID,
			URL:    baseURL + fmt.Sprintf(archiveVolumeBase, volID),
		})
	}
	return dedupeProblems(out)
}

// ListProblems scrapes a Live Archive volume and returns problem stubs.
func (c *Client) ListProblems(ctx context.Context, volumeID string) ([]*Problem, error) {
	url := c.cfg.ArchiveBaseURL + fmt.Sprintf(archiveVolumeBase, volumeID)
	body, err := c.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("list problems (volume %s): %w", volumeID, err)
	}
	return parseProblemList(string(body), volumeID, c.cfg.ArchiveBaseURL), nil
}

// parseProblemList extracts problem stubs from a volume's HTML.
func parseProblemList(html, volumeID, baseURL string) []*Problem {
	var out []*Problem
	for _, href := range hrefRE.FindAllStringSubmatch(html, -1) {
		link := href[1]
		m := problemRE.FindStringSubmatch(link)
		if m == nil {
			continue
		}
		probID := m[1]
		p := &Problem{
			ID:     probID,
			Volume: volumeID,
			URL:    baseURL + fmt.Sprintf(archiveProblemBase, probID),
			PDFURL: pdfURL(baseURL, volumeID, probID),
		}
		out = append(out, p)
	}
	return dedupeProblems(out)
}

// GetProblem scrapes the Live Archive single problem page.
func (c *Client) GetProblem(ctx context.Context, id string) (*Problem, error) {
	url := c.cfg.ArchiveBaseURL + fmt.Sprintf(archiveProblemBase, id)
	body, err := c.Get(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("get problem %s: %w", id, err)
	}
	p := parseProblemPage(string(body), id, c.cfg.ArchiveBaseURL)
	if p.Title == "" {
		p.Title = "Problem " + id
	}
	return p, nil
}

// parseProblemPage extracts problem fields from a single problem's HTML.
func parseProblemPage(html, id, baseURL string) *Problem {
	p := &Problem{
		ID:  id,
		URL: baseURL + fmt.Sprintf(archiveProblemBase, id),
	}
	// Try h2 for title
	if m := h2RE.FindStringSubmatch(html); m != nil {
		p.Title = stripTags(m[1])
	}
	// Try page title as fallback
	if p.Title == "" {
		if m := titleRE.FindStringSubmatch(html); m != nil {
			p.Title = stripTags(m[1])
		}
	}
	// Try to find a PDF link
	for _, href := range hrefRE.FindAllStringSubmatch(html, -1) {
		if strings.HasSuffix(href[1], ".pdf") {
			p.PDFURL = href[1]
			if !strings.HasPrefix(p.PDFURL, "http") {
				p.PDFURL = baseURL + p.PDFURL
			}
			break
		}
	}
	return p
}

// pdfURL builds the expected PDF URL for a problem given its volume and ID.
// The Live Archive stores PDFs at /external/<vol>/<id>.pdf.
func pdfURL(baseURL, volumeID, problemID string) string {
	return fmt.Sprintf("%s/external/%s/%s.pdf", baseURL, volumeID, problemID)
}

// dedupeProblems removes duplicate problems by ID.
func dedupeProblems(ps []*Problem) []*Problem {
	seen := map[string]bool{}
	var out []*Problem
	for _, p := range ps {
		if seen[p.ID] {
			continue
		}
		seen[p.ID] = true
		out = append(out, p)
	}
	return out
}

// clicsScoreboard represents the CLICS 2020 scoreboard JSON.
type clicsScoreboard struct {
	Rows []struct {
		Rank   int    `json:"rank"`
		TeamID string `json:"team_id"`
		Score  struct {
			NumSolved int `json:"num_solved"`
			TotalTime int `json:"total_time"`
		} `json:"score"`
	} `json:"rows"`
}

type clicsTeam struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	OrganizationID string `json:"organization_id"`
	Country        string `json:"country"`
}

type clicsOrganization struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// FetchCLICSStandings fetches and merges standings from a CLICS-compatible
// scoreboard URL. The URL should point to the contest root
// (e.g. https://example.com/contests/wf2023).
func (c *Client) FetchCLICSStandings(ctx context.Context, baseURL string, year int) ([]Standing, error) {
	baseURL = strings.TrimRight(baseURL, "/")

	// Fetch teams
	teamsBody, err := c.Get(ctx, baseURL+"/teams")
	if err != nil {
		return nil, fmt.Errorf("fetch teams: %w", err)
	}
	var teams []clicsTeam
	if err := json.Unmarshal(teamsBody, &teams); err != nil {
		return nil, fmt.Errorf("parse teams: %w", err)
	}
	teamMap := make(map[string]clicsTeam, len(teams))
	for _, t := range teams {
		teamMap[t.ID] = t
	}

	// Fetch organizations
	orgsBody, err := c.Get(ctx, baseURL+"/organizations")
	if err == nil {
		var orgs []clicsOrganization
		if json.Unmarshal(orgsBody, &orgs) == nil {
			for _, o := range orgs {
				// Fill in country from org if team lacks it
				if t, ok := teamMap[o.ID]; ok && t.Country == "" {
					t.Country = o.Country
					teamMap[o.ID] = t
				}
			}
		}
	}

	// Fetch scoreboard
	sbBody, err := c.Get(ctx, baseURL+"/scoreboard")
	if err != nil {
		return nil, fmt.Errorf("fetch scoreboard: %w", err)
	}
	var sb clicsScoreboard
	if err := json.Unmarshal(sbBody, &sb); err != nil {
		return nil, fmt.Errorf("parse scoreboard: %w", err)
	}

	out := make([]Standing, 0, len(sb.Rows))
	for _, row := range sb.Rows {
		t := teamMap[row.TeamID]
		out = append(out, Standing{
			Rank:       row.Rank,
			TeamName:   t.Name,
			University: t.Name,
			Country:    t.Country,
			Solved:     row.Score.NumSolved,
			TotalTime:  row.Score.TotalTime,
			Year:       year,
			Medal:      medal(row.Rank),
		})
	}
	return out, nil
}

// SearchContests searches the embedded World Finals data by query string.
// Matches against year, host, site, winner team, university, and country fields.
func SearchContests(query string) []Contest {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		return nil
	}
	var out []Contest
	for _, c := range worldFinalsData {
		if matchesContest(c, q) {
			out = append(out, c)
		}
	}
	return out
}

func matchesContest(c Contest, q string) bool {
	fields := []string{
		strconv.Itoa(c.Year),
		strings.ToLower(c.Host),
		strings.ToLower(c.Site),
		strings.ToLower(c.WinnerTeam),
		strings.ToLower(c.WinnerUniv),
		strings.ToLower(c.WinnerCountry),
	}
	for _, f := range fields {
		if strings.Contains(f, q) {
			return true
		}
	}
	return false
}
