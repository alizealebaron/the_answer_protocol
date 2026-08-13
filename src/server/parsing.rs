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
use crate::models::character_model::Character;
use crate::models::npc_model::Npc;
use crate::models::questgiver_model::QuestGiver;
use crate::models::questgiver_model::QuestGiverRaw;
use crate::models::monster_model::Monster;
use crate::models::monster_model::MonsterRaw;
use crate::models::room_model::Room;

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

    let lst_questraw: Vec<QuestRaw> = load_data_from_json("data/quest_data.json");

    tap_manager.lst_quest = lst_questraw
        .into_iter()
        .map(|raw| Quest::from_raw(raw, &items_map))
        .collect::<Result<Vec<_>, _>>()
        .unwrap_or_else(|e| panic!("Error while loading quest: {}", e));

    // == Character Object == //

    let quest_map: HashMap<u16, Quest> = tap_manager
        .lst_quest
        .iter()
        .map(|quest| (quest.id, quest.clone()))
        .collect();

    tap_manager.lst_character.extend(load_character("data/npc_data.json", Character::Npc));

    // == QuestGiver Object == //

    let lst_questgiverraw: Vec<QuestGiverRaw> = load_data_from_json("data/questgiver_data.json");

    tap_manager.lst_character.extend(
    lst_questgiverraw
        .into_iter()
        .map(|raw| QuestGiver::from_raw(raw, &quest_map))
        .collect::<Result<Vec<_>, _>>()
        .unwrap_or_else(|e| panic!("Error while loading quest givers: {}", e))
        .into_iter()
        .map(Character::QuestGiver)
    );

    // == Monster Object == //

    let lst_monsterraw: Vec<MonsterRaw> = load_data_from_json("data/monster_data.json");
    tap_manager.lst_character.extend(
    lst_monsterraw
        .into_iter()
        .map(|raw| Monster::from_raw(raw, &items_map))
        .collect::<Result<Vec<_>, _>>()
        .unwrap_or_else(|e| panic!("Error while loading quest givers: {}", e))
        .into_iter()
        .map(Character::Monster)
    );

    // == Room Object == //

    let lst_roomraw: Vec<Room> = load_data_from_json("data/room_data.json");

    println!("Tap: {:?}", lst_roomraw); // Debug

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

fn load_character<T, F>(path: &str, map_fn: F) -> impl Iterator<Item = Character>
where
    T: DeserializeOwned,
    F: Fn(T) -> Character,
{
    let file = File::open(path)
        .unwrap_or_else(|e| panic!("Error: Cannot find or read {}: {}", path, e));
    let items: Vec<T> = serde_json::from_reader(file)
        .unwrap_or_else(|e| panic!("Error while parsing {}: {}", path, e));

    items.into_iter().map(map_fn)
}

fn load_data_from_json<T>(path: &str) -> Vec<T>
where
    T: DeserializeOwned,
{
    let file = File::open(path)
        .unwrap_or_else(|e| panic!("Error: Cannot find or read {}: {}", path, e));

    serde_json::from_reader(file)
        .unwrap_or_else(|e| panic!("Error while parsing {}: {}", path, e))
}