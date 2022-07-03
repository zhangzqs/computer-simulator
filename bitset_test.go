package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBitSet(t *testing.T) {
	bs := NewBitSet(10)
	assert.Equal(t, 10, bs.size)
	assert.Equal(t, 2, len(bs.data))
}

func TestBitSet_Get(t *testing.T) {
	bs := NewBitSet(10)
	for i := 0; i < 10; i++ {
		assert.False(t, bs.Get(i))
		bs.Set(i)
		assert.True(t, bs.Get(i))
		bs.Reset(i)
		assert.False(t, bs.Get(i))
	}

}

func TestBitSet_ResetAll(t *testing.T) {
	bs := NewBitSet(10)
	bs.Fill(0b101001, 0, 6)
	for i := 0; i < bs.Size(); i++ {
		fmt.Println(bs.Get(i))
	}
}
