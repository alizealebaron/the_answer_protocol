/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* room_model.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 15:56:01 by alebaron        #+#    #+#              */
/* Updated: 2026/08/25 16:16:53 by alebaron        ###   ########.fr       */
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
/* |                          Helpfull struct                            | */
/* +---------------------------------------------------------------------+ */

type NeighborRoom struct {
	North int  `json:"north"`
	South int  `json:"south"`
	East  int  `json:"east"`
	West  int  `json:"west"`
}

type FishingEntry struct {
	ItemID   int  `json:"itemId"`
	LootRate int  `json:"lootRate"`
}

/* +---------------------------------------------------------------------+ */
/* |                               Room                                  | */
/* +---------------------------------------------------------------------+ */

type Room struct {
	Id            int            `json:"id"`
	Name          string         `json:"name"`
	AlliesId      []int          `json:"allies"`
	EnnemiesId    []int          `json:"ennemies"`
	ItemsId       []int          `json:"items"`
	NeighborRoom  NeighborRoom   `json:"neighborRoom"`
	Fishing       []FishingEntry `json:"fishing"`
	Allies        []Npc          `json:"-"`
	Ennemies      []Npc          `json:"-"`
	Items         []Item         `json:"-"`
	Lst_Player    []Player       `json:"-"`
}

/* +---------------------------------------------------------------------+ */
/* |                                Get                                  | */
/* +---------------------------------------------------------------------+ */

func (r Room) GetId()   int    { return r.Id }

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func (r Room) ToString() string {
	b, err := json.Marshal(r)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

func (r *Room) AddPlayerToRoom(p Player) {
	r.Lst_Player = append(r.Lst_Player, p)
}

func (r *Room) RemovePlayerToRoom(player Player) {

    for i, p := range r.Lst_Player {
        if p.Id == player.Id {
            r.Lst_Player = append(r.Lst_Player[:i], r.Lst_Player[i+1:]...)
            return
        }
    } 
}