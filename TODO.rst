=================
TODO/Random Notes
=================

TODO: Adjust colors
===================
Adjust colors to match the jellybeans color scheme


TODO: Place this command string result extracting logic in reusable function
============================================================================

	cmdVolumeOut, err = exec.Command("bash", "-c", cmd).Output();

	return strings.TrimSpace(string(cmdMutedOut)) == "yes", strings.TrimSpace(string(cmdVolumeOut))


TODO: Active keyboard layout addon
==================================

Add an addon displaying the currently active keyboard layout (en-us or fr-ca)


TODO: Make project structure match the recommended ghub go structure

TODO: Look into auto running go fmt on project on file save

TODO: Add a "addons generic data" struct, which contains some generic data for
      an addon (e.g. output format, update interval, default color, etc)
      and attach it to every addon
      Add a function that processes an "addonConfig" map and returns a 
      corresponding "addons generic data" struct

TODO: Add possible mouse click handling on go modules
      + Add updating of audio when mouse scroll up/down on it

TODO: Look into adding dropdown menus appearing over bar on click, and handle clicks on those



Possible alternate colors
	ColorWhite   = "#E8E8D3"
	ColorBlack   = "#3B3B3B"
	ColorRed     = "#CF6A4C"
	ColorMaroon  = "#FF9D80"
	ColorYellow  = "#FAD07A"
	ColorGreen   = "#C3E6AD"
	ColorBlue    = "#8198BF"
	ColorMagenta = "#8787AF"
	ColorCyan    = "#8FBFDC"


