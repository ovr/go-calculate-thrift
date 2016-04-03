package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dataProviderPlusTests = []struct {
	a        int32 // inputA
	b        int32 // inputB
	expected int64 // expected result
}{
	{1, 1, 2},
	{2, 1, 3},
	{3, 2, 5},
	{4, 3, 7},
	{5, 5, 10},
	{6, 8, 14},
	{7, 13, 20},
}

func TestPlus(t *testing.T) {
	var (
		result int64
		err    error
	)

	service := NewCalculatorHandler()

	for _, testsFixture := range dataProviderPlusTests {
		result, err = service.Plus(testsFixture.a, testsFixture.b)

		assert.Equal(t, nil, err, "err should be nil")
		assert.Equal(t, testsFixture.expected, result, "they should be equal")
	}
}

var dataProviderMinusTests = []struct {
	a        int32 // inputA
	b        int32 // inputB
	expected int32 // expected result
}{
	{25, 5, 20},
	{2, 1, 1},
	{25, 50, -25},
	{0, 0, 0},
	{25, 0, 25},
	{-25, 0, -25},
	{-1, -1, 0},
	{-1, 1, -2},
}

func TestMinus(t *testing.T) {
	var (
		result int32
		err    error
	)

	service := NewCalculatorHandler()

	for _, testsFixture := range dataProviderMinusTests {
		result, err = service.Minus(testsFixture.a, testsFixture.b)

		assert.Equal(t, nil, err, "err should be nil")
		assert.Equal(t, testsFixture.expected, result, "they should be equal")
	}
}
