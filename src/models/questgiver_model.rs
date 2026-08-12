/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* questgiver_model.rs                               :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/01 16:23:41 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 16:45:46 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importation                                */
/* ----------------------------------------------------------------------- */

use serde::Deserialize;
use std::collections::HashMap;
use crate::models::npc_model::Npc;
use crate::models::quest_model::Quest;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct QuestGiverRaw {
    pub id: u16,
    pub name: String,
    pub dialogue_fr: Vec<String>,
    pub dialogue_en: Vec<String>,
    pub quest: u16,
    pub dialogue_fin_fr: Vec<String>,
    pub dialogue_fin_en: Vec<String>,
}

#[derive(Debug)]
pub struct QuestGiver {
    pub id: u16,
    pub name: String,
    pub dialogue_fr: Vec<String>,
    pub dialogue_en: Vec<String>,
    pub quest: Quest,
    pub dialogue_fin_fr: Vec<String>,
    pub dialogue_fin_en: Vec<String>,
}

impl QuestGiver {
    pub fn from_raw(raw: QuestGiverRaw, quests: &HashMap<u16, Quest>) -> Result<Self, String> {

        let quest = quests
            .get(&raw.quest)
            .cloned()
            .ok_or_else(|| format!("Quête avec l'ID {} introuvable pour le PNJ {}", raw.quest, raw.id))?;

        Ok(Self {
            id: raw.id,
            name: raw.name,
            dialogue_fr: raw.dialogue_fr,
            dialogue_en: raw.dialogue_en,
            quest, // quête récupérée ci-dessus
            dialogue_fin_fr: raw.dialogue_fin_fr,
            dialogue_fin_en: raw.dialogue_fin_en,
        })
    }
}