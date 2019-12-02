#!/bin/bash

curl -o- https://raw.githubusercontent.com/eunchurn/golang-serial-to-tcp-foward/master/scripts/serial-app.service \
> /etc/systemd/system/serial-app.service

systemctl enable serial-app
systemctl start serial-app
