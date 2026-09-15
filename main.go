/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* main-gui.go                                       :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 14:18:03 by alebaron        #+#    #+#              */
/* Updated: 2026/09/15 14:18:26 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                            Package & Import                             */
/* ----------------------------------------------------------------------- */

package main

import (
	"fmt"
	"os"
	"the_answer_protocol/src/gui"
    "the_answer_protocol/src/server"
    "the_answer_protocol/src/server/server_write"
)

/* ----------------------------------------------------------------------- */
/*                                  Main                                   */
/* ----------------------------------------------------------------------- */

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("ERR 904 WRONG_COMMAND_ARG")
        os.Exit(0)
    }

    if os.Args[1] == "server" {
        tapManager := server.ParseJSONFile()
        server_write.CreateLogFolder()
        server.Tcp_server(&tapManager)
    } else if os.Args[1] == "gui" {
        gui.Run(true)
    } else {
        fmt.Printf("ERR 904 WRONG_COMMAND_ARG")
    }
    // gui.Run(false)
}