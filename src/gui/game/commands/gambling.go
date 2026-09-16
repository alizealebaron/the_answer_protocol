/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* gambling.go                                       :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/12 10:02:26 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:47:04 by rruiz           ###   ########.fr       */
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

// Start of GAMBLING. Retrieving information from the "LOOK" command.
func Gambling(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
	lookupPrefixes := []string{"OK {\"id\":", "OK {\"items\":"}
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
		if data.Name != "CASINO" {
			listener.Distribute("You must be in a casino to gamble.")
			wrappedBack()
			return
		}
		askCoinInventory(stdin, listener, subCommandBox, wrappedBack)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

// Retrieves the player's inventory to look for the gambling coins they own.
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
			back()
			return
		}
		listener.Unsubscribe(id)
		showCoinSelection(stdin, listener, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "INVENTORY\n")
}

// Asks how many coins to bet, validates the amount and sends the GAMBLING command.
func showCoinSelection(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, inventory types.InventoryInfo, back func()) {
	subCommandBox.RemoveAll()

	playerCoins := 0

	for _, item := range inventory.Items {
		if item.Name == "Gambling Coin" {
			playerCoins = item.Quantity
			break
		}
	}

	if playerCoins < 10 {
		listener.Distribute("You need at least 10 gambling coins to gamble.")
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
