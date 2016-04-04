package main

import (
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
	"log"
	"os"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	handler := NewCalculatorHandler()
	processor := tutorial.NewCalculatorProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Println("Starting the simple server... on ", addr)
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

	log.Println(*protocol)

	transportFactory := thrift.NewTTransportFactory()
	if err := runServer(transportFactory, protocolFactory, *addr); err != nil {
		log.Println("error running server:", err)
	}
}
