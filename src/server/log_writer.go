/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* log_writer.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/08/24 15:54:02 by alebaron        #+#    #+#              */
/* Updated: 2026/08/24 17:01:20 by alebaron        ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* +---------------------------------------------------------------------+ */
/* |                          Package & Import                           | */
/* +---------------------------------------------------------------------+ */

package server

import (
	"the_answer_protocol/src/utils"
    "os"
    "time"
)

/* +---------------------------------------------------------------------+ */
/* |                             Fonctions                               | */
/* +---------------------------------------------------------------------+ */

func CreateLogFolder() {

    err := os.MkdirAll("log/", 0755)
    if err != nil {
        utils.ExitError("CREATEDIR", err)
    }

    file, err := os.OpenFile("log/server_log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
    if err != nil {
        utils.ExitError("OPENFILE", err)
    }

    defer file.Close()

}

func WriteLog(ip string, level string, texte string) {

    file, err := os.OpenFile("log/server_log.txt", os.O_WRONLY|os.O_APPEND, 0755)
    if err != nil {
        utils.ExitError("OPENFILE", err)
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