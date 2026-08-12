/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* quest_model.rs                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/01 12:34:25 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 15:39:22 by alebaron        ###   ########.fr       */
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
pub struct QuestRaw{
    pub id: u16,
    pub title: String,
    pub description_fr: String,
    pub description_en: String,
    pub reward: u16,
    pub quantity: u16,
}

#[derive(Debug, Clone)]
pub struct Quest{
    pub id: u16,
    pub title: String,
    pub description_fr: String,
    pub description_en: String,
    pub reward: Item,
    pub quantity: u16,
}

impl Quest {

    pub fn from_raw(raw: QuestRaw, items: &HashMap<u16, Item>) -> Result<Self, String> {
        let item = items
            .get(&raw.reward)
            .cloned()
            .ok_or_else(|| format!("Item {} not found for the quest {}", raw.reward, raw.id))?;

        Ok(Self {
            id: raw.id,
            title: raw.title,
            description_fr: raw.description_fr,
            description_en: raw.description_en,
            reward: item,
            quantity: raw.quantity,
        })
    }
}