/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* who.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/26 01:29:35 by emarette        #+#    #+#              */
/* Updated: 2026/09/02 15:37:59 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"fmt"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Who(args []string, tapManager *models.TapManager, player *models.Player) error {
	nb_player := len(tapManager.Lst_Player)
	output := fmt.Sprintf("Ok players=%d\n", nb_player)
	server_write.ServerWrite(player.Conn, output)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + output)

	return nil
}