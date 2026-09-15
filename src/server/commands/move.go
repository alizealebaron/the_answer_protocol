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

	// === Gestion des erreurs potentielles === //

	if len(args) != 1 {
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	room, err := tapManager.FindPlayerRoom(player.Id)

	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}
	
	// === Gestion de la direction du joueurs === //

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
		return errors.New("ERR 408 DIRECTION_INCORRECT")
    }

	q, _ := player.GetQuantityItem(1)
	if room.Name == "CASINO" && q >= 1000 {
		id_nei_room = 15
	}

	if id_nei_room == 0 {
		return errors.New("ERR 301 NO_EXIT")
	} else {

		nei_room, err := tapManager.GetRoomById(id_nei_room)
		if err != nil {
			return errors.New("ERR 404 ROOM_NOT_FOUND")
		}

		// === Ajout dans les nouvelles rooms et envoie du message === //
	
		room.RemovePlayerToRoom(*player)
		nei_room.AddPlayerToRoom(*player)
		server_write.ServerWrite(player.Conn, "OK " + nei_room.Name + "\n")
		server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": OK " + nei_room.Name)
	}

	return nil
}