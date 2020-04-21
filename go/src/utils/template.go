package utils

var SQL_GET_ALL_NEWS            = "select * from NEWSITEM order by sort_id desc"
var SQL_GET_ALL_TOURNAMENTS     = "select * from TOURNAMENT"
var SQL_GET_ALL_TEAM_ID         = "select id, full_name, short_name, logo_id, logo_url, logo_data, logo_timestamp from TEAM where tournament_id = ?"
var SQL_GET_ALL_PLAYER_ID       = "select id, full_name, short_name, logo_id, logo_url, logo_data, logo_timestamp from PLAYER where team_id = ?"
var SQL_GET_ALL_RANKINGTABLE_ID = "select id, name from RANKING_TABLE where tournament_id = ?"
var SQL_GET_ALL_RANKINGITEM_ID  = "select id, place_no, matches_played, won_games, lose_Games, points, team_id from RANKINGITEM where ranking_table_id = ?"
var SQL_GET_ALL_MATCH_ID        = "select id, played_date, match_count, status, bo, home_points, away_points, home_team_id, away_team_id from MATCH where tournament_id = ?"
var SQL_GET_TOP_TOURNAMENTS     = "select * from TOURNAMENT_HIGHLIGHT"
var SQL_GET_TOP_MATCHES         = "select * from MATCH_HIGHLIGHT"
var SQL_GET_TOP_NEWS            = "select * from NEWSITEM  order by sort_id desc LIMIT 10"
var SQL_GET_FILTER_NEWS_        = "select * from NEWSITEM  where sort_id < ? order by sort_id desc limit ?"

var SQL_DELETE_TOP_TOURNAMENTS  = "delete from TOURNAMENT_HIGHLIGHT"
var SQL_DELETE_TOP_MATCHES      = "delete from MATCH_HIGHLIGHT"
var SQL_INSERT_TOP_TOURNAMENTS  = "insert into TOURNAMENT_HIGHLIGHT values(?)"
var SQL_INSERT_TOP_MATCHES      = "insert into MATCH_HIGHLIGHT values(?)"
var SQL_INSERT_NEWS             = "insert into NEWSITEM values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

var SQL_UPDATE_TOURNAMENT           = "update TOURNAMENT set full_name = ?, short_name = ?, logo_data = ?, timestamp = ? where id = ?"
var SQL_UPDATE_MATCH                = "update MATCH set played_date = ?,  match_count = ?, status = ?, bo = ?, home_points = ?, away_points = ?  where id = ?"
var SQL_UPDATE_RANKINGTABLE         = "update RANKING_TABLE set name = ? where id = ?"
var SQL_UPDATE_RANKINGITEM          = "update RANKINGITEM set place_no = ?, matches_played = ?, won_games = ?, lose_games = ?, points = ? where id = ?"
var SQL_UPDATE_NEWSITEM             = "update NEWSITEM set title = ?, created_date = ?, content_url = ?, content_timestamp = ? where id = ?"
var SQL_UPDATE_NEWSITEM_SORTID      = "update NEWSITEM set sort_id = ? where id = ?"
var SQL_UPDATE_NEWSITEM_THUMBNAIL   = "update NEWSITEM set thumbnail_id = ?, thumbnail_url = ?, thumbnail_data = ?, thumbnail_timestamp = ? where id = ?"
var SQL_UPDATE_NEWSITEM_SOURCE      = "update NEWSITEM set source_name = ?, source_id = ?, source_url = ?, source_data = ?, source_timestamp = ? where id = ?"
var SQL_UPDATE_NEWSITEM_STATUS      = "update NEWSITEM set is_active = ? where id = ?"
var SQL_UPDATE_TEAM                 = "update TEAM set full_name = ?, short_name = ? where id = ?"
var SQL_UPDATE_TEAM_LOGO            = "update TEAM set logo_id = ?, logo_url = ?, logo_data = ?, logo_timestamp = ? where id = ?"
var SQL_UPDATE_PLAYER               = "update PLAYER set full_name = ?, short_name = ? where id = ?"
var SQL_UPDATE_PLAYER_LOGO          = "update PLAYER set logo_id = ?, logo_url = ?, logo_data = ?, logo_timestamp = ? where id = ?"


var RESULT_TEMPLATE = `{
    "statusCode": %d,
    "data": %s }`
var ARRAY_TEMPLATE = `[%s]`
var OBJECT_TEMPLATE = `{%s}`
var ITEM_TEMPLATE = `"%s":"%s"`
var ITEM2_TEMPLATE = `"%s":%d`
var ITEM3_TEMPLATE = `"%s":%f`
var ITEM4_TEMPLATE = `"%s": [%s]`
var ITEM5_TEMPLATE = `"%s": {%s}`
var ITEM6_TEMPLATE = `"%s": %s`
var ITEM7_TEMPLATE = `"%s"`

//===========================================
var FORCE_UPDATE = "force_update"
var RECOMMEND = "recommend"
var NONE = "none"