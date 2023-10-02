package view

import (
	"database/sql"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/PyMarcus/rpc_chat/connection"
)

func createTableAndGetConn() *sql.DB{
	conn, _ := connection.ConnectionSQL()
	return conn
}

func buildRoomsList(window fyne.Window, a fyne.App) *widget.Button {
	button := widget.NewButton("Liberar", func() {
		createTableAndGetConn()
		RoomEnabled = true
		buildChatLayout(a, window)
	})

	return button
}
