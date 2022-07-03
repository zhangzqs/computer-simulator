package main

import "bytes"

type BitSet struct {
	data []uint8
	size int
}

func NewBitSet(size int) *BitSet {
	return &BitSet{
		size: size,
		data: make([]uint8, size/8+1),
	}
}

// Set 置1
func (r *BitSet) Set(i int) {
	n := i / 8
	offset := i % 8
	r.data[n] |= 1 << offset
}

// Reset 复位
func (r *BitSet) Reset(i int) {
	n := i / 8
	offset := i % 8
	if r.data[n]>>offset&1 == 1 {
		r.data[n] &= ^(1 << offset)
	}
}
func (r *BitSet) ResetAll() {
	for i := 0; i < len(r.data); i++ {
		r.data[i] = 0
	}
}

// Get 获取
func (r *BitSet) Get(i int) bool {
	n := i / 8
	offset := i % 8
	return r.data[n]>>offset&1 == 1
}

func (r *BitSet) Size() int {
	return r.size
}

func (r *BitSet) MakeLarger(externSize int) {
	r.size += externSize
	externDataSize := len(r.data) - r.size/8 + 1

	for i := 0; i < externDataSize; i++ {
		r.data = append(r.data, 0)
	}
}

func (r *BitSet) Fill(n uint32, offset int, size int) {
	for i := 0; i < size; i++ {
		if n>>i&1 == 1 {
			r.Set(offset + i)
		} else {
			r.Reset(offset + i)
		}
	}
}

func (r *BitSet) Convert() uint32 {
	var res uint32
	for i, e := range r.data {
		res += uint32(e) * (1 << (8 * i))
	}
	return res
}

func (r *BitSet) Not() *BitSet {
	res := NewBitSet(r.Size())
	for i := 0; i < r.Size(); i++ {
		if !r.Get(i) {
			res.Set(i)
		}
	}
	return res
}

func (r *BitSet) toString() string {
	var buf bytes.Buffer
	for i := r.Size() - 1; i >= 0; i-- {
		if r.Get(i) {
			buf.WriteString("1")
		} else {
			buf.WriteString("0")
		}
	}
	return buf.String()
}
