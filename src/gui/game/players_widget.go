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
	"io"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"the_answer_protocol/src/gui/game/types"
)

// Shows the number of players in the room and on the server.
func playerCountLabel(listener *types.Listener, stdin io.WriteCloser, playerName string) *fyne.Container {
	var lenRoom int   // Players in the current room
	var lenServer int // Players on the server
	var firstWho bool // True once a WHO reply has been received

	label := widget.NewLabel(fmt.Sprintf(types.Translate("Players in room: %d\nPlayers on server: %d"), lenRoom, lenServer))
	lookupPrefixes := []string{"OK { \"room\":"}
	types.Mute(lookupPrefixes...)

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

			if !firstWho {
				firstWho = true
			}
		}

		if firstWho {
			if strings.HasPrefix(line, "EVT ROOM PRESENCE ENTER ") {
				if strings.TrimPrefix(line, "EVT ROOM PRESENCE ENTER ") == playerName {
					fmt.Fprintf(stdin, "WHO\n")
					return
				}
				lenRoom += 1
			} else if strings.HasPrefix(line, "EVT ROOM PRESENCE LEAVE ") {
				if strings.TrimPrefix(line, "EVT ROOM PRESENCE LEAVE ") == playerName {
					fmt.Fprintf(stdin, "WHO\n")
					return
				}
				lenRoom -= 1
			} else if strings.HasPrefix(line, "EVT STATS players=") {
				lenServer, _ = strconv.Atoi(strings.TrimPrefix(line, "EVT STATS players="))
			}

			label.SetText(fmt.Sprintf(types.Translate("Players in room: %d\nPlayers on server: %d"), lenRoom, lenServer))
		}

	})
	label.Wrapping = fyne.TextWrapWord
	label.Alignment = fyne.TextAlignCenter

	return container.NewCenter(label)
}
