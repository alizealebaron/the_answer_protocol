/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* goingOn_widget.go                                 :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/17 09:56:18 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 14:53:20 by rruiz           ###   ########.fr       */
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

func placeNameLabel() *widget.Label {
	placeLabel := widget.NewLabel(currentRoomName())
	placeLabel.Alignment = fyne.TextAlignCenter

	return placeLabel
}

func currentRoomName() string {
	id := getCurrentRoom()
	for _, room := range getGameData().Rooms {
		if room.Id == id {
			return room.Name
		}
	}
	return "You're lost!"
}

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
