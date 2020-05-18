package addon

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
)

type volumeAddon struct {
	index int
}

const(
	volumeDefaultFormat = "%s %s"
	volumeCmd           = `awk -F"[][]" '/dB/ { print $2 }' <(amixer sget Master)`
	volumeMuteCmd       = "pacmd list-sinks | awk '/muted/ { print $2 }'"
)

// Gets the current volume.
// Returns whether the volume is muted, and the volume percentage in a string, formatted like "55%".
func (v *volumeAddon) getBlock() *Block {

	// Get whether volume is muted
	var cmdMutedOut []byte

	cmd := volumeMuteCmd
	cmdMutedOut, _ = exec.Command("bash", "-c", cmd).Output();
	muted := strings.TrimSpace(string(cmdMutedOut)) == "yes"

	// Get volume percentage
	var cmdVolumeOut []byte

	cmd = volumeCmd
	cmdVolumeOut, _ = exec.Command("bash", "-c", cmd).Output();
	volume := strings.TrimSpace(string(cmdVolumeOut))

	// Get appropriate icon/color based on whether volume is muted
	var icon, color string

	if muted {
		icon = IconVolumeMuted
		color = ColorRed
	} else {
		icon = IconVolume
		color = ColorWhite
	}

	return &Block{
		FullText: fmt.Sprintf(volumeDefaultFormat, icon, volume),
		Color: color,
		Index: v.index,
	}
}

func (v *volumeAddon) Run(blocks chan *Block, blocksRendered chan *Block) {
	blocks <- v.getBlock()

	// Listen for external volume signal
	sigs := make(chan os.Signal)
	signal.Notify(sigs, SignalVolume)

	for range sigs {
		blocksRendered <- v.getBlock()
	}
}

func NewVolumeAddon(config AddonConfig, index int) Addon {
	v := &volumeAddon{
		index,
	}

	return v
}
