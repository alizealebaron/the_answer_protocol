/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* tcp_server.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/18 10:45:59 by alebaron        #+#    #+#              */
/* Updated: 2026/09/18 11:01:14 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package server

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
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
type CommandFunc func(args []string, tapManager *models.TapManager, player *models.Player) error

// Registre des commandes
var map_commands = map[string]CommandFunc{
	"BUY":       commands.Buy,
	"USE":       commands.Use,
	"WHO":       commands.Who,
	"CHAT":      commands.Chat,
	"DROP":      commands.Drop,
	"LOOK":      commands.Look,
	"MOVE":      commands.Move,
	"SELL":      commands.Sell,
	"TAKE":      commands.Take,
	"TALK":      commands.Talk,
	"GROUP":     commands.Group,
	"QUEST":     commands.Quest,
	"TRADE":     commands.Trade,
	"ATTACK":    commands.Attack,
	"QUESTS":    commands.Quests,
	"SEARCH":    commands.Search,
	"STATUS":    commands.Status,
	"GAMBLING":  commands.Gambling,
	"INVENTORY": commands.Inventory,
}

/* ----------------------------------------------------------------------- */
/*                         Fonctions Principales                           */
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

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("ERR 900 CONNECTION_FAILED")
		}
	}()
	
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// === Ouverture du port d'écoute du serveur === //

	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			log.Printf("ERR 900 CONNECTION_FAILED")
		}
	}()

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

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("ERR 900 DECONNECTION_FAILED")
		}
	}()

	// === Déclaraction des variables === //

	var self_player models.Player
	var code_error string

	// === Envoie du premier message === //

	server_write.ServerWrite(conn, "OK hello proto=1\n")

	// === Gestion des demandes de l'utilisateur === //
	is_connected := false
	reader := bufio.NewReader(conn)
	for {
		// Récupération des commandes envoyées //
		line, err := reader.ReadString('\n')
		if err != nil {
			server_write.WriteLog(conn, "ERROR", "Read error: "+err.Error())
			if err.Error() == "EOF" && is_connected {
				commands.Quit(TapManager, self_player)
			}
			return
		}

		// Séparation des différents arguments //
		command := strings.Split(line, " ")
		for i, arg := range command {
			command[i] = strings.Trim(arg, "\n")
		}

		// Gestion des commandes selon si l'utilisateur est connecté ou non //
		if !is_connected {
			// 2 possbilités CONNECT ou autre
			if command[0] == "CONNECT" && len(command) == 3 {
				self_player, code_error = commands.Connect(TapManager, conn, command[1], command[2])
				if code_error != "" {
					server_write.WriteLog(conn, "ERROR", "900 CONNECTION_FAILED")
				} else {
					server_write.WriteLog(conn, "INFO", "Player "+self_player.Name+" connected")

					// === Envoie de l'évènement de compte de joueur === //
					for _, p := range TapManager.Lst_Player {
						server_write.ServerWrite(p.Conn, "EVT STATS players="+strconv.Itoa(len(TapManager.Lst_Player))+"\n")
					}

					is_connected = true
					// item, _ := TapManager.GetItemById(1)
					// self_player.AddItemToPlayerWQuantity(item, 1000)
				}
			} else {
				server_write.ServerWrite(conn, "ERR 900 CONNECTION_FAILED\n")
				server_write.ServerWrite(conn, "USE \"CONNECT [Name] [Language]\"\n")
			}
		} else {
			switch command[0] {
				case "QUIT":
					server_write.WriteLog(conn, "COMMAND", self_player.Name+" use "+line)
					commands.Quit(TapManager, self_player)
					return
				case "SECRET":
					output, err := json.Marshal(TapManager)
					if err != nil {
						server_write.ServerWrite(conn, err.Error())
						return
					}
					server_write.ServerWrite(conn,"OK SECRET " + string(output) + "\n")
				default:
					// Ecriture de la commande dans les logs
					server_write.WriteLog(conn, "COMMAND", self_player.Name+" use "+line)

					// Reconstruction de la commande sous forme de string pour la comparaison
					rawCommand := strings.Join(command, " ")

					// Vérification anti-spam
					if self_player.AddCommandToHistory(rawCommand) {
						server_write.WriteLog(conn, "WARN", self_player.Name+" is spamming: "+rawCommand)
						server_write.ServerWrite(conn, "ERR 905 TOO_MANY_REQUESTS\n")
						continue // on ignore la commande sans l'exécuter
					}

					// Envoie de la ligne parse dans les différentes commandes
					if err := dispatch(command, TapManager, &self_player); err != nil {
						server_write.WriteLog(conn, "WARN", self_player.Name+" received a warn: "+err.Error())
						server_write.ServerWrite(conn, err.Error()+"\n")
					}
			}
		}
	}
}

/* ----------------------------------------------------------------------- */
/*                       Fonctions Supplémentaires                         */
/* ----------------------------------------------------------------------- */

func dispatch(fields []string, tap *models.TapManager, player *models.Player) error {

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
