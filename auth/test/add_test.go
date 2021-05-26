package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	assert.Equal(t, 5, sum)
}
