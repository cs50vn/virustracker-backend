package app

import (
	"cs50vn/virustracker/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	admin_app "cs50vn/virustracker/admin/app"
	admin_continent "cs50vn/virustracker/admin/continent"
	admin_country "cs50vn/virustracker/admin/country"
	admin_hook "cs50vn/virustracker/admin/hook"
	admin_item "cs50vn/virustracker/admin/item"
	"cs50vn/virustracker/apprepository"
	"github.com/gin-gonic/gin"
)

/////////////////////////////////////////////////////////////////////////////
//Global vars
func InitApp() {
	LoadConfig()
	InitDb()
	LoadData()
	//utils.LoadFakeData()

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
	LoadContinents()

	LoadCountries()
}

func LoadVersions() {
	LoadContinents()

	LoadCountries()
}

func LoadContinents() {

}

func LoadCountries() {

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
	var data = "topall"

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func GetTopHomeHandler(params ...string) string {
	var resultCode = 200
	var data = "GetTopHomeHandler"

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func GetTopCountryHandler(params ...string) string {
	var resultCode = 200
	var data = "GetTopCountryHandler"

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func GetCountryDetailHandler(params ...string) string {
	var resultCode = 200
	var data = "GetCountryDetailHandler"

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func CheckUpdateStatusHandler(params ...string) string {
	var resultCode = 200
	var data = ""

	var versionCode, err = strconv.Atoi(params[1])
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
	var versionCode = c.Query("versionCode")

	var data = GetTopCountryHandler(versionCode)
	c.String(200, data)
}

func GetCountryDetail(c *gin.Context) {
	var versionCode = c.Query("versionCode")

	var data = GetCountryDetailHandler(versionCode)
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
		v1.POST("/admin/continent", admin_continent.CreateNewContinent)
		v1.GET("/admin/continent/:continentId", admin_continent.GetAContinent)
		v1.PATCH("/admin/continent/:continentId", admin_continent.UpdateAContinent)
		v1.DELETE("/admin/continent/:continentId", admin_continent.DeleteAContinent)

		//Country
		v1.GET("/admin/country", admin_country.GetAllCountries)
		v1.POST("/admin/country", admin_country.CreateACountry)
		v1.GET("/admin/country/:countryId", admin_country.GetACountry)
		v1.PATCH("/admin/country/:countryId", admin_country.UpdateACountry)
		v1.DELETE("/admin/country/:countryId", admin_country.DeleteACountry)
		v1.GET("/admin/country/:countryId/item", admin_country.GetAllItemsInCountry)
		v1.PATCH("/admin/country/:countryId/flag", admin_country.UpdateFlagCountry)
		v1.PATCH("/admin/country/:countryId/continent", admin_country.UpdateContinentCountry)

		//Item
		v1.POST("/admin/item", admin_item.CreateAnItem)
		v1.GET("/admin/item/:itemId", admin_item.GetAnItem)
		v1.PATCH("/admin/item/:itemId", admin_item.UpdateAnItem)
		v1.DELETE("/admin/item/:itemId", admin_item.DeleteAnItem)

		//Hook
		v1.POST("/admin/hook/update", admin_hook.UpdateData)
	}

	r.Run(":" + apprepository.Config.Port)
}