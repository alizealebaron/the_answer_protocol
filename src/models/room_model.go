/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* room_model.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: rruiz, alebaron, emarette                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 15:56:01 by alebaron        #+#    #+#              */
/* Updated: 2026/08/30 13:35:37 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import (
	"encoding/json"
	"fmt"
	"errors"
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

type IdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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
	Ennemies      []Monster      `json:"-"`
	Items         []Item         `json:"-"`
	Lst_Player    []Player       `json:"-"`
	Arena		  []Monster
}

/* +---------------------------------------------------------------------+ */
/* |                                Get                                  | */
/* +---------------------------------------------------------------------+ */

func (r Room) GetId()   int    { return r.Id }

func (r Room) GetNpc(id int) (*Npc, error) {
	for _ , a := range r.Allies {
		if a.GetId() == id {
			return &a, nil
		}
	}
	return nil, errors.New("ERR 404 NPC_NOT_FOUND")
}

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

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

func (r *Room) AddItemToRoom(it Item) {
	r.Items = append(r.Items, it)
}

func (r *Room) RemoveItemToRoom(itID int) (*Item, error) {

    for i, p := range r.Items {
        if p.GetId() == itID {
            r.Items = append(r.Items[:i], r.Items[i+1:]...)
            return &p, nil
        }
    }
	return nil, errors.New("ERR 404 ITEM_NOT_FOUND")
}

func (r *Room) AddMonsterToRoom(monster Monster) {
	r.Arena = append(r.Arena, monster)
}

func (r *Room) RemoveMonsterToRoom(monster Monster) (*Monster, error) {

    for i, m := range r.Arena {
        if m.GetId() == monster.GetId() {
            r.Arena = append(r.Arena[:i], r.Arena[i+1:]...)
            return &m, nil
        }
    }
	return nil, errors.New("ERR 404 MONSTER_NOT_FOUND")
}

/* +---------------------------------------------------------------------+ */
/* |                             To_string                               | */
/* +---------------------------------------------------------------------+ */

func (r Room) ToString() string {

	toIdName := func(id int, name string) IdName {
		return IdName{Id: id, Name: name}
	}

	items := make([]IdName, 0, len(r.Items))
	for _, it := range r.Items {
		items = append(items, toIdName(it.GetId(), it.GetName()))
	}

	allies := make([]IdName, 0, len(r.Allies))
	for _, a := range r.Allies {
		allies = append(allies, toIdName(a.GetId(), a.GetName()))
	}

	ennemies := make([]IdName, 0, len(r.Ennemies))
	for _, e := range r.Ennemies {
		ennemies = append(ennemies, toIdName(e.GetId(), e.GetName()))
	}

	out := struct {
		Id           int            `json:"id"`
		Name         string         `json:"name"`
		Allies       []IdName       `json:"allies"`
		Ennemies     []IdName       `json:"ennemies"`
		Items        []IdName       `json:"items"`
		NeighborRoom NeighborRoom   `json:"neighborRoom"`
		Fishing      []FishingEntry `json:"fishing"`
	}{
		Id:           r.Id,
		Name:         r.Name,
		Allies:       allies,
		Ennemies:     ennemies,
		Items:        items,
		NeighborRoom: r.NeighborRoom,
		Fishing:      r.Fishing,
	}

	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}