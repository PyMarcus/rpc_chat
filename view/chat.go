package view

import (
	"log"
	"strings"
	"sync"
	"time"

	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/PyMarcus/rpc_chat/connection"
	"github.com/PyMarcus/rpc_chat/models"
	"github.com/PyMarcus/rpc_chat/repository"
)

var wg sync.WaitGroup
var mutex sync.Mutex

type ReplyListMessages struct {
	Data []*models.Message
}

var savedMessages map[string]bool


func buildChatLayout(a fyne.App, externalW fyne.Window) {
	var messages ReplyListMessages
	var text string

	client := connection.ServerRPCConnection("localhost", "12345")

	log.Println(text)
	w := a.NewWindow(string(UserName))
	chatText := widget.NewMultiLineEntry()
	chatText.SetMinRowsVisible(MAX_LINES_CHAT)

	space := widget.NewSeparator()
	space.Resize(fyne.NewSize(float32(WINDOW_WIDTH), float32(300)))

	inputText := widget.NewMultiLineEntry()
	inputText.SetMinRowsVisible(MAX_LINES_INPUT)
	inputText.PlaceHolder = "Escreva sua mensagem :)"
	inputText.Resize(fyne.Size{Width: float32(WINDOW_WIDTH) - 200, Height: 30})

	buttonSend := widget.NewButton("Enviar", sendMessage(inputText, chatText))
	savedMessages = make(map[string]bool)
	go func() {
		for range time.Tick(5 * time.Second){
			log.Println("Calling...")
			client.Call("MyService.ListMessage", CurrentIdRoom, &messages)
			for _, m := range messages.Data {
				messageID := fmt.Sprintf("%s: %s", m.User, m.Message)
				if _, found := savedMessages[messageID]; !found {
					item := fmt.Sprintf("%s: %s\n\n", m.User, m.Message)
					text += item
					savedMessages[messageID] = true
				}
			}
			item := chatText.Text + "\n"
			if !strings.Contains(text, item) {
				text += item
			}
			chatText.SetText(text)
			text = ""
		}
	}()

	buttonSend.Importance = widget.HighImportance

	buttonExit := widget.NewButton("Sair", exit(externalW, w, string(UserName)))
	buttonExit.Importance = widget.DangerImportance

	top := canvas.NewText("", nil)
	top.Resize(fyne.Size{Width: float32(WINDOW_WIDTH), Height: 100})

	box := container.NewGridWithColumns(2, buttonSend, buttonExit)
	spacer := widget.NewSeparator()
	spacer.Resize(fyne.Size{Width: float32(WINDOW_WIDTH), Height: float32(100)})
	padding := container.NewPadded(container.NewVBox(chatText, space, inputText, box))

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Chat", theme.InfoIcon(), padding),
	)
	tabs.SetTabLocation(container.TabLocationTop)
	w.SetContent(container.NewVBox(tabs))
	w.Resize(fyne.Size{Width: float32(WINDOW_WIDTH), Height: float32(WINDOW_HEIGTH)})
	w.CenterOnScreen()
	w.SetFixedSize(FIXED_SIZE)
	w.Show()
}

type ReplyListUsers struct {
	Data []*models.UsersModel
}

type MessageArgs struct{
	User string 
	Message string 
	RoomId int64
}

type ReplyOk struct{}

// if click in send message, capture the input text and send message to entry
func sendMessage(inputText *widget.Entry, chatText *widget.Entry) func() {
	client := connection.ServerRPCConnection("localhost", "12345")
	return func() {
		message := inputText.Text
		mutex.Lock()
		if message == "/lu" {
			var users ReplyListUsers
			var text string

			client.Call("MyService.ListUsers", CurrentIdRoom, &users)
			log.Println("USRS ", users)
			text += string(UserName) + " chamou método /lu (listar usuarios da sala): \n"
			for i, user := range users.Data {
				text += fmt.Sprintf("%d", i) + ": " + user.Name + "\n"
			}
			chatText.SetText(text)
			inputText.SetText("")
		} else {
			var reply ReplyOk
			m := MessageArgs{ User: string(UserName), Message: message, RoomId: CurrentIdRoom}
			log.Println("A enviar", m)
			client.Call("MyService.SendMessage", m, &reply)
			inputText.SetText("")
		}
		mutex.Unlock()
	}
}

func exit(external fyne.Window, w fyne.Window, user string) func() {
	conn := getConn()
	data := repository.NewRepository(conn)
	return func() {
		data.InsertUser(models.Message{Id: 1, User: user, Message: "Saiu da sala", RoomId: CurrentIdRoom})
		w.Close()
	}
}
