# MCDL UART to TCP

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
