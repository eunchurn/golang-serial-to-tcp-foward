# MCDL UART to TCP

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

## Run main code

```bash
go run main.go
```

## Build

```bash
go build main.go
```

### Run built binary

```bash
./main
```
