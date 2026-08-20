/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quest_model.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/20 08:58:22 by alebaron        #+#    #+#              */
/* Updated: 2026/08/20 09:58:17 by alebaron        ###   ########.fr       */
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

// func (q Quest) ToString() string {
// 	b, err := json.Marshal(q)
// 	if err != nil {
// 		return fmt.Sprintf("erreur: %v", err)
// 	}
// 	return string(b)
// }

// Debug to_string
func (q Quest) ToString() string {
    type QuestJSON struct {
        Id            int    `json:"id"`
        Title         string `json:"title"`
        DescriptionFr string `json:"descriptionFr"`
        DescriptionEn string `json:"descriptionEn"`
        RewardId      int    `json:"rewardId"`
        RewardName    string `json:"rewardName"`
        Quantity      int    `json:"quantity"`
    }

    rewardName := ""
    if q.Reward != nil {
        rewardName = q.Reward.GetName()
    }

    qj := QuestJSON{
        Id:            q.Id,
        Title:         q.Title,
        DescriptionFr: q.DescriptionFr,
        DescriptionEn: q.DescriptionEn,
        RewardId:      q.RewardId,
        RewardName:    rewardName,
        Quantity:      q.Quantity,
    }

    b, err := json.Marshal(qj)
    if err != nil {
        return fmt.Sprintf("erreur: %v", err)
    }
    return string(b)
}