/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* state.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/08 23:23:01 by rruiz           #+#    #+#              */
/* Updated: 2026/09/15 14:35:21 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"encoding/json"
	"strings"
	"sync"

	"the_answer_protocol/src/gui/game/types"
)

var gameData types.Secret
var gameDataMutex sync.RWMutex

var logMessages []string
var globalMessages []string
var roomMessages []string
var groupMessages []string
var logsMutex sync.RWMutex

func setGameData(data types.Secret) {
	gameDataMutex.Lock()
	defer gameDataMutex.Unlock()
	gameData = data
}

func getGameData() types.Secret {
	gameDataMutex.RLock()
	defer gameDataMutex.RUnlock()
	return gameData
}

func subscribeGameData(listener *types.Listener) {
	listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"lst_item\"") {
			return
		}
		raw := strings.TrimPrefix(line, "OK ")
		var data types.Secret
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			return
		}
		setGameData(data)
	})
}

func subscribeLogs(listener *types.Listener) {
	listener.Subscribe(func(line string) {
		category := getTypeLine(line)
		addMessage(category, line)
		onLogLine(category, line)
	})
}

func getTypeLine(line string) string {
	switch {
	case strings.HasPrefix(line, "EVT GLOBAL"):
		return "GLOBAL"
	case strings.HasPrefix(line, "EVT ROOM"):
		return "ROOM"
	case strings.HasPrefix(line, "EVT GROUP"):
		return "GROUP"
	default:
		return "LOG"
	}
}

func addMessage(category string, line string) {
	logsMutex.Lock()
	defer logsMutex.Unlock()
	switch category {
	case "GLOBAL":
		globalMessages = append(globalMessages, line)
	case "ROOM":
		roomMessages = append(roomMessages, line)
	case "GROUP":
		groupMessages = append(groupMessages, line)
	default:
		logMessages = append(logMessages, line)
	}
}

func getMessages(category string) []string {
	logsMutex.RLock()
	defer logsMutex.RUnlock()
	switch category {
	case "GLOBAL":
		return globalMessages
	case "ROOM":
		return roomMessages
	case "GROUP":
		return groupMessages
	default:
		return logMessages
	}
}
