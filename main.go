package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/canvas"
)

import "fyne.io/fyne/v2"

// toolbar consists of 
// rootHBox - horizontal box containing things
// open button - to open different file
// navigation buttons - to move to next or previous image
// in this folder
// scale buttons - reset, increase and decrease scale
type Toolbar struct {
	Root *fyne.Container
	OpenButton *widget.Button
	
	PrevButton *widget.Button
	NextButton *widget.Button
	
	ScaleIncButton *widget.Button
	ScaleDecButton *widget.Button
	ScaleResetButton *widget.Button
}


// MyApp is central global point of 
// component intercommunication, like
// any button from anywhere can change
// image.
type MyApp struct {
	toolbar *Toolbar
	Image *canvas.Image
	Root *fyne.Container
}

var myapp *MyApp

func NewToolbar() *Toolbar {
	OpenButton := widget.NewButton("open", func() {
		
	})
	PrevButton := widget.NewButton("<-", func() {})
	NextButton := widget.NewButton("->", func() {})
	
	ScaleIncButton := widget.NewButton("+", func() {
		oldSize := myapp.Image.MinSize()
		oldSize.Width *= 1.25
		oldSize.Height *= 1.25
		myapp.Image.SetMinSize(oldSize)
	})
	ScaleDecButton := widget.NewButton("-", func() {
		oldSize := myapp.Image.MinSize()
		oldSize.Width *= 0.85
		oldSize.Height *= 0.85
		myapp.Image.SetMinSize(oldSize)
		myapp.Root.Refresh()
	})
	ScaleResetButton := widget.NewButton("=", func() {
		bounds := myapp.Image.Image.Bounds()
		myapp.Image.SetMinSize(fyne.Size{
			Width: float32(bounds.Max.X - bounds.Min.X),
			Height: float32(bounds.Max.Y - bounds.Min.Y),
		})
	})
	RootHBox := container.NewHBox(OpenButton, PrevButton, NextButton, ScaleIncButton, ScaleDecButton, ScaleResetButton)
	return &Toolbar {
		OpenButton: OpenButton,
		PrevButton: PrevButton,
		NextButton: NextButton,
		ScaleIncButton: ScaleIncButton,
		ScaleDecButton: ScaleDecButton,
		ScaleResetButton: ScaleResetButton,
		Root: RootHBox,
	}
}

func NewMyApp() *MyApp {
	app := & MyApp {
		toolbar: NewToolbar(),
		Image: canvas.NewImageFromResource(resourceIconWhitePng),
		Root: container.New(layout.NewVBoxLayout()),
	}
	
	app.Image.FillMode = canvas.ImageFillContain
	//app.image.ScaleMode = canvas.ImageScaleFastest
	
	
	app.Image.SetMinSize(fyne.Size{Width: 512, Height: 512})
	
	app.Root = container.NewBorder(
		app.toolbar.Root,
		nil,
		nil,
		nil,
		container.NewScroll(
			container.New(layout.NewCenterLayout(), app.Image)),
	)
	
	return app
}


func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
//	label2 := widget.NewLabel("Hello")
	
//	toolbar := NewToolbar()
	
	
//	root := container.New(layout.NewVBoxLayout(), 
//		container.New(layout.NewCenterLayout(), toolbar.RootHBox), 
//		label2)
		
	myapp = NewMyApp()

	w.SetContent(myapp.Root)
	w.Resize(fyne.Size{Width: 512, Height: 512})
	w.ShowAndRun()
}

