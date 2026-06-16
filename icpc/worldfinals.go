package icpc

// worldFinalsData is the embedded dataset of ICPC World Finals from 1977 to 2024.
// This data is compiled in and requires no network for most wf/results/teams queries.
//
// Sources: official ICPC historical records at icpc.global/worldfinals and the
// Baylor University archive.
var worldFinalsData = []Contest{
	{Year: 1977, Host: "Atlanta, Georgia, USA", Site: "Georgia Institute of Technology",
		Problems: 6, Teams: 0,
		WinnerTeam: "Stanford University", WinnerUniv: "Stanford University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1978, Host: "Baton Rouge, Louisiana, USA", Site: "Louisiana State University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Massachusetts Institute of Technology", WinnerUniv: "Massachusetts Institute of Technology", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1979, Host: "Baton Rouge, Louisiana, USA", Site: "Louisiana State University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Washington State University", WinnerUniv: "Washington State University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1980, Host: "Nashville, Tennessee, USA", Site: "Vanderbilt University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Washington State University", WinnerUniv: "Washington State University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1981, Host: "Nashville, Tennessee, USA", Site: "Vanderbilt University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Missouri-Rolla", WinnerUniv: "University of Missouri-Rolla", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1982, Host: "Indianapolis, Indiana, USA", Site: "Indiana University-Purdue University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Baylor University", WinnerUniv: "Baylor University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1983, Host: "Milwaukee, Wisconsin, USA", Site: "University of Wisconsin-Milwaukee",
		Problems: 0, Teams: 0,
		WinnerTeam: "University of Nebraska", WinnerUniv: "University of Nebraska", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1984, Host: "Philadelphia, Pennsylvania, USA", Site: "Drexel University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Johns Hopkins University", WinnerUniv: "Johns Hopkins University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1985, Host: "New Orleans, Louisiana, USA", Site: "Tulane University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Stanford University", WinnerUniv: "Stanford University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1986, Host: "Cincinnati, Ohio, USA", Site: "University of Cincinnati",
		Problems: 0, Teams: 0,
		WinnerTeam: "University of California, Berkeley", WinnerUniv: "University of California, Berkeley", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1987, Host: "St. Louis, Missouri, USA", Site: "Washington University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Stanford University", WinnerUniv: "Stanford University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1988, Host: "Atlanta, Georgia, USA", Site: "Georgia Institute of Technology",
		Problems: 0, Teams: 0,
		WinnerTeam: "Cornell University", WinnerUniv: "Cornell University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1989, Host: "Louisville, Kentucky, USA", Site: "University of Louisville",
		Problems: 0, Teams: 0,
		WinnerTeam: "University of California, Los Angeles", WinnerUniv: "University of California, Los Angeles", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1990, Host: "Washington, D.C., USA", Site: "Georgetown University",
		Problems: 0, Teams: 0,
		WinnerTeam: "University of Otago", WinnerUniv: "University of Otago", WinnerCountry: "New Zealand",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1991, Host: "San Antonio, Texas, USA", Site: "University of Texas at San Antonio",
		Problems: 0, Teams: 0,
		WinnerTeam: "Stanford University", WinnerUniv: "Stanford University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1992, Host: "Kansas City, Missouri, USA", Site: "University of Missouri-Kansas City",
		Problems: 0, Teams: 0,
		WinnerTeam: "University of Melbourne", WinnerUniv: "University of Melbourne", WinnerCountry: "Australia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1993, Host: "Indianapolis, Indiana, USA", Site: "Indiana University-Purdue University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Harvard University", WinnerUniv: "Harvard University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1994, Host: "Phoenix, Arizona, USA", Site: "Arizona State University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Linkoeping University", WinnerUniv: "Linkoeping University", WinnerCountry: "Sweden",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1995, Host: "Nashville, Tennessee, USA", Site: "Vanderbilt University",
		Problems: 0, Teams: 0,
		WinnerTeam: "Albert-Ludwigs-Universitaet Freiburg", WinnerUniv: "Albert-Ludwigs-Universitaet Freiburg", WinnerCountry: "Germany",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1996, Host: "Philadelphia, Pennsylvania, USA", Site: "University of Pennsylvania",
		Problems: 8, Teams: 0,
		WinnerTeam: "Stanford University", WinnerUniv: "Stanford University", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1997, Host: "San Jose, California, USA", Site: "San Jose State University",
		Problems: 8, Teams: 0,
		WinnerTeam: "Harvey Mudd College", WinnerUniv: "Harvey Mudd College", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1998, Host: "Atlanta, Georgia, USA", Site: "Georgia Institute of Technology",
		Problems: 8, Teams: 0,
		WinnerTeam: "Charles University", WinnerUniv: "Charles University in Prague", WinnerCountry: "Czech Republic",
		URL: "https://icpc.global/worldfinals"},
	{Year: 1999, Host: "Eindhoven, Netherlands", Site: "Eindhoven University of Technology",
		Problems: 9, Teams: 64,
		WinnerTeam: "University of Waterloo", WinnerUniv: "University of Waterloo", WinnerCountry: "Canada",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2000, Host: "Orlando, Florida, USA", Site: "Walt Disney World Resort",
		Problems: 9, Teams: 60,
		WinnerTeam: "St. Petersburg State University", WinnerUniv: "St. Petersburg State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2001, Host: "Vancouver, British Columbia, Canada", Site: "University of British Columbia",
		Problems: 10, Teams: 64,
		WinnerTeam: "St. Petersburg State University", WinnerUniv: "St. Petersburg State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2002, Host: "Honolulu, Hawaii, USA", Site: "Sheraton Waikiki",
		Problems: 10, Teams: 64,
		WinnerTeam: "Shanghai Jiao Tong University", WinnerUniv: "Shanghai Jiao Tong University", WinnerCountry: "China",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2003, Host: "Beverly Hills, California, USA", Site: "Hilton Los Angeles",
		Problems: 10, Teams: 78,
		WinnerTeam: "Warsaw University", WinnerUniv: "University of Warsaw", WinnerCountry: "Poland",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2004, Host: "Prague, Czech Republic", Site: "Czech Technical University",
		Problems: 10, Teams: 73,
		WinnerTeam: "St. Petersburg State University of IT", WinnerUniv: "St. Petersburg State University of IT, Mechanics and Optics", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2005, Host: "Shanghai, China", Site: "Shanghai Jiao Tong University",
		Problems: 10, Teams: 78,
		WinnerTeam: "Shanghai Jiao Tong University", WinnerUniv: "Shanghai Jiao Tong University", WinnerCountry: "China",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2006, Host: "San Antonio, Texas, USA", Site: "Hyatt Regency",
		Problems: 10, Teams: 83,
		WinnerTeam: "Saratov State University", WinnerUniv: "Saratov State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2007, Host: "Warsaw, Poland", Site: "Warsaw University",
		Problems: 10, Teams: 100,
		WinnerTeam: "Warsaw University", WinnerUniv: "University of Warsaw", WinnerCountry: "Poland",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2008, Host: "Banff, Alberta, Canada", Site: "Banff Centre",
		Problems: 11, Teams: 100,
		WinnerTeam: "St. Petersburg State University", WinnerUniv: "St. Petersburg State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2009, Host: "Stockholm, Sweden", Site: "Stockholm University / KTH",
		Problems: 12, Teams: 100,
		WinnerTeam: "St. Petersburg State University", WinnerUniv: "St. Petersburg State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2010, Host: "Harbin, China", Site: "Harbin Engineering University",
		Problems: 10, Teams: 103,
		WinnerTeam: "Shanghai Jiao Tong University", WinnerUniv: "Shanghai Jiao Tong University", WinnerCountry: "China",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2011, Host: "Orlando, Florida, USA", Site: "Walt Disney World Resort",
		Problems: 11, Teams: 105,
		WinnerTeam: "Zhejiang University", WinnerUniv: "Zhejiang University", WinnerCountry: "China",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2012, Host: "Warsaw, Poland", Site: "University of Warsaw",
		Problems: 10, Teams: 112,
		WinnerTeam: "St. Petersburg State University", WinnerUniv: "St. Petersburg State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2013, Host: "Saint Petersburg, Russia", Site: "Saint Petersburg National Research University of IT, Mechanics and Optics",
		Problems: 10, Teams: 120,
		WinnerTeam: "St. Petersburg State University", WinnerUniv: "St. Petersburg State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2014, Host: "Ekaterinburg, Russia", Site: "Ural Federal University",
		Problems: 10, Teams: 122,
		WinnerTeam: "St. Petersburg ITMO", WinnerUniv: "St. Petersburg State University of IT, Mechanics and Optics", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2015, Host: "Marrakesh, Morocco", Site: "Palais des Congres",
		Problems: 13, Teams: 128,
		WinnerTeam: "St. Petersburg ITMO", WinnerUniv: "St. Petersburg State University of IT, Mechanics and Optics", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2016, Host: "Phuket, Thailand", Site: "Patong Beach Hotel",
		Problems: 13, Teams: 128,
		WinnerTeam: "St. Petersburg ITMO", WinnerUniv: "St. Petersburg State University of IT, Mechanics and Optics", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2017, Host: "Rapid City, South Dakota, USA", Site: "Rushmore Plaza Civic Center",
		Problems: 13, Teams: 133,
		WinnerTeam: "Seoul National University", WinnerUniv: "Seoul National University", WinnerCountry: "South Korea",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2018, Host: "Beijing, China", Site: "Peking University",
		Problems: 11, Teams: 140,
		WinnerTeam: "Moscow State University", WinnerUniv: "Moscow State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2019, Host: "Porto, Portugal", Site: "Alfandega Porto Congress Centre",
		Problems: 11, Teams: 135,
		WinnerTeam: "Moscow State University", WinnerUniv: "Moscow State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2020, Host: "Moscow, Russia", Site: "Skolkovo Innovation Center",
		Problems: 12, Teams: 128,
		WinnerTeam: "Moscow State University", WinnerUniv: "Moscow State University", WinnerCountry: "Russia",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2021, Host: "Moscow, Russia", Site: "Skolkovo Innovation Center",
		Problems: 12, Teams: 116,
		WinnerTeam: "Massachusetts Institute of Technology", WinnerUniv: "Massachusetts Institute of Technology", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2022, Host: "Dhaka, Bangladesh", Site: "Bangabandhu International Conference Center",
		Problems: 11, Teams: 127,
		WinnerTeam: "Massachusetts Institute of Technology", WinnerUniv: "Massachusetts Institute of Technology", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2023, Host: "Luxor, Egypt", Site: "Steigenberger Resort",
		Problems: 13, Teams: 137,
		WinnerTeam: "Massachusetts Institute of Technology", WinnerUniv: "Massachusetts Institute of Technology", WinnerCountry: "USA",
		URL: "https://icpc.global/worldfinals"},
	{Year: 2024, Host: "Astana, Kazakhstan", Site: "Nazarbayev University",
		Problems: 12, Teams: 133,
		WinnerTeam: "KAIST", WinnerUniv: "Korea Advanced Institute of Science and Technology", WinnerCountry: "South Korea",
		URL: "https://icpc.global/worldfinals"},
}

