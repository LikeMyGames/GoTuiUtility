package main

import (
	"os"

	"golang.org/x/term"
)

type (
	Window struct {
		width   int
		height  int
		panels  []Panel
		options PanelOptions
	}

	PanelOptions struct {
		HasBorders     bool
		RoundedCorners bool
	}

	Panel interface {
		GetOptions() PanelOptions
		GetWidth() int
		GetHeight() int
		GetChildren() []Panel
		AddChild(Panel)
		Render()
	}
)

func main() {
	window := NewWindow(PanelOptions{HasBorders: true, RoundedCorners: true})
	window.Render()
}

func NewWindow(options PanelOptions) *Window {
	termWidth, termHeight, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}
	return &Window{
		width:   termWidth,
		height:  termHeight,
		panels:  []Panel{},
		options: options,
	}
}

func (w *Window) Render() {
	if w.options.HasBorders {
		for range w.width {

		}
	}
}

func (w *Window) AddChild(p Panel) {
	w.panels = append(w.panels, p)
}

func (w *Window) GetWidth() int {
	return w.width
}

func (w *Window) GetHeight() int {
	return w.height
}

func (w *Window) GetChildren() []Panel {
	return w.panels
}

func (w *Window) GetOptions() PanelOptions {
	return w.options
}
