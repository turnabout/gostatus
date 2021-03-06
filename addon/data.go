package addon

import (
	"syscall"
)

type Block struct {
	FullText            string                 `json:"full_text"`
	ShortText           string                 `json:"short_text,omitempty"`
	Color               string                 `json:"color,omitempty"`
	BorderColor         string                 `json:"border,omitempty"`
	BackgroundColor     string                 `json:"background,omitempty"`
	Markup              string                 `json:"markup,omitempty"`
	MinWidth            string                 `json:"min_width,omitempty"`
	Align               string                 `json:"align,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Instance            string                 `json:"instance,omitempty"`
	Urgent              bool                   `json:"urgent,omitempty"`
	Separator           *bool                  `json:"separator,omitempty"`
	SeparatorBlockWidth uint16                 `json:"separator_block_width,omitempty"`
	Custom              map[string]interface{} `json:"-"`
	Index               int                    `json:"-"`
}

type ClickEvent struct {
	Name      string `json:"name,omitempty"`
	Instance  string `json:"instance,omitempty"`
	Button    uint8  `json:"button"`
	X         uint16 `json:"x"`
	Y         uint16 `json:"y"`
	RelativeX uint16 `json:"relative_x"`
	RelativeY uint16 `json:"relative_y"`
	Width     uint16 `json:"width"`
	Height    uint16 `json:"height"`
}

// icons
const (
	IconGithub      = "\uf30a"
	IconDisk        = "\uf0a0"
	IconMemory      = "\uf399"
	IconCPU         = "\uf2db"
	IconIP          = "\uf381"
	IconNetwork     = "\uf381"
	IconVolume      = "\uf3b3"
	IconVolumeMuted = "\uf380"
	IconDate        = "\uf073"
	IconTime        = "\uf346"
	IconPomodoro    = "\uf0ae"
	IconWork        = "\uf0e7"
	IconPlay        = "\uf439"
	IconWeather0    = "\uf2cb"
	IconWeather1    = "\uf2ca"
	IconWeather2    = "\uf2c9"
	IconWeather3    = "\uf2c8"
	IconWeather4    = "\uf2c7"
)

const (
	ColorWhite   = "#FFFFFF"
	ColorSilver  = "#C0C0C0"
	ColorGray    = "#808080"
	ColorBlack   = "#000000"
	ColorRed     = "#FF0000"
	ColorMaroon  = "#800000"
	ColorYellow  = "#FFFF00"
	ColorOlive   = "#808000"
	ColorLime    = "#00FF00"
	ColorGreen   = "#008000"
	ColorAqua    = "#00FFFF"
	ColorTeal    = "#008080"
	ColorBlue    = "#0000FF"
	ColorNavy    = "#000080"
	ColorFuchsia = "#FF00FF"
	ColorPurple  = "#800080"
)

// List of extra real-time signals that can be received by addons from the external system
const (
	SIGRTMIN5  = syscall.Signal(0x27)
	SIGRTMIN6  = syscall.Signal(0x28)
	SIGRTMIN7  = syscall.Signal(0x29)
	SIGRTMIN8  = syscall.Signal(0x2a)
	SIGRTMIN9  = syscall.Signal(0x2b)
	SIGRTMIN10 = syscall.Signal(0x2c)
	SIGRTMIN11 = syscall.Signal(0x2d)
	SIGRTMIN12 = syscall.Signal(0x2e)
	SIGRTMIN13 = syscall.Signal(0x2f)
)

// List of signals used by addons
const (
	SignalVolume   = SIGRTMIN5
	SignalKbLayout = SIGRTMIN6
)
