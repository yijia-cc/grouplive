package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := Add(8, 3)
	assert.Equal(t, 11, sum)
}
