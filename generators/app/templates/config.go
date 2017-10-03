package main

import (
	"github.com/tkanos/gonfig"
)

type TopLevelConfig struct {
	Web HttpConfig
}

type HttpConfig struct {
	Port            int
	TemplateBaseDir string
}

var runningConfig *TopLevelConfig = nil

func GetConfig() (*TopLevelConfig, error) {
	if runningConfig != nil {
		return runningConfig, nil
	}

	runningConfig = &TopLevelConfig{}
	err := gonfig.GetConf("config.json", runningConfig)
	if err != nil {
		return nil, err
	}

	return runningConfig, nil
}
