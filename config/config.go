package config

import (
	"fmt"
	"strings"

	"github.com/lsgrep/gostatus/addon"
	"github.com/lsgrep/gostatus/log"
	"github.com/spf13/viper"
)

// Bar configuration struct reflecting the config file's structure
type barConfig struct {
	Addons []addon.AddonConfig `json:"addons"`
}

func LoadAddonsFromConfig(configPath string) []addon.Addon {

	// List of all "New Addon" functions with their corresponding name
	newAddonFuncs := map[string]addon.NewAddonFunc {
		"time": addon.NewTimeAddon,
		"date": addon.NewDateAddon,
		"msg": addon.NewMsgAddon,
	}

	// Read config file
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		log.Error(err)
		return getErrorAddonsList(fmt.Sprintf("Couldn't read config '%s'", configPath))
	}

	// Load config file into variable
	var cfg barConfig

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Error(err)
		return getErrorAddonsList(fmt.Sprintf("Error '%s'", err.Error()))
	}

	// Fill out addons by processing user configs
	var addons []addon.Addon

	for i, config := range cfg.Addons {
		name, _ := config["name"].(string)

		if newAddonFunc, ok := newAddonFuncs[strings.ToLower(name)]; ok {
			addons = append(addons, newAddonFunc(config, i))
			continue
		}

		return []addon.Addon{addon.NewMsgAddonByText(
			fmt.Sprintf("Invalid Addon Name '%s'", name),
			0,
		)}
	}

	return addons
}

// Gets a plain addons list containing a single addon with an error message
func getErrorAddonsList(errMsg string) []addon.Addon {
	return []addon.Addon{addon.NewMsgAddonByText(errMsg, 0)}
}