/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* player_model.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 16:20:31 by alebaron        #+#    #+#              */
/* Updated: 2026/08/26 10:48:03 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import (
    "net"
    "errors"
)

/* +---------------------------------------------------------------------+ */
/* |                          Variable globale                           | */
/* +---------------------------------------------------------------------+ */

var totalPlayer int

/* +---------------------------------------------------------------------+ */
/* |                                Item                                 | */
/* +---------------------------------------------------------------------+ */

type Player struct {
	Id        int
    Name      string
    Pv        int
    MaxPv     int
    Status    string
    Attack    int
    Language  string
    Inventory []Item 
    Conn      net.Conn
    Group     string
}

/* +---------------------------------------------------------------------+ */
/* |                            Constructeur                             | */
/* +---------------------------------------------------------------------+ */

func NewPlayer(name string, language string, conn net.Conn) Player {

    lstItem := []Item{}
    player := Player{totalPlayer, name, 100, 100, "healthy", 5, language, lstItem, conn, ""}
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
	return nil, errors.New("ERR 303 ITEM_NOT_FOUND")
}