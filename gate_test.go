package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnd(t *testing.T) {
	assert.Equal(t, true, And(true, true))
	assert.Equal(t, false, And(true, false))
	assert.Equal(t, false, And(false, true))
	assert.Equal(t, false, And(false, false))
}

func TestNot(t *testing.T) {
	assert.Equal(t, true, Not(false))
	assert.Equal(t, false, Not(true))
}

func TestOr(t *testing.T) {
	assert.Equal(t, true, Or(true, true))
	assert.Equal(t, true, Or(true, false))
	assert.Equal(t, true, Or(false, true))
	assert.Equal(t, false, Or(false, false))
}

func TestXor(t *testing.T) {
	assert.Equal(t, false, Xor(true, true))
	assert.Equal(t, true, Xor(true, false))
	assert.Equal(t, true, Xor(false, true))
	assert.Equal(t, false, Xor(false, false))
}
