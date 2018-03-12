FROM alpine:3.7

LABEL maintainer="phil@neckhair.ch"

COPY config.yml /etc/crontainer.yml
COPY crontainer /usr/local/bin/crontainer

RUN apk add --no-cache --update curl=7.58.0-r1

ENTRYPOINT ["/usr/local/bin/crontainer"]
CMD ["--config", "/etc/crontainer.yml"]
