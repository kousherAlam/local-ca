package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	WELCOME = "Welcome to Local CA"
	SUBTEXT = "Manage your Certificate Easily"
)

func HomeWindow() *fyne.Container {

	welcome := widget.NewRichText(&widget.TextSegment{
		Text: WELCOME,
		Style: widget.RichTextStyle{
			SizeName: widget.RichTextStyleInline.SizeName,
		},
	})
	subText := widget.NewRichText(&widget.TextSegment{
		Text: SUBTEXT,
		Style: widget.RichTextStyle{
			SizeName: widget.RichTextStyleBlockquote.SizeName,
		},
	})

	headingLebel := container.New(layout.NewCenterLayout(), welcome)
	subTextLebel := container.New(layout.NewCenterLayout(), subText)

	button := widget.NewButton("Start CA", func() {
		// welcome.Segments[0].Visual().Show()
		fmt.Println("Showing new window.")
		NewCaWindow()
	})

	grid := container.New(layout.NewGridLayout(3), EmptyBox, button, EmptyBox)

	mainContainer := container.New(layout.NewVBoxLayout(), headingLebel, subTextLebel, grid)

	return mainContainer
}
