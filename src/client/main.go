package main

import (
	"flag"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"os"
)

func main()  {
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

	ransportFactory := thrift.NewTTransportFactory()
	if err := runClient(ransportFactory, protocolFactory, *addr); err != nil {
		fmt.Println("error running client:", err)
	}
}
