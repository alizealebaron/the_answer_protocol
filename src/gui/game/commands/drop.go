/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* drop.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/11 17:37:24 by rruiz           #+#    #+#              */
/* Updated: 2026/09/11 17:51:55 by rruiz           ###   ########.fr       */
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

func Drop(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"items\"") {
			return
		}
		room := strings.TrimPrefix(line, "OK ")
		var data types.InventoryInfo
		if err := json.Unmarshal([]byte(room), &data); err != nil {
			listener.Unsubscribe(id)
			return
		}
		listener.Unsubscribe(id)
		showDropableObjects(stdin, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "INVENTORY\n")
}

func showDropableObjects(stdin io.WriteCloser, subCommandBox *fyne.Container, inventory types.InventoryInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, item := range inventory.Items {
		itemName := item.Name
		itemId := item.Id
		itemButton := widget.NewButton(itemName, func() {
			fmt.Fprintf(stdin, "DROP %d\n", itemId)
			fmt.Printf("DROP %d\n", itemId)
			back()
		})
		itemButton.Importance = widget.LowImportance
		subCommandBox.Add(itemButton)
		len += 1
	}
	if len == 0 {
		back()
	}

	subCommandBox.Refresh()
}
