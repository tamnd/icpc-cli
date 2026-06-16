package icpc

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/tamnd/any-cli/kit"
	"github.com/tamnd/any-cli/kit/errs"
)

// init registers the ICPC domain so ant can load it with a blank import.
func init() { kit.Register(Domain{}) }

// Domain is the ICPC driver. No state; the per-run client is built by the factory.
type Domain struct{}

// Info describes the scheme and identity for this domain.
func (Domain) Info() kit.DomainInfo {
	return kit.DomainInfo{
		Scheme: "icpc",
		Hosts:  []string{ICPCHost, ArchiveHost},
		Identity: kit.Identity{
			Binary: "icpc",
			Short:  "Browse ICPC contest data from the command line",
			Long: `Browse ICPC contest data from the command line.

icpc reads public ICPC data from the embedded World Finals dataset and the
ICPC Live Archive (icpcarchive.ecs.baylor.edu). Most commands work offline
using the built-in dataset.

Note: icpc.global requires a login session for its REST API, so live contest
listings are not available without credentials. Use the embedded dataset (default)
or supply --scoreboard-url for CLICS-compatible live scoreboards.

Quick start:
  icpc wf                    all World Finals (1977-2024)
  icpc wf --year 2024        one year
  icpc results 2024          champion from embedded data
  icpc problems --volume 1   browse Live Archive volume 1
  icpc problem 4234          fetch a single problem
  icpc search MIT            search for MIT in the dataset`,
			Site: ICPCHost,
			Repo: "https://github.com/tamnd/icpc-cli",
		},
	}
}

// Register installs the client factory and all operations onto app.
func (Domain) Register(app *kit.App) {
	app.SetClient(newClient)

	kit.Handle(app, kit.OpMeta{
		Name:    "wf",
		Group:   "contests",
		Summary: "List World Finals contests (embedded dataset, 1977-2024)",
	}, listWorldFinals)

	kit.Handle(app, kit.OpMeta{
		Name:    "contests",
		Group:   "contests",
		Summary: "List contests (alias for wf with --region filter)",
	}, listContests)

	kit.Handle(app, kit.OpMeta{
		Name:    "results",
		Group:   "standings",
		Summary: "Show World Finals standings for a year",
		Single:  true,
		URIType: "results",
		Args:    []kit.Arg{{Name: "year", Help: "World Finals year (e.g. 2024)"}},
	}, getResults)

	kit.Handle(app, kit.OpMeta{
		Name:    "teams",
		Group:   "contests",
		Summary: "List teams in a World Finals year",
		Args:    []kit.Arg{{Name: "year", Help: "World Finals year (optional)", Optional: true}},
	}, listTeams)

	kit.Handle(app, kit.OpMeta{
		Name:    "problems",
		Group:   "archive",
		Summary: "List problems from the ICPC Live Archive",
	}, listProblems)

	kit.Handle(app, kit.OpMeta{
		Name:     "problem",
		Group:    "archive",
		Summary:  "Fetch a single problem by ID from the ICPC Live Archive",
		Single:   true,
		URIType:  "problem",
		Resolver: true,
		Args:     []kit.Arg{{Name: "id", Help: "problem ID (numeric)"}},
	}, getProblem)

	kit.Handle(app, kit.OpMeta{
		Name:    "search",
		Group:   "explore",
		Summary: "Search the embedded dataset for contests, teams, or universities",
		Args:    []kit.Arg{{Name: "query", Help: "search query"}},
	}, search)
}

// newClient builds a Client from the resolved kit Config.
func newClient(_ context.Context, cfg kit.Config) (any, error) {
	c := DefaultConfig()
	if cfg.Rate > 0 {
		c.Rate = cfg.Rate
	}
	if cfg.Retries > 0 {
		c.Retries = cfg.Retries
	}
	if cfg.Timeout > 0 {
		c.Timeout = cfg.Timeout
	}
	if cfg.UserAgent != "" {
		c.UserAgent = cfg.UserAgent
	}
	return NewClient(c), nil
}

