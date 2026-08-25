/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* error_widget.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 14:26:36 by rruiz           #+#    #+#              */
/* Updated: 2026/08/25 15:17:51 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func error_widget(message string, size fyne.Size) fyne.CanvasObject {
	rect := canvas.NewRectangle(color.NRGBA{R: 255, G: 89, B: 89, A: 255})
	rect.Resize(fyne.NewSize(0, 0))
	rect.Move(fyne.NewPos(1, 1))

	return rect
}
