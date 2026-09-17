/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* take.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/11 15:28:14 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:47:26 by rruiz           ###   ########.fr       */
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

// Start of TAKE. Retrieving information from the "LOOK" command.
func Take(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showSelectableObjects(stdin, listener, subCommandBox, data, wrappedBack)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

// Displays the items present in the room as buttons, or returns if there are none.
func showSelectableObjects(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, item := range room.Items {
		itemName := item.Name
		itemId := item.Id
		itemButton := widget.NewButton(itemName, func() {
			fmt.Fprintf(stdin, "TAKE %d\n", itemId)
			fmt.Printf("TAKE %d\n", itemId)
			back()
		})
		itemButton.Importance = widget.LowImportance
		subCommandBox.Add(itemButton)
		len += 1
	}
	if len == 0 {
		listener.Distribute("Nothing to take here.")
		back()
	}

	subCommandBox.Refresh()
}
