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

type Listener struct {
	functions map[int]func(string)
	nextID    int
}

func (l *Listener) Subscribe(functionToAdd func(string)) int {
	if l.functions == nil {
		l.functions = make(map[int]func(string))
	}

	id := l.nextID
	l.functions[id] = functionToAdd
	l.nextID++

	return id
}

func (l *Listener) Unsubscribe(idToRemove int) {
	delete(l.functions, idToRemove)
}

func (l *Listener) Distribute(line string) {
	for _, function := range l.functions {
		if function != nil {
			function(line)
		}
	}
}
