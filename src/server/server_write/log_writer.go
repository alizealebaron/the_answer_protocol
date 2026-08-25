/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* log_writer.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/24 15:54:02 by alebaron        #+#    #+#              */
/* Updated: 2026/08/25 13:20:47 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package server_write

import (
	"the_answer_protocol/src/utils"
    "os"
    "net"
    "time"
    "strings"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func CreateLogFolder() {

    err := os.MkdirAll("log/", 0755)
    if err != nil {
        utils.ExitError("CREATEDIR", err)
    }

    file, err := os.OpenFile("log/server_log.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
    if err != nil {
        utils.ExitError("OPENFILE", err)
    }

    _, err = file.WriteString("")
    defer file.Close()
}

func WriteLog(conn net.Conn, level string, texte string) {

    // Récupération de l'adresse IP //
    remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
    clientIP := remoteAddr.IP.String()

    file, err := os.OpenFile("log/server_log.txt", os.O_WRONLY|os.O_APPEND, 0755)
    if err != nil {
        utils.ExitError("OPENFILE", err)
    }
    defer file.Close()

    now := time.Now()
    formattedTime := now.Format("2006-01-02 15:04:05")

    texte = strings.TrimRight(texte, "\n")
    str := formattedTime + " (" + clientIP + "): [" + level + "] " + texte + "\n"

    _, err = file.WriteString(str)
    if err != nil {
        panic(err)
    }
}