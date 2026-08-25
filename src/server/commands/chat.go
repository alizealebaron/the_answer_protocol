/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* chat.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 22:16:01 by emarette        #+#    #+#              */
/* Updated: 2026/08/25 23:09:08 by emarette        ###   ########.fr       */
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

func Chat(args []string, tapManager *models.TapManager, player models.Player) error {
	scope := args[0]
	message := args[1]
	if scope == "GLOBAL" {
		for _ , p := range tapManager.Lst_Player {
			output := fmt.Sprintf("[Chat Global] %s: %s\n", player.Name, message)
			server_write.ServerWrite(p.Conn, output)
			server_write.WriteLog(player.Conn, "CHAT", output)
		} 
	} else if scope == "ROOM" {
		for _ , room := range tapManager.Lst_Room {
			for _ , p := range room.Lst_Player {
				if p.Id == player.Id {
					for _ , p := range room.Lst_Player {
						output := fmt.Sprintf("[Chat ROOM] %s: %s\n", player.Name, message)
						server_write.ServerWrite(p.Conn, output)
						server_write.WriteLog(player.Conn, "CHAT", output)
					}
				}
			}
		}
	} else if scope == "GROUP" {
		for _ , p := range tapManager.Lst_Player {
			if player.Group == p.Group {
				output := fmt.Sprintf("[Chat Group] %s: %s\n", player.Name, message)
				server_write.ServerWrite(p.Conn, output)
				server_write.WriteLog(player.Conn, "CHAT", output)
			}
		}
	}
	return nil
}