package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	got := Check("345882865")

	assert.Equal(t, Valid, got)
}

func TestInvalid(t *testing.T) {
	got := Check("664371495")

	assert.Equal(t, Invalid, got)
}

func TestUnknown(t *testing.T) {
	got := Check("123??6789")

	assert.Equal(t, Unknown, got)
}
