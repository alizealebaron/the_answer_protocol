/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* main.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/18 19:15:04 by alebaron        #+#    #+#              */
/* Updated: 2026/08/19 17:37:42 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                            Package & Import                             */
/* ----------------------------------------------------------------------- */

package main

import (
	"fmt"
	// "the_answer_protocol/src/server"
	"the_answer_protocol/src/models"
)

/* ----------------------------------------------------------------------- */
/*                                  Main                                   */
/* ----------------------------------------------------------------------- */

func main() {

    // tapManager := server.ParseJSONFile()
    // fmt.Printf("%+v\n", tapManager)

    player1 := models.NewPlayer("alebaron", "en")
    player2 := models.NewPlayer("emarette", "fr")

    fmt.Printf("%+v\n", player1)
    fmt.Printf("%+v\n", player2)
}