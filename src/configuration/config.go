package configuration

import (
	"os"

	"github.com/crud/common"
	"github.com/crud/lib"
	"gopkg.in/yaml.v2"
)

var FullPath = ""
var initialized = false

func LoadConfig(config *ConfigStruct) {
	app_env := os.Getenv(common.VAR_ENVIROMENT)
	relativePath := "configuration/" + app_env + ".yml"
	f, err := os.Open(FullPath + relativePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}
}

func (conf *ConfigStruct) init() {
	LoadConfig(&Cfg)
	lib.Logger.Info("Configuration Initialized!")
	initialized = true
}

func (conf *ConfigStruct) initObservability() {

}

// Load config from /configuration and init Observability(Sentry and Datadog)
func (conf *ConfigStruct) Init() {
	conf.init()
	conf.initObservability()
}

// Get reference to struct configuration.ConfigStruct
func (conf *ConfigStruct) GetConf() *ConfigStruct {
	if !initialized {
		conf.init()
	}
	return &Cfg
}

var Cfg = ConfigStruct{}
