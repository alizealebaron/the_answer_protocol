/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* trade.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/31 11:07:24 by alebaron        #+#    #+#              */
/* Updated: 2026/08/31 13:51:26 by alebaron        ###   ########.fr       */
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
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                              Structure                              | */
/* +---------------------------------------------------------------------+ */

type LittleInv struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

/* +---------------------------------------------------------------------+ */
/* |                               Trade                                 | */
/* +---------------------------------------------------------------------+ */

func Trade(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) != 1 {
		return errors.New("ERR 302 NO_PNJ_SEND")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// === Récupération de l'inventaire du Trader === //
	id, err := strconv.Atoi(args[0])

	inv, err := getTraderInventory(*room, id)
	if err != nil {
		return err
	}

	// === Envoie de l'inventaire === //

	toIdName := func(id int, name string, cost int) LittleInv {
		return LittleInv{Id: id, Name: name, Cost: cost}
	}

	items := make([]LittleInv, 0, len(inv))
	for _, it := range inv {
		items = append(items, toIdName(it.GetId(), it.GetName(), it.GetCost()))
	}

	inv_json, err := json.Marshal(items)
	if err != nil {
		return errors.New("ERR 666 ENCODAGE_ERROR")
	}

	output := "OK " + string(inv_json) + "\n"
	server_write.ServerWrite(player.Conn, output)
	server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+output)

	return nil
}

/* +---------------------------------------------------------------------+ */
/* |                                Buy                                  | */
/* +---------------------------------------------------------------------+ */

func Buy(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) < 1 {
		return errors.New("ERR 302 NO_PNJ_SEND")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// === Récupération de l'inventaire du Trader === //
	id, err := strconv.Atoi(args[0])

	inv, err := getTraderInventory(*room, id)
	if err != nil {
		return err
	}

	// === Vérification de la présence de l'item === //
	id, err = strconv.Atoi(args[1])

	item, err := getItemIventaire(inv, id)
	if err != nil {
		return err
	}

	// === Vérification que le joueur a l'argent pour === //

	quantite, err := strconv.Atoi(args[2])
	if quantite <= 0 {
		return errors.New("ERR 411 INVALID_QUANTITY")
	}

	prix := quantite * item.GetCost()
	if prix > player.Money {
		return errors.New("ERR 411 NOT_ENOUGH_MONEY")
	}

	// === Ajout de l'item à l'inventaire du joueur === //
	player.AddItemToPlayerWQuantity(item, quantite)
	player.Money -= quantite * item.GetCost()

	// === Envoie des messages === //

	str_ret := fmt.Sprintf("OK {\"id\":%d,\"name\":%s,\"quantity\":%d}\n", item.GetId(), item.GetName(), quantite)
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)

	return nil
}

/* +---------------------------------------------------------------------+ */
/* |                      Fonctions Supplémentaires                      | */
/* +---------------------------------------------------------------------+ */

func getTraderInventory(room models.Room, npcId int) ([]models.Item, error) {

	for _, npc := range room.Allies {
		if npc.GetId() == npcId {
			trader, ok := npc.(models.Trader)
			if !ok {
				return nil, errors.New("ERR 405 NPC_NOT_TRADER")
			}

			return trader.Inventory, nil
		}
	}
	return nil, errors.New("ERR 404 NPC_NOT_FOUND")

}

func getItemIventaire(inv []models.Item, item_id int) (models.Item, error) {

	for _, it := range inv {
		if it.GetId() == item_id {
			return it, nil
		}
	}
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")

}
