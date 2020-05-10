package utils

import (
    "cs50vn/virustracker/apprepository"
    "cs50vn/virustracker/apprepository/model"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "strings"
    "time"
)

func GenerateChartTotalCasesJson() string {
    var data = ""

    for _, item := range apprepository.AppModel.TotalCasesChart {
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM_TEMPLATE, "continentId", item.Id)+","+
            fmt.Sprintf(ITEM_TEMPLATE, "continentName", item.Name)+","+
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.Value)) + ","
    }
    data = strings.TrimSuffix(data, ",")

    return data
}

func GenerateChartTotalDeathsJson() string {
    var data = ""

    for _, item := range apprepository.AppModel.TotalDeathsChart {
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM_TEMPLATE, "continentId", item.Id)+","+
            fmt.Sprintf(ITEM_TEMPLATE, "continentName", item.Name)+","+
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.Value)) + ","
    }
    data = strings.TrimSuffix(data, ",")

    return data
}

func GenerateRecentTotalCasesJson() string {
    var data = ""

    for _, item := range apprepository.AppModel.TotalCasesRecent {
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM2_TEMPLATE, "timestamp", item.Timestamp)+","+
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.Value)) + ","
    }
    data = strings.TrimSuffix(data, ",")

    return data
}

func GenerateRecentTotalDeathsJson() string {
    var data = ""

    for _, item := range apprepository.AppModel.TotalDeathsRecent {
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM2_TEMPLATE, "timestamp", item.Timestamp)+","+
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.Value)) + ","
    }
    data = strings.TrimSuffix(data, ",")

    return data
}

func GenerateCountryDetailStatus(countryId string, t int64) string {
    var data = ""

    if country, ok := apprepository.TopCountriesList[countryId]; ok {
        item := GetItemByTimestamp(country, t)
        activeCases := item.TotalCases - item.TotalDeaths - item.TotalRecovered
        data += fmt.Sprintf(ITEM2_TEMPLATE, "totalCases", item.TotalCases) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "newCases", item.TotalCases-GetItemByTimestamp(country, t-86400).TotalCases) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "totalDeaths", item.TotalDeaths) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "newDeaths", item.TotalDeaths-GetItemByTimestamp(country, t-86400).TotalDeaths) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "totalRecovered", item.TotalRecovered) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "activeCases", activeCases) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "seriousCases", item.SeriousCases) + "," +
            fmt.Sprintf(ITEM3_TEMPLATE, "totalCasesPer1Pop", item.TotalCasesPer1Pop) + "," +
            fmt.Sprintf(ITEM3_TEMPLATE, "totalDeathsPer1Pop", item.TotalDeathsPer1Pop) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "totalTests", item.TotalTests) + "," +
            fmt.Sprintf(ITEM3_TEMPLATE, "testsPer1Pop", item.TestsPer1Pop) + "," +
            fmt.Sprintf(ITEM4_TEMPLATE, "totalCasesChart", GenerateTotalCasesChart(country, t)) + "," +
            fmt.Sprintf(ITEM4_TEMPLATE, "totalDeathsChart", GenerateTotalDeathsChart(country, t))

    }

    return data
}

func CheckDayInExist(country *model.Country, timestamp int64) (*model.Item, int, bool) {
    var index = 0

    for _, item := range country.Items {
        tm := time.Unix(timestamp, 0)
        tm1 := time.Unix(item.Timestamp, 0)

        if tm.Day() == tm1.Day() && tm.Month() == tm1.Month() && tm.Year() == tm1.Year() {
            return item, index, true
        }

        index++
    }

    return nil, -1, false
}

func InsertDayInItemList(country *model.Country, item *model.Item) {
    if len(country.Items) == 0 {
        country.Items = append(country.Items, item)
    }

    for index := 0; index < len(country.Items); index++ {
        if item.Timestamp > country.Items[index].Timestamp {
            country.Items = append(country.Items, nil /* use the zero value of the element type */)
            copy(country.Items[index+1:], country.Items[index:])
            country.Items[index] = item
            break
        }
    }
}

func GenerateTotalCasesChart(country *model.Country, t int64) string {
    var data = ""

    for i := 0; i < apprepository.LimitCountryChart; i++ {
        item := GetTotalCasesChartItem(country, t)
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM2_TEMPLATE, "timestamp", t)+","+
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.TotalCases)) + ","

        t = t - 86400
    }
    data = strings.TrimSuffix(data, ",")

    return data
}

func GetTotalCasesChartItem(country *model.Country, timestamp int64) *model.Item {

    if item, _, ok := CheckDayInExist(country, timestamp); ok {
        return item
    } else if timestamp < country.Items[len(country.Items)-1].Timestamp {
        return model.MakeItem(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
    } else {
        return GetTotalDeathsChartItem(country, timestamp-86400)
    }

}

func GenerateTotalDeathsChart(country *model.Country, t int64) string {
    var data = ""

    for i := 0; i < apprepository.LimitCountryChart; i++ {
        //tm := time.Unix(t, 0)
        item := GetTotalDeathsChartItem(country, t)
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM2_TEMPLATE, "timestamp", t)+","+
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.TotalDeaths)) + ","

        t = t - 86400
    }
    data = strings.TrimSuffix(data, ",")

    return data
}

func GetTotalDeathsChartItem(country *model.Country, timestamp int64) *model.Item {

    if item, _, ok := CheckDayInExist(country, timestamp); ok {
        return item
    } else if timestamp < country.Items[len(country.Items)-1].Timestamp {
        return model.MakeItem(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
    } else {
        return GetTotalDeathsChartItem(country, timestamp-86400)
    }

}

func GetItemByTimestamp(country *model.Country, timestamp int64) *model.Item {

    if item, _, ok := CheckDayInExist(country, timestamp); ok {
        return item
    } else if timestamp < country.Items[len(country.Items)-1].Timestamp {
        return model.MakeItem(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)
    } else {
        return GetTotalDeathsChartItem(country, timestamp-86400)
    }

}

///////////////////////////////////////////////////////////////////////////////////////
//DB Access

func QuerySQL(sql string, params ...interface{}) (*sql.Rows, interface{}) {
    return nil, nil
}

func ExecuteSQL(sql string, params ...interface{}) sql.Result {
    tx, err := apprepository.DbConnection.Begin()
    if err != nil {
        fmt.Println(err.Error())
    }

    prepareSt, err := apprepository.DbConnection.Prepare(sql)
    if err != nil {
        log.Fatal(err)
    }

    result, _ := prepareSt.Exec(params...)

    tx.Commit()

    return result
}

func DeleteItem(id int64) {
    result := ExecuteSQL(SQL_DELETE_ITEM, id)
    rowAffected, _ := result.RowsAffected()
    if rowAffected != 1 {
        fmt.Println("Delete failed!")
    }

}

func InsertItem(country *model.Country, item *model.Item) (int64, interface{}) {
    result := ExecuteSQL(SQL_INSERT_ITEM, item.TotalCases, item.TotalDeaths, item.TotalRecovered, item.SeriousCases, item.TotalCasesPer1Pop, item.TotalDeathsPer1Pop, item.TotalTests, item.TestsPer1Pop, item.Timestamp, country.Id)

    rowAffected, _ := result.RowsAffected()
    if rowAffected != 1 {
        fmt.Println("Insert failed!")
    }

    return result.LastInsertId()
}

///////////////////////////////////////////////////////////////////////////////////////
