package addon

type msgAddon struct {
	index int
	text  string
	color string
}

func (m *msgAddon) Run(blocks chan *Block, blocksRendered chan *Block) {
	blocks <- m.getBlock()
}

func (m *msgAddon) getBlock() *Block {
	return &Block{
		FullText: m.text,
		Index:    m.index,
		Color:    m.color,
	}
}

func NewMsgAddon(config AddonConfig, index int) Addon {
	a := &msgAddon{
		index,
		"",
		ColorWhite,
	}

	if text, ok := config["text"].(string); ok {
		a.text = text
	}

	return a
}

func NewCustomMsgAddon(text string, index int, color string) Addon {
	return &msgAddon{
		index,
		text,
		color,
	}
}
