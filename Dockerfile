FROM golang:alpine AS binaryBuilder
# Install build deps
RUN apk --no-cache --no-progress add --virtual build-deps build-base git
WORKDIR /go/src/github.com/okmall/okmall
COPY . .
RUN export GO111MODULE=on && make build

FROM alpine:latest
# Install system utils & Gin-Music runtime dependencies
ADD https://github.com/tianon/gosu/releases/download/1.11/gosu-amd64 /usr/sbin/gosu
RUN chmod +x /usr/sbin/gosu \
  && echo http://dl-2.alpinelinux.org/alpine/edge/community/ >> /etc/apk/repositories \
  && apk --no-cache --no-progress add \
    bash \
    shadow \
    s6

ENV OKMALL_CUSTOM /data/okmall

# Configure LibC Name Service
COPY hack/docker/nsswitch.conf /etc/nsswitch.conf

WORKDIR /app/okmall
COPY hack/docker ./docker
COPY --from=binaryBuilder /go/src/github.com/okmall/okmall/okmall/okmall .

RUN ./docker/finalize.sh

# Configure Docker Container
VOLUME ["/data"]
EXPOSE 8013
ENTRYPOINT ["/app/okmall/docker/start.sh"]
CMD ["/bin/s6-svscan", "/app/okmall/docker/s6/"]