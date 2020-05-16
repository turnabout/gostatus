package addon

// Interface that wraps all addons
//
// Run contains the logic that produces new Blocks over time.
// When it updates is dependant on the addon's own logic.
// blocks and blocksRendered are the channels where the Addon must send new Blocks.
// Blocks sent to blocks will be passed to the status bar, but will not instantly make it render.
// Blocks sent to blocksRendered will be passed to the status bar, and will make it render
// immediately upon reception.
//
// getBlock returns a new, updated block.
type Addon interface {
	Run(blocks chan *Block, blocksRendered chan *Block)
	getBlock() *Block
}
