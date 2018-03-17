package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBazz(t *testing.T) {
	assert.Equal(t, "1", FizzBazz(1), "test number")
	assert.Equal(t, "Fizz", FizzBazz(3), "test fizz")
	assert.Equal(t, "Bazz", FizzBazz(5), "test buzz")
	assert.Equal(t, "FizzBazz", FizzBazz(15), "test fizz-bazz")
}

func BenchmarkFizzBazz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FizzBazz(i)
	}
}
