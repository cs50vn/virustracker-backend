package model

type AppItem struct {
	Timestamp int64
	TotalCases int64
	NewCases int64
	TotalDeaths int64
	NewDeaths int64
	TotalRecovered int64
	TotalCasesChart []*Continent
	TotalDeathsChart []*Continent
	TotalCasesRecent []*RecentItem
	TotalDeathsRecent []*RecentItem
}

func MakeAppItem(timestamp int64, totalCases int64, newCases int64, totalDeaths int64, newDeaths int64, totalRecovered int64, totalCasesChart []*Continent, totalDeathsChart []*Continent, totalCasesRecent []*RecentItem, totalDeathsRecent []*RecentItem) *AppItem {
	return &AppItem{Timestamp: timestamp, TotalCases: totalCases, NewCases: newCases, TotalDeaths: totalDeaths, NewDeaths: newDeaths, TotalRecovered: totalRecovered, TotalCasesChart: totalCasesChart, TotalDeathsChart: totalDeathsChart, TotalCasesRecent: totalCasesRecent, TotalDeathsRecent: totalDeathsRecent}
}
