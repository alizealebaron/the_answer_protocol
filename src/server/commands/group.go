/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* group.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/29 09:47:44 by alebaron        #+#    #+#              */
/* Updated: 2026/08/29 14:00:18 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package commands

import (
	"errors"
	"fmt"
	"strconv"
	"the_answer_protocol/src/models"
	"the_answer_protocol/src/server/server_write"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func Group(args []string, tapManager *models.TapManager, player *models.Player) error {

	// === Vérification de la longueur des arguments === //
	if len(args) < 1 {
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	switch args[0] {
	case "CREATE":
		return create(player, tapManager)
	case "INVITE":
		return invite(args, player, tapManager)
	case "JOIN":
		return join(args, player, tapManager)
	case "LEAVE":
		return leave(player)
	default:
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}
}

/* +---------------------------------------------------------------------+ */
/* |                         Fonctions Secondaires                       | */
/* +---------------------------------------------------------------------+ */

func create(player *models.Player, tap *models.TapManager) error {

	// On vérifie que le joueur n'a pas déjà un groupe
	if player.Group != nil {
		return errors.New("ERR 402 ALREADY_IN_GROUP")
	}

	// Sinon on le met dans un groupe
	group := models.NewGroup(player)
	(*player).Group = &group
	tap.Lst_Group = append(tap.Lst_Group, &group)

	// On envoie les messages
	str_ret := fmt.Sprintf("OK group=%d\n", group.Id)
	group_output := fmt.Sprintf("New Group has been created id=%d\n", group.Id)
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)
	server_write.WriteLog(player.Conn, "GROUP", group_output)

	return nil
}

func invite(args []string, player *models.Player, tap *models.TapManager) error {

	// === Vérification de la longueur des arguments === //
	if len(args) < 2 {
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	// Vérification que le joueur a bien un groupe
	if player.Group == nil {
		return errors.New("ERR 404 GROUP_NOT_FOUND")
	}

	// === Trouver le joueur invité === //
	player_inv, err := tap.GetPlayerByName(args[1])
	if err != nil {
		return errors.New("ERR 404 PLAYER_NOT_FOUND")
	}

	// === Inviter le joueur === //
	err = player.Group.AddPlayerToInvited(*player_inv)
	if err != nil {
		return err
	}

	// === Envoyer les messages === //

	str_ret := "OK\n"
	group_output := fmt.Sprintf("A new invitation has been sent to %s to join the group id=%d\n", player_inv.Name, player.Group.Id)
	output := fmt.Sprintf("EVT GROUP INVITE id=%d\n", player.Group.Id)
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.ServerWrite(player_inv.Conn, output)
	server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)
	server_write.WriteLog(player.Conn, "GROUP", group_output)

	return nil
}

func join(args []string, player *models.Player, tap *models.TapManager) error {

	// === Vérification de la longueur des arguments === //
	if len(args) < 2 {
		return errors.New("ERR 904 WRONG_COMMAND_ARG")
	}

	// Vérification que le joueur n'a pas déjà un groupe
	if player.Group != nil {
		return errors.New("ERR 402 ALREADY_IN_GROUP")
	}

	// Vérification que le groupe existe bien
	id, _ := strconv.Atoi(args[1])
	group, err := tap.GetGroupById(id)
	if err != nil {
		return errors.New("ERR 404 GROUP_NOT_FOUND")
	}

	// Vérification que le joueur est bien invité dans le groupe
	if !(group.IsPlayerInvited(*player)) {
		return errors.New("ERR 402 PLAYER_NOT_INVITED")
	}

	// Vérification que le groupe n'est pas déjà full (4 joueurs)
	if len(group.Lst_Player) == 4 {
		return errors.New("ERR 401 GROUP_ALREADY_FULL")
	}

	// Ajout du joueur dans le groupe
	group.AddPlayerToGroup(player)
	player.Group = group

	// On envoie les messages
	str_ret := fmt.Sprintf("OK group=%d\n", group.Id)
	group_output := fmt.Sprintf("%s has join group id=%d\n", player.Name, group.Id)
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)
	server_write.WriteLog(player.Conn, "GROUP", group_output)

	// Envoie de l'évent à tous les joueurs
	for _, player1 := range group.Lst_Player {
		server_write.ServerWrite(player1.Conn, "EVT GROUP JOIN " + player.Name + "\n")
	}

	return nil
}

func leave(player *models.Player) error {

	// Vérification que le joueur a un groupe
	if player.Group == nil {
		return errors.New("ERR 401 NOT_IN_GROUP")
	}

	player.Group.RemovePlayerFromGroup(*player)

	// On envoie les messages
	str_ret := "OK\n"
	group_output := fmt.Sprintf("%s has leave group id=%d\n", player.Name, player.Group.Id)
	server_write.ServerWrite(player.Conn, str_ret)
	server_write.WriteLog(player.Conn, "SERVER", "To "+player.Name+": "+str_ret)
	server_write.WriteLog(player.Conn, "GROUP", group_output)

	// Envoie de l'évent à tous les joueurs
	for _, player1 := range player.Group.Lst_Player {
		server_write.ServerWrite(player1.Conn, "EVT GROUP LEAVE " + player.Name + "\n")
	}

	player.Group = nil

	return nil
}
