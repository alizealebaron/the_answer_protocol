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

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package utils

import (
	"net"
	"log"
    "os"
    "time"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func ServerWrite(conn net.Conn, message string) {
    _, err := conn.Write([]byte(message)) 
    if err != nil {
        log.Printf("Server write error: %v", err)
    }
}

func CreateLogFolder() {

    err := os.MkdirAll("log/", 0755)
    if err != nil {
        ExitError("CREATEDIR", err)
    }

    file, err := os.OpenFile("log/server_log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
    if err != nil {
        ExitError("OPENFILE", err)
    }

    defer file.Close()

}

func WriteLog(ip string, level string, texte string) {

    file, err := os.OpenFile("log/server_log.txt", os.O_WRONLY|os.O_APPEND, 0755)
    if err != nil {
        ExitError("OPENFILE", err)
    }
    defer file.Close()

    now := time.Now()
    formattedTime := now.Format("2006-01-02 15:04:05")

    str := formattedTime + " (" + ip + "): [" + level + "] " + texte + "\n"

    _, err = file.WriteString(str)
    if err != nil {
        panic(err)
    }
}