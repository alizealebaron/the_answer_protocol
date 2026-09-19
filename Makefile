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

# ------------------------------------------------------------------------- #
#                                 Variables                                 #
# ------------------------------------------------------------------------- #

IP ?=
GOPATH := $(shell go env GOPATH)
PATH := $(GOPATH)/bin:$(PATH)
LINTER := $(GOPATH)/bin/golangci-lint

# ------------------------------------------------------------------------- #
#                                 Commandes                                 #
# ------------------------------------------------------------------------- #

build:
	go build main.go

run-server: build
	./main server

run-client-gui: build
	./main gui

run-cli:
	nc $(IP) 8090

clean:
	rm -rf log
	rm -rf main

lint:
	@which $(LINTER) >/dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
	$(LINTER) run ./...

.PHONY: build run-server run-cli run-client-gui lint