package main

import (
	"github.com/crud/common"
	"github.com/crud/configuration"
	"github.com/crud/lib"
)

func initAll() {
	configuration.Cfg.Init()
}
func main() {
	lib.Logger.Info("Initialing... " + common.APP_NAME)
	initAll()
	lib.Logger.Info("Initialized! " + common.APP_NAME)
}
