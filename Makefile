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
	docker build -f Dockerfile.test . --build-arg CTOKEN=9cb45320-48b3-4fb5-b41f-8886510b5de3 -t goshark && docker run goshark

.PHONY: test
test:
	go test github.com/PacketHelper/goshark/v2/goshark