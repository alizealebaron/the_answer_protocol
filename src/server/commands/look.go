/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* look.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/24 17:28:51 by alebaron        #+#    #+#              */
/* Updated: 2026/09/15 14:26:50 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
    "fmt"
	"errors"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Look(args []string, tapManager *models.TapManager, player *models.Player) error {

	var copy_room models.Room

	room, err := tapManager.FindPlayerRoom(player.Id)
	copy_room = (*room)

	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// Easter-egg pour le CASINO //

	quantity, _ := player.GetQuantityItem(1)
	if room.Name == "CASINO" && quantity >= 1000 {
		copy_room.NeighborRoom.South = 15
	}

	server_write.ServerWrite(player.Conn, "OK " + copy_room.ToString() + "\n")
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + room.ToString())
	return nil
}