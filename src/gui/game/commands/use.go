/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* use.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/12 15:43:01 by rruiz           #+#    #+#              */
/* Updated: 2026/09/12 16:31:37 by rruiz           ###   ########.fr       */
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

func Use(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"items\":") {
			return
		}
		room := strings.TrimPrefix(line, "OK ")
		var data types.InventoryInfo
		if err := json.Unmarshal([]byte(room), &data); err != nil {
			listener.Unsubscribe(id)
			return
		}
		listener.Unsubscribe(id)
		showUsableObjects(stdin, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "INVENTORY\n")
}

func showUsableObjects(stdin io.WriteCloser, subCommandBox *fyne.Container, inventory types.InventoryInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, item := range inventory.Items {
		if item.Is_Usable {
			itemName := item.Name
			itemId := item.Id
			itemButton := widget.NewButton(itemName, func() {
				fmt.Fprintf(stdin, "USE %d\n", itemId)
				fmt.Printf("USE %d\n", itemId)
				back()
			})
			itemButton.Importance = widget.LowImportance
			subCommandBox.Add(itemButton)
			len += 1
		}
	}
	if len == 0 {
		back()
	}

	subCommandBox.Refresh()
}
