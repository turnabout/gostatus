package addon

/*
import (
	"fmt"
	"os"
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
			"%s %.1fGB",
			IconMemory,
			gbAvail,
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

 */
