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

	"fyne.io/fyne/v2"
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

func GameView(window fyne.Window, size fyne.Size) fyne.CanvasObject {
	width, height := size.Width, size.Height
	commandBox := commandWidget(fyne.NewSize(width*28.5/100, height*78.5/100))
	return commandBox
}

func commandWidget(size fyne.Size) *fyne.Container {
	barWidth, barHeight := size.Width/6, size.Height*5/100
	commandList := container.NewVBox()

	refreshCommands := func(category string) {
		commands := commandByCategory[category]

		newObjects := make([]fyne.CanvasObject, 0, len(commands))
		for _, cmd := range commands {
			cmdName := cmd // capture locale pour la closure du bouton
			newObjects = append(newObjects, widget.NewButton(cmdName, func() {
				fmt.Println("Commande exécutée :", cmdName)
			}))
		}

		commandList.Objects = newObjects
		commandList.Refresh()
	}

	firstButton := container.NewGridWrap(fyne.NewSize(barWidth, barHeight),
		widget.NewButton("Environment",
			func() {
				fmt.Println("Environment")
			}))

	secondButton := container.NewGridWrap(fyne.NewSize(barWidth, barHeight),
		widget.NewButton("Social",
			func() {
				fmt.Println("Social")
			}))

	thirdButton := container.NewGridWrap(fyne.NewSize(barWidth, barHeight),
		widget.NewButton("Fight",
			func() {
				fmt.Println("Fight")
			}))

	fourthButton := container.NewGridWrap(fyne.NewSize(barWidth, barHeight),
		widget.NewButton("Quest",
			func() {
				fmt.Println("Quest")
			}))

	fifthButton := container.NewGridWrap(fyne.NewSize(barWidth, barHeight),
		widget.NewButton("Inventory",
			func() {
				fmt.Println("Inventory")
			}))

	sixthButton := container.NewGridWrap(fyne.NewSize(barWidth, barHeight),
		widget.NewButton("Gambling",
			func() {
				fmt.Println("Gambling")
			}))

	buttonBar := container.NewHBox(firstButton, secondButton, thirdButton, fourthButton, fifthButton, sixthButton)
	return buttonBar
}
