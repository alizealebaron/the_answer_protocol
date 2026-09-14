/* ************************************************************************ */
/*      _  _     ____                     ,~~.                              */
/*     | || |   |___  \             ,   (  ^ )>                             */
/*     | || |_    __) |             )\~~'   (       _      _      _         */
/*     |__   _|  / __/             (  .__)   )    >(.)__ <(^)__ =(o)__      */
/*        |_|   |_____| .fr         \_.____,*      (___/  (___/  (___/      */
/*                                                                          */
/* ************************************************************************ */
/* name   : search.go                                                       */
/* author : alebaron <alebaron@student.42.fr>                               */
/*                                                                          */
/* creation : Invalid date        by -----------                            */
/* update   : 2026/09/11 20:22:47 by alebaron                               */
/* ************************************************************************ */

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

				return nil
			}
			server_write.ServerWrite(player.Conn, "ERR 409 FAILED_TO_SUMMON (TRY AGAIN)"+"\n")
			server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+room.ToString())
			return nil
		}
	}

	return errors.New("ERR 404 MONSTER_NOT_FOUND")
}
