/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* who.go                                            :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/26 01:29:35 by emarette        #+#    #+#              */
/* Updated: 2026/09/02 15:37:59 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Who(args []string, tapManager *models.TapManager, player *models.Player) error {

	var lst_player_name []string

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// === Récupération des données === //
	nb_player := len(tapManager.Lst_Player)
	for _, entry := range room.Lst_Player {
		lst_player_name = append(lst_player_name, entry.Name)
	}

	roomJSON, _ := json.Marshal(lst_player_name)
	// === Envoie de la réponse === //
	output := fmt.Sprintf("OK { \"room\": %s, \"server\": %d }\n", roomJSON, nb_player)
	server_write.ServerWrite(player.Conn, output)
	server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+output)

	return nil
}
