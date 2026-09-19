/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* game_view.go                                      :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 18:10:21 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 21:57:00 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"fmt"
	"io"
	"strings"
	"the_answer_protocol/src/gui/game/types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Custom layout that splits the space into two panes, following one ratio.
type ratioLayout struct {
	ratio      float32
	horizontal bool
	gap        float32
}

func (r *ratioLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if r.horizontal {
		// The gap is taken from the total width, the rest is shared by ratio
		usableWidth := size.Width - r.gap
		width1 := usableWidth * r.ratio
		width2 := usableWidth - width1

		objects[0].Resize(fyne.NewSize(width1, size.Height))
		objects[0].Move(fyne.NewPos(0, 0))

		objects[1].Resize(fyne.NewSize(width2, size.Height))
		objects[1].Move(fyne.NewPos(width1+r.gap, 0))
	} else {
		// Same logic but cut vertically instead
		usableHeight := size.Height - r.gap
		height1 := usableHeight * r.ratio
		height2 := usableHeight - height1

		objects[0].Resize(fyne.NewSize(size.Width, height1))
		objects[0].Move(fyne.NewPos(0, 0))

		objects[1].Resize(fyne.NewSize(size.Width, height2))
		objects[1].Move(fyne.NewPos(0, height1+r.gap))
	}
}

// Required by the layout interface, let children decide their own size.
func (r *ratioLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 0)
}

// Builds a container that splits two objects with the given ratio.
func newRatioSplit(ratio float32, horizontal bool, gap float32, a, b fyne.CanvasObject) *fyne.Container {
	return container.New(&ratioLayout{ratio: ratio, horizontal: horizontal, gap: gap}, a, b)
}

// Builds the whole game screen.
func GameView(window fyne.Window, size fyne.Size, stdin io.WriteCloser, listener *types.Listener, playerName string, language string, backToHome func()) fyne.CanvasObject {
	const gap = float32(8)

	types.SetLanguage(language)

	// Registers the listeners that feed the widgets with the server replies
	subscribeGameData(listener)
	subscribeLogs(listener)
	subscribeMapData(listener)

	// First SECRET to get the whole game, then LOOK to know where the player is
	_, _ = fmt.Fprintf(stdin, "SECRET\n")
	subscribeOnce(listener, "OK SECRET ", func() {
		_, _ = fmt.Fprintf(stdin, "LOOK\n")
	})

	commandBox := commandWidget(stdin, listener, playerName, backToHome)
	scrollBox := logWidget()
	right := newRatioSplit(0.72, false, gap, commandBox, scrollBox)

	mapBox := MapWidget()

	whathappened := goingOnWidget(listener)
	playersLabel := playerCountLabel(listener)
	topleft := newRatioSplit(0.83, false, gap, whathappened, playersLabel)

	groupBox := groupWidget(stdin, listener, playerName)
	left := newRatioSplit(0.6, false, gap, topleft, groupBox)

	centerRight := newRatioSplit(0.571, true, gap, mapBox, right)

	return newRatioSplit(0.3, true, gap, left, centerRight)
}

// Runs fn once, on the first line starting with prefix, then unsubscribes.
func subscribeOnce(listener *types.Listener, prefix string, fn func()) {
	var id int
	id = listener.Subscribe(func(line string) {
		if strings.HasPrefix(line, prefix) {
			fn()
			listener.Unsubscribe(id)
		}
	})
}
