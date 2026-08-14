/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* character_model.rs                                :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/01 16:39:14 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 16:40:11 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importation                                */
/* ----------------------------------------------------------------------- */

use crate::models::npc_model::Npc;
use crate::models::monster_model::Monster;
use crate::models::questgiver_model::QuestGiver;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

// Obligatoire en Rust pour simplifier les liaisons de classe sans héritage

#[derive(Debug, Clone)]
pub enum Character{
    Npc(Npc),
    Monster(Monster),
    QuestGiver(QuestGiver),
}

impl Character {
    pub fn id(&self) -> u16 {
        match self {
            Character::Npc(npc) => npc.id,
            Character::QuestGiver(questgiver) => questgiver.id,
            Character::Monster(monster) => monster.id,
        }
    }
}