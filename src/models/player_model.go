/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* player_model.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: rruiz, alebaron, emarette                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 16:20:31 by alebaron        #+#    #+#              */
/* Updated: 2026/09/04 15:15:04 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

/* +---------------------------------------------------------------------+ */
/* |                          Variable globale                           | */
/* +---------------------------------------------------------------------+ */

var totalPlayer int

/* +---------------------------------------------------------------------+ */
/* |                                Item                                 | */
/* +---------------------------------------------------------------------+ */

type Player struct {
	Id               int
	Name             string
	Pv               int
	MaxPv            int
	Status           string
	Attack           int
	Language         string
	Money            int
	Inventory        map[Item]int
	Conn             net.Conn
	Group            *Group
	DialogueProgress map[int]int
}

/* +---------------------------------------------------------------------+ */
/* |                            Constructeur                             | */
/* +---------------------------------------------------------------------+ */

func NewPlayer(name string, language string, conn net.Conn) Player {

	lstItem := make(map[Item]int)
	dialogueProgress := make(map[int]int)
	player := Player{totalPlayer, name, 1, 100, "healthy", 5, language, 10, lstItem, conn, nil, dialogueProgress}
	totalPlayer += 1
	return player
}

/* +---------------------------------------------------------------------+ */
/* |                      Gestion de l'inventaire                        | */
/* +---------------------------------------------------------------------+ */

func (p *Player) GetItem(itID int) (*Item, error) {

	// Parcours des objets de l'inventaire
	for it, _ := range p.Inventory {
		// Gestion des items si on le trouve
		if it.GetId() == itID {
			return &it, nil
		}
	}
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")
}

func (p *Player) AddItemToPlayer(it Item) {
	p.Inventory[it] += 1
}

func (p *Player) AddItemToPlayerWQuantity(it Item, q int) {
	p.Inventory[it] += q
}

func (p *Player) RemoveItemToPlayer(itID int) (*Item, error) {

	// Parcours des objets de l'inventaire
	for it, qty := range p.Inventory {

		// Gestion des items si on le trouve
		if it.GetId() == itID {
			if qty <= 1 {
				delete(p.Inventory, it)
			} else {
				p.Inventory[it] = qty - 1
			}
			itemCopy := it
			return &itemCopy, nil
		}
	}
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")
}

func (p *Player) RemoveItemToPlayerWQuantity(itID int, q int) (*Item, error) {

	// Parcours des objets de l'inventaire
	for it, qty := range p.Inventory {

		// Gestion des items si on le trouve
		if it.GetId() == itID {
			if qty == q {
				delete(p.Inventory, it)
			} else if qty < q {
				return nil, errors.New("ERR 420 NOT_ENOUGH_ITEM")
			} else {
				p.Inventory[it] = qty - q
			}
			itemCopy := it
			return &itemCopy, nil
		}
	}
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")
}

func (p *Player) InventoryToString() string {
	type ItemEntry struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	items := make([]ItemEntry, 0, len(p.Inventory))
	for it, qty := range p.Inventory {
		items = append(items, ItemEntry{
			Id:       it.GetId(),
			Name:     it.GetName(),
			Quantity: qty,
		})
	}

	out := struct {
		Items []ItemEntry `json:"items"`
		Money int         `json:"money"`
	}{
		Items: items,
		Money: p.Money,
	}

	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

func (p *Player) PlayerDeath(tapManager *TapManager) error {
	p.Pv = 30
	group, err := tapManager.GetGroupById(p.Id)
	if err == nil {
		group.RemovePlayerFromGroup(*p)
	}
	p_room, err := tapManager.FindPlayerRoom(p.Id)
	if err != nil {
		return err
	}
	p_room.RemovePlayerToRoom(*p)
	p_room, err = tapManager.GetRoomById(5)
	if err != nil {
		return err
	}
	p_room.AddPlayerToRoom(*p)
	return nil
}

func (p *Player) AddLifePoint(value int) error {
	p.Pv += value
	if p.Pv > 50 {
		p.Status = "healthy"
	} else if p.Pv <= 50 && p.Pv > 0 {
		p.Status = "bloody"
	} else {
		p.Status = "dead"
	}
	return nil
}

/* +---------------------------------------------------------------------+ */
/* |                        Gestion du gambling                          | */
/* +---------------------------------------------------------------------+ */

func (p *Player) IsEnoughGamblingCoin(bet int) (bool, error) {
	
	// Parcours des objets de l'inventaire
	for it := range p.Inventory {
		// Gestion des items si on le trouve
		if it.GetId() == 1 && p.Inventory[it] >= bet {
			return true, nil
		}
	}
	return false, errors.New("ERR 999 NOT_ENOUGH_COIN")
}

/* +---------------------------------------------------------------------+ */
/* |                       Gestion des dialogues                         | */
/* +---------------------------------------------------------------------+ */

func (p *Player) GetNextDialogueLine(npc Npc) (line string) {

	// Récupération de l'ID
	id := npc.GetId()

	// Récupération des dialogues français ou anglais
	lines := npc.GetDialogueFr()
	if p.Language == "EN" {
		lines = npc.GetDialogueEn()
	}

	// Récupération du dialogue à renvoyé
	idx := p.DialogueProgress[id]

	if idx >= len(lines) {
		p.DialogueProgress[id] = 0
		idx = 0
	}

	p.DialogueProgress[id] = idx + 1
	return lines[idx]
}
