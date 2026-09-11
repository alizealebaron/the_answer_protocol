/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* command_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/01 09:04:16 by rruiz           #+#    #+#              */
/* Updated: 2026/09/11 17:11:35 by rruiz           ###   ########.fr       */
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
	"Environment": {"LOOK", "MOVE", "WHO"},
	"Social":      {"TALK", "CHAT", "GROUP"},
	"Fight":       {"ATTACK", "STATUS"},
	"Quest":       {"QUEST", "QUESTS"},
	"Inventory":   {"INVENTORY", "USE", "TRADE", "BUY", "SELL", "TAKE", "DROP"},
	"Gambling":    {"GAMBLING"},
}

func commandWidget(stdin io.WriteCloser, listener *types.Listener) *fyne.Container {
	commandFrame := canvas.NewRectangle(color.Transparent)
	commandFrame.StrokeColor = color.White
	commandFrame.StrokeWidth = float32(2)

	subCommandBox := container.NewVBox()
	subScroll := container.NewScroll(subCommandBox)

	buttonBar := container.NewGridWithColumns(6,
		createButton("Environment", subCommandBox, stdin, listener),
		createButton("Social", subCommandBox, stdin, listener),
		createButton("Fight", subCommandBox, stdin, listener),
		createButton("Quest", subCommandBox, stdin, listener),
		createButton("Inventory", subCommandBox, stdin, listener),
		createButton("Gambling", subCommandBox, stdin, listener),
	)

	layout := container.NewBorder(buttonBar, nil, nil, nil, subScroll)
	return container.NewStack(commandFrame, container.NewPadded(layout))
}

func createButton(category string, subCommandBox *fyne.Container, stdin io.WriteCloser, listener *types.Listener) *widget.Button {
	button := widget.NewButton(category, func() {
		showCommandCategory(category, subCommandBox, stdin, listener)
	})
	return button
}

func showCommandCategory(category string, subCommandBox *fyne.Container, stdin io.WriteCloser, listener *types.Listener) {
	subCommandBox.RemoveAll()

	for _, command := range commandByCategory[category] {
		commandButton := widget.NewButton(command, func() {
			executeCommand(stdin, command, listener, subCommandBox)
		})
		commandButton.Importance = widget.LowImportance
		subCommandBox.Add(commandButton)
	}
	subCommandBox.Refresh()
}

func executeCommand(stdin io.WriteCloser, command string, listener *types.Listener, subCommandBox *fyne.Container) {
	switch command {
	case "LOOK":
		fmt.Fprintf(stdin, "LOOK\n")
		fmt.Println("LOOK")
	case "MOVE":
		commands.Move(stdin, listener, subCommandBox, func() {
			showCommandCategory("Environment", subCommandBox, stdin, listener)
		})
	case "WHO":
		fmt.Fprintf(stdin, "WHO\n")
		fmt.Println("WHO")
	case "TALK":
	case "CHAT":
	case "GROUP":
	case "STATUS":
		fmt.Fprintf(stdin, "STATUS\n")
		fmt.Println("STATUS")
	case "INVENTORY":
		fmt.Fprintf(stdin, "STATUS\n")
		fmt.Println("STATUS")
	case "USE":
	case "TRADE":
	case "BUY":
	case "SELL":
	case "TAKE":
		commands.Take(stdin, listener, subCommandBox, func() {
			showCommandCategory("Inventory", subCommandBox, stdin, listener)
		})
	case "DROP":
	case "GAMBLING":
	}
}
