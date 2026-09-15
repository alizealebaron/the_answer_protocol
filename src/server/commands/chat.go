/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* chat.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 22:16:01 by emarette        #+#    #+#              */
/* Updated: 2026/08/26 16:34:03 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"errors"
	"fmt"
	"strings"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Chat(args []string, tapManager *models.TapManager, player *models.Player) error {
	if len(args) != 2 {
		return errors.New("ERR ARGUMENTS_NOT_FOUND")
	}

	scope := args[0]
	message := strings.Join(args[1:], " ")
	if scope == "GLOBAL" {
		for _, p := range tapManager.Lst_Player {
			output := fmt.Sprintf("EVT GLOBAL CHAT %s %s\n", player.Name, message)
			server_write.ServerWrite(p.Conn, output)
			server_write.WriteLog(player.Conn, "CHAT", output)
		}
	} else if scope == "ROOM" {
		for _, room := range tapManager.Lst_Room {
			for _, p := range room.Lst_Player {
				if p.Id == player.Id {
					for _, p := range room.Lst_Player {
						output := fmt.Sprintf("EVT ROOM CHAT %s %s\n", player.Name, message)
						server_write.ServerWrite(p.Conn, output)
						server_write.WriteLog(player.Conn, "CHAT", output)
					}
				}
			}
		}
	} else if scope == "GROUP" {
		for _, p := range player.Group.Lst_Player {
			output := fmt.Sprintf("EVT GROUP CHAT %s %s\n", player.Name, message)
			server_write.ServerWrite(p.Conn, output)
			server_write.WriteLog(player.Conn, "CHAT", output)
		}
	}
	return nil
}
