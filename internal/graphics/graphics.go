package graphics

import (
	"fmt"

	"github.com/youiux/youix/internal/ui"
)

type Graphics struct {
	labels []*ui.Label
}

func New() *Graphics {
	return &Graphics{}
}

func (g *Graphics) DrawLabel(label *ui.Label) {
	g.labels = append(g.labels, label)
}

func (g *Graphics) Show() {
	for _, label := range g.labels {
		fmt.Println(label.Text)
	}
}
