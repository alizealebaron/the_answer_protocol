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

var currentRoomId int
var visitedRooms = make(map[int]bool)
var knownRooms = make(map[int]bool)
var mapMutex sync.RWMutex

var minX int
var minY int
var maxX int
var maxY int
var boundsComputed bool
var boundsMutex sync.RWMutex

func setGameData(data types.Secret) {
	gameDataMutex.Lock()
	gameData = data
	gameDataMutex.Unlock()

	computeMapBounds(data.Rooms)
	triggerMapRedraw()
}

type roomState int

const (
	// iota simplifies the creation of incremented constants in a const block
	roomHidden roomState = iota
	roomKnown
	roomVisited
	roomCurrent
)

var mapRedrawFunc func()
var redrawMutex sync.RWMutex

func getGameData() types.Secret {
	gameDataMutex.RLock()
	defer gameDataMutex.RUnlock()
	return gameData
}

func subscribeGameData(listener *types.Listener) {
	listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK SECRET ") {
			return
		}
		raw := strings.TrimPrefix(line, "OK SECRET ")
		var data types.Secret
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			return
		}
		setGameData(data)
	})
}

func subscribeLogs(listener *types.Listener) {
	listener.Subscribe(func(line string) {
		if strings.HasPrefix(line, "OK SECRET ") {
			return
		}
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

func setCurrentRoom(id int) {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	currentRoomId = id
	visitedRooms[id] = true
}

func addKnownRoom(id int) {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	if id == 0 {
		return
	}
	knownRooms[id] = true
}

func getCurrentRoom() int {
	mapMutex.RLock()
	defer mapMutex.RUnlock()
	return currentRoomId
}

func getVisitedRooms() map[int]bool {
	mapMutex.RLock()
	defer mapMutex.RUnlock()
	return visitedRooms
}

func getKnownRooms() map[int]bool {
	mapMutex.RLock()
	defer mapMutex.RUnlock()
	return knownRooms
}

func subscribeMapData(listener *types.Listener) {
	listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"id\":") {
			return
		}
		raw := strings.TrimPrefix(line, "OK ")
		var data types.LookInfo
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			return
		}
		setCurrentRoom(data.Id)

		addKnownRoom(data.NeighborRoom.North)
		addKnownRoom(data.NeighborRoom.East)
		addKnownRoom(data.NeighborRoom.South)
		addKnownRoom(data.NeighborRoom.West)

		triggerMapRedraw()
	})
}

func computeMapBounds(rooms []types.RoomInfo) {
	boundsMutex.Lock()
	defer boundsMutex.Unlock()

	if boundsComputed || len(rooms) == 0 {
		return
	}

	minX, maxX = rooms[0].X, rooms[0].X
	minY, maxY = rooms[0].Y, rooms[0].Y

	for _, room := range rooms[1:] {
		if room.X < minX {
			minX = room.X
		}
		if room.X > maxX {
			maxX = room.X
		}
		if room.Y < minY {
			minY = room.Y
		}
		if room.Y > maxY {
			maxY = room.Y
		}
	}

	boundsComputed = true
}

func getMapBounds() (int, int, int, int) {
	boundsMutex.RLock()
	defer boundsMutex.RUnlock()
	return minX, maxX, minY, maxY
}

func areBoundsComputed() bool {
	boundsMutex.RLock()
	defer boundsMutex.RUnlock()
	return boundsComputed
}

func getRoomState(id int) roomState {
	if id == getCurrentRoom() {
		return roomCurrent
	}
	if getVisitedRooms()[id] {
		return roomVisited
	}
	if getKnownRooms()[id] {
		return roomKnown
	}
	return roomHidden
}

func setMapRedrawFunc(f func()) {
	redrawMutex.Lock()
	defer redrawMutex.Unlock()
	mapRedrawFunc = f
}

func triggerMapRedraw() {
	redrawMutex.RLock()
	f := mapRedrawFunc
	redrawMutex.RUnlock()
	if f != nil {
		f()
	}
}
