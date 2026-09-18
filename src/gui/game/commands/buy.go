/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* buy.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/12 16:32:52 by rruiz           #+#    #+#              */
/* Updated: 2026/09/16 19:29:32 by rruiz           ###   ########.fr       */
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

// Start of BUY. Retrieving information from the “LOOK” command.
func Buy(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	lookupPrefixes := []string{"OK {\"id\":", "OK trade=", "OK {\"items\":"}
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
		showTrader(stdin, listener, subCommandBox, data, wrappedBack)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

// Displays the traders present in the room as buttons, then retrieves the trader's inventory when one is selected.
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
						back()
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
		listener.Distribute(types.Translate("No trader here to buy from."))
		back()
	}

	subCommandBox.Refresh()
}

// Displays the trader's saleable items as buttons with their price.
func showTraderInventory(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, traderInventory []types.TradeInfo, traderId int, back func()) {
	subCommandBox.RemoveAll()

	for _, item := range traderInventory {
		itemName := item.Name
		itemId := item.Id
		itemCost := item.Cost
		itemButton := widget.NewButton(fmt.Sprintf("%s   |   %d", itemName, itemCost), func() {
			showQuantityEntryBuy(stdin, listener, subCommandBox, traderId, itemId, itemCost, back)
		})
		itemButton.Importance = widget.LowImportance
		subCommandBox.Add(itemButton)
	}

	subCommandBox.Refresh()
}

// Asks for the quantity to buy, checks the player's money and sends the BUY command.
func showQuantityEntryBuy(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, traderId int, itemId int, itemCost int, back func()) {
	subCommandBox.RemoveAll()

	quantityEntry := widget.NewEntry()
	quantityEntry.SetPlaceHolder(types.Translate("Enter the quantity you would like to purchase."))

	buyButton := widget.NewButton(types.Translate("BUY"), func() {
		quantity, err := strconv.Atoi(strings.TrimSpace(quantityEntry.Text))
		if err != nil {
			quantityEntry.SetPlaceHolder(types.Translate("The quantity must be an integer."))

			quantityEntry.SetText("")
			quantityEntry.Refresh()
			return
		}

		if quantity <= 0 {
			quantityEntry.SetPlaceHolder(types.Translate("The quantity must be at least 1."))

			quantityEntry.SetText("")
			quantityEntry.Refresh()
			return
		}

		getPlayerMoney(stdin, listener, back, func(money int) {
			if money < quantity*itemCost {
				quantityEntry.SetPlaceHolder(types.Translate("You need to have enough money to buy it. "))

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

// Retrieves the player's money via INVENTORY and passes it to the callback.
func getPlayerMoney(stdin io.WriteCloser, listener *types.Listener, onError func(), callback func(int)) {
	var id int
	id = listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"items\":") {
			return
		}
		raw := strings.TrimPrefix(line, "OK ")
		var data types.InventoryInfo
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			listener.Unsubscribe(id)
			onError()
			return
		}
		listener.Unsubscribe(id)
		callback(data.Money)
	})
	fmt.Fprintf(stdin, "INVENTORY\n")
}
