/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* trade.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/14 14:23:20 by rruiz           #+#    #+#              */
/* Updated: 2026/09/16 19:13:05 by rruiz           ###   ########.fr       */
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

// Start of TRADE. Retrieving information from the "LOOK" command.
func Trade(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showTrader1(stdin, listener, subCommandBox, data, wrappedBack)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

// Displays the traders present in the room as buttons, then sends the TRADE command when one is selected.
func showTrader1(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
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
				})
				fmt.Fprintf(stdin, "TRADE %d\n", traderId)
				fmt.Printf("TRADE %d\n", traderId)
				back()
			})
			traderButton.Importance = widget.LowImportance
			subCommandBox.Add(traderButton)
			len += 1
		}
	}
	if len == 0 {
		listener.Distribute("No trader here to trade with.")
		back()
	}

	subCommandBox.Refresh()
}
