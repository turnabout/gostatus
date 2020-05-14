package addon

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// Date addon, used to display the current date
type dateAddon struct {
	format string
	index  int
}

const(
	defaultDateFormat = "Mon Jan 02 2006"
	defaultDateUpdateInterval = 1000 * time.Millisecond
)

func (t *dateAddon) Run(blocks chan *Block) {

	sigs := make(chan os.Signal)
	blocks <- t.getBlock()

	signal.Notify(sigs, SignalDate)

	i := 0

	for {
		<- sigs

		//fmt.Println(sig)

		blocks <- &Block{
			FullText: fmt.Sprintf("SIGNAAAAAAl %d", i),
			Index: t.index,
		};
		i++
	}
}

func (t *dateAddon) getBlock() *Block {
	return &Block{
		FullText: time.Now().Format(t.format),
		Index: t.index,
	}
}

func NewDateAddon(format string, index int) Addon {
	if format == "" {
		format = defaultDateFormat
	}

	return &dateAddon{
		format,
		index,
	}
}
