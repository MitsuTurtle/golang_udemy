package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

// ini

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")

	Config = ConfigList{
		// MustIntで初期値を設定
		Port: cfg.Section("web").Key("Port").MustInt(8080),

		// MustStringで初期値を設定
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),

		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("DbName = %v\n", Config.DbName)
	fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)
}
