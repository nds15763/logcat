package config

import (
	"encoding/json"
	"os"
)

var gCfg *Config

type Config struct {
	Items Items `json:"items"`
}

type Items struct {
	Ip      string    `json:"ip"` //发送至端口的IP
	Port    int       `json:"port"`
	LogType []LogType `json:"log_type"`
}

type LogType struct {
	Paths  string `json:"paths"`
	Topic  string `json:"topic"`
	Type   string `json:"type"`
	RegExp string `json:"reg_exp"` //抓取log的正则
}

func init() {
	f, err := os.Open("conf/config.json")
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err = json.NewDecoder(f).Decode(&cfg); err != nil {
		panic(err)
	}
	gCfg = &cfg
}
