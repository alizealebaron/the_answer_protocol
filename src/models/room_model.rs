/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* room_model.rs                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/01 16:38:17 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 16:50:27 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importation                                */
/* ----------------------------------------------------------------------- */

use serde::{Deserialize, Deserializer};
use std::collections::HashMap;
use crate::models::item_model::Item;
use crate::models::character_model::Character;
use crate::models::monster_model::Monster;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

#[derive(Debug, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Room{
    pub id: u16,
    pub name: String,
    pub allies: Vec<u16>,
    pub ennemies: Vec<u16>,
    pub items: Vec<u16>,
    #[serde(deserialize_with = "deserialize_neighbors")]
    pub lst_neighbor_room: [Option<u16>; 4],
}

// Transforme les 0 en None pour convertir depuis le JSON
fn deserialize_neighbors<'de, D>(deserializer: D) -> Result<[Option<u16>; 4], D::Error>
where
    D: Deserializer<'de>,
{
    let raw: [u16; 4] = Deserialize::deserialize(deserializer)?;
    let mut result = [None; 4];
    for (i, &val) in raw.iter().enumerate() {
        result[i] = if val == 0 { None } else { Some(val) };
    }
    Ok(result)
}