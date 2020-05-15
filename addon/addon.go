package addon

type Addon interface {
	Run(ch chan *Block)
	getBlock() *Block
}
