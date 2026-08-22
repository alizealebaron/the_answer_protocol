/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* connect.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 15:25:33 by emarette        #+#    #+#              */
/* Updated: 2026/08/22 14:50:01 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package commands

import (
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/utils"
	"net"
)

func Connect(tapManager models.TapManager,conn net.Conn, name string, language string) (models.Player, string) {
	player := models.NewPlayer(name, language)
	for _, player := range tapManager.Lst_Player{
		if name == player.Name {
			utils.ServerWrite(conn, "ERR 201 NAME_IN_USE\n")
			return player, "ERR 201 NAME_IN_USE"
		}
	}
	if language != "FR" && language != "EN" {
		utils.ServerWrite(conn, "ERR 202 LANGUAGE_IN_USE\n")
		return player, "ERR 202 LANGUAGE_IN_USE"
	}

	tapManager.Lst_Player = append(tapManager.Lst_Player, player)
	utils.ServerWrite(conn, "Ok connected\n")
	return player, ""
}