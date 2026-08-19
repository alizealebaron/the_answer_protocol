/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* parsing.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/18 19:28:32 by alebaron        #+#    #+#              */
/* Updated: 2026/08/19 18:02:50 by alebaron        ###   ########.fr       */
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

    var lst_item []models.Item
    lst_item = get_all_item()

    // === Initialisation du tapManager ===
    tapManager := models.NewTapManager(lst_item)

    for _, item := range tapManager.Lst_item {
		fmt.Println(item.ToString())
	}

	fmt.Println("[\033[32mSUCCESS\033[0m] (ﾉ◕ヮ◕)ﾉ*:・ﾟ✧ JSON successfully load.")
}

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

func get_data_from_json(filepath string) []uint8{

    data, err := os.ReadFile(filepath)

	if err != nil {
		utils.ExitError("ReadingFileError", err)
	}

    return data
}
