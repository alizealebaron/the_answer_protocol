/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* tapManager_model.rs                               :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/04 08:53:06 by alebaron        #+#    #+#              */
/* Updated: 2026/08/04 12:48:17 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importation                                */
/* ----------------------------------------------------------------------- */

use std::collections::HashMap;
use crate::models::item_model::Item;
use crate::models::character_model::Character;
use crate::models::quest_model::Quest;
use crate::models::room_model::Room;
use crate::models::player_model::Player;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

#[derive(Debug)]
pub struct TapManager<'a> {
    pub lst_item: Vec<Item>,
    pub lst_character: Vec<Character>,
    pub lst_player: Vec<Player<'a>>,
    pub lst_quest: Vec<Quest>,
    pub lst_room: Vec<Room>,
    pub map_item: HashMap<u16, Item>,
    pub map_quest: HashMap<u16, Quest>,
    pub map_character: HashMap<u16, Character>,
}

/* ----------------------------------------------------------------------- */
/*                               Méthodes                                  */
/* ----------------------------------------------------------------------- */

impl<'a> TapManager<'a> {

    /* ----------------------------------------------------------------------- */
    /*                             Constructeur                                */
    /* ----------------------------------------------------------------------- */

    pub fn new() -> Self 
    {
        TapManager {
            lst_item: Vec::new(),
            lst_character: Vec::new(),
            lst_player: Vec::new(),
            lst_quest: Vec::new(),
            lst_room: Vec::new(),
            map_item: HashMap::new(),
            map_quest: HashMap::new(),
            map_character: HashMap::new(),
        }
    }

    /* ----------------------------------------------------------------------- */
    /*                           Méthodes d'ajout                              */
    /* ----------------------------------------------------------------------- */

    pub fn add_player(&mut self, player: Player<'a>) {
        self.lst_player.push(player);
    }
}