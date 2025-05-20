package ui

type Label struct {
	Text string
}

func NewLabel(text string) *Label {
	return &Label{Text: text}
}
