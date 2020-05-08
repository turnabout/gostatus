package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"os/exec"
	"time"
)

const cpuCommand = "top -bn1 | sed -n '/Cpu/p'"

type cpu struct {
}

const(
	cpuColorOk           = ColorWhite
	cpuColorWarning      = ColorYellow
	cpuColorCritical     = ColorRed
	cpuThresholdWarning  = 75
	cpuThresholdCritical = 90
)

func (c *cpu) Update() *Block {
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
		FullText: fmt.Sprintf("%s%3d%%",
			IconCPU,
			int(usageFloat),
		),
		Color: color,
	}
}

func NewCPUAddon() *Addon {
	c := &cpu{}

	return &Addon{
		UpdateInterval: 3000 * time.Millisecond,
		Updater:        c,
	}
}
