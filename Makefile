# ************************************************************************* #
#                                                                           #
#                                                      :::      ::::::::    #
#  Makefile                                          :+:      :+:    :+:    #
#                                                  +:+ +:+         +:+      #
#  By: emarette, rruiz, alebaron                 +#+  +:+       +#+         #
#                                              +#+#+#+#+#+   +#+            #
#  Created: 2026/09/14 16:53:12 by alebaron        #+#    #+#               #
#  Updated: 2026/09/14 16:53:14 by alebaron        ###   ########.fr        #
#                                                                           #
# ************************************************************************* #

IP ?=

build:
	go build main.go

run-server: build
	./main server

run-gui: build
	./main gui

run-cli:
	nc $(IP) 8090

clean:
	rm -rf log
	rm -rf main

.PHONY: build run_server run_cli run_gui