// --- input structs ---

type wfInput struct {
	Year   int     `kit:"flag,inherit" help:"filter to this year (e.g. 2024)" default:"0"`
	Client *Client `kit:"inject"`
}

type contestsInput struct {
	Year   int     `kit:"flag,inherit" help:"filter to this year" default:"0"`
	Region string  `kit:"flag" help:"filter by region (e.g. north-america, asia-pacific, europe)"`
	Client *Client `kit:"inject"`
}

type resultsInput struct {
	Year          int     `kit:"arg" help:"World Finals year (e.g. 2024)"`
	ScoreboardURL string  `kit:"flag" help:"CLICS-compatible scoreboard base URL for live data"`
	Client        *Client `kit:"inject"`
}

type teamsInput struct {
	Year   int     `kit:"arg,optional" help:"World Finals year (optional)"`
	Client *Client `kit:"inject"`
}

type problemsInput struct {
	Volume string  `kit:"flag" help:"Live Archive volume number (default: list all volumes)" default:""`
	Limit  int     `kit:"flag,inherit" help:"max results" default:"50"`
	Client *Client `kit:"inject"`
}

type problemInput struct {
	ID     string  `kit:"arg" help:"problem ID (numeric)"`
	Client *Client `kit:"inject"`
}

type searchInput struct {
	Query  string  `kit:"arg" help:"search query"`
	Client *Client `kit:"inject"`
}

// --- handlers ---

func listWorldFinals(ctx context.Context, in wfInput, emit func(Contest) error) error {
	cs := WorldFinals(in.Year)
	if len(cs) == 0 {
		return errs.NotFound("no World Finals data found for year %d", in.Year)
	}
	for _, c := range cs {
		if err := emit(c); err != nil {
			return err
		}
	}
	return nil
}

func listContests(ctx context.Context, in contestsInput, emit func(Contest) error) error {
	cs := WorldFinals(in.Year)
	if len(cs) == 0 {
		return errs.NotFound("no World Finals data found")
	}
	region := strings.ToLower(in.Region)
	for _, c := range cs {
		if region != "" && !matchesRegion(c, region) {
			continue
		}
		if err := emit(c); err != nil {
			return err
		}
	}
	return nil
}

func getResults(ctx context.Context, in resultsInput, emit func(Standing) error) error {
	year := in.Year
	if year == 0 {
		return errs.Usage("year is required (e.g. icpc results 2024)")
	}

	var standings []Standing
	if in.ScoreboardURL != "" {
		var err error
		standings, err = in.Client.FetchCLICSStandings(ctx, in.ScoreboardURL, year)
		if err != nil {
			return fmt.Errorf("fetch scoreboard: %w", err)
		}
	} else {
		standings = WorldFinalsStandings(year)
		if len(standings) == 0 {
			return errs.NotFound("no data for World Finals %d", year)
		}
	}

	for _, s := range standings {
		if err := emit(s); err != nil {
			return err
		}
	}
	return nil
}

func listTeams(_ context.Context, in teamsInput, emit func(Team) error) error {
	cs := WorldFinals(in.Year)
	if len(cs) == 0 {
		return errs.NotFound("no World Finals data found")
	}
	for _, c := range cs {
		// Emit the winning team as the representative team record from embedded data.
		// Full team lists require a live CLICS scoreboard via `icpc results --scoreboard-url`.
		t := Team{
			ID:         fmt.Sprintf("%d-winner", c.Year),
			Name:       c.WinnerTeam,
			University: c.WinnerUniv,
			Country:    c.WinnerCountry,
			Region:     "",
			Year:       c.Year,
			Rank:       1,
		}
		if err := emit(t); err != nil {
			return err
		}
	}
	return nil
}

