package bar

import (
	"encoding/json"
	"fmt"
	"github.com/lsgrep/gostatus/addon"
	"github.com/lsgrep/gostatus/config"
	"github.com/lsgrep/gostatus/log"
	"os"
	"time"
)

type gostatus struct {
	addons  []addon.Addon
	output  []addon.Block
	encoder *json.Encoder
}

type Bar interface {
	Run(configPath string)
}

// Send the initial bar message to start it off (https://i3wm.org/docs/i3bar-protocol.html)
func sendBarInitMsg() {
	fmt.Print(`{ "version": 1, "stop_signal": 10, "cont_signal": 12, "click_events": true }[`)
}

// Render the status bar's addons' output by sending it (encoded) to stdout
func (gs *gostatus) render() {

	// Encode addons' outputs, sending directly to stdout
	err := gs.encoder.Encode(gs.output)

	// Output "," after the array
	os.Stdout.Write([]byte{44})

	if err != nil {
		log.Error(err)
	}
}

// Initialize & begin running the status bar
func (gs *gostatus) Run() {

	sendBarInitMsg()

	// Continuously run addons, collecting their updated blocks over time
	blocks := make(chan *addon.Block)
	blocksRendered := make(chan *addon.Block)

	for _, a := range gs.addons {
		go a.Run(blocks, blocksRendered)
	}

	// Always render at least once, in case we only have a single error message
	time.AfterFunc(200*time.Millisecond, func() {
		gs.render()
	})

	for {
		select {
		case block := <-blocks:
			gs.output[block.Index] = *block
		case block := <-blocksRendered:
			gs.output[block.Index] = *block
			gs.render()
		}
	}
}

// Creates a new Go Status Bar, using the config file at the given file path.
func NewGoStatusBar(filePath string) *gostatus {

	gs := &gostatus{}

	// Load addons from the config file
	gs.addons = config.LoadAddonsFromConfig(filePath)

	// Set the encoder
	gs.encoder = json.NewEncoder(os.Stdout)
	gs.encoder.SetEscapeHTML(false)

	// Set the addons' block output array
	gs.output = make([]addon.Block, len(gs.addons))

	return gs
}
