/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* talk.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/11 21:52:37 by rruiz           #+#    #+#              */
/* Updated: 2026/09/11 21:59:04 by rruiz           ###   ########.fr       */
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

func Talk(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"id\":") {
			return
		}
		room := strings.TrimPrefix(line, "OK ")
		var data types.LookInfo
		if err := json.Unmarshal([]byte(room), &data); err != nil {
			listener.Unsubscribe(id)
			return
		}
		listener.Unsubscribe(id)
		showTalkableNpc(stdin, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

func showTalkableNpc(stdin io.WriteCloser, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, npc := range room.Allies {
		npcName := npc.Name
		npcId := npc.Id
		npcButton := widget.NewButton(npcName, func() {
			fmt.Fprintf(stdin, "TALK %d\n", npcId)
			fmt.Printf("TALK %d\n", npcId)
			back()
		})
		npcButton.Importance = widget.LowImportance
		subCommandBox.Add(npcButton)
		len += 1
	}

	if len == 0 {
		back()
	}

	subCommandBox.Refresh()
}
