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
    // "fmt"
    "log"
    "net"
    "strings"
    "the_answer_protocol/src/models"
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

    var self_player *models.Player

    for {
        reader := bufio.NewReader(conn)
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Printf("Read error: %v", err)
            return
        }

        command := strings.Split(line, " ") 

        if !self_player {
            if command[0] != "CONNECT" && command[0] != "HELP" {
                response := "use 'CONNECT [Name] [Language]'\n"
                _, err = conn.Write([]byte(response)) 
                if err != nil {
                    log.Printf("Server write error: %v", err)
                }
            } else if command[0] == "HELP" {
                continue
            } else if command[0] == "CONNECT" {
                continue
            }
        }

        // ackMsg := strings.ToUpper(strings.TrimSpace(message))
        // response := fmt.Sprintf("ACK: %s\n", ackMsg)
        // _, err = conn.Write([]byte(response))
        // if err != nil {
        //     log.Printf("Server write error: %v", err)
        // }
    }
}