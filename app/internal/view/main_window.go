package view

import (
	"app/internal/model"
	"app/internal/viewmodel"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// RunApp starts the application
func RunApp() {
	a := app.New()
	w := a.NewWindow("Simple Desktop App")

	model := model.NewModel()
	viewModel := viewmodel.NewViewModel(model)

	label := widget.NewLabelWithData(viewModel.GetMessage())
	button := widget.NewButton("Click Me", func() {
		viewModel.SetMessage("Button Clicked")
	})

	w.SetContent(container.NewVBox(
		label,
		button,
	))

	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
