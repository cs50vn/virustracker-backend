package app

import (
    admin_app "cs50vn/virustracker/admin/app"
    admin_continent "cs50vn/virustracker/admin/continent"
    admin_country "cs50vn/virustracker/admin/country"
    admin_hook "cs50vn/virustracker/admin/hook"
    "cs50vn/virustracker/apprepository"
    "cs50vn/virustracker/apprepository/model"
    "cs50vn/virustracker/utils"
    "database/sql"
    "encoding/json"
    "fmt"
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
    "io/ioutil"
    "runtime"
    "strconv"
    "strings"
    "time"
)

/////////////////////////////////////////////////////////////////////////////
//Global vars
func InitApp() {
    LoadConfig()
    InitDb()
    LoadData()

    MemoryUsage()
}

func LoadConfig() {
    fmt.Println("Loaded config file: ", apprepository.ConfigName)

    data, err := ioutil.ReadFile(apprepository.ConfigName)
    if err != nil {
        fmt.Print(err)
    }

    err = json.Unmarshal(data, &apprepository.Config)
    if err != nil {
        fmt.Print(err)
    }

    fmt.Println(apprepository.Config)
}

func InitDb() {
    var err error
    apprepository.DbConnection, err = sql.Open("sqlite3", apprepository.Config.Dbname)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Println("Connected db: " + apprepository.Config.Dbname)
}

func LoadData() {
    LoadVersions()

    LoadContinents()

    LoadCountries()

    LoadAppItem()

    //admin_hook.Update(admin_hook.FakeData())
    admin_hook.CalculateAppItem()
}

func LoadVersions() {
    prepareSelect := utils.SQL_GET_ALL_VERSIONS

    prepareSt, err := apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ := prepareSt.Query()

    for rows.Next() {
        var versionCode int
        var status string
        var downloadLink string

        err = rows.Scan(&versionCode, &status, &downloadLink)
        if err != nil {
            fmt.Println(err.Error())
        } else {

            apprepository.ListOfVersions[versionCode] = model.MakeVersion(versionCode, status, downloadLink)
            apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, versionCode)
        }
    }

    fmt.Println(apprepository.ListOfVersions)
}

func LoadContinents() {
    prepareSelect := utils.SQL_GET_CONTINENTS

    prepareSt, err := apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ := prepareSt.Query()

    for rows.Next() {
        var id string
        var name string

        err = rows.Scan(&id, &name)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            apprepository.ContinentList[id] = model.MakeContinent(id, name, 0)
        }
    }
}

func LoadCountries() {
    prepareSelect := utils.SQL_GET_COUNTRIES

    prepareSt, err := apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ := prepareSt.Query()

    for rows.Next() {
        var id string
        var name string
        var capitalName string
        var area int64
        var population int64
        var flagId string
        var flagUrl string
        var flagData string
        var flagTimestamp int64
        var timestamp int64
        var continentId string
        var continentName string

        err = rows.Scan(&id, &name, &capitalName, &area, &population, &flagId, &flagUrl,&flagData, &flagTimestamp, &timestamp, &continentId, &continentName)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            continent := apprepository.ContinentList[continentId]
            itemList := LoadItemsInCountry(id)

            countryItem := model.MakeCountry(id, name, capitalName, area, population, flagId, flagUrl,flagData, flagTimestamp, timestamp, continent, itemList)
            apprepository.TopCountriesList[id] = countryItem
            apprepository.TopCountriesListArray = append(apprepository.TopCountriesListArray, countryItem)

            fmt.Println(apprepository.TopCountriesList[id].Id + " - "  + apprepository.TopCountriesList[id].Name, "; items: ", len(itemList))
        }
    }
    fmt.Println("Total countries: %d ", len(apprepository.TopCountriesList))
}

