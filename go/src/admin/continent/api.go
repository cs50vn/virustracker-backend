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

func GetAllContinents(c *gin.Context) {

}

func CreateNewContinent(c *gin.Context) {

}

func GetAContinent(c *gin.Context) {

}

func UpdateAContinent(c *gin.Context) {

}

func DeleteAContinent(c *gin.Context) {

}