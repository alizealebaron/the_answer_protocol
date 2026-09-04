/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* search.go                                         :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: rruiz, alebaron, emarette                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/30 13:19:42 by emarette        #+#    #+#              */
/* Updated: 2026/08/30 13:42:40 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Search(args []string, tapManager *models.TapManager, player *models.Player) error {

	// on verifie le nombre d'argument
	if len(args) != 1 {
		return errors.New("ERR 302 NO_ITEM_SEND")
	}

	// on cherche la room dans lequel se trouve le joueur
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR PLAYER_NOT_FOUND_IN_ANY_ROOM")
	}

	for _, e := range room.Ennemies {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("ERR ATOI_ERROR")
		}
		if id == e.Id {
			luck := rand.IntN(10) + 1

			if luck >= e.SpawnRate {
				e.Entity_id = tapManager.Entity_index
				tapManager.Entity_index += 1
				room.AddMonsterToRoom(e)
				message := fmt.Sprint("Ok ", e.Name,  "[", e.Entity_id, "] summon in the arena \n")
				server_write.ServerWrite(player.Conn, message)
				return nil
			}
			server_write.ServerWrite(player.Conn, "KO failed to summon monster in the arena"+"\n")
			server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + room.ToString())
			return nil
		}
	}

	return errors.New("ERR MOSTER_NOT_FOUND_IN_THIS_ROOM")

}
