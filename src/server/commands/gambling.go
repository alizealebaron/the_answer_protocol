/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* gambling.go                                       :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/04 09:49:49 by alebaron        #+#    #+#              */
/* Updated: 2026/09/04 17:31:33 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"fmt"
	"math"
	"errors"
	"strconv"
	"math/rand/v2"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Constantes                              | */
/* +---------------------------------------------------------------------+ */

type GamblingStat struct {
	LootRate int
	Value    float64
}

var lstloot = []GamblingStat{
	{LootRate: 15, Value: 1.10},
	{LootRate: 15, Value: 0.88},
	{LootRate: 15, Value: 1.00},
	{LootRate: 10, Value: 1.25},
	{LootRate: 10, Value: 0.70},
	{LootRate: 10, Value: 0.65},
	{LootRate: 9, Value: 1.30},
	{LootRate: 5, Value: 2.00},
	{LootRate: 5, Value: 0.45},
	{LootRate: 2, Value: 3.00},
	{LootRate: 2, Value: 0.30},
	{LootRate: 1, Value: 5.00},
	{LootRate: 1, Value: 0.00},
}

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Gambling(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) != 1 {
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	// === Récupération de la room actuelle du Joueur === //
	room, err := tapManager.FindPlayerRoom(player.Id)
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	if room.Name != "CASINO" {
		return errors.New("ERR 408 INVALID_LOCATION")
	}

	// === Récupération de la somme donnée === //
	bet, err := strconv.Atoi(args[0])
	if err != nil || bet < 10 {
		return errors.New("ERR 408 INVALID_BET")
	}

	// === Vérification de la quantité de gambling coin === //
	bool_gc, err := player.IsEnoughGamblingCoin(bet)
	if (!bool_gc) {
		return err
	}

	// === Activation du ✨🎲 GAMBLING 🎲✨ === //

	// Calcul de la somme totale des taux de loot (Si pas égale à 100 plante pas)
	total := 0
	for _, entry := range lstloot {
		total += entry.LootRate
	}

	// Tirage aléatoire entre 0 et total-1
	roll := rand.IntN(total)

	// On parcourt les entrées en cumulant les taux
	cumulative := 0
	for _, entry := range lstloot {
		cumulative += entry.LootRate

		if roll < cumulative {

			// On a trouvé notre probabilité
			// test := fmt.Sprintf("Value obtenue : %f\n", entry.Value)
			// server_write.ServerWrite(player.Conn, test)
			total_gain := int(math.Ceil(float64(bet) * entry.Value))

			item, _ := player.RemoveItemToPlayerWQuantity(1, bet)
			player.AddItemToPlayerWQuantity(*item, total_gain)

			// === Envoie des messages au client et dans les logs === //
			str_ret := fmt.Sprintf("OK gambling={\"bet\":%d,\"result\":%d}\n", bet, total_gain)
			server_write.ServerWrite(player.Conn, str_ret)
			server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)

			return nil
		}
	}

	// Ne devrait normalement jamais arriver mais on sait pas hein
	return errors.New("ERR 999 GAMBLING_ROLL_FAILED")
}
