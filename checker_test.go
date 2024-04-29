package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	got := Check("345882865")

	fmt.Println(GetChecksum("345882865"))

	assert.True(t, got)
}

func TestInvalid(t *testing.T) {
	got := Check("664371495")

	assert.False(t, got)
}
