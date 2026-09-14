/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* sell.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/14 14:44:04 by rruiz           #+#    #+#              */
/* Updated: 2026/09/14 16:07:02 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"the_answer_protocol/src/gui/game/types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Sell(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showTrader2(stdin, listener, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

func showTrader2(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, npc := range room.Allies {
		if npc.Is_trader {
			traderName := npc.Name
			traderId := npc.Id
			traderButton := widget.NewButton(traderName, func() {
				var id int
				id = listener.Subscribe(func(line string) {
					trade := strings.TrimPrefix(line, "OK ")
					var data types.InventoryInfo
					if err := json.Unmarshal([]byte(trade), &data); err != nil {
						listener.Unsubscribe(id)
						return
					}
					listener.Unsubscribe(id)
					showInventory(stdin, subCommandBox, data, traderId, back)
				})
				fmt.Fprintf(stdin, "INVENTORY\n")
			})
			traderButton.Importance = widget.LowImportance
			subCommandBox.Add(traderButton)
			len += 1
		}
	}
	if len == 0 {
		back()
	}

	subCommandBox.Refresh()
}

func showInventory(stdin io.WriteCloser, subCommandBox *fyne.Container, inventory types.InventoryInfo, traderId int, back func()) {
	subCommandBox.RemoveAll()

	if len(inventory.Items) == 0 {
		back()
	}

	for _, item := range inventory.Items {
		itemName := item.Name
		itemId := item.Id
		itemCost := int(float64(item.Cost) * 0.6)
		itemQuantity := item.Quantity
		itemButton := widget.NewButton(fmt.Sprintf("%s   |   %d", itemName, itemCost), func() {
			showQuantityEntrySell(stdin, subCommandBox, traderId, itemId, itemQuantity, back)
		})
		itemButton.Importance = widget.LowImportance
		subCommandBox.Add(itemButton)
	}

	subCommandBox.Refresh()
}

func showQuantityEntrySell(stdin io.WriteCloser, subCommandBox *fyne.Container, traderId int, itemId int, itemQuantity int, back func()) {
	subCommandBox.RemoveAll()

	quantityEntry := widget.NewEntry()
	quantityEntry.SetPlaceHolder("Enter the quantity you would like to sell.")

	buyButton := widget.NewButton("SELL", func() {
		quantity, err := strconv.Atoi(strings.TrimSpace(quantityEntry.Text))
		if err != nil {
			quantityEntry.SetPlaceHolder("The quantity must be an integer.")

			quantityEntry.SetText("")
			quantityEntry.Refresh()
			return
		}

		if quantity <= 0 {
			quantityEntry.SetPlaceHolder("The quantity must be at least 1.")

			quantityEntry.SetText("")
			quantityEntry.Refresh()
			return
		}

		if quantity > itemQuantity {
			quantityEntry.SetPlaceHolder("You must own the item.")

			quantityEntry.SetText("")
			quantityEntry.Refresh()
			return
		}

		fmt.Fprintf(stdin, "SELL %d %d %d\n", traderId, itemId, quantity)
		fmt.Printf("SELL %d %d %d\n", traderId, itemId, quantity)
		back()
	})

	subCommandBox.Add(quantityEntry)
	subCommandBox.Add(buyButton)

	subCommandBox.Refresh()
}
