/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quest.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/14 16:21:56 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:47:15 by rruiz           ###   ########.fr       */
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

// Start of QUEST. Retrieving information from the "LOOK" command.
func Quest(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	lookupPrefixes := []string{"OK {\"id\":"}
	types.Mute(lookupPrefixes...)
	wrappedBack := func() {
		types.Unmute(lookupPrefixes...)
		back()
	}

	// Usage of listener to send the command.
	// Retrieve the information in a dedicated structure, and execute the rest of the command.
	// Used in virtually all commands.
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"id\":") {
			return
		}
		room := strings.TrimPrefix(line, "OK ")
		var data types.LookInfo
		if err := json.Unmarshal([]byte(room), &data); err != nil {
			listener.Unsubscribe(id)
			wrappedBack()
			return
		}
		listener.Unsubscribe(id)
		showQuestGivers(stdin, listener, subCommandBox, data, wrappedBack)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

// Displays the quest givers present in the room as buttons, or returns if there are none.
func showQuestGivers(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, npc := range room.Allies {
		if npc.Is_qg {
			qgName := npc.Name
			qgId := npc.Id
			qgButton := widget.NewButton(qgName, func() {
				fmt.Fprintf(stdin, "QUEST %d\n", qgId)
				fmt.Printf("QUEST %d\n", qgId)
				back()
			})
			qgButton.Importance = widget.LowImportance
			subCommandBox.Add(qgButton)
			len += 1
		}
	}
	if len == 0 {
		listener.Distribute(types.Translate("No quest giver here."))
		back()
	}

	subCommandBox.Refresh()
}
