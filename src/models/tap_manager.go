/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* tap_manager.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/18 19:29:32 by alebaron        #+#    #+#              */
/* Updated: 2026/08/20 11:35:20 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

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