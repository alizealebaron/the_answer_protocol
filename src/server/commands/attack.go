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
	"strconv"
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

	weapon_damage := player.Attack
	if len(args) == 2 {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.New("ERR WEAPON_NOT_FOUND")
		}
		for i, _ := range player.Inventory {
			if i.GetId() == id {
				weapon, err := i.(models.Weapon)
				if !err {
					return errors.New("ERR WEAPON_NOT_FOUND")
				}
				weapon_damage = weapon.Damage
			}
		} 
	}

	
	p_room, err := tapManager.FindPlayerRoom(player.Id)
	if (err != nil) {
		return err
	}
	
	var target *models.Monster
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("ERR TARGET_NOT_FOUND")
	}
	
	target_exist := false
	for _, m := range p_room.Arena {
		if m.Entity_id == id {
			target_exist = true
			target = m
		}
	} 

	if (target_exist == false) {
		return errors.New("ERR TARGET_NOT_FOUND")
	}
		
		t_room, err  := tapManager.FindMonsterRoom(target.GetId())
	if (err != nil) {
		return err
	}
	var damage int
	var message1 string
	var message2 string
	if (p_room == t_room) {
		attack_dice := rand.IntN(20 - 1) + 1
		if attack_dice > target.Defense {
			if attack_dice == 20 {
				damage = rand.IntN(weapon_damage - 1) + rand.IntN(weapon_damage - 1) + 4
			} else {
				damage = rand.IntN(weapon_damage - 1) + 3
			}
		} else {
			damage = 0
		}
		target.Pv -= damage
		var status string
		if target.Pv > 0 {
			status = "healthy"
		} else {
			status = "dead"
			p_room.RemoveMonsterToRoom(*target)
			quantity := rand.IntN(target.QuantityMax - target.QuantityMin) + target.QuantityMax
			for i := 1; i <= quantity; i++ {
				p_room.AddItemToRoom(target.Loot)
			} 
		}
		message1 = fmt.Sprintf("Ok [{\"attacker\": %s, \"attack dice\": %d, \"attacker_hp\": %d, \"target_hp\": %d, \"damage\": %d, \"target_status\": %s}]", player.Name, attack_dice, player.Pv, target.Pv, damage, status)
		
		attack_dice = rand.IntN(20 - 1) + 1
		if attack_dice > 10 {
			if attack_dice == 20 {
				damage = rand.IntN(weapon_damage - 1) + rand.IntN(weapon_damage - 1) + 4
				} else {
					damage = rand.IntN(weapon_damage - 1) + 3
				}
		} else {
			damage = 0
		}
		player.AddLifePoint(-damage)
		
		message2 = fmt.Sprintf(" [{\"attacker\": %s, \"attack dice\": %d, \"attacker_hp\": %d, \"target_hp\": %d, \"damage\": %d, \"target_status\": %s}]", target.Name, attack_dice, target.Pv, player.Pv, damage, player.Status)
		if player.Status == "dead" {
			player.PlayerDeath(tapManager)
		}
	}
	server_write.ServerWrite(player.Conn, message1+message2+"\n")
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + message1+message2+"\n")
	return nil
}