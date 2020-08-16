package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"os/exec"
	"strings"
)

type kbLayoutAddon struct {
	index  int
}

const(
	kbDefaultFormat = "%s"
	kbLayoutCmd = "setxkbmap -query | grep layout | tail -c 3"
)

func (k *kbLayoutAddon) getBlock() *Block {

	cmdOut, err := exec.Command(
		"bash",
		"-c",
		kbLayoutCmd,
	).Output();

	if err != nil {
		log.Error(err)
		return nil
	}

	return &Block{
		FullText: fmt.Sprintf(
			kbDefaultFormat,
			strings.TrimSpace(string(cmdOut)),
		),
		Index: k.index,
	}
}

func (k *kbLayoutAddon) Run(blocks chan *Block, blocksRendered chan *Block) {

	blocks <- k.getBlock()

	/*
	tick := time.NewTicker(weatherDefaultInterval)

	for range tick.C {
		blocksRendered <- k.getBlock()
	}
	 */
}

func NewKbLayoutAddon(config AddonConfig, index int) Addon {

	k := &kbLayoutAddon{index}

	return k
}
