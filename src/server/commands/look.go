/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* look.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/24 17:28:51 by alebaron        #+#    #+#              */
/* Updated: 2026/08/25 13:30:56 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
    // "fmt"
	"net"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Look(args []string, tapManager *models.TapManager, player models.Player, conn net.Conn) error {

	for _, room := range tapManager.Lst_Room {
        for _, player_room := range room.Lst_Player {
			if player_room.Id == player.Id {
				server_write.ServerWrite(conn, "OK " + room.ToString() + "\n")
				server_write.WriteLog(conn, "SERVER", "To " + player.Name + ": " + room.ToString())
			}
		}
    }

	return nil
}