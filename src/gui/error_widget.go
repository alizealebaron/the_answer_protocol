/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* error_widget.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/25 14:26:36 by rruiz           #+#    #+#              */
/* Updated: 2026/08/26 17:56:12 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package gui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func goTo(content *fyne.Container) fyne.CanvasObject {
	// Animation of the errorWidget's descent.
	move := canvas.NewPositionAnimation(fyne.NewPos(1, -100), fyne.NewPos(1, 1), 1*time.Second, content.Move)
	move.Start()

	return content
}

func back(content *fyne.Container) fyne.CanvasObject {
	// Ascent animation for the errorWidget.
	move := canvas.NewPositionAnimation(fyne.NewPos(1, 1), fyne.NewPos(1, -100), 1*time.Second, content.Move)
	move.Start()

	return content
}

func errorWidget(message string) (*fyne.Container, *widget.Label) {
	// A function that creates the errorWidget, with a message inside it.
	rect := canvas.NewRectangle(color.NRGBA{R: 255, G: 89, B: 89, A: 255})
	text := widget.NewLabel(message)
	text.Wrapping = fyne.TextWrapWord
	text.Alignment = fyne.TextAlignCenter
	centeredText := container.NewVBox(layout.NewSpacer(), text, layout.NewSpacer())
	content := container.NewStack(rect, centeredText)
	return content, text
}

func displayError(errText *widget.Label, errContent *fyne.Container, msg string) {
	// Complete animation of the errWidget: descent, pause, and ascent
	fyne.Do(func() {
		errText.Text = msg
		errText.Refresh()
		errContent.Show()
		goTo(errContent)
	})

	go func() {
		time.Sleep(3 * time.Second)
		// Executes the code on the main thread (required to modify the interface).
		fyne.Do(func() {
			back(errContent)
		})
	}()
}
