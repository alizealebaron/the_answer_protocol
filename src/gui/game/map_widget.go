/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* map_widget.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 14:31:47 by alebaron        #+#    #+#              */
/* Updated: 2026/09/15 14:43:07 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package game

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func MapWidget() *fyne.Container {

	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	return container.NewStack(frame)
}