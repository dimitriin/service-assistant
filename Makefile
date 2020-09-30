protoc:
	docker run --rm -it -v `pwd`:/src -w /src jaegertracing/protobuf:latest \
		-I=./ --go_out=./pkg/protocol/cmd ./cmd.proto
.PHONY: protoc

build:
	go build -o bin/service-assistant ./cmd/service-assistant
.PHONY: build

image:
	docker build -t dimitriin/service-assistant .
.PHONY: image