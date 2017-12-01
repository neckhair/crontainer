FROM golang:1.9-alpine as builder

RUN apk add --no-cache --update git curl

ENV DEP_VERSION=0.3.2
RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep && \
    chmod +x $GOPATH/bin/dep

WORKDIR /go/src/github.com/neckhair/crontainer
COPY . .
RUN $GOPATH/bin/dep ensure
RUN GOOS=linux GOARCH=${GOARCH} go-wrapper install


FROM alpine:latest as runtime

COPY examples/no_jobs.yml /etc/crontainer.yml
COPY --from=builder /go/bin/crontainer /usr/local/bin/crontainer

RUN apk add --no-cache --update curl

ENTRYPOINT ["/usr/local/bin/crontainer"]
CMD ["--config", "/etc/crontainer.yml"]
