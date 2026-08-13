/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* main.rs                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/07/30 17:23:43 by alebaron        #+#    #+#              */
/* Updated: 2026/08/04 12:45:04 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                              Importations                               */
/* ----------------------------------------------------------------------- */

// === Utiliser pour load les models depuis le dossier models ===

mod models;
mod server;
use models::tap_manager_model::TapManager;
use serde_json::{Result, Value};

/* ----------------------------------------------------------------------- */
/*                                  Main                                   */
/* ----------------------------------------------------------------------- */

fn main() {
    // println!("--- Hello World ! ---");
    let mut tap_manager = server::parsing::get_tap_manager();
    println!("Tap: {:?}", tap_manager); // Debug

    // untyped_example();
}

fn untyped_example() -> Result<()> {
    // Some JSON input data as a &str. Maybe this comes from the user.
    let data = r#"
        {
            "name": "John Doe",
            "age": 43,
            "phones": [
                "+44 1234567",
                "+44 2345678"
            ]
        }"#;

    // Parse the string of data into serde_json::Value.
    let v: Value = serde_json::from_str(data)?;

    // Access parts of the data by indexing with square brackets.
    println!("Please call {} at the number {}", v["name"], v["phones"][0]);

    Ok(())
}