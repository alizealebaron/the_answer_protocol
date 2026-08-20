/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quest_model.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/20 08:58:22 by alebaron        #+#    #+#              */
/* Updated: 2026/08/20 11:51:31 by alebaron        ###   ########.fr       */
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
/* |                               Quest                                 | */
/* +---------------------------------------------------------------------+ */

type Quest struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	DescriptionFr string `json:"descriptionFr"`
	DescriptionEn string `json:"descriptionEn"`
	RewardId      int    `json:"reward"`
	Reward        Item   `json:"-"`
	Quantity      int    `json:"quantity"`
}

func (q Quest) GetId()   int    { return q.Id   }

func (q Quest) ToString() string {
	b, err := json.Marshal(q)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}