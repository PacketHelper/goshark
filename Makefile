.PHONY: build
build:
	go build -v goshark.main

.PHONY: docker
docker:
	docker -t goshark build .
