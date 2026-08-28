/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* tap_manager.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/18 19:29:32 by alebaron        #+#    #+#              */
/* Updated: 2026/08/28 16:07:32 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import "errors"

/* +---------------------------------------------------------------------+ */
/* |                               Classe                                | */
/* +---------------------------------------------------------------------+ */

type TapManager struct {
    Lst_item   []Item
    Lst_Player []Player
    Lst_Quest  []Quest
    Lst_Npc    []Npc
    Lst_Room   []Room
}

/* +---------------------------------------------------------------------+ */
/* |                            Constructeur                             | */
/* +---------------------------------------------------------------------+ */

func NewTapManager(Lst_item []Item, Lst_Quest []Quest, Lst_Npc []Npc, Lst_Room []Room) TapManager {

    Lst_Player := []Player{}
    tap := TapManager{Lst_item, Lst_Player, Lst_Quest, Lst_Npc, Lst_Room}
    return tap
}

/* +---------------------------------------------------------------------+ */
/* |                             Accesseurs                              | */
/* +---------------------------------------------------------------------+ */

func (tm *TapManager) GetRoomById(id int) (*Room, error) {

	for i := range tm.Lst_Room {
		if tm.Lst_Room[i].Id == id {
			return &tm.Lst_Room[i], nil
		}
	}

	return nil, errors.New("not found")
}

func (tm *TapManager) GetItemById(id int) (Item, error) {

	for i := range tm.Lst_item {
		if tm.Lst_item[i].GetId() == id {
			return tm.Lst_item[i], nil
		}
	}

	return nil, errors.New("not found")
}

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */


func (tap TapManager) ToString() string {
    var tap_str string
    
    for _, item := range tap.Lst_item {
		tap_str += item.ToString()
	}

	tap_str += "\n\n"

    for _, item := range tap.Lst_Quest {
		tap_str += item.ToString()
	}

    tap_str += "\n\n"

    for _, item := range tap.Lst_Npc {
		tap_str += item.ToString()
	}

    return tap_str
}

func (tap TapManager) RemovePlayer(player_id int) {
    for i, p := range tap.Lst_Player {
        if p.Id == player_id {
            tap.Lst_Player = append(tap.Lst_Player[:i], tap.Lst_Player[i+1:]...)
            return
        }
    } 
}

func (tap *TapManager) FindPlayerRoom(player_id int) (*Room, error) {

	for i := range tap.Lst_Room {
		for _, p := range tap.Lst_Room[i].Lst_Player {
			if p.Id == player_id {
				return &tap.Lst_Room[i], nil
			}
		}
	}

	return nil, errors.New("ERR PLAYER_NOT_FOUND_IN_ANY_ROOM")
}
