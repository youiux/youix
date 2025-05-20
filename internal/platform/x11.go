//go:build linux
// +build linux

package platform

import (
	"log"
	"os/exec"
)

type x11Window struct {
	title  string
	width  int
	height int
	text   []drawTextRequest
}

type drawTextRequest struct {
	x, y int
	text string
}

func newX11Window(title string, width, height int) Window {
	return &x11Window{title: title, width: width, height: height}
}

func (w *x11Window) Show() {
	// For demonstration, use xmessage to show a window with all DrawText requests concatenated
	msg := w.title + ":\n"
	for _, t := range w.text {
		msg += t.text + "\n"
	}
	cmd := exec.Command("xmessage", "-center", "-buttons", "OK:0", msg)
	err := cmd.Run()
	if err != nil {
		log.Println("Failed to show window:", err)
	}
}

func (w *x11Window) DrawText(x, y int, text string) {
	w.text = append(w.text, drawTextRequest{x: x, y: y, text: text})
}
