/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* map_widget.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/15 14:31:47 by alebaron        #+#    #+#              */
/* Updated: 2026/09/17 10:01:53 by rruiz           ###   ########.fr       */
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

const cellSize = float32(64)

func roomColor(state roomState) (color.Color, bool) {
	switch state {
	case roomCurrent:
		return color.NRGBA{R: 200, G: 100, B: 100, A: 255}, true
	case roomVisited:
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255}, true
	case roomKnown:
		return color.NRGBA{R: 90, G: 90, B: 90, A: 255}, true
	default:
		return nil, false
	}
}

type gridPos struct{ x, y int }

type mapLayout struct {
	gridPos map[fyne.CanvasObject]gridPos
}

func (l *mapLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	objects[0].Resize(size)
	objects[0].Move(fyne.NewPos(0, 0))

	content := objects[1]
	content.Resize(size)
	content.Move(fyne.NewPos(0, 0))

	centerX := size.Width / 2
	centerY := size.Height / 2
	half := cellSize / 2

	for _, obj := range content.(*fyne.Container).Objects {
		pos, ok := l.gridPos[obj]
		if !ok {
			continue
		}
		obj.Move(fyne.NewPos(centerX+float32(pos.x)*cellSize-half, centerY-float32(pos.y)*cellSize-half))
	}
}

func (l *mapLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.Size{Width: 0, Height: 0}
}

func MapWidget() *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	layout := &mapLayout{gridPos: make(map[fyne.CanvasObject]gridPos)}
	content := container.NewWithoutLayout()
	box := container.New(layout, frame, content)

	redraw := func() {
		content.RemoveAll()
		layout.gridPos = make(map[fyne.CanvasObject]gridPos)

		if !areBoundsComputed() {
			box.Refresh()
			return
		}

		rooms := getGameData().Rooms

		for _, room := range rooms {
			state := getRoomState(room.Id)
			fillColor, visible := roomColor(state)
			if !visible {
				continue
			}

			rect := canvas.NewRectangle(fillColor)
			rect.StrokeColor = color.White
			rect.StrokeWidth = float32(1)
			rect.Resize(fyne.NewSize(cellSize, cellSize))

			layout.gridPos[rect] = gridPos{x: room.X, y: room.Y}
			content.Add(rect)
		}

		box.Refresh()
	}

	setMapRedrawFunc(redraw)
	redraw()

	return box
}
