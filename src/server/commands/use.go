/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* use.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/18 10:46:58 by alebaron        #+#    #+#              */
/* Updated: 2026/09/18 10:47:37 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Use(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) != 1 {
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// === Vérification de la présence de l'item dans l'inventaire === //
	id, err := strconv.Atoi(args[0])

	item, err := player.GetItem(id)
	if err != nil {
		return errors.New("ERR 404 ITEM_NOT_FOUND")
	}

	// === Vérification de si l'item est un edible === //

	edible, ok := (*item).(models.Edible)
	if ok {
		if edible.Effect == "HEAL" {
			player.AddLifePoint(edible.Value)
		}
		if edible.Effect == "DAMAGE" {
			player.AddLifePoint(-edible.Value)
		}

		player.RemoveItemToPlayerWQuantity(edible.Id, 1)
		// === Envoie des messages au client et dans les logs === //

		str_ret := fmt.Sprintf("OK {\"used\": \"%s\", \"effect\": \"%s\", \"value\": %d}\n", edible.Name, edible.Effect, edible.Value)
		server_write.ServerWrite(player.Conn, str_ret)
		server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)

		return nil
	}

	// === Vérification de si l'item est un Usable === //

	usable, ok := (*item).(models.Usable)
	if ok {
		switch usable.Id {
		case 2:
			return fish(player, *room, *tapManager)
		}
	}

	return errors.New("ERR 408 ITEM_NOT_USABLE")
}

func fish(player *models.Player, room models.Room, tap models.TapManager) error {

	if len(room.Fishing) == 0 {
		return errors.New("ERR 408 INVALID_LOCATION")
	}

	// Calcul de la somme totale des taux de loot (Si pas égale à 100 plante pas)
	total := 0
	for _, entry := range room.Fishing {
		total += entry.LootRate
	}

	if total <= 0 {
		return errors.New("ERR 999 INVALID_LOOT_TABLE")
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
			player.AddItemToPlayerWQuantity(item, 1)

			// === Envoie des messages au client et dans les logs === //
			str_ret := "OK fishing=" + (item).GetName() + "\n"
			server_write.ServerWrite(player.Conn, str_ret)
			server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)

			return nil
		}
	}

	// Ne devrait normalement jamais arriver mais on sait pas hein
	return errors.New("ERR 999 FISHING_ROLL_FAILED")
}
