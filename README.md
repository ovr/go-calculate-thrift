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

First you need to change you directory to the server dir:

```sh
cd src/server
```

Next setup SETUP `$GOPATH`, restore deps and build it:

```sh
godep restore
godep go build
```

And run the server:

```sh
./server
```

### How to build Client?

First you need to change you directory to the client dir:

```sh
cd src/client
```

Next setup SETUP `$GOPATH`, restore deps and build it:

```sh
godep restore
godep go build
```

And run the client:

```sh
./client
```

You can change port and address of the server:

```sh
Usage of ./client:
  -P string
    	Protocol (binary, compact, json, simplejson) (default "binary")
  -addr string
    	Address to listen to (default "localhost:9090")
```

# LICENSE

MIT
