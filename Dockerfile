FROM golang:1.9-alpine as builder

RUN apk add --no-cache --update git=2.13.5-r0 curl=7.57.0-r0

ENV DEP_VERSION=0.3.2
RUN curl -L -s "https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64" -o "$GOPATH/bin/dep" && \
    chmod +x "$GOPATH/bin/dep"

WORKDIR /go/src/github.com/neckhair/crontainer
COPY . .
RUN "$GOPATH/bin/dep" ensure
RUN GOOS=linux GOARCH=${GOARCH} go-wrapper install


FROM alpine:3.7 as runtime

LABEL maintainer="phil@neckhair.ch"

COPY --from=builder /go/bin/crontainer /usr/local/bin/crontainer
COPY config.yml /etc/crontainer.yml

RUN apk add --no-cache --update curl=7.57.0-r0

ENTRYPOINT ["/usr/local/bin/crontainer"]
CMD ["--config", "/etc/crontainer.yml"]
