Calculate RPC server on Go/Thrift
=================================

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
