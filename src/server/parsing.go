/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* parsing.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/18 19:28:32 by alebaron        #+#    #+#              */
/* Updated: 2026/08/20 18:06:07 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                            Package & Import                             */
/* ----------------------------------------------------------------------- */

package server

import (
    "os"
    "fmt"
    "encoding/json"
    "the_answer_protocol/src/utils"
    "the_answer_protocol/src/models"
)

/* ----------------------------------------------------------------------- */
/*                                  Main                                   */
/* ----------------------------------------------------------------------- */

func ParseJSONFile() models.TapManager{

    // === Initialisation des Items ===

    var lst_item []models.Item
    lst_item = get_all_item()

    // === Initialisation des Quests ===

    var lst_quest []models.Quest
    lst_quest = get_all_quest()
    resolve_quest_rewards(lst_quest, lst_item)

    // === Initialisation des NPCs ===

    var lst_npc []models.Npc
    lst_npc = get_all_npc(lst_item, lst_quest)

    // === Initialisation des Rooms ===

    var lst_room []models.Room
    lst_room = get_all_room(lst_item, lst_npc)

    // === Initialisation du tapManager ===
    tapManager := models.NewTapManager(lst_item, lst_quest, lst_npc, lst_room)
	fmt.Println("[\033[32mSUCCESS\033[0m] (ﾉ◕ヮ◕)ﾉ*:・ﾟ✧ JSON successfully load.")

    return tapManager
}

/* ----------------------------------------------------------------------- */
/*                         Méthodes de récupération                        */
/* ----------------------------------------------------------------------- */

func get_all_item() []models.Item {

    // === Déclarations des variables === //

    var lst_item   []models.Item
    var lst_loot   []models.Loot
    var lst_weapon []models.Weapon
    var lst_edible []models.Edible
    var lst_usable []models.Usable

    // === Récupérations des loots === //

    data := get_data_from_json("data/loot_data.json")

    err := json.Unmarshal(data, &lst_loot)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    for _, l := range lst_loot {
        lst_item = append(lst_item, l)
    }

    // === Récupérations des weapons === //

    data = get_data_from_json("data/weapon_data.json")

    err = json.Unmarshal(data, &lst_weapon)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    for _, l := range lst_weapon {
        lst_item = append(lst_item, l)
    }

    // === Récupérations des edible === //

    data = get_data_from_json("data/edible_data.json")

    err = json.Unmarshal(data, &lst_edible)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    for _, l := range lst_edible {
        lst_item = append(lst_item, l)
    }

    // === Récupérations des usables === //

    data = get_data_from_json("data/usable_data.json")

    err = json.Unmarshal(data, &lst_usable)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    for _, l := range lst_usable {
        lst_item = append(lst_item, l)
    }

    // === Renvoie des données récupérées === //

    return lst_item
}

func get_all_quest() []models.Quest {

    // === Déclarations des variables === //

    var lst_quest []models.Quest

    // === Récupération des quêtes === //

    data := get_data_from_json("data/quest_data.json")

    err := json.Unmarshal(data, &lst_quest)
    if err != nil {
        utils.ExitError("JSONParsingError", err)
    }

    // === Renvoie des données récupérées === //

    return lst_quest
}

func get_all_npc(lst_item []models.Item, lst_quest []models.Quest) []models.Npc {

    // === Déclarations des variables === //

    var lst_npc         []models.Npc
    var lst_dialoguer   []models.Dialoguer
    var lst_questgiver  []models.QuestGiver
    var lst_monster     []models.Monster
    var lst_room      []models.Trader

    // === Récupérations des loots === //

    data := get_data_from_json("data/dialoguer_data.json")

    err := json.Unmarshal(data, &lst_dialoguer)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    for _, d := range lst_dialoguer {
        lst_npc = append(lst_npc, d)
    }

    // === Récupérations des questgiver === //

    data = get_data_from_json("data/questgiver_data.json")

    err = json.Unmarshal(data, &lst_questgiver)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    resolve_questgiver_quests(lst_questgiver, lst_quest)
    for _, q := range lst_questgiver {
        lst_npc = append(lst_npc, q)
    }

    // === Récupérations des monsters === //

    data = get_data_from_json("data/monster_data.json")

    err = json.Unmarshal(data, &lst_monster)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    resolve_monster_item(lst_monster, lst_item)
    for _, m := range lst_monster {
        lst_npc = append(lst_npc, m)
    }

    // // === Récupérations des traders === //

    data = get_data_from_json("data/trader_data.json")

    err = json.Unmarshal(data, &lst_room)
	if err != nil {
		utils.ExitError("JSONParsingError", err)
	}

    resolve_trader_item(lst_room, lst_item)
    for _, t := range lst_room {
        lst_npc = append(lst_npc, t)
    }

    // === Renvoie des données récupérées === //

    return lst_npc
}

func get_all_room(lst_item []models.Item, lst_npc []models.Npc) []models.Room {

    // === Déclarations des variables === //

    var lst_room []models.Room

    // === Récupération des quêtes === //

    data := get_data_from_json("data/room_data.json")

    err := json.Unmarshal(data, &lst_room)
    if err != nil {
        utils.ExitError("JSONParsingError", err)
    }

    resolve_room_item(lst_room, lst_item)
    resolve_room_allies(lst_room, lst_npc)
    resolve_room_ennemy(lst_room, lst_npc)

    // === Renvoie des données récupérées === //

    return lst_room
}

