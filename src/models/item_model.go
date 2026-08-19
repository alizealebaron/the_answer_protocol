/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* item_model.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 11:04:48 by alebaron        #+#    #+#              */
/* Updated: 2026/08/19 14:40:18 by alebaron        ###   ########.fr       */
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
/* |                                Item                                 | */
/* +---------------------------------------------------------------------+ */

type Item interface {
	ToString() string
}

/* +---------------------------------------------------------------------+ */
/* |                                Loot                                 | */
/* +---------------------------------------------------------------------+ */

type Loot struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	DescriptionFr string `json:"descriptionFr"`
	DescriptionEn string `json:"descriptionEn"`
	Cost          int    `json:"cost"`
	NbCopies      int    `json:"nbCopies"`
	NbAvail       int    `json:"nbAvail"`
}

func (p Loot) ToString() string {
	b, err := json.Marshal(p)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

/* +---------------------------------------------------------------------+ */
/* |                                Weapon                               | */
/* +---------------------------------------------------------------------+ */

type Weapon struct {
	Loot
	Damage int `json:"damage"`
}

func (w Weapon) ToString() string {
	b, err := json.Marshal(w)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

/* +---------------------------------------------------------------------+ */
/* |                                Edible                               | */
/* +---------------------------------------------------------------------+ */

type Edible struct {
	Loot
	Effect string `json:"effect"`
	Value  int    `json:"value"`
}

func (e Edible) ToString() string {
	b, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}