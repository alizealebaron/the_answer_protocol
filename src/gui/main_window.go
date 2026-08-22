/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* main_window.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/06/22 13:23:57 by rruiz           #+#    #+#              */
/* Updated: 2026/08/22 14:24:32 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (
	"the_answer_protocol/src/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func CreateMainWindow(a fyne.App) (fyne.Window, error) {
	w := a.NewWindow("The Gambling Protocol")
	w.SetFullScreen(true)
	res, err := fyne.LoadResourceFromPath("assets/other/icon.png")
	if err != nil {
		return nil, err
	}
	w.SetIcon(res)
	return w, nil
}

func Run() {
	a := app.NewWithID("com.theanswerprotocol.app")
	window, err := CreateMainWindow(a)
	if err != nil {
		utils.ExitError("Error during the window creation", err)
	}
	window.SetContent(HomeView())
	window.ShowAndRun()
}
