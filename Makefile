# ************************************************************************* #
#                                                                           #
#                                                      :::      ::::::::    #
#  Makefile                                          :+:      :+:    :+:    #
#                                                  +:+ +:+         +:+      #
#  By: emarette, rruiz, alebaron                 +#+  +:+       +#+         #
#                                              +#+#+#+#+#+   +#+            #
#  Created: 2026/09/14 16:53:12 by alebaron        #+#    #+#               #
#  Updated: 2026/09/19 15:49:45 by rruiz           ###   ########.fr        #
#                                                                           #
# ************************************************************************* #

# ------------------------------------------------------------------------- #
#                                 Variables                                 #
# ------------------------------------------------------------------------- #

IP ?=127.0.0.1
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

run-client:
	./main cli $(IP)

clean:
	rm -rf log
	rm -rf main

lint:
	@which $(LINTER) >/dev/null 2>&1 || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
	$(LINTER) run ./...

.PHONY: build run-server run-client run-client-gui lint