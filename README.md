![Logo](http://dmtry.me/img/logos/go-thrift.png?v1 "Go Calculate Thrift Tutorial")

Calculate RPC server on Go/Thrift
=================================
[![Build Status](https://travis-ci.org/ovr/go-calculate-thrift.svg?branch=master)](https://travis-ci.org/ovr/go-calculate-thrift)

# Features

- Thrift protocol with methods and own structure
- Client and Sever
- Code style check
- Go tests with data provider for server (very simple)
- Dockerfile
- Vagrant ready

# Installation

### OSX

```sh
brew install thrift
thrift -r --gen go calculate.thrift
```

### How to build Server?

```sh
cd src/server
go build .
./server
```

### How to build Client?

```sh
cd src/client
go build .
./client
```

# LICENSE

MIT
