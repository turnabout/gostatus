package addon

import (
	"time"
)

// Time addon, used to display the current time
type timeAddon struct {
	format string
	index  uint
}

const(
	defaultTimeFormat = "15:04:05"
	defaultTimeUpdateInterval = 1000 * time.Millisecond
)

func (t *timeAddon) Run(blocks chan *Block, blocksRendered chan *Block) {
	blocks <- t.getBlock()

	for range time.Tick(defaultTimeUpdateInterval) {
		blocks <- t.getBlock()
	}
}

func (t *timeAddon) getBlock() *Block {
	return &Block{
		FullText: time.Now().Format(t.format),
		Index: t.index,
	}
}

func NewTimeAddon(format string, index uint) Addon {
	if format == "" {
		format = defaultTimeFormat
	}

	return &timeAddon{
		format,
		index,
	}
}
