/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* use.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/12 15:43:01 by rruiz           #+#    #+#              */
/* Updated: 2026/09/16 19:15:08 by rruiz           ###   ########.fr       */
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

// Start of USE. Retrieving information from the "INVENTORY" command.
func Use(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		if !strings.HasPrefix(line, "OK {\"items\":") {
			return
		}
		inventory := strings.TrimPrefix(line, "OK ")
		var data types.InventoryInfo
		if err := json.Unmarshal([]byte(inventory), &data); err != nil {
			listener.Unsubscribe(id)
			wrappedBack()
			return
		}
		listener.Unsubscribe(id)
		showUsableObjects(stdin, listener, subCommandBox, data, wrappedBack)
	})
	_, _ = fmt.Fprintf(stdin, "INVENTORY\n")
}

// Displays the usable items as buttons, or returns if there are none.
func showUsableObjects(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, inventory types.InventoryInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, item := range inventory.Items {
		if item.Is_Usable {
			itemName := item.Name
			itemId := item.Id
			itemButton := widget.NewButton(itemName, func() {
				_, _ = fmt.Fprintf(stdin, "USE %d\n", itemId)
				_, _ = fmt.Printf("USE %d\n", itemId)
				back()
			})
			itemButton.Importance = widget.LowImportance
			subCommandBox.Add(itemButton)
			len += 1
		}
	}
	if len == 0 {
		listener.Distribute(types.Translate("Nothing to use in your inventory."))
		back()
	}

	subCommandBox.Refresh()
}
