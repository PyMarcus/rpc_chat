package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func buildRoomsList(window fyne.Window, a fyne.App) *widget.Button{
	button := widget.NewButton("Liberar", func() {
		RoomEnabled = true
		buildChatLayout(a, window)
	})

	return button
}