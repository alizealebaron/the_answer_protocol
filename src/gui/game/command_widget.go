/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* command_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/01 09:04:16 by rruiz           #+#    #+#              */
/* Updated: 2026/09/01 14:58:41 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"fmt"
	"image/color"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var commandByCategory = map[string][]string{
	"Environment": {"LOOK", "MOVE", "WHO"},
	"Social":      {"TALK", "CHAT", "GROUP"},
	"Fight":       {"ATTACK", "STATUS"},
	"Quest":       {"QUEST", "QUESTS"},
	"Inventory":   {"INVENTORY", "USE", "TRADE", "BUY", "SELL", "TAKE", "DROP"},
	"Gambling":    {"GAMBLING"},
}

func commandWidget(stdin io.WriteCloser) *fyne.Container {
	commandFrame := canvas.NewRectangle(color.Transparent)
	commandFrame.StrokeColor = color.White
	commandFrame.StrokeWidth = float32(2)

	subCommandBox := container.NewVBox()
	subScroll := container.NewScroll(subCommandBox)

	buttonBar := container.NewGridWithColumns(6,
		createButton("Environment", subCommandBox, stdin),
		createButton("Social", subCommandBox, stdin),
		createButton("Fight", subCommandBox, stdin),
		createButton("Quest", subCommandBox, stdin),
		createButton("Inventory", subCommandBox, stdin),
		createButton("Gambling", subCommandBox, stdin),
	)

	layout := container.NewBorder(buttonBar, nil, nil, nil, subScroll)
	return container.NewStack(commandFrame, container.NewPadded(layout))
}

func createButton(category string, subCommandsBox *fyne.Container, stdin io.WriteCloser) *widget.Button {
	button := widget.NewButton(category, func() {
		subCommandsBox.RemoveAll()

		for _, command := range commandByCategory[category] {
			commandButton := widget.NewButton(command, func() {
				fmt.Fprintf(stdin, "LOOK")
				fmt.Println("LOOK")
				fmt.Println(command)
			})
			commandButton.Importance = widget.LowImportance
			subCommandsBox.Add(commandButton)
		}
		subCommandsBox.Refresh()
	})

	return button
}
