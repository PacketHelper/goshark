PKG := github.com/PacketHelper/goshark

BIN_DIR := $(CURDIR)/bin
# PKGS := $(shell go list ./... | grep -v internal/cmd|grep -v test)
# COVER_PKGS := $(foreach pkg,$(PKGS),$(subst $(PKG),.,$(pkg)))

COMMA := ,
EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
COVERPKG_OPT := $(subst $(SPACE),$(COMMA),$(COVER_PKGS))

.PHONY: build
build:
	go build -v main.go
.PHONY: cover
cover:
	go test -coverprofile=cover.out ./...

.PHONY: docker
docker:
	docker -t goshark build .

