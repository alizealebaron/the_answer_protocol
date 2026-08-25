/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* who.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/26 01:29:35 by emarette        #+#    #+#              */
/* Updated: 2026/08/26 01:36:10 by emarette        ###   ########.fr       */
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

func Who(args []string, tapManager *models.TapManager, player models.Player) error {
	nb_player := len(tapManager.Lst_Player)
	output := fmt.Sprintf("Ok players=%d\n", nb_player)
	server_write.ServerWrite(player.Conn, output)
	server_write.WriteLog(player.Conn, "INFO", "Player " + player.Name + "Send 'WHO'")
	return nil
}