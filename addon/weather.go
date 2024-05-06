package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"os/exec"
	"strings"
	"time"
)

type weatherAddon struct {
	index    int
	location string
}

const (
	weatherDefaultFormat   = "%s %s°C"
	weatherDefaultInterval = 1 * time.Hour
	weatherCmd             = "ansiweather -l %s -a false -p false -h false -w false | grep -oE '\\b[0-9]{1,3}\\b' | awk 'NR==1{print}'\n"
)

func (w *weatherAddon) getBlock() *Block {

	cmdOut, err := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf(weatherCmd, w.location),
	).Output()

	if err != nil {
		log.Error(err)
		return nil
	}

	return &Block{
		FullText: fmt.Sprintf(
			weatherDefaultFormat,
			IconWeather2,
			strings.TrimSpace(string(cmdOut)),
		),
		Index: w.index,
	}
}

func (w *weatherAddon) Run(blocks chan *Block, blocksRendered chan *Block) {

	blocks <- w.getBlock()

	tick := time.NewTicker(weatherDefaultInterval)

	for range tick.C {
		blocksRendered <- w.getBlock()
	}
}

func NewWeatherAddon(config AddonConfig, index int) Addon {

	w := &weatherAddon{
		index,
		dateDefaultFormat,
	}

	if location, ok := config["location"].(string); ok {
		w.location = location
	} else {
		return NewCustomMsgAddon("Weather: no location", index, ColorRed)
	}

	return w
}
