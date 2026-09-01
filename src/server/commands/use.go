/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* use.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/28 09:38:03 by alebaron        #+#    #+#              */
/* Updated: 2026/08/28 16:13:12 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"errors"
	"strconv"
	"math/rand/v2"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Use(args []string, tapManager *models.TapManager, player *models.Player) error {
	
	// === Vérification de la longueur des arguments === //
	if len(args) != 1 {
		return errors.New("ERR 302 NO_ITEM_SEND")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR PLAYER_NOT_FOUND_IN_ANY_ROOM")
	}

	// === Vérification de la présence de l'item dans l'inventaire === //
	id, err := strconv.Atoi(args[0])

	item , err := player.GetItem(id)
	if err != nil {
		return errors.New("ERR 404 ITEM_NOT_FOUND")
	}

	// === Vérification de l'item utilisé === //

	switch (*item).GetId() {
    case 2:
        return fish(player, *room, *tapManager)
	default:
        return errors.New("ERR 407 ITEM_NOT_USABLE_HERE")
    }
}

func fish(player *models.Player, room models.Room, tap models.TapManager) error {
	
	if len(room.Fishing) == 0 {
		return errors.New("ERR 407 ITEM_NOT_USABLE_HERE")
	}
	
	// Calcul de la somme totale des taux de loot (Si pas égale à 100 plante pas)
	total := 0
	for _, entry := range room.Fishing {
		total += entry.LootRate
	}

	if total <= 0 {
		return errors.New("ERR 500 INVALID_LOOT_TABLE")
	}

	// Tirage aléatoire entre 0 et total-1
	roll := rand.IntN(total)

	// On parcourt les entrées en cumulant les taux
	cumulative := 0
	for _, entry := range room.Fishing {
		cumulative += entry.LootRate
		if roll < cumulative {
	
			item, err := tap.GetItemById(entry.ItemID)
			if err != nil {
				return errors.New("ERR 404 ITEM_NOT_FOUND")
			}
			player.AddItemToPlayer(item)

			// === Envoie des messages au client et dans les logs === //
			str_ret := "OK fishing=" + (item).GetName() + "\n"
			server_write.ServerWrite(player.Conn, str_ret)
			server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + str_ret)
			
			return nil
		}
	}

	// Ne devrait normalement jamais arriver mais on sait pas hein
	return errors.New("ERR 500 FISHING_ROLL_FAILED")
}