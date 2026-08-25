/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quit.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 12:44:36 by emarette        #+#    #+#              */
/* Updated: 2026/08/25 13:20:59 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package commands

import (
	"net"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/utils"
)

func Quit(tapManager *models.TapManager, conn net.Conn, player models.Player) {
	tapManager.RemovePlayer(player.Id)
	utils.ServerWrite(conn, "Ok bye\n")
}
