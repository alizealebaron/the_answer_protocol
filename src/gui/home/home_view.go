/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* home_view.go                                      :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/21 18:10:17 by rruiz           #+#    #+#              */
/* Updated: 2026/09/09 10:57:32 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package home

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
	"the_answer_protocol/src/gui/game"
	"the_answer_protocol/src/gui/game/types"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomeView(window fyne.Window, size fyne.Size) fyne.CanvasObject {
	// Definition of widget text to change based on the language selection button.
	var translations = map[string]map[string]string{
		"Français": {
			"selectLang": "Choisir la langue",
			"enterName":  "Entrez votre nom",
			"enterIP":    "Entrez l'IP du serveur",
			"joinServer": "Rejoindre le serveur",
		},
		"English": {
			"selectLang": "Select language",
			"enterName":  "Enter your name",
			"enterIP":    "Enter the ip of the game",
			"joinServer": "Join server",
		},
	}

	// Defining a variable to be used later, done this way to maintain a logical order.
	var langSelect *widget.Select
	var nameField *widget.Entry
	var ipField *widget.Entry
	var joinButton *widget.Button

	// Set variables at the value of the window
	width, height := size.Width, size.Height

	// Set up the error widget once so that it hides immediately, to avoid having to recreate it every time an error occurs.
	errContent, errText := errorWidget("")
	errContent.Resize(fyne.NewSize(width/5, height/10))
	errContent.Move(fyne.NewPos(1, -height))
	errContent.Hide()

	// Language Selection Drop-Down Menu
	// It changes the values of the placeholders in the other widgets based on the value of the selection.
	langSelect = widget.NewSelect([]string{"Français", "English"},
		func(value string) {
			t := translations[value]
			langSelect.PlaceHolder = t["selectLang"]
			nameField.SetPlaceHolder(t["enterName"])
			ipField.SetPlaceHolder(t["enterIP"])
			joinButton.SetText(t["joinServer"])
		})
	langSelect.PlaceHolder = "Select language"

	// Text field for the player's name.
	nameField = widget.NewEntry()
	nameField.SetPlaceHolder("Enter your name")

	// Game Server IP Field.
	ipField = widget.NewEntry()
	ipField.SetPlaceHolder("Enter the ip of the game")

	// Button to join the server
	joinButton = widget.NewButton("Join server", func() {
		// When the button is pressed, this happens.
		// Retrieving the values of nameField and ipField.
		ip := ipField.Text
		name := nameField.Text

		// If any of the fields are empty, the error widget is displayed.
		if ip == "" || name == "" {
			if ip == "" {
				displayError(errText, errContent, "Ip must not be empty.", size)
			} else {
				displayError(errText, errContent, "Name must not be empty.", size)
			}
			return
		}

		matched, _ := regexp.MatchString("^[a-zA-Z0-9_]{1,15}$", name)
		if !matched {
			displayError(errText, errContent, "Invalid name format.", size)
			return
		}

		// A goroutine: It's like a thread, but lighter.
		go func() {
			// Create the command to connect to the server using 'nc'.
			// If an error occurs, it is print properly.
			cmd := exec.Command("nc", ip, "8090")
			stdin, err := cmd.StdinPipe()
			// If StdinPipe return a error
			if err != nil {
				displayError(errText, errContent, "System error: failed to create input pipe.", size)
				return
			}

			stdout, err := cmd.StdoutPipe()
			if err != nil {
				displayError(errText, errContent, "System error: failed to create input pipe.", size)
				return
			}

			listener := &types.Listener{}

			// If Start() return a error.
			if err := cmd.Start(); err != nil {
				displayError(errText, errContent, "Unable to use netcat to connect to the network.", size)
				return
			}
			fmt.Println(strings.Join(cmd.Args, " "))

			// Retrieving the value of langSelect, if it nil, the language is english.
			var language string
			if langSelect.Selected == "Français" {
				language = "FR"
			} else {
				language = "EN"
			}

			reader := bufio.NewReader(stdout)
			line, _ := reader.ReadString('\n')
			line = strings.TrimSpace(line)

			// Sends the command “CONNECT <name> <language>” to the server via stdin.
			fmt.Fprintf(stdin, "CONNECT %s %s\n", name, language)
			//Type the command “CONNECT <name> <language>” in the terminal.
			fmt.Println("CONNECT", name, language)

			line, _ = reader.ReadString('\n')
			line = strings.TrimSpace(line)

			if strings.HasPrefix(line, "OK connected") {
				go stdoutListening(stdout, listener)
				go fyne.Do(func() {
					window.SetContent(game.GameView(window, stdin, listener, name))
				})
			} else {
				displayError(errText, errContent, "Error, during connection to the network.", size)
				fmt.Println(strings.TrimSpace(line))
				cmd.Process.Kill()
				return
			}

			// Blocks the goroutine until netcat finishes. If it returns an error, it means that nc didn't finish properly.
			if err := cmd.Wait(); err != nil {
				displayError(errText, errContent, "Connection to the server failed.", size)
				return
			}
		}()
	})

	// Create a container that holds all the widgets created earlier, resize it, and place it at the bottom center of the window.
	big := container.NewCenter(
		container.NewGridWrap(fyne.NewSize(width/3, height/10), langSelect, nameField, ipField, joinButton),
	)

	// Ensures that errContent is not resized when it is returned.
	err := container.NewWithoutLayout(errContent)

	// Create button and block the resize when it is returned.
	quitButton := quitButton(window, width, height)
	quit := container.NewWithoutLayout(quitButton)

	// return the container at the good place.
	return container.NewStack(container.NewBorder(nil, big, nil, nil), err, quit)
}

func quitButton(window fyne.Window, width float32, height float32) *widget.Button {
	// Function that returns a clickable button that closes the window.
	quitButton := widget.NewButton("X", func() {
		window.Close()
	})
	// Turn the button red.
	quitButton.Importance = widget.DangerImportance
	quitButton.Resize(fyne.NewSize(width/20, height/20))
	quitButton.Move(fyne.NewPos(width-1-width/19, 0))

	return quitButton
}

func stdoutListening(stdout io.ReadCloser, listener *types.Listener) {
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		fyne.Do(func() {
			listener.Distribute(line)
		})
	}
	if err := scanner.Err(); err != nil {
		fyne.Do(func() {
			fmt.Println("Connection lost:", err)
		})
	}
}
