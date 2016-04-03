Calculate RPC server on Go/Thrift
=================================

# Installation

### OSX

```
brew install thrift
thrift -r --gen go calculate.thrift
```

### How to build Server?

```
cd src/server
go build .
./server
```

### How to build Client?

```
cd src/client
go build .
./client
```
