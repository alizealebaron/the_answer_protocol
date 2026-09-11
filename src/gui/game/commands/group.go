/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* group.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/11 18:00:00 by rruiz           #+#    #+#              */
/* Updated: 2026/09/11 22:25:52 by rruiz           ###   ########.fr       */
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

func Group(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	subCommandBox.RemoveAll()

	actions := []groupAction{
		{"CREATE", func() {
			fmt.Fprintf(stdin, "GROUP CREATE\n")
			fmt.Printf("GROUP CREATE\n")
			back()
		}},
		{"INVITE", func() {
			showPlayersToInvite(stdin, listener, subCommandBox, back)
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

func showPlayersToInvite(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK { \"room\":") {
			return
		}
		whoJson := strings.TrimPrefix(line, "OK ")
		var data types.WhoInfo
		if err := json.Unmarshal([]byte(whoJson), &data); err != nil {
			listener.Unsubscribe(id)
			return
		}
		listener.Unsubscribe(id)
		showInvitablePlayers(stdin, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "WHO\n")
}

func showInvitablePlayers(stdin io.WriteCloser, subCommandBox *fyne.Container, who types.WhoInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, player := range who.RoomInfo {
		playerName := player
		playerButton := widget.NewButton(playerName, func() {
			fmt.Fprintf(stdin, "GROUP INVITE %s\n", playerName)
			fmt.Printf("GROUP INVITE %s\n", playerName)
			back()
		})
		playerButton.Importance = widget.LowImportance
		subCommandBox.Add(playerButton)
		len += 1
	}
	if len == 0 {
		back()
	}

	subCommandBox.Refresh()
}

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
