package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SeaveMenuItem *fyne.MenuItem
}

var cfg config

func main() {
	// Create a new instance of the application
	app := app.New()

	// Create a new window with the title "Markdown Editor"
	window := app.NewWindow("Markdown Editor")

	// Create the UI components for editing and previewing markdown content
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(window)

	// Create a vertical box container and add the edit and preview components to it
	

	// Set the content of the window to the vertical box container
	window.SetContent(container.NewVSplit(edit, preview))
	window.Resize(fyne.Size{Width: 800, Height: 600})
	window.CenterOnScreen()
	window.SetIcon(fyne.CurrentApp().Icon())

	// Show the window and start the application event loop
	window.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (app *config) createMenuItems(window fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open...", func() {})
	saveMenuItem := fyne.NewMenuItem("Save", func() {})
	saveAsMenuItem := fyne.NewMenuItem("Save As...", func() {})
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)
	menu := fyne.NewMainMenu(fileMenu)
	window.SetMainMenu(menu)

}
