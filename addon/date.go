package addon

/*
import (
	"time"
)

const defaultDateFormat = "Mon Jan 02 2006";

// Date addon, used to display the current date
type date struct {
	format string
}

func (d *date) Update() *Block {
	return &Block{
		FullText: time.Now().Format(d.format),
	}
}

func NewDateAddon(format string) *Addon {
	if format == "" {
		format = defaultDateFormat
	}

	return &Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater:        &timer{format},
	}
}

 */
