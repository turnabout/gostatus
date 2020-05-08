package addon

import (
	"fmt"
	"os/exec"
	"time"
)

type memory struct {
}

const(
	memoryColorOk           = ColorLime
	memoryColorWarning      = ColorYellow
	memoryColorCritical     = ColorRed
	memoryThresholdWarning  = 5.0
	memoryThresholdCritical = 2.0
)

func (m *memory) Update() *Block {
	var err error

	// Get `free` output
	var cmdOut []byte

	cmd := "free -m | sed -n 2p"
	cmdOut, err = exec.Command("bash", "-c", cmd).Output();

	if err != nil {
		return nil
	}

	// Scan command output, extract MB values from it
	var mbTotal int
	var mbUsed int
	var mbFree int

	_, err = fmt.Sscanf(
		string(cmdOut),
		"Mem: %d %d %d",
		&mbTotal,
		&mbUsed,
		&mbFree,
	)

	if err != nil {
		return nil
	}

	// Translate values from MB => GB
	gbFree := (float64(mbFree) / 1024)

	// Get appropriate color depending on how much memory is used
	var color string

	if gbFree < memoryThresholdCritical {
		color = memoryColorCritical
	} else if gbFree < memoryThresholdWarning {
		color = memoryColorWarning
	} else {
		color = memoryColorOk
	}

	return &Block{
		FullText: fmt.Sprintf(
			"%s %.2fGB",
			IconMemory,
			gbFree,
		),
		Color: color,
	}
}

func NewMemoryAddon() *Addon {
	m := &memory{}
	return &Addon{
		UpdateInterval: 3000 * time.Millisecond,
		Updater:        m}
}
