package app

import (
	"cs50vn/virustracker/apprepository"
	"cs50vn/virustracker/apprepository/model"
	"cs50vn/virustracker/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////
//Handler

func GetAllVersionsHandler(params ...string) string {
	var resultCode = 200
	var data = ""

	for _, key := range apprepository.ListOfVersionsKeys {
		item := apprepository.ListOfVersions[key]
		data += fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM2_TEMPLATE, "versionCode", item.VersionCode)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "status", item.Status)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "downloadLink", item.DownloadLink)) + ","
	}
	data = strings.TrimSuffix(data, ",")
	data = fmt.Sprintf(utils.ARRAY_TEMPLATE, data)

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func GetAVersionHandler(params ...string) string {
	var versionCode, _ = strconv.Atoi(params[0])
	var resultCode = 200
	var data = ""

	if item, ok := apprepository.ListOfVersions[versionCode]; ok {
		data = fmt.Sprintf(utils.OBJECT_TEMPLATE, fmt.Sprintf(utils.ITEM2_TEMPLATE, "versionCode", item.VersionCode)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "status", item.Status)+","+
			fmt.Sprintf(utils.ITEM_TEMPLATE, "downloadLink", item.DownloadLink))
	} else {
		resultCode = 400
		data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid version request")
	}

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func CreateNewVersionHandler(params ...string) string {
	var versionCode, _ = strconv.Atoi(params[0])
	var status = params[1]
	var downloadLink = params[2]
	var resultCode = 200
	var data = ""

	if _, ok := apprepository.ListOfVersions[versionCode]; ok {
		resultCode = 400
		data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request: Duplicate version")
	} else {
		//insert to db
		result := utils.ExecuteSQL(utils.SQL_INSERT_A_VERSION, versionCode, status, downloadLink)
		rowAffected, err := result.RowsAffected()
		if err != nil || rowAffected != 1 {
			fmt.Println(err.Error())
		} else {
			item := model.MakeVersion(versionCode, status, downloadLink)
			apprepository.ListOfVersions[versionCode] = item
			//insert to mem
			if len(apprepository.ListOfVersions) == 0 {

				apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, versionCode)
			} else {
				pos := 0
				for index := 0; index < len(apprepository.ListOfVersionsKeys); index++ {

					if item.VersionCode > apprepository.ListOfVersionsKeys[index] {
						pos = index
						break
					} else {
						pos++
					}
				}
				fmt.Println("pos: ", pos)
				apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys, 0 /* use the zero value of the element type */)
				copy(apprepository.ListOfVersionsKeys[pos+1:], apprepository.ListOfVersionsKeys[pos:])
				apprepository.ListOfVersionsKeys[pos] = item.VersionCode
			}

			data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "OK")
		}

	}

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func UpdateAVersionHandler(params ...string) string {
	var versionCode, _ = strconv.Atoi(params[0])
	var status = params[1]
	var downloadLink = params[2]
	var resultCode = 200
	var data = ""

	if _, ok := apprepository.ListOfVersions[versionCode]; ok {
		//Update in db
		result := utils.ExecuteSQL(utils.SQL_UPDATE_A_VERSION, status, downloadLink, versionCode)
		rowAffected, err := result.RowsAffected()
		if err != nil || rowAffected != 1 {
			fmt.Println(err.Error())
		} else {
			item := model.MakeVersion(versionCode, status, downloadLink)
			apprepository.ListOfVersions[versionCode] = item

			data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "OK")
		}
	} else {
		resultCode = 400
		data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request: wrong version")
	}

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

func DeleteAVersionHandler(params ...string) string {
	var versionCode, _ = strconv.Atoi(params[0])
	var resultCode = 200
	var data = ""

	if _, ok := apprepository.ListOfVersions[versionCode]; ok {
		//Delete in db
		result := utils.ExecuteSQL(utils.SQL_DELETE_A_VERSION, versionCode)
		rowAffected, err := result.RowsAffected()
		if err != nil || rowAffected != 1 {
			fmt.Println(err.Error())
		} else {
			delete(apprepository.ListOfVersions, versionCode)

			if len(apprepository.ListOfVersionsKeys) < 2 {
				apprepository.ListOfVersionsKeys = make([]int, 0)
			} else {
				for index := 0; index < len(apprepository.ListOfVersionsKeys); index++ {
					if apprepository.ListOfVersionsKeys[index] == versionCode {
						apprepository.ListOfVersionsKeys = append(apprepository.ListOfVersionsKeys[:index], apprepository.ListOfVersionsKeys[index+1:]...)
						break
					}

				}
			}

			data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "OK")
		}

	} else {
		resultCode = 400
		data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request: wrong version")
	}

	return fmt.Sprintf(utils.RESULT_TEMPLATE, resultCode, data)
}

/////////////////////////////////////////////////////////////////////////////////
////Rest endpoint - account profile section

func GetAllVersions(c *gin.Context) {
	var data = GetAllVersionsHandler()
	c.String(200, data)
}

func GetAVersion(c *gin.Context) {
	var versionCode = c.Param("versionCode")
	var data = GetAVersionHandler(versionCode)
	c.String(200, data)
}

func CreateNewVersion(c *gin.Context) {
	var versionCode = c.Query("versionCode")
	var status = c.Query("status")
	var downloadLink = c.Query("downloadLink")
	var data = CreateNewVersionHandler(versionCode, status, downloadLink)
	c.String(200, data)
}

func UpdateAVersion(c *gin.Context) {
	var versionCode = c.Param("versionCode")
	var status = c.Query("status")
	var downloadLink = c.Query("downloadLink")
	var data = UpdateAVersionHandler(versionCode, status, downloadLink)
	c.String(200, data)
}

func DeleteAVersion(c *gin.Context) {
	var versionCode = c.Param("versionCode")
	var data = DeleteAVersionHandler(versionCode)
	c.String(200, data)
}