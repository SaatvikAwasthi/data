package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"

	"tester/app/server"
	"tester/crosscutting/util"
)

const (
	configFileKey     = "configFile"
	defaultConfigFile = "./env/local.env"
	configFileUsage   = "this is config file path"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, configFileKey, defaultConfigFile, configFileUsage)
	flag.Parse()

	if os.Getenv("TEST_APP_MODE") != "PROD" {
		path, _ := util.GetAbsPath(configFile)
		_ = godotenv.Load(path)
	}
	server.Init()
}