func listProblems(ctx context.Context, in problemsInput, emit func(*Problem) error) error {
	if in.Volume == "" {
		vols, err := in.Client.ListVolumes(ctx)
		if err != nil {
			return err
		}
		if len(vols) == 0 {
			return errs.NotFound("no volumes found in the ICPC Live Archive")
		}
		for _, p := range vols {
			if err := emit(p); err != nil {
				return err
			}
		}
		return nil
	}

	probs, err := in.Client.ListProblems(ctx, in.Volume)
	if err != nil {
		return err
	}
	if len(probs) == 0 {
		return errs.NotFound("no problems found in volume %s", in.Volume)
	}
	limit := in.Limit
	for i, p := range probs {
		if limit > 0 && i >= limit {
			break
		}
		if err := emit(p); err != nil {
			return err
		}
	}
	return nil
}

func getProblem(ctx context.Context, in problemInput, emit func(*Problem) error) error {
	p, err := in.Client.GetProblem(ctx, in.ID)
	if err != nil {
		return err
	}
	return emit(p)
}

func search(_ context.Context, in searchInput, emit func(Contest) error) error {
	if in.Query == "" {
		return errs.Usage("query is required (e.g. icpc search MIT)")
	}
	results := SearchContests(in.Query)
	if len(results) == 0 {
		return errs.NotFound("no results for %q", in.Query)
	}
	for _, c := range results {
		if err := emit(c); err != nil {
			return err
		}
	}
	return nil
}

// --- Resolver ---

// Classify turns any accepted input into (uriType, id).
func (Domain) Classify(input string) (uriType, id string, err error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", "", errs.Usage("icpc: empty input")
	}
	// Numeric: 4-digit numbers in plausible ICPC year range are years; otherwise problem ID
	if n, e := strconv.Atoi(input); e == nil {
		if len(input) == 4 && n >= 1977 && n <= 2099 {
			return "results", input, nil
		}
		return "problem", input, nil
	}
	return "", "", errs.Usage("icpc: unrecognized reference: %q", input)
}

// Locate returns the canonical URL for a (uriType, id).
func (Domain) Locate(uriType, id string) (string, error) {
	switch uriType {
	case "results":
		return fmt.Sprintf("https://%s/worldfinals", ICPCHost), nil
	case "problem":
		return ArchiveBaseURL + fmt.Sprintf(archiveProblemBase, id), nil
	default:
		return "", errs.Usage("icpc has no resource type %q", uriType)
	}
}

// matchesRegion returns true if c's host matches the region slug.
// Supported slugs: north-america, asia-pacific, europe, latin-america, africa, middle-east.
func matchesRegion(c Contest, region string) bool {
	host := strings.ToLower(c.Host)
	switch region {
	case "north-america", "na":
		return strings.Contains(host, "usa") || strings.Contains(host, "canada")
	case "europe", "eu":
		return strings.Contains(host, "poland") || strings.Contains(host, "russia") ||
			strings.Contains(host, "sweden") || strings.Contains(host, "netherlands") ||
			strings.Contains(host, "portugal") || strings.Contains(host, "czech") ||
			strings.Contains(host, "germany")
	case "asia-pacific", "asia", "ap":
		return strings.Contains(host, "china") || strings.Contains(host, "japan") ||
			strings.Contains(host, "korea") || strings.Contains(host, "thailand") ||
			strings.Contains(host, "singapore") || strings.Contains(host, "australia") ||
			strings.Contains(host, "new zealand") || strings.Contains(host, "bangladesh") ||
			strings.Contains(host, "kazakhstan")
	case "latin-america", "latam":
		return strings.Contains(host, "brazil") || strings.Contains(host, "argentina") ||
			strings.Contains(host, "mexico") || strings.Contains(host, "peru")
	case "africa":
		return strings.Contains(host, "egypt") || strings.Contains(host, "morocco") ||
			strings.Contains(host, "africa")
	case "middle-east", "me":
		return strings.Contains(host, "egypt") || strings.Contains(host, "jordan") ||
			strings.Contains(host, "saudi")
	}
	// Generic substring match
	return strings.Contains(host, region)
}
