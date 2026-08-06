/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* parsing.rs                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/04 10:00:57 by alebaron        #+#    #+#              */
/* Updated: 2026/08/04 12:49:28 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importations                               */
/* ----------------------------------------------------------------------- */

// === Utiliser pour load les models depuis le dossier models ===

use std::fs::File;
use std::path::Path;
use crate::models::tap_manager_model::TapManager;
use crate::models::item_model::Item;
use crate::models::loot_model::Loot;

/* ----------------------------------------------------------------------- */
/*                                  Main                                   */
/* ----------------------------------------------------------------------- */

pub fn get_tap_manager<'a>() -> TapManager<'a> {

    // === Déclaration des variables === //

    let mut tap_manager = TapManager::new();

    // === Récupération des datas du fichier === //

    let json_file_path = Path::new("data/loot_data.json");
    let file = File::open(json_file_path).expect("impossible d'ouvrir le fichier");

    let loots:Vec<Loot> = serde_json::from_reader(file).expect("error while reading or parsing");
    tap_manager.lst_item.extend(loots.into_iter().map(Item::Loot));
    
    return tap_manager;
}