func LoadItemsInCountry(countryId string) []*model.Item {
    list := make([]*model.Item, 0)

    prepareSelect := utils.SQL_GET_ITEMS_IN_COUNTRY

    prepareSt, err := apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ := prepareSt.Query(countryId)

    for rows.Next() {
        var id int64
        var totalCases int64
        var totalDeaths int64
        var totalRecovered int64
        var seriousCases int64
        var totalCasesPer1Pop float64
        var totalDeathsPer1Pop float64
        var totalTests int64
        var testsPer1Pop float64
        var timestamp int64

        err = rows.Scan(&id, &totalCases, &totalDeaths, &totalRecovered, &seriousCases, &totalCasesPer1Pop, &totalDeathsPer1Pop, &totalTests, &testsPer1Pop, &timestamp)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            list = append(list, model.MakeItem(id, totalCases, 0, totalDeaths, 0, totalRecovered, seriousCases, totalCasesPer1Pop, totalDeathsPer1Pop, totalTests, testsPer1Pop, timestamp, 0))
        }
    }

    return list
}

func LoadAppItem() {
    //Load app chart item
    var chartTotalCases []*model.Continent
    var chartTotalDeaths []*model.Continent

    prepareSelect := utils.SQL_GET_APP_CHART_ITEM

    prepareSt, err := apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ := prepareSt.Query()

    for rows.Next() {
        var continentId string
        var continentName string
        var value int64
        var type_ int

        err = rows.Scan(&continentId, &continentName, &value, &type_)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            if type_ == 1 {
                chartTotalCases = append(chartTotalCases, model.MakeContinent(continentId, continentName, value))
            } else if type_ == 2 {
                chartTotalDeaths = append(chartTotalDeaths, model.MakeContinent(continentId, continentName, value))
            }
        }
    }

    //Load app recent item
    var recentTotalCases []*model.RecentItem
    var recentTotalDeaths []*model.RecentItem

    prepareSelect = utils.SQL_GET_APP_RECENT_ITEM

    prepareSt, err = apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ = prepareSt.Query()

    for rows.Next() {
        var timestamp int64
        var value int64
        var type_ int

        err = rows.Scan(&timestamp, &value, &type_)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            if type_ == 1 {
                recentTotalCases = append(recentTotalCases, model.MakeRecentItem(timestamp, value))
            } else if type_ == 2 {
                recentTotalDeaths = append(recentTotalDeaths, model.MakeRecentItem(timestamp, value))
            }
        }
    }

    //Load app item
    prepareSelect = utils.SQL_GET_APP_ITEM

    prepareSt, err = apprepository.DbConnection.Prepare(prepareSelect)
    if err != nil {
        fmt.Println(err.Error())
    }

    rows, _ = prepareSt.Query()

    for rows.Next() {
        var timestamp int64
        var totalCases int64
        var newCases int64
        var totalDeaths int64
        var newDeaths int64
        var totalRecovered int64

        err = rows.Scan(&timestamp, &totalCases, &newCases, &totalDeaths, &newDeaths, &totalRecovered)
        if err != nil {
            fmt.Println(err.Error())
        } else {
            apprepository.AppModel = model.MakeAppItem(timestamp, totalCases, newCases, totalDeaths, newDeaths, totalRecovered, chartTotalCases, chartTotalDeaths, recentTotalCases, recentTotalDeaths)
        }
    }

}

func MemoryUsage() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    // For info on each, see: https://golang.org/pkg/runtime/#MemStats
    fmt.Printf("Alloc = %v kB", (m.Alloc / 1024))
    fmt.Printf("\tTotalAlloc = %v kB\n", (m.TotalAlloc / 1024))
}

///////////////////////////////////////////////////////////////////////////////
//Handler

