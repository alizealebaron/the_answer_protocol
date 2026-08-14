/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* player_model.rs                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/01 15:17:04 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 16:45:24 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importation                                */
/* ----------------------------------------------------------------------- */

use crate::models::item_model::Item;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

#[derive(Debug)]
pub struct Player<'a> {
    pub id: u16,
    pub name: &'a str,
    pub ip: &'a str,
    pub pv: u16,
    pub attack: u16,
    pub inventory: Vec<Item> // Structure à revoir, peut-être utilisé un dico ou équivalent pour gérer la quantité plus facilement
}

/* ----------------------------------------------------------------------- */
/*                               Méthodes                                  */
/* ----------------------------------------------------------------------- */

impl<'a> Player<'a> {
    pub fn new(id: u16, name: &'a str, ip: &'a str, pv: u16, attack:u16) -> Self {
        Player {
            id,
            name,
            ip,
            pv,
            attack,
            inventory: Vec::new(),
        }
    }

    pub fn add_reward(&mut self, reward: Item) {
        self.inventory.push(reward);
    }
}