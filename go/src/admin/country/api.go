package app

import (
	"github.com/gin-gonic/gin"
)

///////////////////////////////////////////////////////////////////////////////
//Handler

func GetAllVersionsHandler(params ...string) string {
	var userToken = params[0]
	var sortId = params[1]

	return userToken + sortId + "api"
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - account profile section

func GetAllCountries(c *gin.Context) {

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