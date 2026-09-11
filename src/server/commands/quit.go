/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quit.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 12:44:36 by emarette        #+#    #+#              */
/* Updated: 2026/08/25 14:38:41 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Quit(tapManager *models.TapManager, player models.Player) {

	// === Récupération de la room actuelle du Joueur === //
	room, _ := tapManager.FindPlayerRoom(player.Id)

	// === Suppression du joueur de partout === //
	room.RemovePlayerToRoom(player)
	tapManager.RemovePlayer(player.Id)
	server_write.ServerWrite(player.Conn, "OK bye\n")

	// === Envoie de l'évènement de compte de joueur === //
	for _, p := range tapManager.Lst_Player {
		server_write.ServerWrite(p.Conn, "EVT STATS players="+strconv.Itoa(len(tapManager.Lst_Player))+"\n")
	}

	server_write.WriteLog(player.Conn, "INFO", "Player "+player.Name+" disconnected")
}
