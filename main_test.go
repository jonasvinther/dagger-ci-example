package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestHello(t *testing.T) {

	expected := "Hello, World"
	actual := Hello()

	assert.Equal(t, expected, actual)
}
