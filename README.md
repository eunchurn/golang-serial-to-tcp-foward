# MCDL UART to TCP

[![Actions Status](https://github.com/eunchurn/golang-serial-to-tcp-foward/workflows/Go/badge.svg)](https://github.com/eunchurn/golang-serial-to-tcp-foward/actions) [![Build Status](https://travis-ci.org/eunchurn/golang-serial-to-tcp-foward.svg?branch=master)](https://travis-ci.org/eunchurn/golang-serial-to-tcp-foward)

Hard coded device `/dev/ttySAC4` baudrate is `2500000` for APROS-MCDL fixed peripheral as DSP to UART. This app is simple UART to TCP server fowarding (fixed port: `8001`) application.

## Install golang in ARTIK710

```bash
wget https://dl.google.com/go/go1.13.4.linux-armv6l.tar.gz
tar -C /usr/local -xzf go1.13.4.linux-armv6l.tar.gz
mkdir $HOME/go
export GOROOT="/usr/local/go"
export GOPATH="/root/go"
export PATH=$PATH:/usr/local/go/bin
go get github.com/tarm/serial
```

### `dep` install

```bash
go get -d -u github.com/golang/dep
cd $(go env GOPATH)/src/github.com/golang/dep
DEP_LATEST=$(git describe --abbrev=0 --tags)
git checkout $DEP_LATEST
go install -ldflags="-X main.version=$DEP_LATEST" ./cmd/dep
git checkout master
```

### Install service

```bash
curl -o- https://raw.githubusercontent.com/eunchurn/golang-serial-to-tcp-foward/master/scripts/serial-app.service \
| bash
```

## Run main code

```bash
go run serial-app.go
```

## Build

```bash
go build serial-app.go
```

### Run built binary

```bash
./serial-app
```
