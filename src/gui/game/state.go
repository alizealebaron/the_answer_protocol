/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* state.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/08 23:23:01 by rruiz           #+#    #+#              */
/* Updated: 2026/09/17 21:51:38 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"encoding/json"
	"strings"
	"sync"

	"the_answer_protocol/src/gui/game/types"
)

// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |                                                            Game Data                                                            |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

var gameData types.Secret

var gameDataMutex sync.RWMutex

// Saves the new game data, then refreshes the widgets that use it.
func setGameData(data types.Secret) {
	gameDataMutex.Lock()
	gameData = data
	gameDataMutex.Unlock()

	// The room list is only used to find the map limits.
	computeMapBounds(data.Rooms)
	triggerMapRedraw()
	triggerGroupRedraw()
}

// Returns a copy of the game data.
func getGameData() types.Secret {
	gameDataMutex.RLock()
	defer gameDataMutex.RUnlock()
	return gameData
}

// Listens to the "OK SECRET" lines and stores the data they contain.
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

// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |                                                              Logs                                                               |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

// One list of stored messages per category.
var logMessages []string
var globalMessages []string
var roomMessages []string
var groupMessages []string
var logsMutex sync.RWMutex

// Forwards every incoming line to the log widget.
func subscribeLogs(listener *types.Listener) {
	listener.Subscribe(func(line string) {
		if strings.HasPrefix(line, "OK SECRET ") {
			return
		}
		if types.IsMuted(line) {
			return
		}
		category := getTypeLine(line)
		addMessage(category, line)
		onLogLine(category, line)
	})
}

// Returns the category of a line based on its prefix.
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

// Appends a line to the correct message list.
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

// Returns one category of the stored messages.
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

// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |                                                               Map                                                               |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

var currentRoomId int
var visitedRooms = make(map[int]bool)
var knownRooms = make(map[int]bool)
var mapMutex sync.RWMutex

// Map limits, found the first time the room list is received.
var minX int
var minY int
var maxX int
var maxY int
var boundsComputed bool
var boundsMutex sync.RWMutex

// How well the player knows a room, from unknown to current.
type roomState int

const (
	// iota starts at 0 and adds 1 to each following value
	roomHidden roomState = iota
	roomKnown
	roomVisited
	roomCurrent
)

// Saves the room the player is in and marks it as visited.
func setCurrentRoom(id int) {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	currentRoomId = id
	visitedRooms[id] = true
}

// Marks a neighbor room as known on the map.
func addKnownRoom(id int) {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	if id == 0 {
		return
	}
	knownRooms[id] = true
}

// Returns the id of the room the player is in.
func getCurrentRoom() int {
	mapMutex.RLock()
	defer mapMutex.RUnlock()
	return currentRoomId
}

// Returns the rooms the player has already seen.
func getVisitedRooms() map[int]bool {
	mapMutex.RLock()
	defer mapMutex.RUnlock()
	return visitedRooms
}

// Returns the rooms the player knows about.
func getKnownRooms() map[int]bool {
	mapMutex.RLock()
	defer mapMutex.RUnlock()
	return knownRooms
}

// Listens to the LOOK replies to update the player position.
func subscribeMapData(listener *types.Listener) {
	listener.Subscribe(func(line string) {
		if !isLookReply(line) {
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

// Finds the map limits from the room list, only once.
func computeMapBounds(rooms []types.RoomInfo) {
	boundsMutex.Lock()
	defer boundsMutex.Unlock()

	if boundsComputed || len(rooms) == 0 {
		return
	}

	// Start from the first room, then widen the box with every other room
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

// Tells if the map limits are known, so the widget knows if it can draw.
func areBoundsComputed() bool {
	boundsMutex.RLock()
	defer boundsMutex.RUnlock()
	return boundsComputed
}

// Returns how well the player knows a room, from hidden to current.
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

var mapRedrawFunc func()
var redrawMapMutex sync.RWMutex

// Stores the redraw function the map widget gives us.
func setMapRedrawFunc(f func()) {
	redrawMapMutex.Lock()
	defer redrawMapMutex.Unlock()
	mapRedrawFunc = f
}

// Calls the saved redraw function if one exists.
func triggerMapRedraw() {
	redrawMapMutex.RLock()
	f := mapRedrawFunc
	redrawMapMutex.RUnlock()
	if f != nil {
		f()
	}
}

// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |                                                              Group                                                              |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

var groupRedrawFunc func()
var redrawGroupMutex sync.RWMutex

// Stores the redraw function the group widget gives us.
func setGroupRedrawFunc(f func()) {
	redrawGroupMutex.Lock()
	defer redrawGroupMutex.Unlock()
	groupRedrawFunc = f
}

// Calls the saved redraw function if one exists.
func triggerGroupRedraw() {
	redrawGroupMutex.RLock()
	f := groupRedrawFunc
	redrawGroupMutex.RUnlock()
	if f != nil {
		f()
	}
}

// Returns the group the player is in, matching his id to skip ghost members.
func getPlayerGroup(playerName string) (types.GroupInfo, bool) {
	data := getGameData()
	myId := -1

	for _, p := range data.Players {
		if p.Name == playerName {
			myId = p.Id
			break
		}
	}

	if myId == -1 {
		return types.GroupInfo{}, false
	}
	// A new connection gets a new id, so old disconnected members do not match
	for _, group := range data.Groups {
		for _, player := range group.LstPlayer {
			if player.Id == myId {
				return group, true
			}
		}
	}
	return types.GroupInfo{}, false
}

// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+
// |                                                              Utils                                                              |
// +=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+

// Forgets the previous session.
func resetState() {
	logsMutex.Lock()
	logMessages = nil
	globalMessages = nil
	roomMessages = nil
	groupMessages = nil
	logsMutex.Unlock()

	mapMutex.Lock()
	currentRoomId = 0
	visitedRooms = make(map[int]bool)
	knownRooms = make(map[int]bool)
	mapMutex.Unlock()

	boundsMutex.Lock()
	boundsComputed = false
	boundsMutex.Unlock()

	gameDataMutex.Lock()
	gameData = types.Secret{}
	gameDataMutex.Unlock()
}

// True only if line is LOOK reply.
func isLookReply(line string) bool {
	return strings.HasPrefix(line, "OK {\"id\":") &&
		strings.Contains(line, "\"neighborRoom\"")
}
