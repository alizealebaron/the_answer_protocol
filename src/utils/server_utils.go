/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* server_utils.go                                   :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: alebaron, ruiz, emarette                  +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/22 14:03:24 by emarette        #+#    #+#              */
/* Updated: 2026/08/22 14:04:14 by emarette        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

package utils

import (
	"net"
	"log"
)

func ServerWrite(conn net.Conn, message string) {
    _, err := conn.Write([]byte(message)) 
    if err != nil {
        log.Printf("Server write error: %v", err)
    }
}