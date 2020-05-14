package addon

import (
	"time"
)

// Timer addon, used to display the current time
type timer struct {
	format string
	index  int
}

const(
	defaultTimeFormat = "15:04:05"
	defaultTimeUpdateInterval = 1000 * time.Millisecond
)

func (t *timer) Run(blocks chan *Block) {
	blocks <- t.getBlock()

	for range time.Tick(defaultTimeUpdateInterval) {
		blocks <- t.getBlock()
	}
}

func (t *timer) getBlock() *Block {
	return &Block{
		FullText: time.Now().Format(t.format),
		Index: t.index,
	}
}

func NewTimeAddon(format string, index int) Addon {
	if format == "" {
		format = defaultTimeFormat
	}

	return &timer{
		format,
		index,
	}
}
