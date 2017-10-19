package main

import (
	"flag"
	sPlugin "github.com/xenolog/go-play-with-plugins/main/plugin"
	logger "gopkg.in/xenolog/go-tiny-logger.v1"
	"os"
	"plugin"
)

var version string // this string setting up at compile time

func main() {
	var (
		SetValue       bool
		KeyName        string
		Value          string
		Meta           string
		PluginFileName string
		// pingTimeout int
		// err error
	)

	// config := utils.GetOrCreateConfig()
	log := logger.New()
	flag.StringVar(&KeyName, "key", "", "key in the tree")
	flag.StringVar(&Value, "val", "", "value of key")
	flag.StringVar(&PluginFileName, "plugin", "", "Path to storage plugin binary")
	flag.StringVar(&Meta, "meta", "", "Metadata, which will be passed to plugin")
	// flag.BoolVar(&SetValue, "", false, "Control initialization kubernetes client set for connectivity check")
	flag.Parse()
	log.Info("Version: %s", version)
	if Value == "" {
		SetValue = true
	}
	// Load plugin from file
	plg, err := plugin.Open(PluginFileName)
	if err != nil {
		log.Fail("%s", err)
		os.Exit(1)
	}
	plgFabric, err := plg.Lookup("New")
	if err != nil {
		log.Fail("%s", err)
		os.Exit(1)
	}
	storage := plgFabric.(func() sPlugin.StoragePlugin)()
	storage.Init(log, Meta)

	log.Info("PL: %s", storage.Ping())

	if SetValue {
		// set value of ...
	}
}
