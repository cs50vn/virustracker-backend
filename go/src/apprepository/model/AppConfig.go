package model

type AppConfig struct {
	Adminkey  string `json: "adminkey"`
	Clientkey string `json: "clientkey"`
	Dbname    string `json: "dbname"`
	Port      string `json: "port"`
}
