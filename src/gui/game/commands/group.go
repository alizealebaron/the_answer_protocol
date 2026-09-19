/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* group.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/11 18:00:00 by rruiz           #+#    #+#              */
/* Updated: 2026/09/18 12:00:00 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"the_answer_protocol/src/gui/game/types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type groupAction struct {
	name string
	run  func()
}

// Start of GROUP. Displays the group actions (CREATE, INVITE, JOIN, LEAVE) as a button menu.
func Group(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, playerName string, back func()) {
	lookupPrefixes := []string{"OK SECRET"}
	types.Mute(lookupPrefixes...)
	wrappedBack := func() {
		types.Unmute(lookupPrefixes...)
		back()
	}

	subCommandBox.RemoveAll()

	actions := []groupAction{
		{"CREATE", func() {
			_, _ = fmt.Fprintf(stdin, "GROUP CREATE\n")
			_, _ = fmt.Printf("GROUP CREATE\n")
			refreshGroupAfterReply(stdin, listener, "OK group=")
			wrappedBack()
		}},
		{"INVITE", func() {
			showPlayersToInvite(stdin, listener, subCommandBox, playerName, wrappedBack)
		}},
		{"JOIN", func() {
			showJoinForm(stdin, listener, subCommandBox, wrappedBack)
		}},
		{"LEAVE", func() {
			_, _ = fmt.Fprintf(stdin, "GROUP LEAVE\n")
			_, _ = fmt.Printf("GROUP LEAVE\n")
			refreshGroupAfterReply(stdin, listener, "OK")
			wrappedBack()
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

// Requests the SECRET data to list the players available for an invitation.
func showPlayersToInvite(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, playerName string, back func()) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK SECRET") {
			return
		}
		raw := strings.TrimPrefix(line, "OK SECRET ")
		var data types.Secret
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			listener.Unsubscribe(id)
			back()
			return
		}

		listener.Unsubscribe(id)

		showInvitablePlayers(stdin, listener, subCommandBox, data, playerName, back)
	})
	_, _ = fmt.Fprintf(stdin, "SECRET\n")

}

// Displays the other players as buttons to invite them to the group, or returns if there is no one.
func showInvitablePlayers(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, who types.Secret, playerName string, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, player := range who.Players {

		otherPlayerName := player.Name
		if otherPlayerName == playerName {
			continue
		}
		playerButton := widget.NewButton(otherPlayerName, func() {
			_, _ = fmt.Fprintf(stdin, "GROUP INVITE %s\n", otherPlayerName)
			_, _ = fmt.Printf("GROUP INVITE %s\n", otherPlayerName)
			refreshGroupAfterReply(stdin, listener, "OK")
			back()
		})
		playerButton.Importance = widget.LowImportance
		subCommandBox.Add(playerButton)
		len += 1
	}
	if len == 0 {
		listener.Distribute(types.Translate("No one to invite to your group."))
		back()
	}

	subCommandBox.Refresh()
}

// Asks for a group id and sends the GROUP JOIN command with it.
func showJoinForm(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	subCommandBox.RemoveAll()

	groupEntry := widget.NewEntry()
	groupEntry.SetPlaceHolder(types.Translate("Enter the group id."))

	joinButton := widget.NewButton(types.Translate("JOIN"), func() {
		groupID := strings.TrimSpace(groupEntry.Text)
		if groupID == "" {
			return
		}
		_, _ = fmt.Fprintf(stdin, "GROUP JOIN %s\n", groupID)
		_, _ = fmt.Printf("GROUP JOIN %s\n", groupID)
		refreshGroupAfterReply(stdin, listener, "OK group=")
		back()
	})

	subCommandBox.Add(groupEntry)
	subCommandBox.Add(joinButton)

	subCommandBox.Refresh()
}

// Sends SECRET once the server replies to a group action, so the widget get fresh game data.
func refreshGroupAfterReply(stdin io.WriteCloser, listener *types.Listener, prefixes ...string) {
	var id int
	id = listener.Subscribe(func(line string) {
		if strings.HasPrefix(line, "ERR ") {
			listener.Unsubscribe(id)
			return
		}
		for _, prefix := range prefixes {
			if strings.HasPrefix(line, prefix) {
				listener.Unsubscribe(id)
				_, _ = fmt.Fprintf(stdin, "SECRET\n")
				return
			}
		}
	})
}
