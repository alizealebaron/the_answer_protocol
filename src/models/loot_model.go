/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* loot_model.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 11:04:48 by alebaron        #+#    #+#              */
/* Updated: 2026/08/19 12:12:44 by alebaron        ###   ########.fr       */
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
/* |                                Main                                 | */
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