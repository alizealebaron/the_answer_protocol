/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* players_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/07 11:17:14 by rruiz           #+#    #+#              */
/* Updated: 2026/09/11 22:27:54 by rruiz           ###   ########.fr       */
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
	label := widget.NewLabel("Use the “WHO” command to view information about players number.")
	listener.Subscribe(func(line string) {
		if strings.HasPrefix(line, "OK { \"room\":") {
			whoJson := strings.TrimPrefix(line, "OK ")
			var data types.WhoInfo
			err := json.Unmarshal([]byte(whoJson), &data)
			if err != nil {
				return
			}
			label.SetText(fmt.Sprintf("Nombre de joueur dans la room: %d\nNombre de joueur global: %d", len(data.RoomInfo), data.ServerInfo))
			fmt.Println(line)
		}
	})
	label.Wrapping = fyne.TextWrapWord
	label.Alignment = fyne.TextAlignCenter

	return container.NewCenter(label)
}
