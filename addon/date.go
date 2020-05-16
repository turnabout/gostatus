package addon

import (
	"time"
)

// Date addon, used to display the current date
type dateAddon struct {
	format string
	index  uint
}

const(
	defaultDateFormat = "Mon Jan 02 2006"
)

// Returns the duration from now until tomorrow
func getDurationTillTomorrow() time.Duration {
	now := time.Now()

	tomorrow := time.Date(
		now.Year(),
		now.Month(),
		now.Day() + 1,
		0,
		0,
		0,
		0,
		now.Location(),
	)

	return tomorrow.Sub(now)
}

func (d *dateAddon) Run(blocks chan *Block, blocksRendered chan *Block) {

	blocks <- d.getBlock()

	// Send new block when the date changes
	for {
		<- time.NewTimer(getDurationTillTomorrow()).C
		blocks <- d.getBlock()
	}
}

func (d *dateAddon) getBlock() *Block {
	return &Block{
		FullText: time.Now().Format(d.format),
		Index: d.index,
	}
}

func NewDateAddon(format string, index uint) Addon {
	if format == "" {
		format = defaultDateFormat
	}

	return &dateAddon{
		format,
		index,
	}
}
