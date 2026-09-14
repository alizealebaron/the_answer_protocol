/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* buy.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/12 16:32:52 by rruiz           #+#    #+#              */
/* Updated: 2026/09/12 16:32:52 by rruiz           ###   ########.fr       */
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

func GetMoney(stdin io.WriteCloser, listener *types.Listener, callback func(int)) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"items\":") {
			return
		}
		raw := strings.TrimPrefix(line, "OK ")
		var data types.InventoryInfo
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			listener.Unsubscribe(id)
			return
		}
		listener.Unsubscribe(id)
		callback(data.Money)
	})
	fmt.Fprintf(stdin, "INVENTORY\n")
}

func Buy(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showTrader(stdin, listener, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

func showTrader(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, npc := range room.Allies {
		if npc.Is_trader {
			traderName := npc.Name
			traderId := npc.Id
			traderButton := widget.NewButton(traderName, func() {
				var id int
				id = listener.Subscribe(func(line string) {
					trade := strings.TrimPrefix(line, "OK trade=")
					var data []types.TradeInfo
					if err := json.Unmarshal([]byte(trade), &data); err != nil {
						listener.Unsubscribe(id)
						return
					}
					listener.Unsubscribe(id)
					showTraderInventory(stdin, listener, subCommandBox, data, traderId, back)
				})
				fmt.Fprintf(stdin, "TRADE %d\n", traderId)
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

func showTraderInventory(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, traderInventory []types.TradeInfo, traderId int, back func()) {
	subCommandBox.RemoveAll()

	for _, item := range traderInventory {
		itemName := item.Name
		itemId := item.Id
		itemCost := item.Cost
		itemButton := widget.NewButton(fmt.Sprintf("%s   |   %d", itemName, itemCost), func() {
			showQuantityEntry(stdin, listener, subCommandBox, traderId, itemId, itemCost, back)
		})
		itemButton.Importance = widget.LowImportance
		subCommandBox.Add(itemButton)
	}

	subCommandBox.Refresh()
}

func showQuantityEntry(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, traderId int, itemId int, itemCost int, back func()) {
	subCommandBox.RemoveAll()

	quantityEntry := widget.NewEntry()
	quantityEntry.SetPlaceHolder("Enter the quantity you would like to purchase.")

	buyButton := widget.NewButton("BUY", func() {
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

		getPlayerMoney(stdin, listener, func(money int) {
			if money < quantity*itemCost {
				quantityEntry.SetPlaceHolder("You need to have enough money to buy it. ")

				quantityEntry.SetText("")
				quantityEntry.Refresh()
				return
			}
			fmt.Fprintf(stdin, "BUY %d %d %d\n", traderId, itemId, quantity)
			fmt.Printf("BUY %d %d %d\n", traderId, itemId, quantity)
			back()
		})
	})

	subCommandBox.Add(quantityEntry)
	subCommandBox.Add(buyButton)

	subCommandBox.Refresh()

}

func getPlayerMoney(stdin io.WriteCloser, listener *types.Listener, callback func(int)) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"items\":") {
			return
		}
		raw := strings.TrimPrefix(line, "OK ")
		var data types.InventoryInfo
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			listener.Unsubscribe(id)
			return
		}
		listener.Unsubscribe(id)
		callback(data.Money)
	})
	fmt.Fprintf(stdin, "INVENTORY\n")
}