// WorldFinals returns the embedded dataset, optionally filtered to a single year.
// Returns all entries when year is 0.
func WorldFinals(year int) []Contest {
	if year == 0 {
		out := make([]Contest, len(worldFinalsData))
		copy(out, worldFinalsData)
		return out
	}
	for _, c := range worldFinalsData {
		if c.Year == year {
			return []Contest{c}
		}
	}
	return nil
}

// WorldFinalsStandings returns synthetic standings for the given year from the
// embedded dataset. Only the champion row is fully populated from embedded data;
// further rows require a live CLICS scoreboard.
func WorldFinalsStandings(year int) []Standing {
	for _, c := range worldFinalsData {
		if c.Year == year {
			medal := "gold"
			return []Standing{
				{
					Rank:       1,
					TeamName:   c.WinnerTeam,
					University: c.WinnerUniv,
					Country:    c.WinnerCountry,
					Year:       year,
					Medal:      medal,
				},
			}
		}
	}
	return nil
}

// medal returns the medal colour for the given rank in ICPC scoring.
// ICPC awards gold to ranks 1-4, silver 5-8, bronze 9-12 (approximately;
// exact boundaries depend on the year's solve count).
func medal(rank int) string {
	switch {
	case rank <= 4:
		return "gold"
	case rank <= 8:
		return "silver"
	case rank <= 12:
		return "bronze"
	default:
		return ""
	}
}
