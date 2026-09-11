/* ************************************************************************ */
/*      _  _     ____                     ,~~.                              */
/*     | || |   |___  \             ,   (  ^ )>                             */
/*     | || |_    __) |             )\~~'   (       _      _      _         */
/*     |__   _|  / __/             (  .__)   )    >(.)__ <(^)__ =(o)__      */
/*        |_|   |_____| .fr         \_.____,*      (___/  (___/  (___/      */
/*                                                                          */
/* ************************************************************************ */
/* name   : item_model.go                                                   */
/* author : alebaron <alebaron@student.42.fr>                               */
/*                                                                          */
/* creation : Invalid date        by -----------                            */
/* update   : 2026/09/11 20:02:04 by alebaron                               */
/* ************************************************************************ */


package models

import (
	"encoding/json"
	"fmt"
)

/* +---------------------------------------------------------------------+ */
/* |                                Item                                 | */
/* +---------------------------------------------------------------------+ */

type Item interface {
	ToString()        string
	GetId()           int
	GetName()         string
	IsItemAvailable() bool
	GetCost()         int
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

func (l Loot) GetId()   int    { return l.Id   }
func (l Loot) GetName() string { return l.Name }
func (l Loot) GetCost() int    { return l.Cost }

func (l Loot) IsItemAvailable() bool {
	return (l.NbAvail > 0)
}

func (l Loot) ToString() string {
	b, err := json.Marshal(l)
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

/* +---------------------------------------------------------------------+ */
/* |                                Usable                               | */
/* +---------------------------------------------------------------------+ */

type Usable struct {
	Loot
	Location string `json:"location"`
}

func (u Usable) ToString() string {
	b, err := json.Marshal(u)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}