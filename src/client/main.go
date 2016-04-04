package main

import (
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
	"reflect"
"bytes"
"strings"
"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
)

func Usage() {
	flag.PrintDefaults()
	fmt.Fprintln(os.Stdout, "\nFunctions:")
	fmt.Fprintln(os.Stdout, "  void ping()")
	fmt.Fprintln(os.Stdout, "  i64 plus(i32 num1, i32 num2)")
	fmt.Fprintln(os.Stdout, "  i32 minus(i32 num1, i32 num2)")
	fmt.Fprintln(os.Stdout, "  i64 mul(i32 num1, i32 num2)")
	fmt.Fprintln(os.Stdout, "  i32 div(i32 num1, i32 num2)")
	fmt.Fprintln(os.Stdout, "  double mod(double num1, double num2)")
	fmt.Fprintln(os.Stdout, "  i32 pow(i32 num1, i32 num2)")
	fmt.Fprintln(os.Stdout, "  i32 calculate(Work w)")
	fmt.Fprintln(os.Stdout)

	os.Exit(0)
}

func MakeFirstUpperCase(s string) string {

	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)

	lc := bytes.ToUpper([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}

func main() {
	flag.Usage = Usage

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

	client, err := runClient(thrift.NewTTransportFactory(), protocolFactory, *addr)
	if  err != nil {
		fmt.Println("error running client:", err)
		os.Exit(1)
	}
	defer client.Transport.Close()

	var (
		command string
		work string
		num1 int32
		num2 int32
	)

	for {
		fmt.Print("Enter command: ")
		fmt.Scanln(&command)

		switch command {
		case "ping":
			fmt.Println("Ping send")
			responseError := client.Ping();
			if responseError != nil {
				fmt.Fprintln(os.Stderr, "Error was caused")
				fmt.Fprintln(os.Stderr, responseError.Error())
				continue
			}
		case "plus", "minus", "mul", "div":
			fmt.Println("Enter 1 number: ")
			fmt.Scanln(&num1)

			fmt.Println("Enter 2 number: ")
			fmt.Scanln(&num2)

			inputs := []reflect.Value{reflect.ValueOf(num1), reflect.ValueOf(num2)}
			returnValues := reflect.ValueOf(client).MethodByName(MakeFirstUpperCase(command)).Call(inputs)

			if !returnValues[1].IsNil() {
				responseError := returnValues[1].Interface().(error)

				fmt.Fprintln(os.Stderr, "Error was caused")
				fmt.Fprintln(os.Stderr, responseError.Error())
				continue
			}

			result := returnValues[0].Int()
			fmt.Println("Result ", result)
		case "calculate", "work":
			fmt.Println("Enter work: ")
			fmt.Scanln(&work)

			workRequest := tutorial.NewWork()

			switch work {
			case "plus", "add":
				workRequest.Op = tutorial.Operation_ADD
			case "minus":
				workRequest.Op = tutorial.Operation_SUBTRACT
			case "mul":
				workRequest.Op = tutorial.Operation_MULTIPLY
			case "div":
				workRequest.Op = tutorial.Operation_DIVIDE
			default:
				fmt.Fprintln(os.Stderr, "Invalid work type", work)
				os.Exit(1)

			}

			fmt.Println("Enter 1 number: ")
			fmt.Scanln(&workRequest.Num1)

			fmt.Println("Enter 2 number: ")
			fmt.Scanln(&workRequest.Num2)

			result, err := client.Calculate(workRequest)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println("Result := ", result)
		default:
			fmt.Fprintln(os.Stderr, "Invalid command", command)
			os.Exit(1)
		}
	}
}
