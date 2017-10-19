package main

import (
	"fmt"
	plBase "github.com/xenolog/go-play-with-plugins/main/plugin"
	logger "gopkg.in/xenolog/go-tiny-logger.v1"
)

type StoragePluginType struct {
	plBase.BaseStoragePlugin
}

var version string // this string setting up at compile time

func (p *StoragePluginType) SetLogger(lg *logger.Logger) {
	p.Log = lg
}

func (p *StoragePluginType) Set(key, val string) error {
	return nil
}

func (p *StoragePluginType) Get(key string) (string, error) {
	return "", nil
}

func (p *StoragePluginType) Ping() string {
	return fmt.Sprintf("OK, Plugin for File storage, version: %s", version)
}

func (p *StoragePluginType) Init(log *logger.Logger, metadata string) error {
	if p.AlreadyInit == false {
		p.Log = log
		p.Metadata = metadata
		p.AlreadyInit = true
	}
	return nil
}

func New() plBase.StoragePlugin {
	return &StoragePluginType{}
}
