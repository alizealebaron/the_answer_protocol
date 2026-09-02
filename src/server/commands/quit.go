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
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Quit(tapManager *models.TapManager, player models.Player) {
	tapManager.RemovePlayer(player.Id)
	server_write.ServerWrite(player.Conn, "Ok bye\n")
	server_write.WriteLog(player.Conn, "INFO", "Player " + player.Name + " disconnected")
}
