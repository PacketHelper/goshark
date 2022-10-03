.PHONY: build
build:
	go build -v main.go

.PHONY: cover
cover:
	go test -coverprofile=cover.out github.com/PacketHelper/goshark/v2/goshark

.PHONY: docker
docker:
	docker build . -t goshark 

.PHONY: docker-test
docker-test:
	docker build -f Dockerfile.test .  -t goshark

.PHONY: test
test:
	go test github.com/PacketHelper/goshark/v2/goshark