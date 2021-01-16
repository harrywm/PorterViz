package gui

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// Main gui func
func Main() {

	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
