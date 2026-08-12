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

use crate::models::item_model::Item;
use crate::models::character_model::Character;
use crate::models::monster_model::Monster;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

#[derive(Debug)]
pub struct Room<'a> {
    pub id: u16,
    pub name: &'a str,
    pub allies: Vec<&'a Character>,
    pub ennemies: Vec<&'a Monster>,
    pub items: Vec<&'a Item>,
    pub lst_neighbor_room: [Option<&'a Room<'a>>; 4], // Option permet d'avoir 'None' s'il n'y a pas de pièce voisine
}
