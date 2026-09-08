/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* group_model.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/29 09:04:12 by alebaron        #+#    #+#              */
/* Updated: 2026/08/29 13:58:08 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import (
	"errors"
	"fmt"
)

/* +---------------------------------------------------------------------+ */
/* |                          Variable globale                           | */
/* +---------------------------------------------------------------------+ */

var totalGroup int

/* +---------------------------------------------------------------------+ */
/* |                               Group                                 | */
/* +---------------------------------------------------------------------+ */

type Group struct {
	Id            int
	Lst_Player    []*Player
	Lst_Invited   []Player
}

/* +---------------------------------------------------------------------+ */
/* |                            Constructeur                             | */
/* +---------------------------------------------------------------------+ */

func NewGroup(player *Player) Group {

	lst_inv := []Player{}
	lst_player := []*Player{player}
	group := Group{totalGroup, lst_player, lst_inv}
	totalGroup += 1
	return group
}

/* +---------------------------------------------------------------------+ */
/* |                         Gestion des joueurs                         | */
/* +---------------------------------------------------------------------+ */

func (g *Group) AddPlayerToGroup(player *Player) error {

	g.Lst_Player = append(g.Lst_Player, player)
	g.RemovePlayerFromInvit(*player)

	return nil
}

func (g *Group) RemovePlayerFromGroup(player Player) {

    for i, p := range g.Lst_Player {
        if p.Id == player.Id {
            g.Lst_Player = append(g.Lst_Player[:i], g.Lst_Player[i+1:]...)
            return
        }
    } 

}

/* +---------------------------------------------------------------------+ */
/* |                       Gestion des invitations                       | */
/* +---------------------------------------------------------------------+ */

func (g *Group) IsPlayerInvited(player Player) bool {

	// Parcours des joueurs invités
	for _, p := range g.Lst_Invited {
        // Gestion des items si on le trouve
		if p.Id == player.Id {
			return true
		}
	}

	return false
}

func (g *Group) IsPlayerInGroup(player Player) bool {

	// Parcours des joueurs dans le groupe
	for _, p := range g.Lst_Player {
        // Gestion des items si on le trouve
		if p.Id == player.Id {
			return true
		}
	}

	return false
}

func (g *Group) AddPlayerToInvited(player Player) error {
	
	if g.IsPlayerInvited(player) {
		return errors.New("ERR 409 PLAYER_ALREADY_INVITED")
	}

	if g.IsPlayerInGroup(player) {
		return errors.New("ERR 402 ALREADY_IN_GROUP")
	}

	g.Lst_Invited = append(g.Lst_Invited, player)
	fmt.Println(g.Lst_Invited)
	return nil
}

func (g *Group) RemovePlayerFromInvit(player Player) {

    for i, p := range g.Lst_Invited {
        if p.Id == player.Id {
            g.Lst_Invited = append(g.Lst_Invited[:i], g.Lst_Invited[i+1:]...)
            return
        }
    } 

}