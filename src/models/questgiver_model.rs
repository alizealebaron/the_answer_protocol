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

#[derive(Debug, Deserialize, Clone)]
#[serde(rename_all = "camelCase")]
pub struct QuestGiver {
    pub id: u16,
    pub name: String,
    pub dialogue_fr: Vec<String>,
    pub dialogue_en: Vec<String>,
    pub quest: u16,
    pub dialogue_fin_fr: Vec<String>,
    pub dialogue_fin_en: Vec<String>,
}
