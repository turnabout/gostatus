package addon

import (
	"fmt"
	"os"
	"time"
)

type memoryAddon struct {
	index int
}

const(
	memoryDefaultFormat     = "%s %.1fGB"
	memoryDefaultInterval   = 5 * time.Second
	memoryColorOk           = ColorLime
	memoryColorWarning      = ColorYellow
	memoryColorCritical     = ColorRed
	memoryThresholdWarning  = 5.0
	memoryThresholdCritical = 2.0
)

func (m *memoryAddon) getBlock() *Block {
	var err error

	// Get available memory
	var memAvail, memTotal, memFree int32
	r, err := os.Open("/proc/meminfo")

	if err != nil {
		return nil
	}

	defer r.Close()

	_, err = fmt.Fscanf(
		r,
		"MemTotal: %d kB\nMemFree: %d kB\nMemAvailable: %d ",
		&memTotal,
		&memFree,
		&memAvail,
	)

	gbAvail := float32(memAvail) / 1024 / 1024

	// Get appropriate color depending on how much memory is used
	var color string

	if gbAvail < memoryThresholdCritical {
		color = memoryColorCritical
	} else if gbAvail < memoryThresholdWarning {
		color = memoryColorWarning
	} else {
		color = memoryColorOk
	}

	return &Block{
		FullText: fmt.Sprintf(
			memoryDefaultFormat,
			IconMemory,
			gbAvail,
		),
		Color: color,
		Index: m.index,
	}
}

func (m *memoryAddon) Run(blocks chan *Block, blocksRendered chan *Block) {
	blocks <- m.getBlock()

	tick := time.NewTicker(memoryDefaultInterval)

	for range tick.C {
		blocks <- m.getBlock()
	}
}

func NewMemoryAddon(config AddonConfig, index int) Addon {
	m := &memoryAddon{
		index,
	}

	return m
}
