/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* log_widget.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 13:07:38 by rruiz           #+#    #+#              */
/* Updated: 2026/09/16 22:49:09 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var logSubBox *fyne.Container
var logSubScroll *container.Scroll
var currentCategory string

func logWidget() *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	subBox := container.NewVBox()
	subScroll := container.NewScroll(subBox)

	logSubBox = subBox
	logSubScroll = subScroll
	currentCategory = "LOG"

	buttonBar := container.NewGridWithColumns(4,
		createLogButton("LOG", subBox, subScroll),
		createLogButton("GLOBAL", subBox, subScroll),
		createLogButton("ROOM", subBox, subScroll),
		createLogButton("GROUP", subBox, subScroll),
	)

	layout := container.NewBorder(buttonBar, nil, nil, nil, subScroll)
	return container.NewStack(frame, container.NewPadded(layout))

}

func createLogButton(category string, subBox *fyne.Container, subScroll *container.Scroll) *widget.Button {
	button := widget.NewButton(category, func() {
		showCategory(category, subBox, subScroll)
	})
	return button
}

func showCategory(category string, subBox *fyne.Container, subScroll *container.Scroll) {
	subBox.RemoveAll()
	currentCategory = category

	for _, line := range getMessages(category) {
		subBox.Add(newWrappingLabel(line))
	}
	subBox.Refresh()
	fyne.Do(func() {
		subScroll.ScrollToBottom()
	})
}

func newWrappingLabel(line string) *widget.Label {
	label := widget.NewLabel(line)
	label.Wrapping = fyne.TextWrapWord
	return label
}

func onLogLine(category, line string) {
	if category != currentCategory || logSubBox == nil {
		return
	}
	logSubBox.Add(newWrappingLabel(line))
	logSubBox.Refresh()
	fyne.Do(func() {
		logSubScroll.ScrollToBottom()
	})
}