/* ----------------------------------------------------------------------- */
/*                      Méthodes de liaison des objets                     */
/* ----------------------------------------------------------------------- */

func resolve_room_item(lst_room []models.Room, lst_item []models.Item) []models.Room {

    // On construit une map pour un accès rapide O(1)
    itemById := make(map[int]models.Item)
    for _, it := range lst_item {
        itemById[it.GetId()] = it
    }

    for i := range lst_room {
        for j := range lst_room[i].ItemsId {
            if item, ok := itemById[lst_room[i].ItemsId[j]]; ok {
                lst_room[i].Items = append(lst_room[i].Items, item)
            } else {
                utils.ExitError("UnknownRewardId", fmt.Errorf("reward id %d not found", lst_room[i].ItemsId[j]))
            }
        }
    }

    return lst_room
}

func resolve_room_allies(lst_room []models.Room, lst_npc []models.Npc) []models.Room {

    // On construit une map pour un accès rapide O(1)
    npcById := make(map[int]models.Npc)
    for _, npc := range lst_npc {
        npcById[npc.GetId()] = npc
    }

    for i := range lst_room {
        for j := range lst_room[i].AlliesId {
            if npc, ok := npcById[lst_room[i].AlliesId[j]]; ok {
                lst_room[i].Allies = append(lst_room[i].Allies, npc)
            } else {
                utils.ExitError("UnknownRewardId", fmt.Errorf("reward id %d not found", lst_room[i].AlliesId[j]))
            }
        }
    }

    return lst_room
}

func resolve_room_ennemy(lst_room []models.Room, lst_npc []models.Npc) []models.Room {

    // On construit une map pour un accès rapide O(1)
    npcById := make(map[int]models.Npc)
    for _, npc := range lst_npc {
        npcById[npc.GetId()] = npc
    }

    for i := range lst_room {
        for j := range lst_room[i].EnnemiesId {
            if npc, ok := npcById[lst_room[i].EnnemiesId[j]]; ok {
                lst_room[i].Ennemies = append(lst_room[i].Ennemies, npc)
            } else {
                utils.ExitError("UnknownRewardId", fmt.Errorf("reward id %d not found", lst_room[i].EnnemiesId[j]))
            }
        }
    }

    return lst_room
}

func resolve_trader_item(lst_room []models.Trader, lst_item []models.Item) []models.Trader {

    // On construit une map pour un accès rapide O(1)
    itemById := make(map[int]models.Item)
    for _, it := range lst_item {
        itemById[it.GetId()] = it
    }

    for i := range lst_room {
        for j := range lst_room[i].InventoryId {
            if item, ok := itemById[lst_room[i].InventoryId[j]]; ok {
                lst_room[i].Inventory = append(lst_room[i].Inventory, item)
            } else {
                utils.ExitError("UnknownRewardId", fmt.Errorf("reward id %d not found", lst_room[i].InventoryId[j]))
            }
        }
    }

    return lst_room
}

func resolve_monster_item(lst_monster []models.Monster, lst_item []models.Item) []models.Monster {

    // On construit une map pour un accès rapide O(1)
    itemById := make(map[int]models.Item)
    for _, it := range lst_item {
        itemById[it.GetId()] = it
    }

    for i := range lst_monster {
        if item, ok := itemById[lst_monster[i].LootId]; ok {
            lst_monster[i].Loot = item
        } else {
            utils.ExitError("UnknownRewardId", fmt.Errorf("reward id %d not found", lst_monster[i].LootId))
        }
    }

    return lst_monster
}

func resolve_quest_rewards(lst_quest []models.Quest, lst_item []models.Item) []models.Quest {

    // On construit une map pour un accès rapide O(1)
    itemById := make(map[int]models.Item)
    for _, it := range lst_item {
        itemById[it.GetId()] = it
    }

    for i := range lst_quest {
        if item, ok := itemById[lst_quest[i].RewardId]; ok {
            lst_quest[i].Reward = item
        } else {
            utils.ExitError("UnknownRewardId", fmt.Errorf("reward id %d not found", lst_quest[i].RewardId))
        }
    }

    return lst_quest
}

func resolve_questgiver_quests(lst_questgiver []models.QuestGiver, lst_quest []models.Quest) []models.QuestGiver {

    // On construit une map pour un accès rapide O(1)
    questById := make(map[int]models.Quest)
    for _, q := range lst_quest {
        questById[q.Id] = q
    }

    for i := range lst_questgiver {
        if quest, ok := questById[lst_questgiver[i].QuestId]; ok {
            lst_questgiver[i].Quest = quest
        } else {
            utils.ExitError("UnknownQuestId", fmt.Errorf("quest id %d introuvable", lst_questgiver[i].QuestId))
        }
    }

    return lst_questgiver
}

/* ----------------------------------------------------------------------- */
/*                       Méthodes de récupération JSON                     */
/* ----------------------------------------------------------------------- */

func get_data_from_json(filepath string) []uint8{

    data, err := os.ReadFile(filepath)

	if err != nil {
		utils.ExitError("ReadingFileError", err)
	}

    return data
}
