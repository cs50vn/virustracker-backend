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

func GetAllContinentsHandler(params ...string) string {
	var resultCode = 200
	var data = ""

	for _, continent := range apprepository.ContinentList {

		data += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM_TEMPLATE, "id", continent.Id)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "name", continent.Name)) + ","
	}
	data = strings.TrimSuffix(data, ",")
	data = fmt.Sprintf(utils.ARRAY_TEMPLATE, data)

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - account profile section

func GetAllContinents(c *gin.Context) {
	var data = GetAllContinentsHandler()
	c.String(200, data)
}

func CreateNewContinent(c *gin.Context) {

}

func GetAContinent(c *gin.Context) {

}

func UpdateAContinent(c *gin.Context) {

}

func DeleteAContinent(c *gin.Context) {

}