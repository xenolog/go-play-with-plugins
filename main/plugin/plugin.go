package plugin

import (
	logger "gopkg.in/xenolog/go-tiny-logger.v1"
)

type BaseStoragePlugin struct {
	log      *logger.Logger
	init     bool
	metadata string
}

// func (p *BaseStoragePlugin) SetLogger(lg *logger.Logger) {
// 	p.log = lg
// }

type StoragePlugin interface {
	Init(*logger.Logger, string) error
	SetLogger(*logger.Logger)
	Set(string, string) error
	Get(string) (string, error)
	Ping() string
}

// This is a template for InitStoragePlugin() function, which should be
// defined into each plugin, and return something,
// implemented StoragePlugin interface
// func InitStoragePlugin(log *logger.Logger, metadata string) StoragePlugin {
// 	// .....
// }
