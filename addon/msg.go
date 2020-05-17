package addon

type msgAddon struct {
	index int
	text  string
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

func NewMsgAddon(config AddonConfig, index int) Addon {
	a := &msgAddon{
		index,
		"",
	}

	if format, ok := config["text"].(string); ok {
		a.text = format
	}

	return a
}

func NewMsgAddonByText(text string, index int) Addon {
	return &msgAddon{
		index,
		text,
	}
}
