package addon

import (
	"fmt"
	"syscall"
	"time"
)

type diskStatus struct {
	Path string
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

const(
	diskColorOk           = ColorLime
	diskColorWarning      = ColorYellow
	diskColorCritical     = ColorRed
	diskThresholdWarning  = 50.0
	diskThresholdCritical = 20.0
)

func (ds *diskStatus) Update() *Block {

	// Gather disk data
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(ds.Path, &fs)

	if err != nil {
		return nil
	}

	freeMem := fs.Bfree * uint64(fs.Bsize)

	// Get appropriate color
	var color string

	if freeMem < diskThresholdCritical {
		color = diskColorCritical
	} else if freeMem < diskThresholdWarning {
		color = diskColorWarning
	} else {
		color = diskColorOk
	}

	// Get addon text
	text := fmt.Sprintf(
		"[%s] %.2fGB",
		ds.Path,
		float64(freeMem) / float64(GB),
	)

	return &Block{
		FullText: fmt.Sprintf("%s %s", IconDisk, text),
		Color: color,
	}
}

func NewDiskAddon(path string) *Addon {
	ds := &diskStatus{Path: path}

	return &Addon{
		UpdateInterval: 5000 * time.Millisecond,
		Updater:        ds,
	}
}
