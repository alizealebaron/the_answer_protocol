/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* item_model.rs                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/01 14:59:37 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 15:37:47 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importation                                */
/* ----------------------------------------------------------------------- */

use crate::models::loot_model::Loot;
use crate::models::weapon_model::Weapon;
use crate::models::edible_model::Edible;

/* ----------------------------------------------------------------------- */
/*                               Structure                                 */
/* ----------------------------------------------------------------------- */

// Obligatoire en Rust pour simplifier les liaisons de classe sans héritage

#[derive(Debug)]
pub enum Item<'a> {
    Loot(Loot),
    Weapon(Weapon),
    Edible(Edible<'a>),
}