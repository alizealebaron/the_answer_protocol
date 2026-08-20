/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* parsing.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/18 19:28:32 by alebaron        #+#    #+#              */
/* Updated: 2026/08/20 09:55:14 by alebaron        ###   ########.fr       */
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

func ParseJSONFile() {

    // === Initialisation des Items ===

    var lst_item []models.Item
    lst_item = get_all_item()

    // === Initialisation des Quests ===

    var lst_quest []models.Quest
    lst_quest = get_all_quest()
    resolve_quest_rewards(lst_quest, lst_item)

    // === Initialisation du tapManager ===
    tapManager := models.NewTapManager(lst_item, lst_quest)

    for _, item := range tapManager.Lst_item {
		fmt.Println(item.ToString())
	}

	fmt.Println("\n\n")

    for _, item := range tapManager.Lst_Quest {
		fmt.Println(item.ToString())
	}

	fmt.Println("[\033[32mSUCCESS\033[0m] (ﾉ◕ヮ◕)ﾉ*:・ﾟ✧ JSON successfully load.")
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

/* ----------------------------------------------------------------------- */
/*                      Méthodes de liaison des objets                     */
/* ----------------------------------------------------------------------- */

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
            utils.ExitError("UnknownRewardId", fmt.Errorf("reward id %d introuvable", lst_quest[i].RewardId))
        }
    }

    return lst_quest
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
