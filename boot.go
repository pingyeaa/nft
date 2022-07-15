package main

import (
	"gopkg.in/ini.v1"
	"log"
)

var cfg *ini.File
var err error

func init() {

	// init config
	cfg, err = ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetConfig(section string, key string) string {
	k, err := cfg.Section(section).GetKey(key)
	if err != nil {
		return ""
	}
	return k.String()
}