func GetTopAllHandler(params ...string) string {
    var resultCode = 200
    var data = ""
    var str = ""
    var str1 = ""
    var str2 = ""
    var str3 = ""
    var t = time.Now().Unix()

    //Top home
    str += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM2_TEMPLATE, "timestamp", apprepository.AppModel.Timestamp)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalCases", apprepository.AppModel.TotalCases)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "newCases", apprepository.AppModel.NewCases)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalDeaths", apprepository.AppModel.TotalDeaths)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "newDeaths", apprepository.AppModel.NewDeaths)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalRecovered", apprepository.AppModel.TotalRecovered)+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalCasesChart", utils.GenerateChartTotalCasesJson())+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalDeathsChart", utils.GenerateChartTotalDeathsJson())+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalCasesRecent", utils.GenerateRecentTotalCasesJson())+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalDeathsRecent", utils.GenerateRecentTotalDeathsJson()))

    //Top country
    for _, continent := range apprepository.ContinentList {

        str1 += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM_TEMPLATE, "id", continent.Id)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "name", continent.Name)) + ","
    }
    str1 = strings.TrimSuffix(str1, ",")
    str1 = fmt.Sprintf(utils.ARRAY_TEMPLATE, str1)

    for _, item := range apprepository.TopCountriesListArray {
        str2 += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM_TEMPLATE, "id", item.Id)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "name", item.Name)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagId", item.FlagId)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagUrl", item.FlagUrl)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagData", item.FlagData)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "flagTimestamp", item.FlagTimestamp)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "timestamp", item.Timestamp)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalCases", utils.GetItemByTimestamp(item, t).TotalCases)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "newCases", utils.GetItemByTimestamp(item, t).TotalCases - utils.GetItemByTimestamp(item, t - 86400).TotalCases)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalDeaths", utils.GetItemByTimestamp(item, t).TotalDeaths) + "," +
            fmt.Sprintf(utils.ITEM_TEMPLATE, "continentId", item.Continent.Id)) + ","
    }
    str2 = fmt.Sprintf(utils.ARRAY_TEMPLATE, strings.TrimSuffix(str2, ","))
    str3 = fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM6_TEMPLATE, "continentList", str1) + "," + fmt.Sprintf(utils.ITEM6_TEMPLATE, "countryList", str2))

    data = fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM6_TEMPLATE, "tophome", str) + "," + fmt.Sprintf(utils.ITEM6_TEMPLATE, "topcountry", str3))

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func GetTopHomeHandler(params ...string) string {
    var resultCode = 200
    var data = ""

    data += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM2_TEMPLATE, "timestamp", apprepository.AppModel.Timestamp)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalCases", apprepository.AppModel.TotalCases)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "newCases", apprepository.AppModel.NewCases)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalDeaths", apprepository.AppModel.TotalDeaths)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "newDeaths", apprepository.AppModel.NewDeaths)+","+
        fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalRecovered", apprepository.AppModel.TotalRecovered)+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalCasesChart", utils.GenerateChartTotalCasesJson())+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalDeathsChart", utils.GenerateChartTotalDeathsJson())+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalCasesRecent", utils.GenerateRecentTotalCasesJson())+","+
        fmt.Sprintf(utils.ITEM4_TEMPLATE, "totalDeathsRecent", utils.GenerateRecentTotalDeathsJson()))

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}


