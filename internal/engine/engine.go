package engine

import (
	"github.com/youiux/youix/internal/api"
	"github.com/youiux/youix/internal/ui"
)

func Run() {
	window := api.NewWindow("Hello Window", 800, 500)
	label := ui.NewLabel("Hello, World!")
	window.DrawText(50, 100, label.Text)
	window.Show()
}
