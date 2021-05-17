package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := Add(5, 3)
	assert.Equal(t, 8, sum)
}
