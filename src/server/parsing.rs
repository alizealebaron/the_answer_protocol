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
use std::collections::HashMap;
use serde::de::DeserializeOwned;
use crate::models::tap_manager_model::TapManager;
use crate::models::item_model::Item;
use crate::models::loot_model::Loot;
use crate::models::weapon_model::Weapon;
use crate::models::edible_model::Edible;
use crate::models::quest_model::QuestRaw;
use crate::models::quest_model::Quest;

/* ----------------------------------------------------------------------- */
/*                             Main Fonction                               */
/* ----------------------------------------------------------------------- */

pub fn get_tap_manager<'a>() -> TapManager<'a> {

    // === Déclaration des variables === //

    let mut tap_manager = TapManager::new();

    // === Récupération des datas du fichier === //

    // == Item Object == //

    tap_manager.lst_item.extend(load_items("data/loot_data.json", Item::Loot));
    tap_manager.lst_item.extend(load_items("data/weapon_data.json", Item::Weapon));
    tap_manager.lst_item.extend(load_items("data/edible_data.json", Item::Edible));

    // == Quest Object == //

    let items_map: HashMap<u16, Item> = tap_manager
        .lst_item
        .iter()
        .map(|item| (item.id(), item.clone()))
        .collect();

    let path = "data/quest_data.json";
    let file = File::open(path).unwrap_or_else(|e| panic!("Error: Cannot find or read {}: {}", path, e));
    let lst_questraw: Vec<QuestRaw> = serde_json::from_reader(file)
        .unwrap_or_else(|e| panic!("Error while parsing {}: {}", path, e));

    tap_manager.lst_quest = lst_questraw
        .into_iter()
        .map(|raw| Quest::from_raw(raw, &items_map))
        .collect::<Result<Vec<_>, _>>()
        .unwrap_or_else(|e| panic!("Error while loading quest: {}", e));

    return tap_manager;
}

/* ----------------------------------------------------------------------- */
/*                             Helpful Fonction                            */
/* ----------------------------------------------------------------------- */

fn load_items<T, F>(path: &str, map_fn: F) -> impl Iterator<Item = Item>
where
    T: DeserializeOwned,
    F: Fn(T) -> Item,
{
    let file = File::open(path)
        .unwrap_or_else(|e| panic!("Error: Cannot find or read {}: {}", path, e));
    let items: Vec<T> = serde_json::from_reader(file)
        .unwrap_or_else(|e| panic!("Error while parsing {}: {}", path, e));

    items.into_iter().map(map_fn)
}

