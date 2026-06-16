package icpc

// Contest represents one ICPC World Finals event.
type Contest struct {
	Year          int    `json:"year"`
	Host          string `json:"host"`           // city and country
	Site          string `json:"site"`           // venue name
	Problems      int    `json:"problems"`       // number of problems
	Teams         int    `json:"teams"`          // number of competing teams
	WinnerTeam    string `json:"winner_team"`    // team name of winner
	WinnerUniv    string `json:"winner_univ"`    // university of winner
	WinnerCountry string `json:"winner_country"` // country of winner
	URL           string `json:"url"`
}

// Team represents one participating team in a World Finals.
type Team struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	University string `json:"university"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	Year       int    `json:"year"`
	Rank       int    `json:"rank"`
}

// Standing represents one team's final position in a contest.
type Standing struct {
	Rank       int    `json:"rank"`
	TeamName   string `json:"team_name"`
	University string `json:"university"`
	Country    string `json:"country"`
	Solved     int    `json:"solved"`
	TotalTime  int    `json:"total_time"`
	Year       int    `json:"year"`
	Medal      string `json:"medal"` // gold, silver, bronze, or ""
}

// Problem represents a problem from the ICPC Live Archive.
type Problem struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Volume string `json:"volume"`
	Number string `json:"number"`
	PDFURL string `json:"pdf_url"`
	URL    string `json:"url"`
}