func GetTopCountryHandler(params ...string) string {
    var resultCode = 200
    var data = ""
    var str = ""
    var str1 = ""
    var t = time.Now().Unix()

    for _, continent := range apprepository.ContinentList {

        str += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM_TEMPLATE, "id", continent.Id)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "name", continent.Name)) + ","
    }
    str = strings.TrimSuffix(str, ",")
    str = fmt.Sprintf(utils.ARRAY_TEMPLATE, str)

    for _, item := range apprepository.TopCountriesListArray {
        str1 += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM_TEMPLATE, "id", item.Id)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "name", item.Name)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagId", item.FlagId)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagUrl", item.FlagUrl)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagData", item.FlagData)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "flagTimestamp", item.FlagTimestamp)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "timestamp", item.Timestamp)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalCases", utils.GetItemByTimestamp(item, t).TotalCases)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "newCases", utils.GetItemByTimestamp(item, t).TotalCases - utils.GetItemByTimestamp(item, t - 86400).TotalCases)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalDeaths", utils.GetItemByTimestamp(item, t).TotalDeaths) + "," +
            fmt.Sprintf(utils.ITEM_TEMPLATE, "continentId", item.Continent.Id)) + ","
    }

    str1 = fmt.Sprintf(utils.ARRAY_TEMPLATE, strings.TrimSuffix(str1, ","))
    data = fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM6_TEMPLATE, "continentList", str) + "," + fmt.Sprintf(utils.ITEM6_TEMPLATE, "countryList", str1))

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func GetCountryDetailHandler(params ...string) string {
    var resultCode = 200
    var countryId = params[0]
    var data = ""
    var t = time.Now().Unix()

    if item, ok := apprepository.TopCountriesList[countryId]; ok {
        data += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM_TEMPLATE, "id", item.Id)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "name", item.Name)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "capitalName", item.CapitalName)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "area", item.Area)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "population", item.Population)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagId", item.FlagId)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagUrl", item.FlagUrl)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "flagData", item.FlagData)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "flagTimestamp", item.FlagTimestamp)+","+
            fmt.Sprintf(utils.ITEM2_TEMPLATE, "timestamp", item.Timestamp)+","+
            fmt.Sprintf(utils.ITEM_TEMPLATE, "continentId", item.Continent.Id)+","+
            fmt.Sprintf(utils.ITEM5_TEMPLATE, "status", utils.GenerateCountryDetailStatus(countryId, t)))
    } else {
        resultCode = 400;
        data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid Country Detail Request")

    }

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func CheckUpdateStatusHandler(params ...string) string {
    var resultCode = 200
    var data = ""

    var versionCode, err = strconv.Atoi(params[0])
    if err != nil {
        resultCode = 400
        data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid Request")
    } else if _, ok := apprepository.ListOfVersions[versionCode]; ok {
        var itemCode = apprepository.ListOfVersionsKeys[0]
        if versionCode < itemCode {
            oldItem := apprepository.ListOfVersions[versionCode]
            item := apprepository.ListOfVersions[itemCode]
            //Have a new app version
            resultCode = 200
            data = fmt.Sprintf(utils.ITEM2_TEMPLATE, "requestVersionCode", versionCode) + "," +
                fmt.Sprintf(utils.ITEM2_TEMPLATE, "hasNewVersion", 1) + "," +
                fmt.Sprintf(utils.ITEM_TEMPLATE, "status", oldItem.Status) + "," +
                fmt.Sprintf(utils.ITEM_TEMPLATE, "downloadLink", item.DownloadLink)
            data = fmt.Sprintf(utils.OBJECT_TEMPLATE, data)
        } else {
            resultCode = 200
            data = fmt.Sprintf(utils.ITEM2_TEMPLATE, "requestVersionCode", versionCode) + "," +
                fmt.Sprintf(utils.ITEM2_TEMPLATE, "hasNewVersion", 0)
            data = fmt.Sprintf(utils.OBJECT_TEMPLATE, data)
        }
    } else {
        resultCode = 400
        data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid Request")
    }

    fmt.Println()

    return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - App section

func GetTopAll(c *gin.Context) {
    var data = GetTopAllHandler()
    c.String(200, data)
}

func GetTopHome(c *gin.Context) {
    var data = GetTopHomeHandler()
    c.String(200, data)
}

func GetTopCountry(c *gin.Context) {
    var data = GetTopCountryHandler()
    c.String(200, data)
}

func GetCountryDetail(c *gin.Context) {
    var countryId = c.Param("countryId")

    var data = GetCountryDetailHandler(countryId)
    c.String(200, data)
}

func CheckUpdateStatus(c *gin.Context) {
    var versionCode = c.Query("versionCode")

    var data = CheckUpdateStatusHandler(versionCode)
    c.String(200, data)
}

/////////////////////////////////////////////////////////////////////////////////
//Main func

func Run() {
    r := gin.Default()

    v1 := r.Group("/v1")
    {
        //App
        v1.GET("/app/topall", GetTopAll)
        v1.GET("/app/tophome", GetTopHome)
        v1.GET("/app/topcountry", GetTopCountry)
        v1.GET("/app/country/:countryId", GetCountryDetail)
        v1.GET("/app/version/check", CheckUpdateStatus)

        //*********************************************************************//
        //Admin
        //App
        v1.GET("/admin/app/version", admin_app.GetAllVersions)
        v1.POST("/admin/app/version", admin_app.CreateNewVersion)
        v1.GET("/admin/app/version/:versionCode", admin_app.GetAVersion)
        v1.PATCH("/admin/app/version/:versionCode", admin_app.UpdateAVersion)
        v1.DELETE("/admin/app/version/:versionCode", admin_app.DeleteAVersion)

        //Continent
        v1.GET("/admin/continent", admin_continent.GetAllContinents)

        //Country
        v1.GET("/admin/country", admin_country.GetAllCountries)

        //Hook
        v1.POST("/admin/hook/update", admin_hook.UpdateData)
    }

    r.Run(":" + apprepository.Config.Port)
}
