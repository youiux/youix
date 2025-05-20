package platform

// Window is the interface for a windowing system.
type Window interface {
	Show()
	DrawText(x, y int, text string)
}

// NewWindow creates a new platform-specific window.
func NewWindow(title string, width, height int) Window {
	return newX11Window(title, width, height)
}
