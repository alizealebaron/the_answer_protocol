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

	SetReward(item Item)
	SetStatus(s string)

	ToString()         string
}

/* +---------------------------------------------------------------------+ */
/* |                             QuestModel                              | */
/* +---------------------------------------------------------------------+ */

// === Constructeur === //

type QuestModel struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	DescriptionFr string `json:"descriptionFr"`
	DescriptionEn string `json:"descriptionEn"`
	RewardId      int    `json:"reward"`
	Reward        Item   `json:"-"`
	Quantity      int    `json:"quantity"`
	Status        string `json:"-"`
}

// === Accesseurs === //

func (q QuestModel) GetId()            int    { return q.Id            }
func (q QuestModel) GetTitle()         string { return q.Title         }
func (q QuestModel) GetDescriptionFr() string { return q.DescriptionFr }
func (q QuestModel) GetDescriptionEn() string { return q.DescriptionEn }
func (q QuestModel) GetReward()        Item   { return q.Reward        }
func (q QuestModel) GetRewardId()      int    { return q.RewardId      }

// === Modificateurs === //

func (q *QuestModel) SetReward(item Item)   {q.Reward = item}
func (q *QuestModel) SetStatus(s    string) {q.Status = s}

// === ToString === //

func (q QuestModel) ToString() string {
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

type QuestItem struct {
	QuestModel
	ItemNeededId int   `json:"ItemNeededId"`
	ItemNeededQu int   `json:"ItemNeededQu"`
}

func (q *QuestItem) SetReward(item Item) {q.Reward = item}

// === ToString === //

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
	MonsterNeededId int      `json:"MonsterNeededId"`
	MonsterNeededQu int      `json:"MonsterNeededQu"`
	MonsterSlay     int
}

func (q *QuestMonster) SetReward(item Item) {q.Reward = item}

// === ToString === //

func (q QuestMonster) ToString() string {
	b, err := json.Marshal(q)
	if err != nil {
		return fmt.Sprintf("erreur: %v", err)
	}
	return string(b)
}