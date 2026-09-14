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

	// on declare les variable \\
	var target *models.Monster
	var new_target *models.Player
	var damage int
	var status string
	var message1 string
	var message2 string

	if len(args) != 2 && len(args) != 1{
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	// on récupère l'arme du joueur depuis son inventaire grace a l'id de l'objet \\
	weapon_damage := player.Attack
	if len(args) == 2 {
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return errors.New("ERR 404 WEAPON_NOT_FOUND")
		}
		for i, _ := range player.Inventory {
			if i.GetId() == id {
				weapon, err := i.(models.Weapon)
				if !err {
					return errors.New("ERR 404 WEAPON_NOT_FOUND")
				}
				weapon_damage = weapon.Damage
			}
		} 
	}

	// on recherche la room du joueur \\
	p_room, err := tapManager.FindPlayerRoom(player.Id)
	if (err != nil) {
		return err
	}
	
	// On converti le deuxieme argument en int \\
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("ERR 404 TARGET_NOT_FOUND")
	}
	
	// Grace a l'id precedant on essaie de recuperer la cible depuis la room du joueur \\
	target_exist := false
	for _, m := range p_room.Arena {
		if m.Entity_id == id {
			target_exist = true
			target = m
		}
	} 

	// si on ne trouve pas la cible on renvoie une erreur \\
	if (target_exist == false) {
		return errors.New("ERR 404 TARGET_NOT_FOUND")
	}

	// le joueur attaque la cible \\
	// il lance un de d'attaque \\
	attack_dice := rand.IntN(20 - 1) + 1
	//si le jet d'attacke est > a la classe d'armure de la cible
	if attack_dice > target.Defense {
	// si le jets est egal a 20 coup critique sinon coup simple
		if attack_dice == 20 {
			damage = rand.IntN(weapon_damage - 1) + rand.IntN(weapon_damage - 1) + 4
		} else {
			damage = rand.IntN(weapon_damage - 1) + 3
		}
		target.Pv -= damage
	}
	
	// on verifie et modifie si besoin le status de la cible
	if target.Pv > 50 {
		status = "healthy"
	} else if target.Pv <= 50 && target.Pv > 0 {
		status = "bloody"
	} else {
		status = "dead"
		// si la cible est morte on la retire de l'arene et on drop son loot au sol
		p_room.RemoveMonsterToRoom(*target)
		quantity := rand.IntN(target.QuantityMax - target.QuantityMin) + target.QuantityMax
		for i := 1; i <= quantity; i++ {
			p_room.AddItemToRoom(target.Loot)
		}

		// On udpate les quêtes du joueurs si besoin
		player.UpdateQuestMonster(*target)
	}
	// on ecris la premiere moitier du message
	message1 = fmt.Sprintf("OK [{\"attacker\": %s, \"attack dice\": %d, \"attacker_hp\": %d, \"target_hp\": %d, \"damage\": %d, \"target_status\": %s}]", player.Name, attack_dice, player.Pv, target.Pv, damage, status)
	
	// la cible attack le joueur

	if player.Group != nil {
		index := rand.IntN(len(player.Group.Lst_Player))
		new_target = player.Group.Lst_Player[index]
	} else {
		new_target = player
	}
	
	attack_dice = rand.IntN(20 - 1) + 1
	if attack_dice > new_target.Defense {
		if attack_dice == 20 {
			damage = rand.IntN(target.Attack - 1) + rand.IntN(target.Attack - 1) + 4
		} else {
			damage = rand.IntN(target.Attack - 1) + 3
		}
		new_target.AddLifePoint(-damage)
	}
	
	// on ecris la deuxieme moitier du message
	message2 = fmt.Sprintf(" [{\"attacker\": %s, \"attack dice\": %d, \"attacker_hp\": %d, \"target\": %s, \"target_hp\": %d, \"damage\": %d, \"target_status\": %s}]", target.Name, attack_dice, target.Pv, new_target.Name, new_target.Pv, damage, new_target.Status)
	
	// on ecris le resultat de l'attaque
	server_write.ServerWrite(player.Conn, message1+message2+"\n")
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + message1+message2+"\n")

	// on verfie si le joueur est mort (Je le mets ici pour que le message de changement de room soit dans le bon ordre ~Alizéa)
	if new_target.Status == "dead" {
		new_target.PlayerDeath(tapManager)
	}

	return nil
}
