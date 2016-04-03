package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
)

func handleClient(client *tutorial.CalculatorClient) (err error) {
	var (
		requestErr error
		result int32
	)

	fmt.Println("Send ping()")

	requestErr = client.Ping()
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}

	fmt.Println("Send Plus(10, 10)")
	result, requestErr = client.Plus(10, 10)
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result)

	fmt.Println("Send Minus(25, 10)")
	result, requestErr = client.Minus(25, 10)
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result)


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
