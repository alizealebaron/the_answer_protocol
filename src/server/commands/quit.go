/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quit.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 12:44:36 by emarette        #+#    #+#              */
/* Updated: 2026/08/25 22:25:38 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package commands

import (
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

func Quit(tapManager *models.TapManager, player models.Player) {
	tapManager.RemovePlayer(player.Id)
	server_write.ServerWrite(player.Conn, "Ok bye\n")
}
