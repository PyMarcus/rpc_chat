package view

import (
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/PyMarcus/rpc_chat/connection"
	"github.com/PyMarcus/rpc_chat/models"
	"github.com/PyMarcus/rpc_chat/repository"
)

var CurrentIdRoom int64

func getConn() *sql.DB {
	conn, _ := connection.ConnectionSQL()
	return conn
}

func buildRoomsList(window fyne.Window, a fyne.App, user string) *widget.Table {
	conn := getConn()
	data := repository.NewRepository(conn)

	tableContent, err := data.GetAllRooms()

	if err != nil {
		panic(err)
	}

	table := widget.NewTable(
		func() (int, int) {
			return len(tableContent), 3
		},
		func() fyne.CanvasObject {
			return container.NewVBox(widget.NewLabel(""))
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			room := tableContent[i.Row]

			if i.Col == 2 {
				b := widget.NewButtonWithIcon("Entrar", theme.LoginIcon(), func() {
					RoomEnabled = true
					buildChatLayout(a, window)
					id := int64(i.Row + 1)
					CurrentIdRoom = id
					data.InsertUser(models.Message{Id:1, User: user, Message: "Entrou na sala", RoomId: id})
				})

				b.Importance = widget.SuccessImportance

				o.(*fyne.Container).Objects = []fyne.CanvasObject{b}

			} else {
				if(i.Col == 0){
					o.(*fyne.Container).Objects = []fyne.CanvasObject{
						widget.NewLabel(fmt.Sprintf("%d", room.Id)),
					}
				}else if(i.Col == 1){
					o.(*fyne.Container).Objects = []fyne.CanvasObject{
						widget.NewLabel(room.Name),
					}
				}
		}
		})
	// style
	for i := 0; i < 3; i++ {
		table.SetColumnWidth(i, 250)
	}

	return table

}
