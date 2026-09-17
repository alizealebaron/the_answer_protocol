/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* mute.go                                           :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/16 19:02:45 by rruiz           #+#    #+#              */
/* Updated: 2026/09/16 19:22:27 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */
package types

import "strings"

var mutedPrefixes []string

func Mute(prefixes ...string) {
	mutedPrefixes = append(mutedPrefixes, prefixes...)
}

func Unmute(prefixes ...string) {
	for _, target := range prefixes {
		for i, cur := range mutedPrefixes {
			if cur == target {
				mutedPrefixes = append(mutedPrefixes[:i], mutedPrefixes[i+1:]...)
				break
			}
		}
	}
}

func IsMuted(line string) bool {
	for _, prefix := range mutedPrefixes {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func ResetMute() {
	mutedPrefixes = mutedPrefixes[:0]
}
