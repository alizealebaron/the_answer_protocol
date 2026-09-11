/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quest.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/08 14:48:34 by alebaron        #+#    #+#              */
/* Updated: 2026/09/08 14:53:39 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	// "fmt"
	"errors"
	"strconv"
	// "encoding/json"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Quest(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) != 1 {
		return errors.New("ERR 302 NO_PNJ_SEND")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// === Récupération de la quête du QuestGiver === //
	id, err := strconv.Atoi(args[0])

	quest, err := getGiverQuest(*room, id)
	if err != nil {
		return err
	}

	// === Ajout de la quête au joueur === //
	err = player.AddQuestToPlayer(quest)
	if err != nil {
		return err
	}

	// === Envoie des messages au client et dans les logs === //
	str_ret := "OK " + quest.ToStringQuest(*player) + "\n"
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + str_ret)

	return nil
}

func Quests(args []string, tapManager *models.TapManager, player *models.Player) error {

	format_lst_quest := getPrettyLstQuest(player.Lst_Quest, *player)

	// === Envoie des messages au client et dans les logs === //
	str_ret := "OK " + format_lst_quest + "\n"
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To " + player.Name + ": " + str_ret)

	return nil
}

/* +---------------------------------------------------------------------+ */
/* |                      Fonctions Supplémentaires                      | */
/* +---------------------------------------------------------------------+ */

func getPrettyLstQuest(lst_quest []models.Quest, p models.Player) string {

	ret_str := "["

	for i, q := range lst_quest {
		ret_str = ret_str + q.ToStringQuest(p)
		if i != len(lst_quest) - 1 {
			ret_str += ", "
		}
	}
	
	return ret_str + "]"
}

func getGiverQuest(room models.Room, npcId int) (models.Quest, error) {

	for _, npc := range room.Allies {
		if npc.GetId() == npcId {
			trader, ok := npc.(models.QuestGiver)
			if !ok {
				return nil, errors.New("ERR 406 NO_QUEST_AVAILABLE")
			}

			return trader.Quest, nil
		}
	}
	return nil, errors.New("ERR 404 NPC_NOT_FOUND")
}