/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* command_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/01 09:04:16 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:34:47 by rruiz           ###   ########.fr       */
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

	"the_answer_protocol/src/gui/game/commands"
	"the_answer_protocol/src/gui/game/types"
)

// All commands sorted by category.
var commandByCategory = map[string][]string{
	"Environment": {"LOOK", "MOVE", "SEARCH", "WHO"},
	"Social":      {"CHAT", "GROUP", "TALK"},
	"Fight":       {"ATTACK", "STATUS"},
	"Quest":       {"QUEST", "QUESTS"},
	"Inventory":   {"BUY", "DROP", "INVENTORY", "SELL", "TAKE", "TRADE", "USE"},
	"Gambling":    {"GAMBLING"},
}

// Creating the widget with the scrollable frame and the button bar.
func commandWidget(stdin io.WriteCloser, listener *types.Listener, playerName string) *fyne.Container {
	commandFrame := canvas.NewRectangle(color.Transparent)
	commandFrame.StrokeColor = color.White
	commandFrame.StrokeWidth = float32(2)

	subCommandBox := container.NewVBox()
	subScroll := container.NewScroll(subCommandBox)

	buttonBar := container.NewGridWithColumns(6,
		createButton("Environment", subCommandBox, stdin, listener, playerName, subScroll),
		createButton("Social", subCommandBox, stdin, listener, playerName, subScroll),
		createButton("Fight", subCommandBox, stdin, listener, playerName, subScroll),
		createButton("Quest", subCommandBox, stdin, listener, playerName, subScroll),
		createButton("Inventory", subCommandBox, stdin, listener, playerName, subScroll),
		createButton("Gambling", subCommandBox, stdin, listener, playerName, subScroll),
	)

	layout := container.NewBorder(buttonBar, nil, nil, nil, subScroll)
	return container.NewStack(commandFrame, container.NewPadded(layout))
}

// Creating category buttons with a command to generate game commands.
func createButton(category string, subCommandBox *fyne.Container, stdin io.WriteCloser, listener *types.Listener, playerName string, subScroll *container.Scroll) *widget.Button {
	button := widget.NewButton(category, func() {
		showCommandCategory(category, subCommandBox, stdin, listener, playerName, subScroll)
	})
	return button
}

// Create the category buttons when the category is clicked, along with code to define their functionality.
func showCommandCategory(category string, subCommandBox *fyne.Container, stdin io.WriteCloser, listener *types.Listener, playerName string, subScroll *container.Scroll) {
	subCommandBox.RemoveAll()

	for _, command := range commandByCategory[category] {
		commandButton := widget.NewButton(command, func() {
			executeCommand(stdin, command, listener, subCommandBox, playerName, subScroll)
		})
		commandButton.Importance = widget.LowImportance
		subCommandBox.Add(commandButton)
	}
	subScroll.ScrollToTop()
	subCommandBox.Refresh()
}

// A clunky switch-case statement to handle button functionality in a simple way.
func executeCommand(stdin io.WriteCloser, command string, listener *types.Listener, subCommandBox *fyne.Container, playerName string, subScroll *container.Scroll) {
	switch command {

	// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
	// |                                                       Environment                                                               |
	// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

	case "LOOK":
		fmt.Fprintf(stdin, "LOOK\n")
		fmt.Println("LOOK")

	case "MOVE":
		commands.Move(stdin, listener, subCommandBox, func() {
			showCommandCategory("Environment", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "WHO":
		fmt.Fprintf(stdin, "WHO\n")
		fmt.Println("WHO")

	case "SEARCH":
		commands.Search(stdin, listener, subCommandBox, func() {
			showCommandCategory("Environment", subCommandBox, stdin, listener, playerName, subScroll)
		})

		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		// |                                                            Social                                                               |
		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

	case "TALK":
		commands.Talk(stdin, listener, subCommandBox, func() {
			showCommandCategory("Social", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "CHAT":
		commands.Chat(stdin, listener, subCommandBox, func() {
			showCommandCategory("Social", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "GROUP":
		commands.Group(stdin, listener, subCommandBox, playerName, func() {
			showCommandCategory("Social", subCommandBox, stdin, listener, playerName, subScroll)
		})

		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		// |                                                            Fight                                                                |
		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

	case "ATTACK":
		commands.Attack(stdin, listener, subCommandBox, func() {
			showCommandCategory("Fight", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "STATUS":
		fmt.Fprintf(stdin, "STATUS\n")
		fmt.Println("STATUS")

		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		// |                                                            Quest                                                                |
		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

	case "QUEST":
		commands.Quest(stdin, listener, subCommandBox, func() {
			showCommandCategory("Quest", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "QUESTS":
		fmt.Fprintf(stdin, "QUESTS\n")
		fmt.Println("QUESTS")

		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		// |                                                         Inventory                                                               |
		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

	case "INVENTORY":
		fmt.Fprintf(stdin, "INVENTORY\n")
		fmt.Println("INVENTORY")

	case "USE":
		commands.Use(stdin, listener, subCommandBox, func() {
			showCommandCategory("Inventory", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "TRADE":
		commands.Trade(stdin, listener, subCommandBox, func() {
			showCommandCategory("Inventory", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "BUY":
		commands.Buy(stdin, listener, subCommandBox, func() {
			showCommandCategory("Inventory", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "SELL":
		commands.Sell(stdin, listener, subCommandBox, func() {
			showCommandCategory("Inventory", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "TAKE":
		commands.Take(stdin, listener, subCommandBox, func() {
			showCommandCategory("Inventory", subCommandBox, stdin, listener, playerName, subScroll)
		})

	case "DROP":
		commands.Drop(stdin, listener, subCommandBox, func() {
			showCommandCategory("Inventory", subCommandBox, stdin, listener, playerName, subScroll)
		})

		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		// |                                                          Gambling                                                               |
		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

	case "GAMBLING":
		commands.Gambling(stdin, listener, subCommandBox, func() {
			showCommandCategory("Gambling", subCommandBox, stdin, listener, playerName, subScroll)
		})
	}
}
