/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* chat.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/14 17:07:41 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:42:18 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package commands

import (
	"fmt"
	"io"
	"the_answer_protocol/src/gui/game/types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Start of CHAT. Displays the chat scopes (GLOBAL, ROOM, GROUP) as a button menu.
func Chat(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	subCommandBox.RemoveAll()

	actions := []groupAction{
		{"GLOBAL", func() {
			sendMessage(stdin, subCommandBox, "GLOBAL", back)
		}},

		{"ROOM", func() {
			sendMessage(stdin, subCommandBox, "ROOM", back)
		}},

		{"GROUP", func() {
			sendMessage(stdin, subCommandBox, "GROUP", back)
		}},
	}
	for _, action := range actions {
		actionName := action.name
		actionRun := action.run
		actionButton := widget.NewButton(types.Translate(actionName), actionRun)
		actionButton.Importance = widget.LowImportance
		subCommandBox.Add(actionButton)
	}

	subCommandBox.Refresh()
}

// Asks for a message and sends the CHAT command in the given scope.
func sendMessage(stdin io.WriteCloser, subCommandBox *fyne.Container, scope string, back func()) {
	subCommandBox.RemoveAll()

	messageEntry := widget.NewEntry()
	messageEntry.SetPlaceHolder(types.Translate("Enter your message."))

	chatButton := widget.NewButton(types.Translate("Send message"), func() {
		if messageEntry.Text == "" {
			return
		}

		_, _ = fmt.Fprintf(stdin, "CHAT %s %s\n", scope, messageEntry.Text)
		_, _ = fmt.Printf("CHAT %s %s\n", scope, messageEntry.Text)
		back()
	})

	subCommandBox.Add(messageEntry)
	subCommandBox.Add(chatButton)

	subCommandBox.Refresh()
}
