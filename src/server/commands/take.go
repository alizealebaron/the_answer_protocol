/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* take.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/26 10:05:36 by alebaron        #+#    #+#              */
/* Updated: 2026/08/26 10:59:39 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
    // "fmt"
	"errors"
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Take(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) != 1 {
		return errors.New("ERR 302 NO_ITEM_SEND")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR PLAYER_NOT_FOUND_IN_ANY_ROOM")
	}

	// === Vérification de la présence de l'item dans la room === //
	id, err := strconv.Atoi(args[0])

	item, err := room.RemoveItemToRoom(id)
	if err != nil {
		return errors.New("ERR 404 ITEM_NOT_FOUND")
	}

	// === Ajout de l'item à l'inventaire du joueur === //
	player.AddItemToPlayerWQuantity(*item, 1)
	
	// === Envoie des messages au client et dans les logs === //
	str_ret := "OK taken=" + (*item).GetName() + "\n"
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + str_ret)

	return nil
}