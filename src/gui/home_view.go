/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* home_view.go                                      :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 18:10:17 by rruiz           #+#    #+#              */
/* Updated: 2026/08/22 18:01:50 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (
	// "the_answer_protocol/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomeView() fyne.CanvasObject {
	// MENU DEROULANT DE LA LANGUE
	langSelect := widget.NewSelect([]string{"French", "English"}, func(value string) {})
	langSelect.PlaceHolder = "Select langage"

	content := container.NewVBox(
		langSelect,
		// tes autres champs (nom, ip serveur, bouton rejoindre) viendront ici
	)

	return container.NewCenter(content)
}
