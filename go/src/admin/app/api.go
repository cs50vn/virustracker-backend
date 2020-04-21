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

func GetAllVersions(c *gin.Context) {

}

func GetAVersion(c *gin.Context) {

}

func CreateNewVersion(c *gin.Context) {

}

func UpdateAVersion(c *gin.Context) {

}

func DeleteAVersion(c *gin.Context) {

}