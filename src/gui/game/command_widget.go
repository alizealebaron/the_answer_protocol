/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* command_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/01 09:04:16 by rruiz           #+#    #+#              */
/* Updated: 2026/09/07 16:26:05 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"image/color"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
				executeCommand(stdin, command)
			})
			commandButton.Importance = widget.LowImportance
			subCommandsBox.Add(commandButton)
		}
		subCommandsBox.Refresh()
	})

	return button
}

func executeCommand(stdin io.WriteCloser, command string) {
	commandInfo := allCommand[command]
}

func startCommand(command string, subCommandsBox *fyne.Container, stdin io.WriteCloser) {
	commandInfo := allCommand[command]

	if commandInfo.parameters.name == "" {
		executeCommand(stdin, command)
		return
	}

}

func renderParams(parameter *parameter, command string, box *fyne.Container, stdin io.WriteCloser) {
	box.RemoveAll()

	if parameter.form == list {
		renderList()
	} else if parameter.form == field {
		renderField()
	}

	box.Refresh()
}

func renderList() {}

func renderField() {}
