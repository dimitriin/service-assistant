protoc:
	docker run --rm -it -v `pwd`:/src -w /src jaegertracing/protobuf:latest \
		-I=./.proto --go_out=./pkg/protocol/payload --php_out=./php-service-assistant-client/gen payload.proto
.PHONY: protoc

build:
	go build -o bin/service-assistant ./cmd/service-assistant
.PHONY: build

image:
	docker build -t dimitriin/service-assistant -f ./.docker/service-assistant/Dockerfile .
.PHONY: image

build-assisted-service-example:
	go build -o bin/assisted-service-example ./cmd/assisted-service-example
.PHONY: build-assisted-service-example

image-assisted-service-example:
	docker build -t dimitriin/assisted-service-example -f ./.docker/assisted-service-example/Dockerfile .
.PHONY: image-assisted-service-example