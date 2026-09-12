/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* look.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/24 17:28:51 by alebaron        #+#    #+#              */
/* Updated: 2026/08/26 10:59:01 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
    // "fmt"
	"errors"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Look(args []string, tapManager *models.TapManager, player *models.Player) error {

	room, err := tapManager.FindPlayerRoom(player.Id)

	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	server_write.ServerWrite(player.Conn, "OK " + room.ToString() + "\n")
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + room.ToString())
	return nil
}