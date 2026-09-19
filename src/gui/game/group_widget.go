/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* group_widget.go                                    :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/17 20:59:38 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 21:50:20 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"fmt"
	"image/color"
	"io"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"the_answer_protocol/src/gui/game/types"
)

// Displays whether the player is in a group and its members.
func groupWidget(stdin io.WriteCloser, listener *types.Listener, playerName string) *fyne.Container {
	frame := canvas.NewRectangle(color.Transparent)
	frame.StrokeColor = color.White
	frame.StrokeWidth = float32(2)

	content := container.NewVBox()
	box := container.NewStack(frame, content)

	listener.Subscribe(func(line string) {
		if strings.HasPrefix(line, "EVT GROUP") {
			_, _ = fmt.Fprintf(stdin, "SECRET\n")
		}
	})

	// Rebuilds the whole content with the current group data.
	redraw := func() {
		content.RemoveAll()

		group, isOk := getPlayerGroup(playerName)
		if !isOk {
			label := widget.NewLabel(types.Translate("Not in a group"))
			label.Alignment = fyne.TextAlignCenter
			content.Add(label)
			content.Refresh()
			return
		}

		// Title with a bigger font than the members
		title := widget.NewRichText(&widget.TextSegment{
			Style: widget.RichTextStyle{
				Alignment: fyne.TextAlignCenter,
				SizeName:  theme.SizeNameHeadingText,
			},
			Text: fmt.Sprintf(types.Translate("In group (id=%d)"), group.Id),
		})
		title.Wrapping = fyne.TextWrapWord
		content.Add(title)

		// One label per member, with "(you)" next to the player himself
		for _, member := range group.LstPlayer {
			memberName := member.Name
			text := fmt.Sprintf("  - %s", memberName)
			if memberName == playerName {
				text += types.Translate("(you)")
			}
			memberLabel := widget.NewRichText(&widget.TextSegment{
				Style: widget.RichTextStyle{SizeName: theme.SizeNameSubHeadingText},
				Text:  text,
			})
			memberLabel.Wrapping = fyne.TextWrapWord
			content.Add(memberLabel)
		}

		content.Refresh()
	}

	setGroupRedrawFunc(redraw)
	redraw()

	return box
}
