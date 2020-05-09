package utils

import (
    "cs50vn/virustracker/apprepository"
    "cs50vn/virustracker/apprepository/model"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    "fmt"
    "log"
    "strconv"
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

func GenerateCountryDetailStatus(countryId string) string {
    var data = ""

    if item, ok := apprepository.TopCountriesList[countryId]; ok {
       activeCases := item.Items[0].TotalCases - item.Items[0].TotalDeaths - item.Items[0].TotalRecovered
       data += fmt.Sprintf(ITEM2_TEMPLATE, "totalCases", item.Items[0].TotalCases)+","+
           fmt.Sprintf(ITEM2_TEMPLATE, "newCases", item.Items[0].NewCases) + "," +
           fmt.Sprintf(ITEM2_TEMPLATE, "totalDeaths", item.Items[0].TotalDeaths) + "," +
           fmt.Sprintf(ITEM2_TEMPLATE, "newDeaths", item.Items[0].NewDeaths) + "," +
           fmt.Sprintf(ITEM2_TEMPLATE, "totalRecovered", item.Items[0].TotalRecovered) + "," +
           fmt.Sprintf(ITEM2_TEMPLATE, "activeCases", activeCases) + "," +
           fmt.Sprintf(ITEM2_TEMPLATE, "seriousCases", item.Items[0].SeriousCases) + "," +
           fmt.Sprintf(ITEM_TEMPLATE, "totalCasesPer1Pop", strconv.FormatFloat(item.Items[0].TotalCasesPer1Pop, 'f', 2, 64)) + "," +
           fmt.Sprintf(ITEM_TEMPLATE, "totalDeathsPer1Pop", strconv.FormatFloat(item.Items[0].TotalDeathsPer1Pop, 'f', 2, 64)) + "," +
           fmt.Sprintf(ITEM2_TEMPLATE, "totalTests", item.Items[0].TotalTests) + "," +
           fmt.Sprintf(ITEM_TEMPLATE, "testsPer1Pop", strconv.FormatFloat(item.Items[0].TestsPer1Pop, 'f', 2, 64)) + "," +
           fmt.Sprintf(ITEM4_TEMPLATE, "totalCasesChart", GenerateTotalCasesChart(item)) + "," +
           fmt.Sprintf(ITEM4_TEMPLATE, "totalDeathsChart", GenerateTotalDeathsChart(item))

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

func GenerateTotalCasesChart(country *model.Country) string {
    var data = ""

    t := time.Now().Unix()

    for i := 0; i < apprepository.LimitCountryChart; i++ {
        item := GetTotalCasesChartItem(country, t)
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM2_TEMPLATE, "timestamp", t) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.TotalCases)) + ","

        t = t - 86400
    }
    data = strings.TrimSuffix(data, ",")

    return data
}

func GetTotalCasesChartItem(country *model.Country, timestamp int64) *model.Item {

    if item, _, ok := CheckDayInExist(country, timestamp); ok {
        return item
    } else if timestamp < country.Items[len(country.Items) - 1].Timestamp {
        return model.MakeItem(0,0,0,0,0,0,0,0,0,0,0, 0, 0)
    } else {
        return GetTotalDeathsChartItem(country, timestamp - 86400)
    }

}

func GenerateTotalDeathsChart(country *model.Country) string {
    var data = ""

    t := time.Now().Unix()

    for i := 0; i < apprepository.LimitCountryChart; i++ {
        //tm := time.Unix(t, 0)
        item := GetTotalDeathsChartItem(country, t)
        data += fmt.Sprintf(OBJECT_TEMPLATE, fmt.Sprintf(ITEM2_TEMPLATE, "timestamp", t) + "," +
            fmt.Sprintf(ITEM2_TEMPLATE, "value", item.TotalDeaths)) + ","

        t = t - 86400
    }
    data = strings.TrimSuffix(data, ",")

    return data
}


func GetTotalDeathsChartItem(country *model.Country, timestamp int64) *model.Item {

    if item, _, ok := CheckDayInExist(country, timestamp); ok {
        return item
    } else if timestamp < country.Items[len(country.Items) - 1].Timestamp {
        return model.MakeItem(0,0,0,0,0,0,0,0,0,0,0, 0, 0)
    } else {
        return GetTotalDeathsChartItem(country, timestamp - 86400)
    }

}

func GetItemByTimestamp(country *model.Country, timestamp int64) *model.Item {

    if item, _, ok := CheckDayInExist(country, timestamp); ok {
        return item
    } else if timestamp < country.Items[len(country.Items) - 1].Timestamp {
        return model.MakeItem(0,0,0,0,0,0,0,0,0,0,0, 0, 0)
    } else {
        return GetTotalDeathsChartItem(country, timestamp - 86400)
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

func InsertItem(country *model.Country,  item *model.Item) (int64, interface{}) {
    result := ExecuteSQL(SQL_INSERT_ITEM, item.TotalCases, item.TotalDeaths, item.TotalRecovered, item.SeriousCases, item.TotalCasesPer1Pop, item.TotalDeathsPer1Pop, item.TotalTests, item.TestsPer1Pop, item.Timestamp, country.Id)

    rowAffected, _ := result.RowsAffected()
    if rowAffected != 1 {
        fmt.Println("Insert failed!")
    }

    return result.LastInsertId()
}

///////////////////////////////////////////////////////////////////////////////////////
