package addon

import (
	"time"
)

// Time addon, used to display the current time
type timeAddon struct {
	index  int
	format string
}

const(
	defaultTimeFormat = "15:04:05"
	defaultTimeUpdateInterval = 1 * time.Second
)

func (t *timeAddon) Run(blocks chan *Block, blocksRendered chan *Block) {

	// Delay start by 1 millisecond to make sure this addon starts slightly after the others
	time.AfterFunc(1 * time.Millisecond, func() {
		blocksRendered <- t.getBlock()

		t.innerRun(blocks, blocksRendered)
	})
}

func (t *timeAddon) innerRun(blocks chan *Block, blocksRendered chan *Block) {

	tick := time.NewTicker(defaultTimeUpdateInterval)

	for range tick.C {
		blocksRendered <- t.getBlock()
	}
}

func (t *timeAddon) getBlock() *Block {
	return &Block{
		FullText: time.Now().Format(t.format),
		Index: t.index,
	}
}

func NewTimeAddon(config AddonConfig, index int) Addon {
	a := &timeAddon{
		index,
		defaultTimeFormat,
	}

	if format, ok := config["format"].(string); ok {
		a.format = format
	}

	return a
}
