FROM golang:1.9.1 as builder

COPY /cmd /go/cmd
COPY /crontainer /go/crontainer
COPY main.go /go/main.go

# go get triggers a go install that errors, but the required libs are still installed
# cgo parameters here are for building a statically linked binary
RUN cd /go && go get; CGO_ENABLED=0 GOOS=linux  go build -a -ldflags '-extldflags "-static"' -o /crontainer .

FROM alpine:3.6

ARG GCRON_VERSION=0.2.0

COPY examples/no_jobs.yml /etc/crontainer.yml
COPY --from=builder /crontainer /usr/local/bin/crontainer

RUN apk add --no-cache --update curl && \
    chmod +x /usr/local/bin/crontainer

ENTRYPOINT ["crontainer"]
CMD ["--config", "/etc/gcron.yml"]
