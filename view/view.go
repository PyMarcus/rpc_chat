package view

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var wg sync.WaitGroup

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

func chatWindow(a fyne.App){
	// create window
	log.Println("Start second app...")
	window := a.NewWindow(fmt.Sprintf("RPC CHAT - %v", UserName))
	window.SetMaster()
	// items
	tabs := buildTopMenuItems(window)
	window.SetContent(container.NewVBox(tabs))
	// settings
	window.Resize(fyne.Size{Width: float32(WINDOW_WIDTH), Height: float32(WINDOW_HEIGTH)})
	window.CenterOnScreen()
	window.SetFixedSize(FIXED_SIZE)
	window.Show()
}

// add tab on top from screen
func buildTopMenuItems(window fyne.Window) *container.AppTabs {
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Salas", theme.HomeIcon(), canvas.NewText("salas aq", nil)),
		container.NewTabItemWithIcon("Chat", theme.InfoIcon(), buildChatLayout()),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	return tabs
}
