/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* inventory.go                                      :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/27 10:29:28 by alebaron        #+#    #+#              */
/* Updated: 2026/08/27 10:38:08 by alebaron        ###   ########.fr       */
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

func Inventory(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Envoie de l'inventaire === //

	inventaire := player.InventoryToString()

	str_ret := "OK " + inventaire + "\n"
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + str_ret)

	return nil
}