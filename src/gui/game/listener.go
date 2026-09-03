/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* listener.go                                       :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/03 16:14:59 by rruiz           #+#    #+#              */
/* Updated: 2026/09/03 16:33:07 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package game

type Listener struct {
	functions []func(string)
}

func (l *Listener) Subscribe(functionToAdd func(string)) {
	l.functions = append(l.functions, functionToAdd)
}

func (l *Listener) Distribute(line string) {
	for _, function := range l.functions {
		function(line)
	}
}
