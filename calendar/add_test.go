package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := Add(7, 3)
	assert.Equal(t, 10, sum)
}
