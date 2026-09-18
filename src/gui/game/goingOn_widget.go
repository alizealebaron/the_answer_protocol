/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* goingOn_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/17 09:56:18 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 21:58:49 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"encoding/json"
	"image/color"
	"os"
	"strings"
	"the_answer_protocol/src/gui/game/types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Shows the room the player is in: its image in the middle, its name on top.
func goingOnWidget(listener *types.Listener) *fyne.Container {
	label := placeNameLabel()
	image := placeImage()

	subscribeRoomChange(listener, label, image)

	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	roomSquare := container.NewStack(image, frame)

	return container.NewBorder(label, nil, nil, nil, roomSquare)
}

// Label that shows the current room name.
func placeNameLabel() *widget.Label {
	placeLabel := widget.NewLabel(currentRoomName())
	placeLabel.Alignment = fyne.TextAlignCenter

	return placeLabel
}

// Returns the name of the room the player is in.
func currentRoomName() string {
	id := getCurrentRoom()
	for _, room := range getGameData().Rooms {
		if room.Id == id {
			return room.Name
		}
	}
	return types.Translate("You're lost!") // Room not found in the data
}

// On each LOOK reply, updates the label and the room image.
func subscribeRoomChange(listener *types.Listener, label *widget.Label, image *canvas.Image) {
	listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"id\":") {
			return
		}

		var data types.LookInfo
		if err := json.Unmarshal([]byte(strings.TrimPrefix(line, "OK ")), &data); err != nil {
			return
		}

		label.SetText(data.Name)
		setImage(image, data.Id)
	})
}

// Room image built from the current room type.
func placeImage() *canvas.Image {
	image := canvas.NewImageFromFile(currentRoomImage(getCurrentRoom()))
	image.FillMode = canvas.ImageFillStretch
	image.SetMinSize(fyne.NewSize(100, 100))

	return image
}

func setImage(image *canvas.Image, id int) {
	image.File = currentRoomImage(id)
	image.Refresh()
}

// Returns the asset path matching the room type, or a placeholder if missing.
func currentRoomImage(id int) string {
	for _, room := range getGameData().Rooms {
		if room.Id == id {
			path := "assets/" + room.Type + ".png"
			if _, err := os.Stat(path); err != nil {
				return "assets/placeholder.png"
			}
			return path
		}
	}
	return "assets/placeholder.png"
}
