package app

import (
	"cs50vn/virustracker/apprepository/model"
	"cs50vn/virustracker/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func isValidDoc(doc []model.Item) bool {

	return true
}

func Update(doc []model.Item) {
	//Extract to country item
	//for _, item := range doc {
	//	//Write to mem
	//
	//
	//	//Write to db
	//}

	//Calculate to app model
		//Write to mem


		//Write to db

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
		var arr []model.Item
		json.Unmarshal([]byte(data), &arr)

		if len(arr) == 0 {
			resultCode = 400
			data = fmt.Sprintf(utils.ITEM7_TEMPLATE, "Invalid request")
		} else {
			if isValidDoc(arr) {
				//Recreate new metadata
				Update(arr)

				fmt.Println(arr)
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