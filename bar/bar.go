package bar

import (
	"fmt"
	"time"

	"github.com/lsgrep/gostatus/config"

	"encoding/json"

	"bufio"
	"os"

	"github.com/lsgrep/gostatus/addon"
	"github.com/lsgrep/gostatus/log"
)

// https://i3wm.org/docs/i3bar-protocol.html
var initMsg = `{ "version": 1, "stop_signal": 10, "cont_signal": 12, "click_events": true }`

type gostatus struct {
	addons []*addon.Addon
}

type Bar interface {
	Run(configPath string)
	Add(addon *addon.Addon)
}

func setupBar() {
	fmt.Print(initMsg)
	// let's start the endless array
	fmt.Print("[")

	// first array as empty
	fmt.Print("[]")
}

func (gs *gostatus) processInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			log.Error(err)
			break
		}
	}
}

// Continuously render the status bar
func (gs *gostatus) render() {

	// Make JSON encoder that writes to stdout
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false)

	// Allocate space for addons' output
	addonsOutput := make([]addon.Block, len(gs.addons))

	for {
		// Update addons' output
		for i, a := range gs.addons {
			if a.NewData != nil {
				addonsOutput[i] = *a.NewData
			}
		}

		if len(addonsOutput) == 0 {
			continue
		}

		// Start with comma written to stdout
		os.Stdout.Write([]byte{ 44 })

		// Encode addons' outputs, sending directly to stdout
		err := encoder.Encode(addonsOutput)

		if err != nil {
			log.Error(err)
			break
		}

		time.Sleep(time.Second)
	}
}

// Initialize & begin running the status bar
func (gs *gostatus) Run(filePath string) {
	// 0. load config
	gs.loadConfig(filePath)

	// 1. setup i3bar
	setupBar()

	// 2. process events
	go gs.processInput()

	// 3. Continually run addons, making them update over time
	for _, a := range gs.addons {
		go a.Run()
	}

	// 3. render addons
	gs.render()
}

func NewGoStatusBar() *gostatus {
	return &gostatus{}
}

func (gs *gostatus) Add(a *addon.Addon) {
	gs.addons = append(gs.addons, a)
}

func (gs *gostatus) loadConfig(filePath string) {
	addons, err := config.ReadConfig(filePath)
	if err != nil {
		gs.Add(addon.NewMessageAddon(err.Error()))
	}
	for _, a := range addons {
		gs.Add(a)
	}
}
