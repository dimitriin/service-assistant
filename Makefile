gen:
	docker run --rm -it -v `pwd`:/src -w /src jaegertracing/protobuf:latest \
		-I=./ --go_out=./pkg/protocol/models ./models.proto
.PHONY: gen