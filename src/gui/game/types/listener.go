/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* listener.go                                       :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/03 16:14:59 by rruiz           #+#    #+#              */
/* Updated: 2026/09/09 12:32:31 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */
package types

// Listener is a simple pub/sub system that lets the GUI subscribe to server lines.
type Listener struct {
	functions map[int]func(string)
	nextID    int
}

// Subscribe registers a callback and returns its ID (used later to unsubscribe).
func (l *Listener) Subscribe(functionToAdd func(string)) int {
	if l.functions == nil {
		l.functions = make(map[int]func(string))
	}

	id := l.nextID
	l.functions[id] = functionToAdd
	l.nextID++

	return id
}

// Unsubscribe removes the callback associated with the given ID.
func (l *Listener) Unsubscribe(idToRemove int) {
	delete(l.functions, idToRemove)
}

// Distribute sends a server line to every registered callback.
func (l *Listener) Distribute(line string) {
	for _, function := range l.functions {
		if function != nil {
			function(line)
		}
	}
}
