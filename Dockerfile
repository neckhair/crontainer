FROM frolvlad/alpine-glibc:alpine-3.7

LABEL maintainer="phil@neckhair.ch"
LABEL INSTALL="docker run --rm --privileged -v /:/host -e HOST=/host -e LOGDIR=${LOGDIR} -e CONFDIR=${CONFDIR} -e DATADIR=${DATADIR} --entrypoint /bin/sh -e NAME=NAME -e IMAGE=IMAGE IMAGE /root/atomic/install.sh"
LABEL UNINSTALL="docker run --rm --privileged -v /:/host -e HOST=/host -e LOGDIR=${LOGDIR} -e CONFDIR=${CONFDIR} -e DATADIR=${DATADIR} --entrypoint /bin/sh -e NAME=NAME -e IMAGE=IMAGE IMAGE /root/atomic/uninstall.sh"


COPY config.yml /etc/crontainer/crontainer.yml
COPY bin/crontainer /usr/local/bin/crontainer
COPY atomic /root/atomic

RUN apk add --no-cache --update curl=7.58.0-r1

ENTRYPOINT ["/usr/local/bin/crontainer"]
CMD ["--config", "/etc/crontainer/crontainer.yml"]
