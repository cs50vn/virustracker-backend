package utils

var SQL_GET_CONTINENTS = "select * from CONTINENT"
var SQL_GET_COUNTRIES = "select a.id, a.name, a.capital_name, a.area, a.population, a.flag_id, a.flag_url, a.flag_data, a.flag_timestamp, a.timestamp, a.continent_id, b.name from COUNTRY a, CONTINENT b where a.continent_id = b.id order by a.name asc"

var SQL_GET_ITEMS_IN_COUNTRY = "select id, total_cases, total_deaths, total_recovered, serious_cases, total_cases_per_1pop, total_deaths_per_1pop, total_tests, tests_per_1pop, timestamp from ITEM where country_id = ? order by timestamp desc"
var SQL_DELETE_ITEM = "delete from ITEM where id = ?"
var SQL_INSERT_ITEM = "insert into ITEM(total_cases, total_deaths, total_recovered, serious_cases, total_cases_per_1pop, total_deaths_per_1pop, total_tests, tests_per_1pop, timestamp, country_id) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

var SQL_GET_APP_ITEM = "select * from APP_ITEM"
var SQL_GET_APP_CHART_ITEM = "select continent_id, continent_name, value, type from APP_CHART_ITEM order by value desc"
var SQL_GET_APP_RECENT_ITEM = "select timestamp, value, type from APP_RECENT_ITEM order by timestamp desc"
var SQL_DELETE_APP_ITEM = "delete from APP_ITEM"
var SQL_DELETE_APP_CHART_ITEM = "delete from APP_CHART_ITEM"
var SQL_DELETE_APP_RECENT_ITEM = "delete from APP_RECENT_ITEM"
var SQL_INSERT_APP_ITEM = "insert into APP_ITEM(timestamp, total_cases, new_cases, total_deaths, new_deaths, total_recovered) values(?, ?, ?, ?, ?, ?)"
var SQL_INSERT_APP_CHART_ITEM = "insert into APP_CHART_ITEM(continent_id, continent_name, value, type) values(?, ?, ?, ?)"
var SQL_INSERT_APP_RECENT_ITEM = "insert into APP_RECENT_ITEM(timestamp, value, type) values(?, ?, ?)"

var SQL_GET_ALL_VERSIONS = "select * from VERSION order by version_code desc"

var RESULT_TEMPLATE = `{
    "statusCode": %d,
    "data": %s }`
var ARRAY_TEMPLATE = `[%s]`
var OBJECT_TEMPLATE = `{%s}`
var ITEM_TEMPLATE = `"%s":"%s"`
var ITEM2_TEMPLATE = `"%s":%d`
var ITEM3_TEMPLATE = `"%s":%.2f`
var ITEM4_TEMPLATE = `"%s": [%s]`
var ITEM5_TEMPLATE = `"%s": {%s}`
var ITEM6_TEMPLATE = `"%s": %s`
var ITEM7_TEMPLATE = `"%s"`

//===========================================
var FORCE_UPDATE = "force_update"
var RECOMMEND = "recommend_update"
var NONE = "none"
