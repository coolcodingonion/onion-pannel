package utils

import "gopkg.in/ini.v1"

var (
	AppMode  string
	HttpPort string

	Db     string
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
)

func init() { 
	file, err := ini.Load("config/app.ini")
	if err != nil {
		panic("Failed to load 'config/app.ini': " + err.Error())
	}
	LoadServer(file)
	LoadDatabase(file)
}


func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
}

func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("onion")
	DbPass = file.Section("database").Key("DbPass").MustString("password")
	DbName = file.Section("database").Key("DbName").MustString("onion")
}
