package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)
import "fyne.io/fyne/v2"


func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

        label := widget.NewLabel("Hallo")
        label.TextSize = 24

	w.SetContent(label)
	w.ShowAndRun()
}

