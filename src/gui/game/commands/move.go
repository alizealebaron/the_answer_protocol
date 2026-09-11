/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* move.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/09 10:10:18 by rruiz           #+#    #+#              */
/* Updated: 2026/09/09 17:28:01 by rruiz           ###   ########.fr       */
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

func Move(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showDirections(stdin, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

func showDirections(stdin io.WriteCloser, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

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
				back()
			})
			directionButton.Importance = widget.LowImportance
			subCommandBox.Add(directionButton)
			len += 1
		}
		if len == 0 {
			back()
		}
	}
	subCommandBox.Refresh()
}
