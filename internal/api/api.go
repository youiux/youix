package api

import (
	"github.com/youiux/youix/internal/platform"
)

type Window interface {
	Show()
	DrawText(x, y int, text string)
}

func NewWindow(title string, width, height int) Window {
	return platform.NewWindow(title, width, height)
}
