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
        }
    }

    /* ----------------------------------------------------------------------- */
    /*                           Méthodes d'ajout                              */
    /* ----------------------------------------------------------------------- */

    pub fn add_item(&mut self, reward: Item) {
        self.lst_item.push(reward);
    }

    pub fn add_character(&mut self, charac: Character) {
        self.lst_character.push(charac);
    }

    pub fn add_player(&mut self, player: Player<'a>) {
        self.lst_player.push(player);
    }

    pub fn add_quest(&mut self, quest: Quest) {
        self.lst_quest.push(quest);
    }

    pub fn add_room(&mut self, room: Room) {
        self.lst_room.push(room);
    }
}