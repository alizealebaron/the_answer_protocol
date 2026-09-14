/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* command_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/01 09:04:16 by rruiz           #+#    #+#              */
/* Updated: 2026/09/14 16:23:04 by rruiz           ###   ########.fr       */
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

var list = "list"
var field = "field"

type parameter struct {
	name        string
	form        string
	next        *parameter
	placeholder string
}

type commandModel struct {
	name       string
	parameters parameter
}

var allCommand = map[string]commandModel{
	"LOOK": {name: "LOOK"},
	"MOVE": {name: "MOVE", parameters: parameter{name: "direction", form: list, placeholder: "Select a direction."}},
	"WHO":  {name: "WHO"},

	"TALK":  {name: "TALK", parameters: parameter{name: "npc", form: list, placeholder: "Choose an NPC to talk to."}},
	"CHAT":  {name: "CHAT", parameters: parameter{name: "scope", form: list, placeholder: "Choose the scope of the message.", next: &parameter{name: "text", form: field, placeholder: "Enter your message"}}},
	"GROUP": {name: "GROUP", parameters: parameter{name: "action", form: list, placeholder: "Choose a group action."}},

	// "ATTACK":    {name: "ATTACK"},
	"STATUS": {name: "STATUS"},

	// "QUEST":     {name: "QUEST"},
	// "QUESTS":    {name: "QUESTS"},

	"INVENTORY": {name: "INVENTORY"},
	"USE":       {name: "USE", parameters: parameter{name: "item", form: list, placeholder: "Choose an object to use."}},
	"TRADE":     {name: "TRADE", parameters: parameter{name: "trade things", form: list, placeholder: "Choose a merchant."}},
	"BUY":       {name: "BUY", parameters: parameter{name: "trader", form: list, placeholder: "Choose a merchant.", next: &parameter{name: "item", form: list, placeholder: "Choose an item to buy.", next: &parameter{name: "quantity", form: field, placeholder: "Enter the quantity."}}}},
	"SELL":      {name: "SELL", parameters: parameter{name: "trader", form: list, placeholder: "Choose a merchant.", next: &parameter{name: "item", form: list, placeholder: "Choose an item to sell.", next: &parameter{name: "quantity", form: field, placeholder: "Enter the quantity."}}}},
	"TAKE":      {name: "TAKE", parameters: parameter{name: "take item", form: list, placeholder: "Choose an item to pick-up"}},
	"DROP":      {name: "DROP", parameters: parameter{name: "drop item", form: list, placeholder: "Choose an item to drop"}},

	"GAMBLING": {name: "GAMBLING", parameters: parameter{name: "quantity", form: "field", placeholder: "Enter the amount (in gambling coins)."}},
}

var commandByCategory = map[string][]string{
	"Environment": {"LOOK", "MOVE", "SEARCH", "WHO"},
	"Social":      {"CHAT", "GROUP", "TALK"},
	"Fight":       {"ATTACK", "STATUS"},
	"Quest":       {"QUEST", "QUESTS"},
	"Inventory":   {"BUY", "DROP", "INVENTORY", "SELL", "TAKE", "TRADE", "USE"},
	"Gambling":    {"GAMBLING"},
}

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

func createButton(category string, subCommandBox *fyne.Container, stdin io.WriteCloser, listener *types.Listener, playerName string, subScroll *container.Scroll) *widget.Button {
	button := widget.NewButton(category, func() {
		showCommandCategory(category, subCommandBox, stdin, listener, playerName, subScroll)
	})
	return button
}

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

	case "GROUP":
		commands.Group(stdin, listener, subCommandBox, playerName, func() {
			showCommandCategory("Social", subCommandBox, stdin, listener, playerName, subScroll)
		})

		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
		// |                                                            Fight                                                                |
		// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

	case "ATTACK":

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
