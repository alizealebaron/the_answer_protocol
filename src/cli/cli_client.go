/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* cli_client.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/19 10:00:00 by alebaron        #+#    #+#              */
/* Updated: 2026/09/19 13:12:59 by rruiz           ###   ########.fr       */
/*                                                                         */
/* *********************************************************************** */

/* ----------------------------------------------------------------------- */
/*                            Package & Import                             */
/* ----------------------------------------------------------------------- */

package cli

import (
	"fmt"
	"io"
	"net"
	"os"
)

const Port = "8090"

/* ----------------------------------------------------------------------- */
/*                             Fonction Principal                          */
/* ----------------------------------------------------------------------- */

func Client(adress string) {
	conn, err1 := net.Dial("tcp", net.JoinHostPort(adress, Port))
	if err1 != nil {
		fmt.Println("Cannot connect to", adress+":"+Port)
		os.Exit(1)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			return
		}
	}()

	fmt.Println("Connected at", adress+":"+Port)
	fmt.Println("Use QUIT to exit.")

	channel := make(chan error, 2)

	go func() {
		_, err2 := io.Copy(os.Stdout, conn)
		channel <- err2
	}()

	go func() {
		_, err3 := io.Copy(conn, os.Stdin)
		channel <- err3
	}()

	<-channel
	_ = conn.Close()
}
