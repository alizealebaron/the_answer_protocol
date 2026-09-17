/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* main_window.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/06/22 13:23:57 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 21:06:32 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (
	"the_answer_protocol/src/gui/home"
	"the_answer_protocol/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func CreateMainWindow(a fyne.App, fullscreen bool) (fyne.Window, error) {
	// Create the Window with his title
	w := a.NewWindow("The Gambling Protocol")

	// Whether to open the window in full-screen mode or not, used only during testing
	if fullscreen {
		w.SetFullScreen(true)
	} else {
		w.Resize(fyne.NewSize(1024, 768))
		w.CenterOnScreen()
	}

	// Load icon.png if error occurs, return an error
	res, err := fyne.LoadResourceFromPath("assets/icon.png")

	// Returns nil and the error, or the window without an error
	if err != nil {
		return nil, err
	}
	// Set the icon
	w.SetIcon(res)
	return w, nil
}

func Run(fullscreen bool) {
	// Create a new app instance use for all window
	a := app.NewWithID("theanswerprotocol")

	// Create the window, if an error occurs exit with a clear error message
	window, err := CreateMainWindow(a, fullscreen)
	if err != nil {
		utils.ExitError("Error during the window creation", err)
	}

	var size fyne.Size
	// Set the window size because Go has trouble with fullscreen mode
	if fullscreen {
		size = fyne.NewSize(1920, 1080)
	} else {
		size = window.Canvas().Size()
	}

	// Displays the window's default content on the home screen
	window.SetContent(home.HomeView(window, size))

	window.ShowAndRun()
}
