package app

import (
    "cs50vn/virustracker/apprepository"
    "cs50vn/virustracker/apprepository/model"
    "cs50vn/virustracker/utils"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "time"
)

func isValidDoc(doc []*model.InputItem) bool {

    return true
}

func FakeData() []*model.InputItem {
    doc := make([]*model.InputItem, 0)
    doc = append(doc, model.MakeInputItem("", "USA", 0, 0, 0,0, 0 ,0,0,0, 1588032000))
    //doc = append(doc, model.MakeInputItem("", "USA", 0, 0, 0,0, 0 ,0 ,0,0, 1588698000))
    //doc = append(doc, model.MakeInputItem("", "SAU", 0, 0, 0,0, 0 ,0 ,0,0, 1588698000))

    return doc
}

func CalculateAppItem() {
    fmt.Println("Begin Calculate app item =================")
    timestamp := time.Now().Unix()
    var totalCases int64
    var totalDeaths int64
    var totalRecovered int64
    var newCases int64
    var newDeaths int64

    for _, country := range apprepository.TopCountriesList {
        totalCases = totalCases + utils.GetItemByTimestamp(country, timestamp).TotalCases
        totalDeaths = totalDeaths + utils.GetItemByTimestamp(country, timestamp).TotalDeaths
        totalRecovered = totalRecovered + utils.GetItemByTimestamp(country, timestamp).TotalRecovered

        newCases = newCases + utils.GetItemByTimestamp(country, timestamp).TotalCases - utils.GetItemByTimestamp(country, timestamp - 86400).TotalCases
        newDeaths = newDeaths + utils.GetItemByTimestamp(country, timestamp).TotalDeaths - utils.GetItemByTimestamp(country, timestamp - 86400).TotalDeaths
    }

    totalCasesChart := make([]*model.Continent, 0)
    totalDeathsChart := make([]*model.Continent, 0)
    for _, item := range apprepository.ContinentList {
        var value int64
        var value2 int64
        for _, country := range apprepository.TopCountriesList {
            if item.Id == country.Continent.Id {
                value = value + country.Items[0].TotalCases
                value2 = value2 + country.Items[0].TotalDeaths
            }
        }
        totalCasesChart = append(totalCasesChart, model.MakeContinent(item.Id, item.Name, value))
        totalDeathsChart = append(totalDeathsChart, model.MakeContinent(item.Id, item.Name, value2))
    }

    totalCasesRecent := make([]*model.RecentItem, 0)
    totalDeathsRecent := make([]*model.RecentItem, 0)
    t := timestamp
    for i := 0; i < apprepository.LimitCountryChart; i++ {
        var item *model.Item
        var totalCases int64
        var totalDeaths int64
        for _, country := range apprepository.TopCountriesList {
            item = utils.GetItemByTimestamp(country, t)
            totalCases = totalCases + item.TotalCases
            totalDeaths = totalDeaths + item.TotalDeaths
        }

        totalCasesRecent = append(totalCasesRecent, model.MakeRecentItem(t, totalCases))
        totalDeathsRecent = append(totalDeathsRecent, model.MakeRecentItem(t, totalDeaths))

        t = t - 86400
    }

    appItem := model.MakeAppItem(timestamp, totalCases, newCases, totalDeaths, newDeaths, totalRecovered, totalCasesChart, totalDeathsChart, totalCasesRecent, totalDeathsRecent)

    apprepository.AppModel = appItem

    //Write to db
    fmt.Println("Begin write=================")

    utils.ExecuteSQL(utils.SQL_DELETE_APP_ITEM)
    utils.ExecuteSQL(utils.SQL_DELETE_APP_CHART_ITEM)
    utils.ExecuteSQL(utils.SQL_DELETE_APP_RECENT_ITEM)

    utils.ExecuteSQL(utils.SQL_INSERT_APP_ITEM, appItem.Timestamp, appItem.TotalCases, appItem.NewCases, appItem.TotalDeaths, appItem.NewDeaths, appItem.TotalRecovered)

    for _, item := range totalCasesChart {
        utils.ExecuteSQL(utils.SQL_INSERT_APP_CHART_ITEM, item.Id, item.Name, item.Value, 1)
    }

    for _, item := range totalDeathsChart {
        utils.ExecuteSQL(utils.SQL_INSERT_APP_CHART_ITEM, item.Id, item.Name, item.Value, 2)
    }

    for _, item := range totalCasesRecent {
        utils.ExecuteSQL(utils.SQL_INSERT_APP_RECENT_ITEM, item.Timestamp, item.Value, 1)
    }

    for _, item := range totalDeathsRecent {
        utils.ExecuteSQL(utils.SQL_INSERT_APP_RECENT_ITEM, item.Timestamp, item.Value, 2)
    }


    fmt.Println("End =================")
}

func Update(doc []*model.InputItem) {
    currentTime := time.Now().Unix()

    //Extract to country item
    for _, inputItem := range doc {
        //Write to mem
        if country, ok := apprepository.TopCountriesList[inputItem.RightName]; ok {
            fmt.Println("Before 1")
            fmt.Println(len(country.Items))
            if item, index, ok := utils.CheckDayInExist(country, currentTime); ok {
                //Delete old day
                utils.DeleteItem(item.Id)
                country.Items = append(country.Items[:index], country.Items[index+1:]...)
                fmt.Println("Before 2")
            }

            //Add a new day
            fmt.Println(inputItem.TotalCasesPer1Pop, inputItem.TotalDeathsPer1Pop, inputItem.TotalDeathsPer1Pop)

            it := model.MakeItem(0, inputItem.TotalCases, 0, inputItem.TotalDeaths, 0, inputItem.TotalRecovered, inputItem.SeriousCases, inputItem.TotalCasesPer1Pop, inputItem.TotalDeathsPer1Pop, inputItem.TotalTests, inputItem.TestsPer1Pop, currentTime,0)
            id, _ := utils.InsertItem(country, it)
            it.Id = id
            fmt.Println("new item id: ", id)
            utils.InsertDayInItemList(country, it)

            fmt.Println("After")
            fmt.Println(len(country.Items))
            for _, i := range apprepository.TopCountriesList[inputItem.RightName].Items {
                fmt.Println(i)
            }
        }
    }

    //Calculate to app model
    //Write to mem
    CalculateAppItem()

}


///////////////////////////////////////////////////////////////////////////////
//Handler

func UpdateDataHandler(params ...string) string {
    var resultCode = 200
    var data = params[0]

    if len(data) <= 0 {
        resultCode = 400
        data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request")
    } else {
        var arr []*model.InputItem
        json.Unmarshal([]byte(data), &arr)

        for _, item := range arr {
            fmt.Println(item)
        }

        if len(arr) == 0 {
            resultCode = 400
            data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request")
        } else {
            if isValidDoc(arr) {
                //Recreate new metadata

                Update(arr)

                //fmt.Println(arr)
                data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "OK")
            } else {
                resultCode = 400
                data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request")
            }
        }
    }

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - account profile section

func UpdateData(c *gin.Context) {
    dataJson, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        fmt.Println(err.Error())
    }

    var data = UpdateDataHandler(string(dataJson))
    c.String(200, data)
}