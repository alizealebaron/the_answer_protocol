/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* map_widget.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 14:31:47 by alebaron        #+#    #+#              */
/* Updated: 2026/09/17 21:56:00 by rruiz           ###   ########.fr       */
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
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

const cellSize = float32(64) // Side of one room square on the map

// Returns the color for a room state, and if the room must be drawn.
func roomColor(state roomState) (color.Color, bool) {
	switch state {
	case roomCurrent:
		return color.NRGBA{R: 200, G: 100, B: 100, A: 255}, true // Red: player is here
	case roomVisited:
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255}, true // Light gray: seen
	case roomKnown:
		return color.NRGBA{R: 90, G: 90, B: 90, A: 255}, true // Dark gray: known
	default:
		return nil, false // Hidden room, we do not draw it
	}
}

// Grid coordinates (x, y) of one room on the map.
type gridPos struct{ x, y int }

// Custom layout that places each room at a fixed position on the map.
type mapLayout struct {
	gridPos map[fyne.CanvasObject]gridPos
}

func (l *mapLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	// Both layers (border and rooms) cover the whole box
	objects[0].Resize(size)
	objects[0].Move(fyne.NewPos(0, 0))

	content := objects[1]
	content.Resize(size)
	content.Move(fyne.NewPos(0, 0))

	// The box center is the origin, each room is placed around it
	centerX := size.Width / 2
	centerY := size.Height / 2
	half := cellSize / 2

	for _, obj := range content.(*fyne.Container).Objects {
		pos, ok := l.gridPos[obj]
		if !ok {
			continue // Not a room (frame, ...), we leave it alone
		}
		// Center the square on its grid spot, y grows upward so it is negative
		obj.Move(fyne.NewPos(centerX+float32(pos.x)*cellSize-half, centerY-float32(pos.y)*cellSize-half))
	}
}

// Required by the layout interface, we let children decide their own size.
func (l *mapLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.Size{Width: 0, Height: 0}
}

// Builds the map widget, which draws one square per known room.
func MapWidget() *fyne.Container {
	// Transparent rectangle that only draws its white outline, the widget border
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	layout := &mapLayout{gridPos: make(map[fyne.CanvasObject]gridPos)}
	content := container.NewWithoutLayout()
	box := container.New(layout, frame, content)

	// Rebuilds all the squares with the current game data.
	redraw := func() {
		content.RemoveAll()
		layout.gridPos = make(map[fyne.CanvasObject]gridPos)

		if !areBoundsComputed() {
			box.Refresh() // Map size not known yet, nothing to draw
			return
		}

		rooms := getGameData().Rooms

		for _, room := range rooms {
			state := getRoomState(room.Id)
			fillColor, visible := roomColor(state)
			if !visible {
				continue // Hidden room, we skip it
			}

			// One square, centered on its grid position by the layout
			rect := canvas.NewRectangle(fillColor)
			rect.StrokeColor = color.White
			rect.StrokeWidth = float32(1)
			rect.Resize(fyne.NewSize(cellSize, cellSize))

			layout.gridPos[rect] = gridPos{x: room.X, y: room.Y}
			content.Add(rect)
		}

		box.Refresh()
	}

	// Gives the redraw function to the state, so new data triggers it
	setMapRedrawFunc(redraw)
	redraw() // First draw of the starting state

	return box
}
