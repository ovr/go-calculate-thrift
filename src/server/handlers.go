package main

import (
	"errors"
	"github.com/ovr/go-calculate-thrift/gen-go/tutorial"
	"math"
	"log"
)

type CalculatorHandler struct {
	tutorial.Calculator
}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{}
}

func (p *CalculatorHandler) Ping() (err error) {
	log.Println("Ping()")
	return nil
}

func (p *CalculatorHandler) Plus(num1 int32, num2 int32) (r int64, err error) {
	log.Println("Plus(", num1, ",", num2, ")")
	return int64(num1 + num2), nil
}

func (p *CalculatorHandler) Minus(num1 int32, num2 int32) (val int32, err error) {
	log.Println("Minus(", num1, ",", num2, ")")
	return num1 - num2, nil
}

func (p *CalculatorHandler) Mul(num1 int32, num2 int32) (r int64, err error) {
	log.Println("Mul(", num1, ",", num2, ")")
	return int64(num1 * num2), nil
}

func (p *CalculatorHandler) Div(num1 int32, num2 int32) (val int32, err error) {
	log.Println("Div(", num1, ",", num2, ")")

	if num2 == 0 {
		return 0, errors.New("You cannot div on zero")
	}

	return num1 / num2, nil
}

func (p *CalculatorHandler) Mod(num1 float64, num2 float64) (r float64, err error) {
	log.Println("Mod(", num1, ",", num2, ")")
	return math.Mod(num1, num2), nil
}

func (p *CalculatorHandler) Pow(num1 int32, num2 int32) (val int32, err error) {
	log.Println("Pow(", num1, ",", num2, ")")
	return 0, nil
}

func (p *CalculatorHandler) Calculate(w *tutorial.Work) (val int32, err error) {
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
