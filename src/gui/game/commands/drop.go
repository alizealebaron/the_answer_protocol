/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* drop.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/11 17:37:24 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:46:59 by rruiz           ###   ########.fr       */
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

// Start of DROP. Retrieving information from the "INVENTORY" command.
func Drop(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	lookupPrefixes := []string{"OK {\"items\":"}
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
		if !strings.HasPrefix(line, "OK {\"items\"") {
			return
		}
		room := strings.TrimPrefix(line, "OK ")
		var data types.InventoryInfo
		if err := json.Unmarshal([]byte(room), &data); err != nil {
			listener.Unsubscribe(id)
			wrappedBack()
			return
		}
		listener.Unsubscribe(id)
		showDropableObjects(stdin, listener, subCommandBox, data, wrappedBack)
	})
	_, _ = fmt.Fprintf(stdin, "INVENTORY\n")
}

// Displays the items present in the player's inventory as buttons, or returns if nothing can be dropped.
func showDropableObjects(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, inventory types.InventoryInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, item := range inventory.Items {
		itemName := item.Name
		itemId := item.Id
		itemButton := widget.NewButton(itemName, func() {
			_, _ = fmt.Fprintf(stdin, "DROP %d\n", itemId)
			_, _ = fmt.Printf("DROP %d\n", itemId)
			back()
		})
		itemButton.Importance = widget.LowImportance
		subCommandBox.Add(itemButton)
		len += 1
	}
	if len == 0 {
		listener.Distribute(types.Translate("Nothing to drop."))
		back()
	}

	subCommandBox.Refresh()
}
