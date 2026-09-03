/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* attack.go                                         :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: rruiz, alebaron, emarette                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/30 13:18:29 by emarette        #+#    #+#              */
/* Updated: 2026/08/30 13:19:32 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"fmt"
	// "strings"
	"errors"
	"math/rand/v2"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Attack(args []string, tapManager *models.TapManager, player *models.Player) error {
	if len(args) != 2 && len(args) != 1{
		return errors.New("ERR 302 NO_ITEM_SEND")
	}

	weapon_damage := 4
	if len(args) == 2 {
		for i, _ := range player.Inventory {
			if i.GetName() == args[1] {
				weapon_damage = i.GetDamage() 
			}
		} 
	}

	var target models.Monster
	target_exist := false
	for _, m := range tapManager.Lst_Monster {
		if m.Name == args[0] {
			target_exist = true
			target = m
		}
	} 
	if (target_exist == false) {
		return errors.New("ERR TARGET_NOT_FOUND")
	}

	p_room, err := tapManager.FindPlayerRoom(player.Id)
	if (err != nil) {
		return err
	}
	t_room, err  := tapManager.FindMonsterRoom(target.GetId())
	if (err != nil) {
		return err
	}
	var damage int
	var message1 string
	var message2 string
	if (p_room == t_room) {
		attack_dice := rand.IntN(20)
		if attack_dice > target.Defense {
			if attack_dice == 20 {
				damage = rand.IntN(weapon_damage) + rand.IntN(weapon_damage) + 2
			} else {
				damage = rand.IntN(weapon_damage) + 2
			}
		} else {
			damage = 0
		}
		target.Pv -= damage
		message1 = fmt.Sprint("OK [{\"attacker\": ", player.Name, "\"attacker_hp\": ", player.Pv, ", \"target_hp\":", target.Pv, "\"damage\": ", damage, "}, ")
		attack_dice = rand.IntN(20)
		if attack_dice > 10 {
			if attack_dice == 20 {
				damage = rand.IntN(weapon_damage) + rand.IntN(weapon_damage) + 2
				} else {
					damage = rand.IntN(weapon_damage) + 2
				}
		} else {
			damage = 0
		}
		player.Pv -= damage
		message2 = fmt.Sprint("[{\"attacker\": ", target.Name, "\"attacker_hp\": ", target.Pv, ", \"target_hp\":", player.Pv, "\"damage\": ", damage, "}]")
	}
	server_write.ServerWrite(player.Conn, message1+message2+"\n")
	// server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + room.ToString())
	return nil
}