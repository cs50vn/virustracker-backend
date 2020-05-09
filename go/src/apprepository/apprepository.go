package apprepository

import (
    "cs50vn/virustracker/apprepository/model"
    "database/sql"
)

//App
var ConfigName = "config.json"
var Config model.AppConfig
var DbConnection *sql.DB
var LimitCountryChart = 7
var ListOfVersions = make(map[int]*model.Version)
var ListOfVersionsKeys = make([]int, 0)

var AppModel = model.MakeAppItem(0,0, 0, 0, 0, 0, nil, nil, nil, nil)
var ContinentList = make(map[string]*model.Continent)
var TopCountriesList = make(map[string]*model.Country)
var TopCountriesListArray = make([]*model.Country, 0)

