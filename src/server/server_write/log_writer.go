/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* log_writer.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/24 15:54:02 by alebaron        #+#    #+#              */
/* Updated: 2026/08/25 14:59:55 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package server_write

import (
    "os"
    "net"
    "fmt"
    "log"
    "time"
    "strings"
	"the_answer_protocol/src/utils"
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

    _, _ = file.WriteString("")

    defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Erreur lors de la création du fichier de log")
		}
	}()
}

func WriteLog(conn net.Conn, level string, texte string) {

    // Récupération de l'adresse IP //
    remoteAddr := conn.RemoteAddr().(*net.TCPAddr)
    clientIP := remoteAddr.IP.String()

    file, err := os.OpenFile("log/server_log.txt", os.O_WRONLY|os.O_APPEND, 0755)
    if err != nil {
        utils.ExitError("OPENFILE", err)
    }

    defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Erreur lors de la fermeture du fichier")
		}
	}()

    now := time.Now()
    formattedTime := now.Format("2006-01-02 15:04:05")

    texte = strings.TrimRight(texte, "\n")
    str := fmt.Sprintf("%s %-12s: %-9s %s\n", formattedTime, "(" + clientIP + ")", "[" + level + "]", texte)

    fmt.Println(str[:len(str)-1])
    _, err = file.WriteString(str)
    if err != nil {
        panic(err)
    }
}