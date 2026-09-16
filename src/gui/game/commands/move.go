/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* move.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/09 10:10:18 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:51:02 by rruiz           ###   ########.fr       */
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

// Start of MOVE. Retrieving information from the "LOOK" command.
func Move(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
			return
		}
		listener.Unsubscribe(id)
		showDirections(stdin, listener, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

// Displays the available directions as buttons, or returns to the previous menu if there are none.
func showDirections(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	// A map of directions link with the corresponding room IDs.
	directions := map[string]int{
		"NORTH": room.NeighborRoom.North,
		"EAST":  room.NeighborRoom.East,
		"SOUTH": room.NeighborRoom.South,
		"WEST":  room.NeighborRoom.West,
	}

	len := 0

	for direction, id := range directions {
		if id != 0 {
			dir := direction
			directionButton := widget.NewButton(dir, func() {
				fmt.Fprintf(stdin, "MOVE %s\n", dir)
				fmt.Printf("MOVE %s\n", dir)

				// Send LOOK only after the server has processed MOVE.
				var listenerId int
				listenerId = listener.Subscribe(func(line string) {
					if strings.HasPrefix(line, "OK {\"id\":") {
						return
					}
					if strings.HasPrefix(line, "OK ") || strings.HasPrefix(line, "ERR ") {
						fmt.Fprintf(stdin, "LOOK\n")
						listener.Unsubscribe(listenerId)
						back()
					}
				})
			})

			directionButton.Importance = widget.LowImportance
			subCommandBox.Add(directionButton)
			len += 1
		}
	}
	if len == 0 {
		back()
	}
	subCommandBox.Refresh()
}
