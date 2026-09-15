/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* group.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/11 18:00:00 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 14:45:20 by rruiz           ###   ########.fr       */
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
	subCommandBox.RemoveAll()

	actions := []groupAction{
		{"CREATE", func() {
			fmt.Fprintf(stdin, "GROUP CREATE\n")
			fmt.Printf("GROUP CREATE\n")
			back()
		}},
		{"INVITE", func() {
			showPlayersToInvite(stdin, listener, subCommandBox, playerName, back)
			fmt.Println("1")
		}},
		{"JOIN", func() {
			showJoinForm(stdin, subCommandBox, back)
		}},
		{"LEAVE", func() {
			fmt.Fprintf(stdin, "GROUP LEAVE\n")
			fmt.Printf("GROUP LEAVE\n")
			back()
		}},
	}

	for _, action := range actions {
		actionName := action.name
		actionRun := action.run
		actionButton := widget.NewButton(actionName, actionRun)
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
		var data types.Secret
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			listener.Unsubscribe(id)
			fmt.Println("3")
			return
		}
		fmt.Println("4")

		listener.Unsubscribe(id)
		fmt.Println("5")

		showInvitablePlayers(stdin, subCommandBox, data, playerName, back)
	})
	fmt.Fprintf(stdin, "SECRET\n")
	fmt.Println("2")

}

// Displays the other players as buttons to invite them to the group, or returns if there is no one.
func showInvitablePlayers(stdin io.WriteCloser, subCommandBox *fyne.Container, who types.Secret, playerName string, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, player := range who.Players {
		fmt.Printf("%d", 6+len)

		otherPlayerName := player.Name
		if otherPlayerName == playerName {
			continue
		}
		playerButton := widget.NewButton(otherPlayerName, func() {
			fmt.Println("TEST")
			fmt.Fprintf(stdin, "GROUP INVITE %s\n", otherPlayerName)
			fmt.Printf("GROUP INVITE %s\n", otherPlayerName)
			back()
		})
		playerButton.Importance = widget.LowImportance
		subCommandBox.Add(playerButton)
		len += 1
	}
	if len == 0 {
		fmt.Println("CACA")
		back()
	}

	subCommandBox.Refresh()
}

// Asks for a group id and sends the GROUP JOIN command with it.
func showJoinForm(stdin io.WriteCloser, subCommandBox *fyne.Container, back func()) {
	subCommandBox.RemoveAll()

	groupEntry := widget.NewEntry()
	groupEntry.SetPlaceHolder("Enter the group id.")

	joinButton := widget.NewButton("JOIN", func() {
		groupID := strings.TrimSpace(groupEntry.Text)
		if groupID == "" {
			return
		}
		fmt.Fprintf(stdin, "GROUP JOIN %s\n", groupID)
		fmt.Printf("GROUP JOIN %s\n", groupID)
		back()
	})

	subCommandBox.Add(groupEntry)
	subCommandBox.Add(joinButton)

	subCommandBox.Refresh()
}
