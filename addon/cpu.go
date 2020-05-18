package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"os/exec"
	"time"
)

type cpuAddon struct {
	index int
}

const(
	cpuDefaultFormat     = "%s%3d%%"
	cpuDefaultInterval   = 1 * time.Second
	cpuCommand           = "top -bn1 | sed -n '/Cpu/p'"
	cpuColorOk           = ColorWhite
	cpuColorWarning      = ColorYellow
	cpuColorCritical     = ColorRed
	cpuThresholdWarning  = 75
	cpuThresholdCritical = 90
)

func (c *cpuAddon) Run(blocks chan *Block, blocksRendered chan *Block) {

	tick := time.NewTicker(cpuDefaultInterval)

	for range tick.C {
		blocks <- c.getBlock()
	}
}

func (c *cpuAddon) getBlock() *Block {
	var err error

	// Get command output
	var cmdOut []byte

	cmd := "top -bn1 | sed -n '/Cpu/p'"
	cmdOut, err = exec.Command("bash", "-c", cmd).Output();

	if err != nil {
		log.Error(err)
		return nil
	}

	// Extract percentage from command output
	var usageFloat float32

	_, err = fmt.Sscanf(
		string(cmdOut),
		"%%Cpu(s): %f us",
		&usageFloat,
	)

	if err != nil {
		return nil
	}

	usage := int(usageFloat)

	// Change color depending on usage percentage
	var color string

	if usage > cpuThresholdCritical {
		color = cpuColorCritical
	} else if usage > cpuThresholdWarning {
		color = cpuColorWarning
	} else {
		color = cpuColorOk
	}

	return &Block{
		FullText: fmt.Sprintf(
			cpuDefaultFormat,
			IconCPU,
			int(usageFloat),
		),
		Color: color,
		Index: c.index,
	}
}

func NewCpuAddon(config AddonConfig, index int) Addon {
	return &cpuAddon{
		index,
	}
}
