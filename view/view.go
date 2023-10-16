package view

import (
	"fmt"
	"log"
	"strings"
	"time"
	"github.com/PyMarcus/rpc_chat/models"
	"github.com/PyMarcus/rpc_chat/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var RoomEnabled bool = false

type User string

var UserName User

func Start() {
	log.Println("Starting app...")
	buildScreen()
}

// create the screen and build the settings from window
func buildScreen() {
	buildTheFirstWindow()
}

// the first window to appear
func buildTheFirstWindow() {
	a := app.New()
	log.Println("Starting first app")
	window := a.NewWindow("RPC Chat")

	text := widget.NewLabel("RPC CHAT")
	text.Importance = widget.DangerImportance
	text.TextStyle.Bold = true
	text.Alignment = fyne.TextAlignCenter

	textContent := widget.NewEntry()
	f := widget.NewForm(widget.NewFormItem("Entrar como: ", textContent))

	button := widget.NewButton("Entrar", func() {
		func() {
			UserName = User(strings.ToUpper(textContent.Text))
			window.Hide()
			chatWindow(a)
			return
		}()
	})

	button.Resize(fyne.Size{Width: 100, Height: 50})
	button.Importance = widget.SuccessImportance
	content := container.NewVBox(text, f, button)

	window.SetContent(content)
	window.Resize(fyne.Size{Width: float32(500), Height: float32(100)})
	window.CenterOnScreen()
	window.SetFixedSize(FIXED_SIZE)
	window.ShowAndRun()
}

func chatWindow(a fyne.App) {
	// create window
	log.Println("Start second app...")
	window := a.NewWindow(fmt.Sprintf("RPC CHAT - %v", UserName))
	window.SetMaster()
	// items
	tabs := buildTopMenuItems(window, a)
	window.SetContent(tabs)
	// settings
	window.Resize(fyne.Size{Width: float32(WINDOW_WIDTH), Height: float32(WINDOW_HEIGTH)})
	window.CenterOnScreen()
	window.SetFixedSize(FIXED_SIZE)
	window.Show()
}

// add tab on top from screen
func buildTopMenuItems(window fyne.Window, a fyne.App) *container.AppTabs {
	rooms := buildRoomsList(window, a, string(UserName))
	pvChat := buildPvChat(window, a)
	chat := buildChatPV(window, a)

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Salas", theme.HomeIcon(), rooms),
		container.NewTabItemWithIcon("Chat privado", theme.MailSendIcon(), pvChat),
		container.NewTabItemWithIcon("Mensagens recebidas", theme.MailReplyIcon(), chat),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	return tabs
}

func buildPvChat(window fyne.Window, a fyne.App) *fyne.Container {
	conn := getConn()
	data := repository.NewRepository(conn)

	text := widget.NewLabel("Enviar mensagem")
	text.Importance = widget.WarningImportance
	text.TextStyle.Bold = true
	text.Alignment = fyne.TextAlignCenter

	textContent := widget.NewEntry()
	f := widget.NewForm(widget.NewFormItem("Mensagem para: ", textContent))
	textContentMessage := widget.NewEntry()
	f2 := widget.NewForm(widget.NewFormItem("Mensagem: ", textContentMessage))

	button := widget.NewButton("Enviar", func() {
		pv := &models.Pv{Name: textContent.Text, Message: textContentMessage.Text, From: string(UserName)}
		data.InsertPvMessage(pv)
	})

	return container.NewVBox(f, f2, button)
}

func buildChatPV(window fyne.Window, a fyne.App) *fyne.Container {
	var txt string

	conn := getConn()
	data := repository.NewRepository(conn)

	text := widget.NewLabel("Mensagens recebidas")
	text.Importance = widget.WarningImportance
	text.TextStyle.Bold = true
	text.Alignment = fyne.TextAlignCenter

	result, err := data.GetPvMessage(string(UserName))

	if err != nil{
		log.Println("Fail to read pv message ", err)
		result = nil
	}

	chatText := widget.NewMultiLineEntry()
	chatText.SetMinRowsVisible(MAX_LINES_CHAT)

	savedMessages = make(map[string]bool)
	go func() {
		for range time.Tick(5 * time.Second){
			for _, m := range result {
				messageID := fmt.Sprintf("%s: %s", m.From, m.Message)
				if _, found := savedMessages[messageID]; !found {
					item := fmt.Sprintf("%s: %s\n\n", m.From, m.Message)
					txt += item
					savedMessages[messageID] = true
				}
			}
			item := chatText.Text + "\n"
			if !strings.Contains(txt, item) {
				txt += item
			}
			chatText.SetText(txt)
			txt = ""
		}
	}()

	btn := widget.NewButton("Atualizar", func() {
		// sim fiz isso às pressas :D
		go func() {
			result, _ := data.GetPvMessage(string(UserName))

			for range time.Tick(5 * time.Second){
				for _, m := range result {
					messageID := fmt.Sprintf("%s: %s", m.From, m.Message)
					if _, found := savedMessages[messageID]; !found {
						item := fmt.Sprintf("%s: %s\n\n", m.From, m.Message)
						txt += item
						savedMessages[messageID] = true
					}
				}
				item := chatText.Text + "\n"
				if !strings.Contains(txt, item) {
					txt += item
				}
				chatText.SetText(txt)
				txt = ""
			}
		}()
	})
	
	return container.NewVBox(chatText, btn)
}
