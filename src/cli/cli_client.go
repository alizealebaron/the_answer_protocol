/* *********************************************************************** */
/*                                                                         */
/*                                                     :::      ::::::::   */
/* cli_client.go                                     :+:      :+:    :+:   */
/*                                                 +:+ +:+         +:+     */
/* By: emarette, rruiz, alebaron                 +#+  +:+       +#+        */
/*                                             +#+#+#+#+#+   +#+           */
/* Created: 2026/09/19 10:00:00 by alebaron        #+#    #+#              */
/* Updated: 2026/09/19 12:56:48 by rruiz           ###   ########.fr       */
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
	conn, err := net.Dial("tcp", net.JoinHostPort(adress, Port))
	if err != nil {
		fmt.Println("Cannot connect to", adress+":"+Port)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connected at", adress+":"+Port)
	fmt.Println("Use QUIT to exit.")

	channel := make(chan error, 2)

	go func() {
		_, err = io.Copy(os.Stdout, conn)
		channel <- err
	}()

	go func() {
		_, err = io.Copy(conn, os.Stdin)
		channel <- err
	}()

	<-channel
	conn.Close()
}
