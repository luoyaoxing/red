package config

import (
	"github.com/micro/go-micro/config"
)

var basePath = "/code/gocode/kay"

func LoadFile(path string) error {
	fileName := basePath + path
	return config.LoadFile(fileName)
}

func Map() map[string]interface{}  {
	return config.Map()
}

func Get(path ...string) string {
	return config.Get(path...).String("")
}