package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koding/multiconfig"
)

type Config struct {
	ListenAddress string `default:":80"`
	HomeUrl       string `default:"/"`
	StaticFS      []string
	Accounts      gin.Accounts
}

var (
	config = new(Config)
)

//read config file
func readConfig() {
	m := multiconfig.NewWithPath("config.toml") // supports TOML and JSON
	// Populated the serverConf struct
	m.MustLoad(config) // Check for error
}
