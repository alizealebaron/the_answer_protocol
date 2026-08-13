/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* monster_model.rs                                  :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/01 16:26:59 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 16:31:50 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importation                                */
/* ----------------------------------------------------------------------- */

use serde::Deserialize;
use std::collections::HashMap;
use crate::models::item_model::Item;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct MonsterRaw {
    pub id: u16,
    pub name: String,
    pub dialogue_fr: Vec<String>,
    pub dialogue_en: Vec<String>,
    pub pv: u16,
    pub attack: u16,
    pub defense: u16,
    pub is_boss: bool,
    pub loot: u16,
    pub quantity_min: u16,
    pub quantity_max: u16,
}

#[derive(Debug)]
pub struct Monster {
    pub id: u16,
    pub name: String,
    pub dialogue_fr: Vec<String>,
    pub dialogue_en: Vec<String>,
    pub pv: u16,
    pub attack: u16,
    pub defense: u16,
    pub is_boss: bool,
    pub loot: Item,
    pub quantity_min: u16,
    pub quantity_max: u16,
}

impl Monster {
    pub fn from_raw(raw: MonsterRaw, items: &HashMap<u16, Item>) -> Result<Self, String> {

        let item = items
            .get(&raw.loot)
            .cloned()
            .ok_or_else(|| format!("Quête avec l'ID {} introuvable pour le PNJ {}", raw.loot, raw.id))?;

        Ok(Self {
            id: raw.id,
            name: raw.name,
            dialogue_fr: raw.dialogue_fr,
            dialogue_en: raw.dialogue_en,
            pv: raw.pv,
            attack: raw.attack,
            defense: raw.defense,
            is_boss: raw.is_boss,
            loot: item,
            quantity_min: raw.quantity_min,
            quantity_max: raw.quantity_max,
        })
    }
}