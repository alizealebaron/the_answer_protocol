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

type Quest interface {
	GetId()            int
	GetTitle()         string
	GetDescriptionFr() string
	GetDescriptionEn() string
	GetReward()        Item
	GetRewardId()      int
	GetStatus()        string
	GetQuantity()      int

	SetReward(item Item)
	SetStatus(s string)

	UpdateProgress(quantity int)

	ToStringQuest(p Player) string
	ToString()              string
}

/* +---------------------------------------------------------------------+ */
/* |                             QuestModel                              | */
/* +---------------------------------------------------------------------+ */

// === Constructeur === //

type QuestModel struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	DescriptionFr  string `json:"descriptionFr"`
	DescriptionEn  string `json:"descriptionEn"`
	RewardId       int    `json:"reward"`
	Reward         Item   `json:"-"`
	Quantity       int    `json:"quantity"`
	SearchQuantity int    `json:"src_quantity"`
	Status         string `json:"-"`
	Progress       int    `json:"-"`
}

// === Accesseurs === //

func (q QuestModel) GetId()            int    { return q.Id            }
func (q QuestModel) GetTitle()         string { return q.Title         }
func (q QuestModel) GetDescriptionFr() string { return q.DescriptionFr }
func (q QuestModel) GetDescriptionEn() string { return q.DescriptionEn }
func (q QuestModel) GetReward()        Item   { return q.Reward        }
func (q QuestModel) GetRewardId()      int    { return q.RewardId      }
func (q QuestModel) GetStatus()        string { return q.Status        }
func (q QuestModel) GetQuantity()      int    { return q.Quantity      }

// === Modificateurs === //

func (q *QuestModel) SetReward(item   Item  ) {q.Reward   = item}
func (q *QuestModel) SetStatus(s      string) {q.Status   = s   }

// === ToString === //

func (q QuestModel) ToString() string {
	b, err := json.Marshal(q)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

func (q QuestModel) ToStringQuest(p Player) string {

	var desc string
	if p.Language == "FR" {
		desc = q.DescriptionFr
	} else {
		desc = q.DescriptionEn
	}

	progress := fmt.Sprintf("%d/%d", q.Progress, q.SearchQuantity)

	out := struct {
		Id           int            `json:"id"`
		Title        string         `json:"title"`
		Description  string         `json:"desc"`
		Reward       string         `json:"reward"`
		Quantity     int            `json:"quantity"`
		Progress     string         `json:"progress"`
		Status       string         `json:"status"`
	}{
		Id:           q.Id,
		Title:        q.Title,
		Description:  desc,
		Reward:       q.Reward.GetName(),
		Quantity:     q.Quantity,
		Progress:     progress,
		Status:       q.Status,
	}

	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

// Update Progress

func (q *QuestModel) UpdateProgress(quantity int) {

	// Mise à jour de la quantité récupérée pour la quête
	q.UpdateProgressBrut(q.Progress + quantity)

}

func (q *QuestModel) UpdateProgressBrut(quantity int) {

	if (q.Status != "rewarded") {
		q.Progress = quantity
	} else {
		return
	}

	// Vérification que la quantité est entre 0 et q.SearchQuantity
	if q.Progress > q.SearchQuantity {
		q.Progress = q.SearchQuantity
	}
	if q.Progress < 0 {
		q.Progress = 0
	}

	// Mise à jour du statut de la quête

	if q.Progress == 0 {
		q.Status = "active"
	} else if (q.Progress > 0 && q.Progress < q.SearchQuantity) {
		q.Status = "in progress"
	} else {
		q.Status = "completed"
	}
}

/* +---------------------------------------------------------------------+ */
/* |                             QuestItem                               | */
/* +---------------------------------------------------------------------+ */

// === Constructeur === //

type QuestItem struct {
	QuestModel
	ItemNeededId int   `json:"itemNeededId"`
}

func (q QuestItem) ToString() string {
	b, err := json.Marshal(q)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}

/* +---------------------------------------------------------------------+ */
/* |                             QuestItem                               | */
/* +---------------------------------------------------------------------+ */

// === Constructeur === //

type QuestMonster struct {
	QuestModel
	MonsterNeededId int      `json:"monsterNeededId"`
}

func (q QuestMonster) ToString() string {
	b, err := json.Marshal(q)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}