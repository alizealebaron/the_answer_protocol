/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* player_model.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: rruiz, alebaron, emarette                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 16:20:31 by alebaron        #+#    #+#              */
/* Updated: 2026/08/29 10:17:43 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import (
	"errors"
	"net"
	"fmt"
	"encoding/json"
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
	player := Player{totalPlayer, name, 100, 100, "healthy", 5, language, lstItem, conn, nil, dialogueProgress}
	totalPlayer += 1
	return player
}

/* +---------------------------------------------------------------------+ */
/* |                      Gestion de l'inventaire                        | */
/* +---------------------------------------------------------------------+ */

func (p *Player) GetItem(itID int) (*Item, error) {

    // Parcours des objets de l'inventaire
	for it := range p.Inventory {
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
	}{
		Items: items,
	}

	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
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
