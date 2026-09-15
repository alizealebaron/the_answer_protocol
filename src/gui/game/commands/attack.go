/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* attack.go                                         :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 08:53:30 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 09:26:22 by rruiz           ###   ########.fr       */
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

func Attack(stdin io.WriteCloser, listener *types.Listener, subCommandBox *fyne.Container, back func()) {
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
		showEnemy(stdin, listener, subCommandBox, data, back)
	})
	fmt.Fprintf(stdin, "LOOK\n")
}

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
		back()
	}

	subCommandBox.Refresh()
}

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
