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
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// === Vérification de la présence du NPC dans la room === //
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("904 WRONG_COMMAND_ARG")
	}

	npc, err := room.GetNpc(id)
	if err != nil {
		return errors.New("ERR 404 NPC_NOT_FOUND")
	}

	// === Envoie du dialogue du NPC === //

	// = Vérification qu'une quête n'est pas complétée = //

	isRewarded := false
	dialogue := ""

	quest_giver, ok := (*npc).(models.QuestGiver)
	if ok {
		isRewarded = player.IsNpcQuestCompleted(quest_giver)
	}

	if (!isRewarded) {
		dialogue = player.GetNextDialogueLine(*npc)
	} else {
		if player.Language == "FR" {
			dialogue = quest_giver.DialogueFinFr
		} else {
			dialogue = quest_giver.DialogueFinEn
		}
	}

	str_ret := "OK " + dialogue + "\n"
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + str_ret)

	return nil
}