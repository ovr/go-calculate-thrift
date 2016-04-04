package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
)

func runClient(tFactory thrift.TTransportFactory, pFactory thrift.TProtocolFactory, addr string) (*tutorial.CalculatorClient, error) {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return nil, err
	}

	if transport == nil {
		return nil, fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}

	transport = tFactory.GetTransport(transport)
	if transport == nil {
		return nil, fmt.Errorf("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}

	err = transport.Open()
	if err != nil {
		return nil, err
	}

	return tutorial.NewCalculatorClientFactory(transport, pFactory), nil
}
