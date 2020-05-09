package addon

import (
	"time"
)

const defaultTimeFormat = "15:04:05";

// Timer addon, used to display the current time
type timer struct {
	format string
}

func (t *timer) Update() *Block {
	return &Block{
		FullText: time.Now().Format(t.format),
	}
}

func NewTimeAddon(format string) *Addon {
	if format == "" {
		format = defaultTimeFormat
	}

	return &Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater: &timer{format},
	}
}
