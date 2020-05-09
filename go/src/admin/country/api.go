package app

import (
	"cs50vn/virustracker/apprepository"
	"cs50vn/virustracker/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////
//Handler

func GetAllCountriesHandler(params ...string) string {
	var resultCode = 200
	var data = ""

	for _, country := range apprepository.TopCountriesListArray {

		data += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM_TEMPLATE, "id", country.Id)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "name", country.Name)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "capitalName", country.CapitalName)+","+
			fmt.Sprintf(utils.ITEM2_TEMPLATE, "area", country.Area)+","+
			fmt.Sprintf(utils.ITEM2_TEMPLATE, "population", country.Population)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "flagId", country.FlagId)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "flagUrl", country.FlagUrl)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "flagData", country.FlagData)+","+
			fmt.Sprintf(utils.ITEM2_TEMPLATE, "flagTimestamp", country.FlagTimestamp)+","+
			fmt.Sprintf(utils.ITEM2_TEMPLATE, "timestamp", country.Timestamp)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "continentId", country.Continent.Id)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "continentName", country.Continent.Name)+","+
			fmt.Sprintf(utils.ITEM2_TEMPLATE, "totalItems", len(country.Items))) + ","
	}
	data = strings.TrimSuffix(data, ",")
	data = fmt.Sprintf(utils.ARRAY_TEMPLATE, data)

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - account profile section

func GetAllCountries(c *gin.Context) {
	var data = GetAllCountriesHandler()
	c.String(200, data)
}

func CreateACountry(c *gin.Context) {

}

func GetACountry(c *gin.Context) {

}

func UpdateACountry(c *gin.Context) {

}

func DeleteACountry(c *gin.Context) {

}

func GetAllItemsInCountry(c *gin.Context) {

}

func UpdateFlagCountry(c *gin.Context) {

}

func UpdateContinentCountry(c *gin.Context) {

}