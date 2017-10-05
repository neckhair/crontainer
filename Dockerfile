FROM alpine:3.5

ARG GCRON_VERSION=0.1

COPY examples/no_jobs.yml /etc/crontainer.yml

RUN    apk add --no-cache --update curl \
    && curl -L -o /usr/local/bin/crontainer "https://github.com/neckhair/crontainer/releases/download/${GCRON_VERSION}/crontainer-linux-386" \
    && chmod +x /usr/local/bin/crontainer

ENTRYPOINT ["crontainer"]
CMD ["--config", "/etc/gcron.yml"]
