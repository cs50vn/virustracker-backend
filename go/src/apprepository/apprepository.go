package apprepository

import (
	"cs50vn/virustracker/apprepository/model"
	"database/sql"
)

//App
var ConfigName = "config.json"
var Config model.AppConfig
var DbConnection *sql.DB
var ListOfVersions = make(map[int]*model.Version)
var ListOfVersionsKeys = make([]int, 0)

var ContinentList = make([]*model.Continent, 0)
var TopCountriesList = make(map[string]*model.Country)
var TopCountriesListArray = make([]*model.Country, 0)


