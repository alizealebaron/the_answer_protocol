/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* attack.go                                         :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 08:53:30 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 10:52:15 by rruiz           ###   ########.fr       */
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

// Start of ATTACK. Retrieving information from the “LOOK” command.
func Attack(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showEnemy(stdin, listener, subCommandBox, data, wrappedBack)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

// Displays the enemies available in the arena, selecting one retrieves the player's inventory to choose a weapon.
func showEnemy(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, room types.LookInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0
	for _, enemy := range room.Arena {
		enemyName := enemy.Name
		enemyId := enemy.Id
		enemyButton := widget.NewButton(enemyName, func() {
			var id int
			id = listener.Subscribe(func(line string) {
				if !strings.HasPrefix(line, "OK {\"items\":") {
					return
				}
				room := strings.TrimPrefix(line, "OK ")
				var inventory types.InventoryInfo
				if err := json.Unmarshal([]byte(room), &inventory); err != nil {
					listener.Unsubscribe(id)
					back()
					return
				}
				listener.Unsubscribe(id)
				showWeapon(stdin, subCommandBox, enemyId, inventory, back)
			})
			fmt.Fprintf(stdin, "INVENTORY\n")
		})
		enemyButton.Importance = widget.LowImportance
		subCommandBox.Add(enemyButton)
		len += 1
	}
	if len == 0 {
		listener.Distribute("No enemy to attack here.")
		back()
	}

	subCommandBox.Refresh()
}

// Shows the available weapons to launch an ATTACK, or fights bare-handed if none are owned.
func showWeapon(stdin io.WriteCloser, subCommandBox *fyne.Container, enemyId int, inventory types.InventoryInfo, back func()) {
	subCommandBox.RemoveAll()

	len := 0

	for _, item := range inventory.Items {
		if item.Is_Weapon {
			if len == 0 {
				handButton := widget.NewButton("Hand", func() {
					fmt.Fprintf(stdin, "ATTACK %d\n", enemyId)
					fmt.Printf("ATTACK %d\n", enemyId)
					back()
				})
				handButton.Importance = widget.LowImportance
				subCommandBox.Add(handButton)
			}

			weaponName := item.Name
			weaponId := item.Id
			weaponButton := widget.NewButton(weaponName, func() {
				fmt.Fprintf(stdin, "ATTACK %d %d\n", enemyId, weaponId)
				fmt.Printf("ATTACK %d %d\n", enemyId, weaponId)
				back()
			})
			weaponButton.Importance = widget.LowImportance
			subCommandBox.Add(weaponButton)
			len += 1
		}
	}

	if len == 0 {
		fmt.Fprintf(stdin, "ATTACK %d\n", enemyId)
		fmt.Printf("ATTACK %d\n", enemyId)
		back()
	}

	subCommandBox.Refresh()
}
