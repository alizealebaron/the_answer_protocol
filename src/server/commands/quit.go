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

package commands

import (
	"net"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

func Quit(tapManager *models.TapManager, conn net.Conn, player models.Player) {
	tapManager.RemovePlayer(player.Id)
	server_write.ServerWrite(conn, "Ok bye\n")
	server_write.WriteLog(conn, "INFO", "Player " + player.Name + " disconnected")
}
