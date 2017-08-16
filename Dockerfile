FROM alpine:3.5

ENV GCRON_VERSION=0.1

RUN    apk add --update curl \
    && curl -L -o /usr/local/bin/crontainer "https://github.com/neckhair/crontainer/releases/download/${GCRON_VERSION}/crontainer-linux-386" \
    && chmod +x /usr/local/bin/crontainer \
    && rm -rf /var/cache/apk/*

ENTRYPOINT ["crontainer"]
