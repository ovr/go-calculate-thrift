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
		result64 int64
		result64f float64
	)

	fmt.Println("Send ping()")

	requestErr = client.Ping()
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}

	fmt.Println("Send Plus(10, 10)")
	result64, requestErr = client.Plus(10, 10)
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result64)

	fmt.Println("Send Minus(25, 10)")
	result, requestErr = client.Minus(25, 10)
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result)

	fmt.Println("Send Div(25, 5)")
	result, requestErr = client.Div(25, 5)
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result)

	fmt.Println("Send Mul(5, 5)")
	result64, requestErr = client.Mul(5, 5)
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result64)

	fmt.Println("Send Mod(5, 5)")
	result64f, requestErr = client.Mod(5.0, 5.0)
	if requestErr != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result64f)

	work := tutorial.NewWork()

	fmt.Println("Send Work(Operation_ADD, 10, 25)")
	work.Op = tutorial.Operation_ADD
	work.Num1 = 10
	work.Num2 = 15

	result, err = client.Calculate(work)
	if err != nil {
		fmt.Println(requestErr.Error())
		return err
	}
	fmt.Println("Result := ", result)


	fmt.Println("Send Work(Operation_DIVIDE, 25, 5)")
	work.Op = tutorial.Operation_DIVIDE
	work.Num1 = 25
	work.Num2 = 5

	result, err = client.Calculate(work)
	if err != nil {
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
