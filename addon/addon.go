package addon

type Addon interface {
	Run(ch chan *Block)
	getBlock() *Block
}


/*
type Addon struct {
	UpdateInterval time.Duration
	Updater        Updater

	//Icon           string
	//Run            AddonRunner
}
 */

// Runs the Addon continuously, at its given interval
/*
func (a *Addon) Run(ch chan *Block) {

	for range time.Tick(a.UpdateInterval) {
		// generating data should not be locked
		newData := a.Updater.Update()
		a.NewData = newData
	}
}
 */
