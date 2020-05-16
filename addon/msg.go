package addon

type msgAddon struct {
	text  string
	index uint
}

func (m *msgAddon) Run(blocks chan *Block, blocksRendered chan *Block) {
	blocks <- m.getBlock()
}

func (m *msgAddon) getBlock() *Block {
	return &Block{
		FullText: m.text,
		Index: m.index,
	}
}

func NewMsgAddon(text string, index uint) Addon {
	return &msgAddon{
		text,
		index,
	}
}

func NewMsgAddonByText(text string, index uint) Addon {
	return &msgAddon{
		text,
		index,
	}
}
