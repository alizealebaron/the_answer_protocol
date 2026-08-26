/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* move.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 14:26:01 by alebaron        #+#    #+#              */
/* Updated: 2026/08/26 10:59:33 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
    // "fmt"
	"errors"
	"strings"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Move(args []string, tapManager *models.TapManager, player *models.Player) error {

	if len(args) != 1 {
		return errors.New("ERR 302 NO_DIRECTION_SEND")
	}

	room, err := tapManager.FindPlayerRoom(player.Id)

	if err != nil {
		return errors.New("ERR PLAYER_NOT_FOUND_IN_ANY_ROOM")
	}

	direction := strings.ToLower(args[0])
	id_nei_room := 0

	switch direction {
    case "north":
        id_nei_room = room.NeighborRoom.North
    case "east":
        id_nei_room = room.NeighborRoom.East
    case "south":
        id_nei_room = room.NeighborRoom.South
	case "west":
        id_nei_room = room.NeighborRoom.West
	default:
		return errors.New("ERR 303 DIRECTION_INCORRECT")
    }

	if id_nei_room == 0 {
		return errors.New("ERR 301 NO_EXIT")
	} else {

		nei_room, err := tapManager.GetRoomById(id_nei_room)
		if err != nil {
			return errors.New("ERR ROOM_NOT_FOUND")
		}
	
		room.RemovePlayerToRoom(*player)
		nei_room.AddPlayerToRoom(*player)
		server_write.ServerWrite(player.Conn, "OK " + nei_room.Name + "\n")
		server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": OK " + nei_room.Name)
	}

	return nil
}