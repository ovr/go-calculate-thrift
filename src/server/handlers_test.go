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
