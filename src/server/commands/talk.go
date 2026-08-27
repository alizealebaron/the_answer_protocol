/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* talk.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/26 17:42:26 by alebaron        #+#    #+#              */
/* Updated: 2026/08/26 17:47:56 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"errors"
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Talk(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) != 1 {
		return errors.New("ERR 302 NO_NPC_SEND")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR PLAYER_NOT_FOUND_IN_ANY_ROOM")
	}

	// === Vérification de la présence du NPC dans la room === //
	id, err := strconv.Atoi(args[0])

	npc, err := room.GetNpc(id)
	if err != nil {
		return errors.New("ERR 404 NPC_NOT_FOUND")
	}

	// === Envoie du dialogue du NPC === //
	dialogue := player.GetNextDialogueLine(*npc)

	str_ret := "OK " + dialogue + "\n"
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + str_ret)

	return nil
}