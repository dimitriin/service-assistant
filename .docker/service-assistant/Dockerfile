ARG GOLANG_VERSION=1.15.2
FROM golang:${GOLANG_VERSION} as builder
COPY . /src
WORKDIR /src
RUN make build


ARG HTTP_PORT=8181
ARG UDP_PORT=8282
FROM debian:buster-slim
RUN groupadd --gid 2000 service-assistant \
  && useradd --uid 2000 --gid service-assistant --shell /bin/bash --create-home service-assistant

WORKDIR /home/service-assistant

COPY --chown=service-assistant:service-assistant --from=builder /src/bin/service-assistant ./.service-assistant/bin/service-assistant
CMD ["./.service-assistant/bin/service-assistant"]

EXPOSE ${HTTP_PORT} ${UDP_PORT}