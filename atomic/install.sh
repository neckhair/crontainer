#!/usr/bin/env sh


chroot ${HOST} mkdir /etc/crontainer
# Install systemd unit file for running container
sed -e "s/NAME/${NAME}/g" /root/atomic/crontainer.service &> ${HOST}/etc/systemd/system/crontainer_${NAME}.service

# Enabled systemd unit file
chroot ${HOST} /usr/bin/systemctl daemon-reload
chroot ${HOST} /usr/bin/systemctl enable /etc/systemd/system/crontainer_${NAME}.service