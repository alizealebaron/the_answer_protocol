/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* tcp_server.go                                       :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 13:30:33 by emarette        #+#    #+#              */
/* Updated: 2026/08/19 13:35:18 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package server

import (
	"bufio"
	"fmt"
	// "fmt"
	"log"
	"net"
	"strings"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/utils"
	"the_answer_protocol/src/server/commands"
)

var TapManager models.TapManager

func Tcp_server(tapManager models.TapManager) {

    TapManager = tapManager
    // fmt.Printf("%+v\n", TapManager)
    listener, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatal("Error listening:", err)
    }

    defer listener.Close()

    for {

        conn, err := listener.Accept()
        if err != nil {
            log.Println("Error accepting conn:", err)
            continue
        }

        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {

    defer conn.Close()

    var self_player models.Player
    var code_error string

    is_connected := false 
    for {
        reader := bufio.NewReader(conn)
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Printf("Read error: %v", err)
            return
        }

        command := strings.Split(line, " ")
        for i, arg := range command {
            command[i] = strings.Trim(arg, "\n")
            fmt.Printf("%d\n", i)
        } 
        fmt.Printf("%d\n", len(command))
        if is_connected == false {
            if command[0] == "HELP" {
                continue
            } else if command[0] == "CONNECT" && len(command) == 3 {
                self_player, code_error = commands.Connect(TapManager, conn, command[1], command[2])
                if code_error != "" {
                    fmt.Print("on trouvera un truc a dire\n")
                    } else {
                        fmt.Printf("%s connected\n", command[1])
                        is_connected = true
                    }
            } else {
                utils.ServerWrite(conn, "use 'CONNECT [Name] [Language]' or 'HELP' for more information\n")
            }
        } else {
             utils.ServerWrite(conn, "attend 2s\n")
             fmt.Printf("Bonjour %s\n", self_player.Name)
        }

        // ackMsg := strings.ToUpper(strings.TrimSpace(message))
        // response := fmt.Sprintf("ACK: %s\n", ackMsg)
        // _, err = conn.Write([]byte(response))
        // if err != nil {
        //     log.Printf("Server write error: %v", err)
        // }
    }
}