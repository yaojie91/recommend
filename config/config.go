package config

import (
	"encoding/json"
	"os"
)

type RedisCfg struct {
	Host string
	Port string
}

var Conf *RedisCfg

func init() {
	file, err := os.Open("./cfg.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Conf)
	if err != nil {
		panic(err)
	}
}