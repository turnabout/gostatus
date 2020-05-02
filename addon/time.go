package addon

import (
	"fmt"
	"time"
)

const defaultTimeFormat = "Mon Jan 02 2006 15:04:05";

type timer struct {
	format string
}

func (t *timer) Update() *Block {
	fullTxt := fmt.Sprintf(
		" %s  %s",
		IconTime,
		time.Now().Format(t.format),
	)
	return &Block{FullText: fullTxt}
}

func NewTimeAddon(format string) *Addon {
	if format == "" {
		format = defaultTimeFormat
	}

	t := &timer{
		format,
	}
	aa := Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater: t,
	}
	return &aa
}
