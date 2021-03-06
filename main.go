package main

import (
	"github.com/lsgrep/gostatus/bar"
	"github.com/lsgrep/gostatus/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	// Parse flags / get configs
	pflag.String("config", "gostatus.yml", "config file")
	pflag.String("log", "/tmp/gostatus.log", "log file")
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	l := viper.GetString("log")
	log.ConfigureLogger(l)
	log.Debug("gostatus has been started")

	// Create & run status bar
	statusBar := bar.NewGoStatusBar(viper.GetString("config"))
	statusBar.Run()
}
