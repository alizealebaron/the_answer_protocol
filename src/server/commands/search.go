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
	"math/rand/v2"
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
		if args[0] == e.GetName() {
			luck := rand.IntN(10)
			if luck >= 5 {
				room.AddMonsterToRoom(e)
				server_write.ServerWrite(player.Conn, "OK "+e.GetName()+" summon in the arena"+"\n")
				return nil
			}
			server_write.ServerWrite(player.Conn, "KO failed to summon in the arena"+"\n")
			// server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + room.ToString())
			return nil
		}
	}

	return errors.New("ERR MOSTER_NOT_FOUND_IN_THIS_ROOM")

}
