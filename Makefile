PKG := github.com/PacketHelper/goshark

BIN_DIR := $(CURDIR)/bin
PKGS := $(shell go list ./... | grep -v internal/cmd|grep -v test)
COVER_PKGS := $(foreach pkg,$(PKGS),$(subst $(PKG),.,$(pkg)))

COMMA := ,
EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
COVERPKG_OPT := $(subst $(SPACE),$(COMMA),$(COVER_PKGS))

.PHONY: build
build:
	go build -v goshark.main

.PHONY: cover
cover:
	go test -coverpkg=$(COVERPKG_OPT) -coverprofile=cover.out ./...

.PHONY: docker
docker:
	docker -t goshark build .

