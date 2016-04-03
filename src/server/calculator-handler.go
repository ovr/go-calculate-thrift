package main

import (
	"fmt"
	"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
	"errors"
	"math"
)

type CalculatorHandler struct {
	tutorial.Calculator
}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{}
}

func (p *CalculatorHandler) Ping() (err error) {
	fmt.Print("ping()\n")
	return nil
}

func (p *CalculatorHandler) Plus(num1 int32, num2 int32) (r int64, err error) {
	fmt.Print("Plus(", num1, ",", num2, ")\n")
	return int64(num1 + num2), nil
}

func (p *CalculatorHandler) Minus(num1 int32, num2 int32) (val int32, err error) {
	fmt.Print("Minus(", num1, ",", num2, ")\n")
	return num1 - num2, nil
}

func (p *CalculatorHandler) Mul(num1 int32, num2 int32) (r int64, err error) {
	fmt.Print("Mul(", num1, ",", num2, ")\n")

	return int64(num1 * num2), nil
}

func (p *CalculatorHandler) Div(num1 int32, num2 int32) (val int32, err error) {
	fmt.Print("Div(", num1, ",", num2, ")\n")

	if num2 == 0 {
		return 0, errors.New("You cannot div on zero");
	}

	return num1 / num2, nil
}

func (p *CalculatorHandler) Mod(num1 float64, num2 float64) (r float64, err error) {
	fmt.Print("Mod(", num1, ",", num2, ")\n")
	return math.Mod(num1, num2), nil
}

func (p *CalculatorHandler) Pow(num1 int32, num2 int32) (val int32, err error) {
	fmt.Print("Pow(", num1, ",", num2, ")\n")
	return 0, nil
}

func (p *CalculatorHandler) Calculate(logid int32, w *tutorial.Work) (val int32, err error) {
	switch w.Op {
	case tutorial.Operation_ADD:
		val = w.Num1 + w.Num2
		break
	case tutorial.Operation_SUBTRACT:
		val = w.Num1 - w.Num2
		break
	case tutorial.Operation_MULTIPLY:
		val = w.Num1 * w.Num2
		break
	case tutorial.Operation_DIVIDE:
		if w.Num2 == 0 {
			ouch := tutorial.NewInvalidOperation()
			ouch.WhatOp = int32(w.Op)
			ouch.Why = "Cannot divide by 0"

			return 0, ouch
		}

		val = w.Num1 / w.Num2
		break
	default:
		ouch := tutorial.NewInvalidOperation()
		ouch.WhatOp = int32(w.Op)
		ouch.Why = "Unknown operation"

		return 0, ouch
	}

	return val, err
}
