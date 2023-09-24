package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"

)

func buildChatLayout(a fyne.App, externalW fyne.Window) {
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
	buttonSend.Importance = widget.HighImportance

	buttonExit := widget.NewButton("Sair", exit(externalW, w))
	buttonExit.Importance = widget.DangerImportance

	top := canvas.NewText("", nil)
	top.Resize(fyne.Size{Width: float32(WINDOW_WIDTH), Height: 100})

	box := container.NewGridWithColumns(2, buttonSend, buttonExit)
	spacer :=  widget.NewSeparator()
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

// if click in send message, capture the input text and send message to entry
func sendMessage(inputText *widget.Entry, chatText *widget.Entry) func() {
	return func() {
		message := inputText.Text
		chatText.SetText("Eu: " + message) 
		inputText.SetText("")
	}
}

func exit(external fyne.Window, w fyne.Window) func(){
	return func(){
		w.Close()
	}
}
