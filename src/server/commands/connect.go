/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* connect.go                                        :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/19 15:25:33 by emarette        #+#    #+#              */
/* Updated: 2026/08/19 15:25:38 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package connect

import (
	"the_answer_protocol/src/server"
	"the_answer_protocol/src/models"
	// "the_answer_protocol/src/utils"
)

func Connect(name string, language string) models.Player {
	for _, player := range server.TapManager.Lst_Player{
		if name == player.Name {
			panic("ERR 201 NAME_IN_USE")
		}
	}
	if language != "FR" && language != "EN" {
		panic("ERR 202 LANGUAGE_IN_USE")
	}

}