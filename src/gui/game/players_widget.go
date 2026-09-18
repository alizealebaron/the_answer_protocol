/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* players_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/07 11:17:14 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 22:01:57 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"encoding/json"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"the_answer_protocol/src/gui/game/types"
)

// Shows the number of players in the room and on the server.
func playerCountLabel(listener *types.Listener) *fyne.Container {
	var lenRoom int   // Players in the current room
	var lenServer int // Players on the server
	var firstWho bool // True once a WHO reply has been received

	label := widget.NewLabel(types.Translate("Use the \u2018WHO\u2019 command to view information about players number."))
	listener.Subscribe(func(line string) {
		// The WHO reply gives the starting room and server counts
		if strings.HasPrefix(line, "OK { \"room\":") {
			whoJson := strings.TrimPrefix(line, "OK ")
			var data types.WhoInfo
			err := json.Unmarshal([]byte(whoJson), &data)
			if err != nil {
				return
			}
			lenRoom = len(data.RoomInfo)
			lenServer = data.ServerInfo
			fmt.Println(line)

			if !firstWho {
				firstWho = true
			}
		}

		if firstWho {
			// A presence event changes the room count without asking the server
			if strings.HasPrefix(line, "EVT ROOM PRESENCE ENTER") {
				lenRoom += 1
			} else if strings.HasPrefix(line, "EVT ROOM PRESENCE LEAVE") {
				lenRoom -= 1
			}

			label.SetText(fmt.Sprintf(types.Translate("Players in room: %d\nPlayers on server: %d"), lenRoom, lenServer))
		}

	})
	label.Wrapping = fyne.TextWrapWord
	label.Alignment = fyne.TextAlignCenter

	return container.NewCenter(label)
}
