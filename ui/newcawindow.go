package ui

import "fyne.io/fyne/v2"

func NewCaWindow() {
	w := fyne.CurrentApp().NewWindow("Test New Widnow")
	w.SetFixedSize(true)
}
