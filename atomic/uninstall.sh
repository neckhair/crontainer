#!/usr/bin/env sh

# Disable and remove systemd unit file
chroot ${HOST} /usr/bin/systemctl stop crontainer_${NAME}.service
chroot ${HOST} /usr/bin/systemctl disable crontainer_${NAME}.service
chroot ${HOST} rm -f /etc/systemd/system/crontainer_${NAME}.service