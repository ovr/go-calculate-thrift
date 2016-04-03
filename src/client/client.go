package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
)

func handleClient(client *tutorial.CalculatorClient) (err error) {
	client.Ping()
	fmt.Println("ping()")
	//
	//sum, _ := client.Add(1, 1)
	//fmt.Print("1+1=", sum, "\n")
	//
	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}

	if transport == nil {
		return fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}

	transport = transportFactory.GetTransport(transport)
	if transport == nil {
		return fmt.Errorf("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}

	err = transport.Open()
	if err != nil {
		return err
	}
	defer transport.Close()

	return handleClient(tutorial.NewCalculatorClientFactory(transport, protocolFactory))
}
