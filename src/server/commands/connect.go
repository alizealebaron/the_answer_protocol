/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* connect.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 15:25:33 by emarette        #+#    #+#              */
/* Updated: 2026/08/25 13:13:04 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package commands

import (
	"net"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

func Connect(tapManager models.TapManager, conn net.Conn, name string, language string) (models.Player, string) {
	player := models.NewPlayer(name, language)
	for _, player := range tapManager.Lst_Player{
		if name == player.Name {
			server_write.ServerWrite(conn, "ERR 201 NAME_IN_USE\n")
			return player, "ERR 201 NAME_IN_USE"
		}
	}
	if language != "FR" && language != "EN" {
		server_write.ServerWrite(conn, "ERR 202 LANGUAGE_IN_USE\n")
		return player, "ERR 202 LANGUAGE_IN_USE"
	}

	tapManager.Lst_Player = append(tapManager.Lst_Player, player)
	tapManager.Lst_Room[1].Lst_Player = append(tapManager.Lst_Room[1].Lst_Player, player)
	server_write.ServerWrite(conn, "Ok connected\n")
	return player, ""
}