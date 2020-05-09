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

	var err error
	var cmdOut []byte
	var cmd string

	// Get whether volume is muted
	cmd = "pacmd list-sinks | awk '/muted/ { print $2 }'"
	cmdOut, err = exec.Command("bash", "-c", cmd).Output();

	if err != nil {
		return true, ""
	}

	if strings.TrimSpace(string(cmdOut)) == "yes" {
		return true, ""
	}

	// Get volume percentage
	cmd = `awk -F"[][]" '/dB/ { print $2 }' <(amixer sget Master)`
	cmdOut, err = exec.Command("bash", "-c", cmd).Output();

	return false, strings.TrimSpace(string(cmdOut))
}

func (vs *volumeStatus) Update() *Block {

	muted, volume := GetVolume()

	// Get appropriate text/color based on whether volume is muted
	var text string
	var color string

	if muted {
		text = IconVolumeMuted
		color = ColorRed
	} else {
		text = fmt.Sprintf("%s %s", IconVolume, volume)
		color = ColorWhite
	}

	return &Block{
		FullText: text,
		Color: color,
	}
}

func NewVolumeAddon() *Addon {
	return &Addon{
		UpdateInterval: 1000 * time.Millisecond,
		Updater:        &volumeStatus{},
	}
}
