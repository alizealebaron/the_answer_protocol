/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* npc_models.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/20 11:03:02 by alebaron        #+#    #+#              */
/* Updated: 2026/08/20 18:04:54 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import (
	"encoding/json"
	"fmt"
)

/* +---------------------------------------------------------------------+ */
/* |                                NPC                                  | */
/* +---------------------------------------------------------------------+ */

type Npc interface {
	ToString()        string
	GetId()           int
	GetName()         string
}

/* +---------------------------------------------------------------------+ */
/* |                             Dialoguer                               | */
/* +---------------------------------------------------------------------+ */

type Dialoguer struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	DialogueFr []string `json:"dialogueFr"`
	DialogueEn []string `json:"dialogueEn"`
}

func (d Dialoguer) GetId()   int    { return d.Id   }
func (d Dialoguer) GetName() string { return d.Name }

func (d Dialoguer) ToString() string {
	b, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

/* +---------------------------------------------------------------------+ */
/* |                             QuestGiver                              | */
/* +---------------------------------------------------------------------+ */

type QuestGiver struct {
	Dialoguer
	QuestId       int      `json:"quest"`
	Quest         Quest    `json:"-"`
	DialogueFinFr []string `json:"dialogueFinFr"`
	DialogueFinEn []string `json:"dialogueFinEn"`
}

func (q QuestGiver) ToString() string {
	b, err := json.Marshal(q)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

/* +---------------------------------------------------------------------+ */
/* |                               Monster                               | */
/* +---------------------------------------------------------------------+ */

type Monster struct {
	Dialoguer
	Pv           int  `json:"pv"`
	Attack       int  `json:"attack"`
	Defense      int  `json:"Defense"`
	IsBoss       bool `json:"isBoss"`
	LootId       int  `json:"loot"`
	Loot         Item `json:"-"`
	QuantityMin  int  `json:"quantityMin"`
	QuantityMax  int  `json:"quantityMax"`
}

func (m Monster) ToString() string {
	b, err := json.Marshal(m)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

/* +---------------------------------------------------------------------+ */
/* |                                Trader                               | */
/* +---------------------------------------------------------------------+ */

type Trader struct {
	Dialoguer
	Inventory    []Item  `json:"-"`
	InventoryId  []int   `json:"inventory"`
}

func (t Trader) ToString() string {
	b, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}