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
use crate::models::item_model::Item;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

#[derive(Debug, Deserialize, Clone)]
#[serde(rename_all = "camelCase")]
pub struct Monster {
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
