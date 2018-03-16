package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBazz(t *testing.T) {
	assert.Equal(t, FizzBazz(1), "1")
	assert.Equal(t, FizzBazz(3), "Fizz")
	assert.Equal(t, FizzBazz(5), "Bazz")
	assert.Equal(t, FizzBazz(15), "FizzBazz")
}

func BenchmarkFizzBazz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBazz(i)
	}
}
