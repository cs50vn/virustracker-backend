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

func CreateAnItem(c *gin.Context) {

}

func GetAnItem(c *gin.Context) {

}

func UpdateAnItem(c *gin.Context) {

}

func DeleteAnItem(c *gin.Context) {

}
