/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* error_widget.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 14:26:36 by rruiz           #+#    #+#              */
/* Updated: 2026/08/26 17:42:44 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func aller(content *fyne.Container) fyne.CanvasObject {
	move := canvas.NewPositionAnimation(fyne.NewPos(1, -100), fyne.NewPos(1, 1), 1*time.Second, content.Move)
	move.Start()

	return content
}

func retour(content *fyne.Container) fyne.CanvasObject {
	move := canvas.NewPositionAnimation(fyne.NewPos(1, 1), fyne.NewPos(1, -100), 1*time.Second, content.Move)
	move.Start()

	return content
}

func error_widget(message string) (*fyne.Container, *canvas.Text) {
	rect := canvas.NewRectangle(color.NRGBA{R: 255, G: 89, B: 89, A: 255})
	text := canvas.NewText(message, color.White)
	text.Alignment = fyne.TextAlignCenter
	content := container.NewStack(rect, text)
	return content, text
}

func displayError(errText *canvas.Text, errContent *fyne.Container, msg string) {
	fyne.Do(func() {
		errText.Text = msg
		errText.Refresh()
		errContent.Show()
		aller(errContent)
	})

	go func() {
		time.Sleep(3 * time.Second)
		fyne.Do(func() {
			retour(errContent)
		})
	}()
}
