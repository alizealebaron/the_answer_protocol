/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* game_view.go                                      :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 18:10:21 by rruiz           #+#    #+#              */
/* Updated: 2026/08/22 13:31:19 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var commandByCategory = map[string][]string{
	"Environment": {"LOOK", "MOVE", "WHO"},
	"Social":      {"TALK", "CHAT", "GROUP"},
	"Fight":       {"ATACK", "STATUS"},
	"Quest":       {"QUEST", "QUESTS"},
	"Inventory":   {"INVENTORY", "USE", "TRADE", "BUY", "SELL", "TAKE", "DROP"},
	"Gambling":    {"GAMBLING"},
}

// func GameView(window fyne.Window, size fyne.Size) fyne.CanvasObject {
// 	width, height := size.Width, size.Height
// 	commandBox := commandWidget(fyne.NewSize(width*30/100, height*72/100))
// 	scrollBox := actionWidget(fyne.NewSize(width*30/100, height*27/100))

// 	mapBox := mapWidget(fyne.NewSize(width*40/100, height))

// 	leftWidth := width * 30 / 100

// 	enemyBox := enemyWidget(fyne.NewSize(leftWidth, height*50/100))
// 	playersLabel := playerCountLabel(fyne.NewSize(leftWidth, height*10/100))
// 	groupBox := groupWidget(fyne.NewSize(leftWidth, height*32/100))

// 	right := container.NewVBox(commandBox, scrollBox)
// 	left := container.NewVBox(enemyBox, playersLabel, groupBox)
// 	return container.NewBorder(nil, nil, left, right, mapBox)
// }

func GameView(window fyne.Window, size fyne.Size) fyne.CanvasObject {
	width, height := size.Width, size.Height
	commandBox := commandWidget(fyne.NewSize(width*30/100, height*72/100))
	scrollBox := actionWidget(fyne.NewSize(width*30/100, height))

	mapBox := mapWidget(fyne.NewSize(width*40/100, height))

	leftWidth := width * 30 / 100
	enemyBox := enemyWidget(fyne.NewSize(leftWidth, height*50/100))
	playersLabel := playerCountLabel(fyne.NewSize(leftWidth, height*10/100))
	groupBox := groupWidget(fyne.NewSize(leftWidth, height))

	right := container.NewBorder(commandBox, nil, nil, nil, scrollBox)
	left := container.NewBorder(container.NewVBox(enemyBox, playersLabel), nil, nil, nil, groupBox)
	return container.NewBorder(nil, nil, left, right, mapBox)
}

func commandWidget(size fyne.Size) *fyne.Container {
	width, height := size.Width, size.Height

	commandFrame := canvas.NewRectangle(color.Transparent)
	commandFrame.StrokeColor = color.White
	commandFrame.StrokeWidth = float32(2)
	commandFrame.SetMinSize(size)

	buttonBar := container.NewHBox(
		createButton(width, height, "Environment"),
		createButton(width, height, "Social"),
		createButton(width, height, "Fight"),
		createButton(width, height, "Quest"),
		createButton(width, height, "Inventory"),
		createButton(width, height, "Gambling"),
	)
	return container.NewStack(commandFrame, container.NewPadded(buttonBar))
}

func createButton(width float32, height float32, category string) *fyne.Container {
	buttonWidth, buttonHeight := (width-4)/6, height*5/100
	button := container.NewGridWrap(fyne.NewSize(buttonWidth, buttonHeight),
		widget.NewButton(category,
			func() {
				fmt.Println(category)
			}))
	return button
}

func actionWidget(size fyne.Size) *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)
	frame.SetMinSize(size)

	content := container.NewVBox()
	scroll := container.NewScroll(content)
	scroll.SetMinSize(size)

	return container.NewStack(frame, scroll)
}

func mapWidget(size fyne.Size) *fyne.Container {

	mapFrame := canvas.NewRectangle(color.Transparent)
	mapFrame.StrokeColor = color.White
	mapFrame.StrokeWidth = float32(2)
	mapFrame.SetMinSize(size)

	return container.NewStack(mapFrame)
}

func enemyWidget(size fyne.Size) *fyne.Container {
	width, height := size.Width, size.Height

	label := widget.NewLabel("Un simple test pour voir si c'est bien aligné, bien centré et si c'est assez gros.")
	label.Wrapping = fyne.TextWrapWord
	label.Alignment = fyne.TextAlignCenter

	label.Resize(fyne.NewSize(width, 0))
	wrappedHeight := label.MinSize().Height

	tightBox := container.NewGridWrap(fyne.NewSize(width, wrappedHeight), label)

	labelHeight := height * 5 / 100
	labelZone := container.NewGridWrap(fyne.NewSize(width, labelHeight), container.NewCenter(tightBox))

	side := width
	square := canvas.NewRectangle(color.Transparent)
	square.StrokeColor = color.White
	square.StrokeWidth = float32(2)
	square.SetMinSize(fyne.NewSize(side, side))

	return container.NewVBox(labelZone, container.NewStack(square))
}

func playerCountLabel(size fyne.Size) *fyne.Container {
	width, height := size.Width, size.Height

	label := widget.NewLabel(fmt.Sprintf("Nombre de joueur dans la room: %d\nNombre de joueur global: %d", 1, 42))
	label.Wrapping = fyne.TextWrapWord
	label.Alignment = fyne.TextAlignCenter

	label.Resize((fyne.NewSize(width, 0)))
	wrappedHeight := label.MinSize().Height

	tightBox := container.NewGridWrap(fyne.NewSize(width, wrappedHeight), label)
	labelZone := container.NewGridWrap(fyne.NewSize(width, height), container.NewCenter(tightBox))

	return labelZone
}

func groupWidget(size fyne.Size) *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)
	frame.SetMinSize(size)

	return container.NewStack(frame)
}
