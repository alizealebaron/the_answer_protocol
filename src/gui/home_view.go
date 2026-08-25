/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* home_view.go                                      :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 18:10:17 by rruiz           #+#    #+#              */
/* Updated: 2026/08/22 18:09:10 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (

	// "the_answer_protocol/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomeView(size fyne.Size) fyne.CanvasObject {
	// Definition of widget text to change based on the language selection button
	var translations = map[string]map[string]string{
		"Français": {
			"selectLang": "Choisir la langue",
			"enterName":  "Entrez votre nom",
			"enterIP":    "Entrez l'IP du serveur",
			"joinServer": "Rejoindre le serveur",
		},
		"English": {
			"selectLang": "Select language",
			"enterName":  "Enter your name",
			"enterIP":    "Enter the ip of the game",
			"joinServer": "Join server",
		},
	}

	// Defining a variable to be used later, done this way to maintain a logical order
	var langSelect *widget.Select
	var nameField *widget.Entry
	var ipField *widget.Entry
	var joinButton *widget.Button

	// Language Selection Drop-Down Menu
	// It changes the values of the placeholders in the other widgets based on the value of the selection
	langSelect = widget.NewSelect([]string{"Français", "English"},
		func(value string) {
			t := translations[value]
			langSelect.PlaceHolder = t["selectLang"]
			nameField.SetPlaceHolder(t["enterName"])
			ipField.SetPlaceHolder(t["enterIP"])
			joinButton.SetText(t["joinServer"])
		})
	langSelect.PlaceHolder = "Select langage"

	// Text field for the player's name
	nameField = widget.NewEntry()
	nameField.SetPlaceHolder("Enter your name")

	// Game Server IP Field
	ipField = widget.NewEntry()
	ipField.SetPlaceHolder("Enter the ip of the game")

	// Button to join the server
	joinButton = widget.NewButton("Join server", func() {})

	// Set variables at the value of the window
	width, height := size.Width, size.Height

	// Create a container that holds all the widgets created earlier, resize it, and place it at the bottom center of the window
	big := container.NewCenter(
		container.NewGridWrap(fyne.NewSize(width/3, height/10), langSelect, nameField, ipField, joinButton),
	)

	// return the container at the good place
	return container.NewBorder(nil, big, nil, nil)
}
