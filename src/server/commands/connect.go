/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* connect.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 15:25:33 by emarette        #+#    #+#              */
/* Updated: 2026/08/25 22:24:42 by emarette        ###   ########.fr       */
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

func Connect(tapManager *models.TapManager, conn net.Conn, name string, language string) (models.Player, string) {
	player := models.NewPlayer(name, language, conn)
	for _, player := range tapManager.Lst_Player {
		if name == player.Name {
			server_write.ServerWrite(conn, "ERR 201 NAME_IN_USE\n")
			return player, "ERR 201 NAME_IN_USE"
		}
	}
	if language != "FR" && language != "EN" {
		server_write.ServerWrite(conn, "ERR 202 INCORRECT_LANGUAGE\n")
		return player, "ERR 202 INCORRECT_LANGUAGE"
	}

	tapManager.Lst_Player = append(tapManager.Lst_Player, player)
	server_write.ServerWrite(conn, "OK connected\n")
	tapManager.Lst_Room[1].AddPlayerToRoom(player)

	return player, ""
}
