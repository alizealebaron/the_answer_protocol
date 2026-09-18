/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* log_widget.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 13:07:38 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 21:55:50 by rruiz           ###   ########.fr       */
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

// Package variables so the state can reach the log widget from anywhere.
var logSubBox *fyne.Container
var logSubScroll *container.Scroll
var currentCategory string

// Builds the log: category buttons on top, the scrollable message list below.
func logWidget() *fyne.Container {
	// Transparent rectangle that only draws its white outline, the widget border
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	subBox := container.NewVBox()
	subScroll := container.NewScroll(subBox)

	// Keeps the widget reachable for onLogLine
	logSubBox = subBox
	logSubScroll = subScroll
	currentCategory = "LOG" // Start on the default log

	buttonBar := container.NewGridWithColumns(4,
		createLogButton("LOG", subBox, subScroll),
		createLogButton("GLOBAL", subBox, subScroll),
		createLogButton("ROOM", subBox, subScroll),
		createLogButton("GROUP", subBox, subScroll),
	)

	layout := container.NewBorder(buttonBar, nil, nil, nil, subScroll)
	// The stack paints the border below the content, padded to not overlap it
	return container.NewStack(frame, container.NewPadded(layout))

}

// Creates a filter button for one log category.
func createLogButton(category string, subBox *fyne.Container, subScroll *container.Scroll) *widget.Button {
	button := widget.NewButton(category, func() {
		showCategory(category, subBox, subScroll)
	})
	return button
}

// Clears the list and shows every stored message of this category.
func showCategory(category string, subBox *fyne.Container, subScroll *container.Scroll) {
	subBox.RemoveAll()
	currentCategory = category

	for _, line := range getMessages(category) {
		subBox.Add(newWrappingLabel(line))
	}
	subBox.Refresh()
	// ScrollToBottom is wrapped in fyne.Do: just after Refresh the size is
	// still outdated, running it on the next frame gives the real size
	fyne.Do(func() {
		subScroll.ScrollToBottom()
	})
}

// Label with word wrapping so long lines do not get cut.
func newWrappingLabel(line string) *widget.Label {
	label := widget.NewLabel(line)
	label.Wrapping = fyne.TextWrapWord
	return label
}

// Adds one new line to the log, only if the current category is displayed.
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
