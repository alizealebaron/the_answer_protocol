/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* player_model.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: rruiz, alebaron, emarette                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 16:20:31 by alebaron        #+#    #+#              */
/* Updated: 2026/08/30 13:49:50 by emarette        ###   ########.fr       */
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
	Inventory        []Item
	Conn             net.Conn
	Group            *string
	DialogueProgress map[int]int
}

/* +---------------------------------------------------------------------+ */
/* |                            Constructeur                             | */
/* +---------------------------------------------------------------------+ */

func NewPlayer(name string, language string, conn net.Conn) Player {

	lstItem := []Item{}
	dialogueProgress := make(map[int]int)
	player := Player{totalPlayer, name, 100, 100, "healthy", 5, language, lstItem, conn, nil, dialogueProgress}
	totalPlayer += 1
	return player
}

/* +---------------------------------------------------------------------+ */
/* |                      Gestion de l'inventaire                        | */
/* +---------------------------------------------------------------------+ */

func (p *Player) AddItemToPlayer(it Item) {
	p.Inventory = append(p.Inventory, it)
}

func (p *Player) RemoveItemToPlayer(itID int) (*Item, error) {

	for i, it := range p.Inventory {
		if it.GetId() == itID {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			return &it, nil
		}
	}
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")
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

func (p *Player) InventoryToString() string {

    toIdName := func(id int, name string) IdName {
		return IdName{Id: id, Name: name}
	}

	items := make([]IdName, 0, len(p.Inventory))
	for _, it := range p.Inventory {
		items = append(items, toIdName(it.GetId(), it.GetName()))
	}

    out := struct {
		Items        []IdName       `json:"items"`
	}{
		Items:        items,
	}

    b, err := json.Marshal(out)
    if err != nil {
        return fmt.Sprintf("erreur: %v", err)
    }
    return string(b)
}