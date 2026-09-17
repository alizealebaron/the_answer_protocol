/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* players_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/07 11:17:14 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 09:50:06 by rruiz           ###   ########.fr       */
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

func playerCountLabel(listener *types.Listener) *fyne.Container {
	var lenRoom int
	var lenServer int
	var firstWho bool

	label := widget.NewLabel("Use the “WHO” command to view information about players number.")
	listener.Subscribe(func(line string) {
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
			if strings.HasPrefix(line, "EVT ROOM PRESENCE ENTER") {
				lenRoom += 1
			} else if strings.HasPrefix(line, "EVT ROOM PRESENCE LEAVE") {
				lenRoom -= 1
			}

			label.SetText(fmt.Sprintf("Nombre de joueur dans la room: %d\nNombre de joueur global: %d", lenRoom, lenServer))
		}

	})
	label.Wrapping = fyne.TextWrapWord
	label.Alignment = fyne.TextAlignCenter

	return container.NewCenter(label)
}
