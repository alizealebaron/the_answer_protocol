/* ************************************************************************ */
/*      _  _     ____                     ,~~.                              */
/*     | || |   |___  \             ,   (  ^ )>                             */
/*     | || |_    __) |             )\~~'   (       _      _      _         */
/*     |__   _|  / __/             (  .__)   )    >(.)__ <(^)__ =(o)__      */
/*        |_|   |_____| .fr         \_.____,*      (___/  (___/  (___/      */
/*                                                                          */
/* ************************************************************************ */
/* name   : room_model.go                                                   */
/* author : alebaron <alebaron@student.42.fr>                               */
/*                                                                          */
/* creation : Invalid date        by -----------                            */
/* update   : 2026/09/11 20:07:34 by alebaron                               */
/* ************************************************************************ */

package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                          Helpfull struct                            | */
/* +---------------------------------------------------------------------+ */

type NeighborRoom struct {
	North int `json:"north"`
	South int `json:"south"`
	East  int `json:"east"`
	West  int `json:"west"`
}

type FishingEntry struct {
	ItemID   int `json:"itemId"`
	LootRate int `json:"lootRate"`
}

type IdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type IdNameNpc struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Is_trader bool   `json:"is_trader"`
	Is_qg     bool   `json:"is_qg"`
}

/* +---------------------------------------------------------------------+ */
/* |                               Room                                  | */
/* +---------------------------------------------------------------------+ */

type Room struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	AlliesId     []int          `json:"allies"`
	EnnemiesId   []int          `json:"ennemies"`
	ItemsId      []int          `json:"items"`
	NeighborRoom NeighborRoom   `json:"neighborRoom"`
	Fishing      []FishingEntry `json:"fishing"`
	Type         string         `json:"type"`
	X            int            `json:"x"`
	Y            int            `json:"y"`

	Allies     []Npc      `json:"-"`
	Ennemies   []Monster  `json:"-"`
	Items      []Item     `json:"-"`
	Lst_Player []Player   `json:"-"`
	Arena      []*Monster `json:"-"`
}

/* +---------------------------------------------------------------------+ */
/* |                                Get                                  | */
/* +---------------------------------------------------------------------+ */

func (r Room) GetId() int { return r.Id }

func (r Room) GetNpc(id int) (*Npc, error) {
	for _, a := range r.Allies {
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

	// == Envoie à la room d'arrivée == //
	for _, pl := range r.Lst_Player {
		output := fmt.Sprintf("EVT ROOM PRESENCE ENTER %s\n", p.Name)
		server_write.ServerWrite(pl.Conn, output)
	}
}

func (r *Room) RemovePlayerToRoom(player Player) {

	for i, p := range r.Lst_Player {
		if p.Id == player.Id {
			r.Lst_Player = append(r.Lst_Player[:i], r.Lst_Player[i+1:]...)

			// == Envoie à la room de départ == //
			for _, pl := range r.Lst_Player {
				output := fmt.Sprintf("EVT ROOM PRESENCE LEAVE %s\n", p.Name)
				server_write.ServerWrite(pl.Conn, output)
			}
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

	// Ajout du monstre à la room
	r.Arena = append(r.Arena, &monster)

	// Envoie d'un évènement à tous les joueurs de la room
	for _, p := range r.Lst_Player {
		server_write.ServerWrite(p.Conn, "EVT ROOM MONSTER HAS SPAWN\n")
	}
}

func (r *Room) RemoveMonsterToRoom(monster Monster) (*Monster, error) {

	for i, m := range r.Arena {
		if m.Entity_id == monster.Entity_id {
			r.Arena = append(r.Arena[:i], r.Arena[i+1:]...)
			return m, nil
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

	toIdNameNpc := func(id int, name string, b1 bool, b2 bool) IdNameNpc {
		return IdNameNpc{Id: id, Name: name, Is_trader: b1, Is_qg: b2}
	}

	items := make([]IdName, 0, len(r.Items))
	for _, it := range r.Items {
		items = append(items, toIdName(it.GetId(), it.GetName()))
	}

	allies := make([]IdNameNpc, 0, len(r.Allies))
	for _, a := range r.Allies {
		_, ok1 := a.(Trader)
		_, ok2 := a.(QuestGiver)
		allies = append(allies, toIdNameNpc(a.GetId(), a.GetName(), ok1, ok2))
	}

	ennemies := make([]IdName, 0, len(r.Ennemies))
	for _, e := range r.Ennemies {
		ennemies = append(ennemies, toIdName(e.GetId(), e.GetName()))
	}

	arena := make([]IdName, 0, len(r.Arena))
	for _, e := range r.Arena {
		arena = append(arena, toIdName(e.Entity_id, e.GetName()))
	}

	out := struct {
		Id           int            `json:"id"`
		Name         string         `json:"name"`
		Allies       []IdNameNpc    `json:"allies"`
		Ennemies     []IdName       `json:"ennemies"`
		Items        []IdName       `json:"items"`
		NeighborRoom NeighborRoom   `json:"neighborRoom"`
		Fishing      []FishingEntry `json:"fishing"`
		Arena        []IdName       `json:"arena"`
	}{
		Id:           r.Id,
		Name:         r.Name,
		Allies:       allies,
		Ennemies:     ennemies,
		Items:        items,
		NeighborRoom: r.NeighborRoom,
		Fishing:      r.Fishing,
		Arena:        arena,
	}

	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}
