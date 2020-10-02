ARG GOLANG_VERSION=1.15.2
FROM golang:${GOLANG_VERSION} as builder
COPY . /src
WORKDIR /src
RUN make build-assisted-service-example

FROM debian:buster-slim
RUN groupadd --gid 2000 service-assistant \
  && useradd --uid 2000 --gid service-assistant --shell /bin/bash --create-home service-assistant

WORKDIR /home/service-assistant

COPY --chown=service-assistant:service-assistant --from=builder /src/bin/assisted-service-example ./.service-assistant/bin/assisted-service-example
CMD ["./.service-assistant/bin/assisted-service-example"]

EXPOSE 8080