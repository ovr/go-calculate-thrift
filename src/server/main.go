package main

import (
	"flag"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
	"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
)


func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	processor := tutorial.NewCalculatorProcessor(NewCalculatorHandler())
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}

func main() {
	protocol := flag.String("P", "binary", "Protocol (binary, compact, json, simplejson)")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")

	flag.Parse()

	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "binary":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol", protocol)
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	if err := runServer(transportFactory, protocolFactory, *addr); err != nil {
		fmt.Println("error running server:", err)
	}
}
