/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* player_model.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 16:20:31 by alebaron        #+#    #+#              */
/* Updated: 2026/08/25 22:54:33 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package models

import (
    "net"
)

/* +---------------------------------------------------------------------+ */
/* |                          Variable globale                           | */
/* +---------------------------------------------------------------------+ */

var totalPlayer int

/* +---------------------------------------------------------------------+ */
/* |                                Item                                 | */
/* +---------------------------------------------------------------------+ */

type Player struct {
	Id       int
    Name     string
    Pv       int
    Attack   int
    Language string
    Inventory []Item 
    Conn net.Conn
    Group string
}

func NewPlayer(name string, language string, conn net.Conn) Player {

    lstItem := []Item{}
    player := Player{totalPlayer, name, 100, 5, language, lstItem, conn, ""}
    totalPlayer += 1
    return player
}
