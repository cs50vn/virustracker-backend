package model

type InputItem struct {
    LeftName string `json: "LeftName"`
    RightName string `json: "RightName"`
    TotalCases int64 `json: "TotalCases"`
    TotalDeaths int64 `json: "TotalDeaths"`
    TotalRecovered int64 `json: "TotalRecovered"`
    SeriousCases int64 `json: "SeriousCases"`
    TotalCasesPer1Pop float64 `json: "TotalCasesPer1Pop"`
    TotalDeathsPer1Pop float64 `json: "TotalDeathsPer1Pop"`
    TotalTests int64 `json: "TotalTests"`
    TestsPer1Pop float64 `json: "TestsPer1Pop"`
    Timestamp int64 `json: "Timestamp"`
}

func MakeInputItem(leftName string, rightName string, totalCases int64, totalDeaths int64, totalRecovered int64, seriousCases int64, totalCasesPer1Pop float64, totalDeathsPer1Pop float64, totalTests int64, testsPer1Pop float64, timestamp int64) *InputItem {
    return &InputItem{LeftName: leftName, RightName: rightName, TotalCases: totalCases, TotalDeaths: totalDeaths, TotalRecovered: totalRecovered, SeriousCases: seriousCases, TotalCasesPer1Pop: totalCasesPer1Pop, TotalDeathsPer1Pop: totalDeathsPer1Pop, TotalTests: totalTests, TestsPer1Pop: testsPer1Pop, Timestamp: timestamp}
}

