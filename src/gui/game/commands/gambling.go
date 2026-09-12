/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* gambling.go                                       :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/12 10:02:26 by rruiz           #+#    #+#              */
/* Updated: 2026/09/12 15:23:52 by rruiz           ###   ########.fr       */
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

func Gambling(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		if data.Name != "CASINO" {
			return
		}
		askCoinInventory(stdin, listener, subCommandBox, back)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

func askCoinInventory(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showCoinSelection(stdin, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "INVENTORY\n")
}

func showCoinSelection(stdin io.WriteCloser, subCommandBox *fyne.Container, inventory types.InventoryInfo, back func()) {
	subCommandBox.RemoveAll()

	playerCoins := 0

	for _, item := range inventory.Items {
		if item.Name == "Gambling Coin" {
			playerCoins = item.Quantity
			break
		}
	}

	if playerCoins < 10 {
		back()
		return
	}

	gamblingEntry := widget.NewEntry()
	gamblingEntry.SetPlaceHolder("Enter the number of gambling coins you want to bet.")

	gamblingButton := widget.NewButton("GAMBLE", func() {
		quantityToGamble, err := strconv.Atoi(strings.TrimSpace(gamblingEntry.Text))
		if err != nil || quantityToGamble < 10 {
			gamblingEntry.SetPlaceHolder("The amount must be at least 10.")

			gamblingEntry.SetText("")
			gamblingEntry.Refresh()
			return
		}

		if quantityToGamble > playerCoins {
			gamblingEntry.SetPlaceHolder("You don't have that many gambling coins.")
			gamblingEntry.SetText("")
			gamblingEntry.Refresh()
			return
		}

		fmt.Fprintf(stdin, "GAMBLING %d\n", quantityToGamble)
		fmt.Printf("GAMBLING %d\n", quantityToGamble)
		back()
	})

	subCommandBox.Add(gamblingEntry)
	subCommandBox.Add(gamblingButton)

	subCommandBox.Refresh()
}
