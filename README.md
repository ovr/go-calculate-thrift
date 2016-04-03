Calculate RPC server on Go/Thrift
=================================
[![Build Status](https://travis-ci.org/ovr/go-calculate-thrift.svg?branch=master)](https://travis-ci.org/ovr/go-calculate-thrift)

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
