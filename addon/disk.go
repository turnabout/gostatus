package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"syscall"
	"time"
)

type diskAddon struct {
	index int
	path string
}

const(
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	diskColorOk           = ColorLime
	diskColorWarning      = ColorYellow
	diskColorCritical     = ColorRed
	diskThresholdWarning  = 50.0
	diskThresholdCritical = 20.0
	diskDefaultInterval   = 30 * time.Second
)

func (d *diskAddon) getBlock() *Block {
	// Gather disk data
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(d.path, &fs)

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
		d.path,
		float64(freeMem) / float64(GB),
	)

	return &Block{
		FullText: fmt.Sprintf("%s %s", IconDisk, text),
		Color: color,
		Index: d.index,
	}
}

func (d *diskAddon) Run(blocks chan *Block, blocksRendered chan *Block) {
	blocks <- d.getBlock()

	tick := time.NewTicker(diskDefaultInterval)

	for range tick.C {
		blocks <- d.getBlock()
	}
}

func NewDiskAddon(config AddonConfig, index int) Addon {
	d := &diskAddon{
		index,
		"",
	}

	if path, ok := config["path"].(string); ok {
		d.path = path
	} else {
		log.Error("NewDiskAddon: no path given")

		return NewCustomMsgAddon(
			"No path!",
			index,
			ColorRed,
		)
	}

	return d
}
