package addon

import (
	"time"
)

// for each data fetcher
type Updater interface {
	Update() *Block
}

type Addon struct {
	// guard lastData
	NewData *Block

	UpdateInterval time.Duration
	Icon           string
	Updater        Updater
}

// Runs the Addon continuously, at its given interval
func (a *Addon) Run() {

	for range time.Tick(a.UpdateInterval) {
		// generating data should not be locked
		newData := a.Updater.Update()
		a.NewData = newData
	}
}
