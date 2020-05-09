package addon

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type volumeStatus struct {
}

// Gets the current volume.
// Returns whether the volume is muted, and the volume percentage in a string, formatted like "55%".
func GetVolume() (bool, string) {

	var cmd string

	// Get whether volume is muted
	var cmdMutedOut []byte

	cmd = "pacmd list-sinks | awk '/muted/ { print $2 }'"
	cmdMutedOut, _ = exec.Command("bash", "-c", cmd).Output();

	// Get volume percentage
	var cmdVolumeOut []byte
	cmd = `awk -F"[][]" '/dB/ { print $2 }' <(amixer sget Master)`
	cmdVolumeOut, _ = exec.Command("bash", "-c", cmd).Output();

	return strings.TrimSpace(string(cmdMutedOut)) == "yes", strings.TrimSpace(string(cmdVolumeOut))
}

func (vs *volumeStatus) Update() *Block {

	muted, volume := GetVolume()

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
		FullText: fmt.Sprintf("%s %s", icon, volume),
		Color: color,
	}
}

func NewVolumeAddon() *Addon {
	return &Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater:        &volumeStatus{},
	}
}
