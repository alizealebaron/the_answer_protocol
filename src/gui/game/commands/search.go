/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* search.go                                         :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/14 16:10:58 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:47:18 by rruiz           ###   ########.fr       */
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

// Start of SEARCH. Retrieving information from the "LOOK" command.
func Search(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showSearchableEnemy(stdin, listener, subCommandBox, data, wrappedBack)
	})
	_, _ = fmt.Fprintf(stdin, "LOOK\n")
}

// Displays the enemies present in the room as buttons, or returns if there are none.
func showSearchableEnemy(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, enemy := range room.Ennemies {
		enemyName := enemy.Name
		enemyId := enemy.Id
		enemyButton := widget.NewButton(enemyName, func() {
			_, _ = fmt.Fprintf(stdin, "SEARCH %d\n", enemyId)
			_, _ = fmt.Printf("SEARCH %d\n", enemyId)
			back()
		})
		enemyButton.Importance = widget.LowImportance
		subCommandBox.Add(enemyButton)
		len += 1
	}
	if len == 0 {
		listener.Distribute(types.Translate("Nothing to search here."))
		back()
	}

	subCommandBox.Refresh()
}
