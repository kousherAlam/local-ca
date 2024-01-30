package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/kousheralam/local-ca/ui"
)

const (
	APP_NAME = "Local CA"
)

func main() {
	a := app.New()
	w := a.NewWindow(APP_NAME)
	w.Resize(fyne.NewSize(500, 400))

	w.SetContent(ui.HomeWindow())

	w.ShowAndRun()
}
