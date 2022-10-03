.PHONY: build
build:
	go build -v main.go

.PHONY: cover
cover:
	go test -coverprofile=cover.out github.com/PacketHelper/goshark/v2/goshark

.PHONY: docker
docker:
	docker -t goshark build .

.PHONY: test
test:
	go test github.com/PacketHelper/goshark/v2/goshark