/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* state.go                                          :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/08 23:23:01 by rruiz           #+#    #+#              */
/* Updated: 2026/09/08 23:39:36 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

import (
	"encoding/json"
	"strings"
	"sync"
)

var gameData Secret
var gameDataMutex sync.RWMutex

func setGameData(data Secret) {
	gameDataMutex.Lock()
	defer gameDataMutex.Unlock()
	gameData = data
}

func getGameData() Secret {
	gameDataMutex.RLock()
	defer gameDataMutex.RUnlock()
	return gameData
}

func subscribeGameData(listener *Listener) {
	listener.Subscribe(func(line string) {
		if !strings.HasPrefix(line, "OK {\"lst_item\"") {
			return
		}
		raw := strings.TrimPrefix(line, "OK ")
		var data Secret
		if err := json.Unmarshal([]byte(raw), &data); err != nil {
			return
		}
		setGameData(data)
	})
}
