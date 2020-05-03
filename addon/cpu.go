package addon

import (
	"fmt"
	"github.com/lsgrep/gostatus/log"
	"os/exec"
	"strings"
	"time"
)

const cpuCommand = "top -bn1 | sed -n '/Cpu/p'"

type cpu struct {
}

func (c *cpu) Update() *Block {
	out, err := exec.Command("bash", "-c", cpuCommand).Output();

	if err != nil {
		log.Error(err)
		return nil
	}

	// Extract percentage from command output
	pStart := strings.Index(string(out), ":") + 1
	pEnd := strings.Index(string(out), " us")

	formattedOut := strings.TrimSpace(string(out)[pStart: pEnd])

	// Convert to a float value so we can change output color based on how high it is
	// TODO
	// floatVal, _ := strconv.ParseFloat(formattedOut, 16)

	return &Block{FullText: fmt.Sprintf("%s %s%%", IconCPU, formattedOut)}
}

func NewCPUAddon() *Addon {
	c := &cpu{}
	return &Addon{
		UpdateInterval: 3000 * time.Millisecond,
		Updater:        c}
}
