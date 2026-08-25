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

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package server

import (
	"fmt"
	"log"
	"net"
    "bufio"
    "errors"
	"strings"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/commands"
	"the_answer_protocol/src/server/server_write"
)

/* ----------------------------------------------------------------------- */
/*                           Variables Globales                            */
/* ----------------------------------------------------------------------- */

var TapManager *models.TapManager

// Signature commune à toutes les commandes
type CommandFunc func(args []string, tapManager *models.TapManager, player models.Player) error

// Registre des commandes
var map_commands = map[string]CommandFunc{
	"LOOK":   commands.Look,
    "CHAT":   commands.Chat,
	"MOVE":   commands.Move,
}

/* ----------------------------------------------------------------------- */
/*                                Fonctions                                */
/* ----------------------------------------------------------------------- */

func Tcp_server(tapManager *models.TapManager) {

    // === Récupération du tapmanager === //

    TapManager = tapManager
    // fmt.Printf("%+v\n", TapManager)

    // === Récupération de l'adresse IP de la machine === //

    conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)

    // === Ouverture du port d'écoute du serveur === //

    listener, err := net.Listen("tcp", ":8090")
    if err != nil {
        log.Fatal("Error listening:", err)
    }

    defer listener.Close()

	fmt.Println("[\033[32mSUCCESS\033[0m] (⊃｡•́‿•̀｡)⊃━☆ﾟ* Server started ! Use nc", localAddr.IP.String(), "8090")
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

    // === Déclaraction des variables === //

    var self_player models.Player
    var code_error string

    // === Gestion des demandes de l'utilisateur === //
    is_connected := false 
    for {

        // Récupération des commandes envoyées //
        reader := bufio.NewReader(conn)
        line, err := reader.ReadString('\n')
        if err != nil {
            log.Printf("Read error: %v", err)
            return
        }

        // Séparation des différents arguments //
        command := strings.Split(line, " ")
        for i, arg := range command {
            command[i] = strings.Trim(arg, "\n")
        }

        // Gestion des commandes selon si l'utilisateur est connecté ou non //
        if is_connected == false {
            // 3 possbilités HELP, CONNECT ou autre
            if command[0] == "HELP" {
                continue
            } else if command[0] == "CONNECT" && len(command) == 3 {
                self_player, code_error = commands.Connect(TapManager, conn, command[1], command[2])
                if code_error != "" {
                    fmt.Print("Connection attempt failed\n")
                    } else {
                        server_write.WriteLog(conn, "INFO", "Player " + self_player.Name + " connected")
                        is_connected = true
                    }
            } else {
                server_write.ServerWrite(conn, "use 'CONNECT [Name] [Language]' or 'HELP' for more information\n")
            }
        } else {
            if command[0] == "QUIT" {
                commands.Quit(TapManager, self_player)
                return
            }
            // Ecriture de la commande dans les logs
            server_write.WriteLog(conn, "COMMAND", self_player.Name + " use " + line)

            // Envoie de la ligne parse dans les différentes commandes
            if err := dispatch(command, TapManager, self_player, conn); err != nil {
                server_write.WriteLog(conn, "WARN", self_player.Name + " received a warn: " + err.Error())
                server_write.ServerWrite(conn, err.Error() + "\n")
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

func dispatch(fields []string, tap *models.TapManager, player models.Player, conn net.Conn) error {

    // Vérification de la longueur de la commande
	if len(fields) == 0 {
		return errors.New("ERR 903 COMMAND_EMPTY")
	}

    // Récupération des variables
	cmdName := fields[0]
	args := fields[1:]

    // Lancement de la commande
	fn, ok := map_commands[cmdName]
	if !ok {
		return fmt.Errorf("ERR 902 COMMAND_UNKNOWN %q", cmdName)
	}

	return fn(args, tap, player)
}