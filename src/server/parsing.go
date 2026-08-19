/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* parsing.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/18 19:28:32 by alebaron        #+#    #+#              */
/* Updated: 2026/08/19 12:12:34 by alebaron        ###   ########.fr       */
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
    "the_answer_protocol/src/models"
)

/* ----------------------------------------------------------------------- */
/*                                  Main                                   */
/* ----------------------------------------------------------------------- */

func ParseJSONFile() {

    // === Lecture du fichier json ===

    data, err := os.ReadFile("data/loot_data.json")
	if err != nil {
		fmt.Println("Erreur lecture fichier:", err)
		return
	}

    // === Initialisation des données ===
    var liste []models.Loot
    err = json.Unmarshal(data, &liste)
	if err != nil {
		fmt.Println("Erreur parsing JSON:", err)
		return
	}

    // 3. Utiliser les données
	for _, p := range liste {
        fmt.Println(p.ToString())
	}
}