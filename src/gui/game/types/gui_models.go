/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* gui_models.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/08 21:22:30 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 09:21:32 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package types

// The JSON returned by SECRET.
type Secret struct {
	Items       []ItemInfo    `json:"lst_item"`
	Players     []PlayerInfo  `json:"lst_player"`
	Quests      []QuestInfo   `json:"lst_quest"`
	Npcs        []NpcInfo     `json:"lst_npc"`
	Monsters    []MonsterInfo `json:"lst_monster"`
	Rooms       []RoomInfo    `json:"lst_room"`
	Groups      []GroupInfo   `json:"lst_group"`
	EntityIndex int           `json:"Entity_index"`
}

// Is a mirror of models.Item.
type ItemInfo struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	DescriptionFr string `json:"descriptionFr"`
	DescriptionEn string `json:"descriptionEn"`
	Cost          int    `json:"cost"`
	NbCopies      int    `json:"nbCopies"`
	NbAvail       int    `json:"nbAvail"`
	Damage        int    `json:"damage"`
	Effect        string `json:"effect"`
	Value         int    `json:"value"`
	Location      string `json:"location"`
}

// Is a mirror of models.Player.
type PlayerInfo struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	Pv       int    `json:"Pv"`
	MaxPv    int    `json:"MaxPv"`
	Status   string `json:"Status"`
	Attack   int    `json:"Attack"`
	Defense  int    `json:"Defense"`
	Language string `json:"Language"`
	Money    int    `json:"Money"`
}

// Is a mirror of models.Quest.
type QuestInfo struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	DescriptionFr   string `json:"descriptionFr"`
	DescriptionEn   string `json:"descriptionEn"`
	RewardId        int    `json:"reward"`
	Quantity        int    `json:"quantity"`
	SearchQuantity  int    `json:"src_quantity"`
	ItemNeededId    int    `json:"ItemNeededId"`
	ItemNeededQu    int    `json:"ItemNeededQu"`
	MonsterNeededId int    `json:"MonsterNeededId"`
	MonsterNeededQu int    `json:"MonsterNeededQu"`
	MonsterSlay     int    `json:"MonsterSlay"`
}

// Is a mirror of models.Npc.
type NpcInfo struct {
	Id            int      `json:"id"`
	Name          string   `json:"name"`
	DialogueFr    []string `json:"dialogueFr"`
	DialogueEn    []string `json:"dialogueEn"`
	QuestId       int      `json:"quest"`
	DialogueFinFr []string `json:"dialogueFinFr"`
	DialogueFinEn []string `json:"dialogueFinEn"`
	InventoryId   []int    `json:"inventory"`
}

// Is a mirror of models.Monster.
type MonsterInfo struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	DialogueFr  []string `json:"dialogueFr"`
	DialogueEn  []string `json:"dialogueEn"`
	Pv          int      `json:"pv"`
	Attack      int      `json:"attack"`
	Defense     int      `json:"Defense"`
	IsBoss      bool     `json:"isBoss"`
	LootId      int      `json:"loot"`
	QuantityMin int      `json:"quantityMin"`
	QuantityMax int      `json:"quantityMax"`
	SpawnRate   int      `json:"spawnRate"`
	EntityId    int      `json:"entity_id"`
}

// Is a mirror of models.NeighborRoom.
type NeighborRoomInfo struct {
	North int `json:"north"`
	South int `json:"south"`
	East  int `json:"east"`
	West  int `json:"west"`
}

// Is a mirror of models.FishingEntry.
type FishingEntryInfo struct {
	ItemId   int `json:"itemId"`
	LootRate int `json:"lootRate"`
}

// Is a mirror of models.Room.
type RoomInfo struct {
	Id           int                `json:"id"`
	Name         string             `json:"name"`
	AlliesId     []int              `json:"allies"`
	EnnemiesId   []int              `json:"ennemies"`
	ItemsId      []int              `json:"items"`
	NeighborRoom NeighborRoomInfo   `json:"neighborRoom"`
	Fishing      []FishingEntryInfo `json:"fishing"`
	Type         string             `json:"type"`
	X            int                `json:"x"`
	Y            int                `json:"y"`
}

// Is a mirror of models.IdName.
type IdNameInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Is a mirror of models.IdNameNpc (allies inside the LOOK reply).
type IdNameNpcInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Is_trader bool   `json:"is_trader"`
	Is_qg     bool   `json:"is_qg"`
}

type WhoInfo struct {
	RoomInfo   []string `json:"room"`
	ServerInfo int      `json:"server"`
}

// Is a mirror of the JSON returned by models.Room.ToString() (reply to LOOK).
type LookInfo struct {
	Id           int                `json:"id"`
	Name         string             `json:"name"`
	Allies       []IdNameNpcInfo    `json:"allies"`
	Ennemies     []IdNameInfo       `json:"ennemies"`
	Items        []IdNameInfo       `json:"items"`
	NeighborRoom NeighborRoomInfo   `json:"neighborRoom"`
	Fishing      []FishingEntryInfo `json:"fishing"`
	Arena        []IdNameInfo       `json:"arena"`
}

// Is a mirror of models.Group.
type GroupInfo struct {
	Id         int          `json:"Id"`
	LstPlayer  []PlayerInfo `json:"Lst_Player"`
	LstInvited []PlayerInfo `json:"Lst_Invited"`
}

// Is a mirror of the struct returned by models.Player.InventoryToString() (reply to INVENTORY).
type InventoryInfo struct {
	Items []InventoryItemInfo `json:"items"`
	Money int                 `json:"money"`
}

// Is a mirror of each item entry inside the INVENTORY reply.
type InventoryItemInfo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	Is_Usable bool   `json:"is_usable"`
	Is_Weapon bool   `json:"is_weapon"`
	Cost      int    `json:"cost"`
}

// Is a mirror of commands.LittleInv (reply to TRADE).
type TradeInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Cost int    `json:"cost"`
}

// Is a mirror of the JSON returned by the STATUS command.
type StatusInfo struct {
	Hp     int    `json:"hp"`
	MaxHp  int    `json:"max_hp"`
	Status string `json:"status"`
}
