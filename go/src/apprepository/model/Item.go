package model

type Item struct {
    Id int64
    TotalCases int64
    NewCases int64
    TotalDeaths int64
    NewDeaths int64
    TotalRecovered int64
    SeriousCases int64
    TotalCasesPer1Pop float64
    TotalDeathsPer1Pop float64
    TotalTests int64
    TestsPer1Pop float64
    Timestamp int64
    ComparePreviousDay int
}

func MakeItem(id int64, totalCases int64, newCases int64, totalDeaths int64, newDeaths int64, totalRecovered int64, seriousCases int64, totalCasesPer1Pop float64, totalDeathsPer1Pop float64, totalTests int64, testsPer1Pop float64, timestamp int64, comparePreviousDay int) *Item {
    return &Item{Id: id, TotalCases: totalCases, NewCases: newCases, TotalDeaths: totalDeaths, NewDeaths: newDeaths, TotalRecovered: totalRecovered, SeriousCases: seriousCases, TotalCasesPer1Pop: totalCasesPer1Pop, TotalDeathsPer1Pop: totalDeathsPer1Pop, TotalTests: totalTests, TestsPer1Pop: testsPer1Pop, Timestamp: timestamp, ComparePreviousDay: comparePreviousDay}
}