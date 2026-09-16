/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* game_view.go                                      :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 18:10:21 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 14:54:07 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"fmt"
	"io"
	"strings"
	"the_answer_protocol/src/gui/game/types"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ratioLayout struct {
	ratio      float32
	horizontal bool
	gap        float32
}

func (r *ratioLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if r.horizontal {
		usableWidth := size.Width - r.gap
		width1 := usableWidth * r.ratio
		width2 := usableWidth - width1

		objects[0].Resize(fyne.NewSize(width1, size.Height))
		objects[0].Move(fyne.NewPos(0, 0))

		objects[1].Resize(fyne.NewSize(width2, size.Height))
		objects[1].Move(fyne.NewPos(width1+r.gap, 0))
	} else {
		usableHeight := size.Height - r.gap
		height1 := usableHeight * r.ratio
		height2 := usableHeight - height1

		objects[0].Resize(fyne.NewSize(size.Width, height1))
		objects[0].Move(fyne.NewPos(0, 0))

		objects[1].Resize(fyne.NewSize(size.Width, height2))
		objects[1].Move(fyne.NewPos(0, height1+r.gap))
	}
}

func (r *ratioLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}

func newRatioSplit(ratio float32, horizontal bool, gap float32, a, b fyne.CanvasObject) *fyne.Container {
	return container.New(&ratioLayout{ratio: ratio, horizontal: horizontal, gap: gap}, a, b)
}

func GameView(window fyne.Window, size fyne.Size, stdin io.WriteCloser, listener *types.Listener, playerName string, backToHome func()) fyne.CanvasObject {
	const gap = float32(8)

	subscribeGameData(listener)
	subscribeLogs(listener)
	subscribeMapData(listener)

	fmt.Fprintf(stdin, "SECRET\n")
	subscribeOnce(listener, "OK SECRET ", func() {
		fmt.Fprintf(stdin, "LOOK\n")
	})

	commandBox := commandWidget(stdin, listener, playerName, backToHome)
	scrollBox := logWidget()
	right := newRatioSplit(0.72, false, gap, commandBox, scrollBox)

	mapBox := MapWidget()

	whathappened := goingOnWidget()
	playersLabel := playerCountLabel(listener)
	topleft := newRatioSplit(0.83, false, gap, whathappened, playersLabel)

	groupBox := groupWidget()
	left := newRatioSplit(0.6, false, gap, topleft, groupBox)

	centerRight := newRatioSplit(0.571, true, gap, mapBox, right)

	return newRatioSplit(0.3, true, gap, left, centerRight)
}

func subscribeOnce(listener *types.Listener, prefix string, fn func()) {
	var id int
	id = listener.Subscribe(func(line string) {
		if strings.HasPrefix(line, prefix) {
			fn()
			listener.Unsubscribe(id)
		}
	})
}

func actionWidget() *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	content := container.NewVBox()
	scroll := container.NewScroll(content)

	return container.NewStack(frame, scroll)
}

func goingOnWidget() *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	label := widget.NewLabel("Texte juste pour tester que la longueur de ma chaine fasse bien et que ca wrap bien :)")
	label.Alignment = fyne.TextAlignCenter

	square := canvas.NewRectangle(color.Transparent)
	square.StrokeColor = color.White
	square.StrokeWidth = float32(2)

	border := container.NewBorder(label, nil, nil, nil, square)
	return container.NewStack(frame, border)
}

func groupWidget() *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	return container.NewStack(frame)
}
