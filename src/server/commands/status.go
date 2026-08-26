/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* status.go                                         :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/26 17:19:50 by alebaron        #+#    #+#              */
/* Updated: 2026/08/26 17:25:09 by alebaron        ###   ########.fr       */
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

func Status(args []string, tapManager *models.TapManager, player *models.Player) error {

	output := "OK {\"hp\": " + strconv.Itoa(player.Pv) + ", \"max_hp\": " + strconv.Itoa(player.MaxPv) + ", \"status\": " + player.Status + "}\n"
	server_write.ServerWrite(player.Conn, output)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + output)

	return nil
}