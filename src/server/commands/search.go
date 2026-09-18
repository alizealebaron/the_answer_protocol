/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* search.go                                         :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/18 10:46:52 by alebaron        #+#    #+#              */
/* Updated: 2026/09/18 10:47:32 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Search(args []string, tapManager *models.TapManager, player *models.Player) error {

	// on verifie le nombre d'argument
	if len(args) != 1 {
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	// on cherche la room dans lequel se trouve le joueur
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// On vérifie l'ennemie que l'on veut faire spawn
	for _, e := range room.Ennemies {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("ERR 904 WRONG_COMMAND_ARG")
		}

		// Quand l'ennemi trouvé est le bon on tente de le faire spawn
		if id == e.Id {
			if e.IsBoss == true {
				for _, mob := range room.Arena {
					if e.Name == mob.Name {
						return errors.New("ERR BOSS_ALREADY_SPAWN")
					}
				}
			}
			luck := rand.IntN(10) + 1

			if luck >= e.SpawnRate {

				// Récupération de l'index du monstre
				e.Entity_id = tapManager.Entity_index
				tapManager.Entity_index += 1

				// Envoie du message au joueur
				message := fmt.Sprintf("OK summon={\"monster\": \"%s\", \"id\": %d}\n", e.Name, e.Entity_id)
				server_write.ServerWrite(player.Conn, message)
				server_write.WriteLog(player.Conn, "SERVER", message)

				// Ajout du monstre à la room (Dans cet ordre pour un joli rendu côté client ~Alizéa)
				room.AddMonsterToRoom(e)
				server_write.WriteLog(player.Conn, "WORLD", player.Name + " summoned a \""+ e.GetName() +"\" in \""+ room.Name +"\"\n")
				return nil
			}
			server_write.ServerWrite(player.Conn, "ERR 409 FAILED_TO_SUMMON (TRY AGAIN)"+"\n")
			server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+room.ToString())
			return nil
		}
	}

	return errors.New("ERR 404 MONSTER_NOT_FOUND")
}
