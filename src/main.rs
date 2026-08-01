/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* main.rs                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/07/30 17:23:43 by alebaron        #+#    #+#              */
/* Updated: 2026/08/01 11:15:12 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importations                               */
/* ----------------------------------------------------------------------- */

mod models;

use models::item_model::Item;
use models::edible_model::Edible;

/* ----------------------------------------------------------------------- */
/*                                  Main                                   */
/* ----------------------------------------------------------------------- */

fn main() {
    println!("--- Test de la structure Item ---");

    // 1. Création d'un Item
    let potion = Item {
        id: 1,
        name: "Jambon Beurre",
        description_fr: "Oui oui baguette.",
        description_en: "Yes yes baguette.",
        nb_copies: 10,
        nb_avail: 10,
    };

    // Affichage des champs de Item
    println!("ID: {}", potion.id);
    println!("Nom: {}", potion.name);
    println!("Description (FR): {}", potion.description_fr);
    println!("Disponibles: {}/{}", potion.nb_avail, potion.nb_copies);

    println!("\n--- Test de la structure Edible ---");

    // 2. Création d'un Edible qui englobe l'Item
    let pomme = Item {
        id: 2,
        name: "Délicieux bâtonnets de poisson",
        description_fr: "Miam.",
        description_en: "Yummy.",
        nb_copies: 5,
        nb_avail: 5,
    };

    let pomme_edible = Edible {
        item: pomme,
        effect: "HEAL",
        var_nb: 15,
    };

    // Accès imbriqué aux champs : .item.name
    println!("Objet consommable: {}", pomme_edible.item.name);
    println!("Effet: {}", pomme_edible.effect);
    println!("Valeur d'effet: {}", pomme_edible.var_nb);
}