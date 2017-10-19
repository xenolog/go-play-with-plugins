package plugin

import (
	"github.com/xenolog/go-play-with-plugins/main/plugin"
	logger "gopkg.in/xenolog/go-tiny-logger.v1"
)

type StoragePluginType struct {
	BaseStoragePlugin
}

var version string // this string setting up at compile time

func (p *BaseStoragePlugin) SetLogger(lg *logger.Logger) {
	p.Log = lg
}

func (p *BaseStoragePlugin) Set(key, val string) error {
	return nil
}

func (p *BaseStoragePlugin) Get(key string) (string, error) {
	return "", nil
}

func (p *BaseStoragePlugin) Ping() {
	return "OK, Plugin for File storage, version: %s", version
}

func (p *BaseStoragePlugin) Init(log *logger.Logger, metadata string) error {
	if p.alreadyInit == false {
		p.log = log
		p.metadata = metadata
	}
	return nil
}

func New() StoragePlugin {
	return &StoragePluginType{}
}
