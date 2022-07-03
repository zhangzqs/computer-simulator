package main

type FullAdder struct {
	InputA     bool
	InputB     bool
	InputCarry bool
}

func (r *FullAdder) GetS() bool {
	return Xor(r.InputA, Xor(r.InputB, r.InputCarry))
}

func (r *FullAdder) GetNextCarry() bool {
	g := And(r.InputA, r.InputB)
	p := Xor(r.InputA, r.InputB)
	return Or(g, And(p, r.InputCarry))
}

type NBitFullAdder struct {
	InputA     *BitSet
	InputB     *BitSet
	InputCarry bool
}

func (r *NBitFullAdder) GetResult() (bs *BitSet, cf bool) {
	if r.InputA.Size() != r.InputB.Size() {
		panic("")
	}
	bs = NewBitSet(r.InputA.Size())
	c := r.InputCarry
	for i := 0; i < bs.Size(); i++ {
		fa := FullAdder{
			InputA:     r.InputA.Get(i),
			InputB:     r.InputB.Get(i),
			InputCarry: c,
		}
		if fa.GetS() {
			bs.Set(i)
		}
		c = fa.GetNextCarry()
	}
	cf = c
	return
}

// MUX 多路选择器
func MUX(inputBits *BitSet, subSignal bool) (outputBits *BitSet) {
	if subSignal {
		return inputBits.Not()
	}
	return inputBits
}

type NBitComplementAddSubComputer struct {
	InputA    *BitSet
	InputB    *BitSet
	SubSignal bool
}

func (r *NBitComplementAddSubComputer) GetResult() (bs *BitSet, cf bool) {
	return (&NBitFullAdder{
		InputA:     r.InputA,
		InputB:     MUX(r.InputB, r.SubSignal),
		InputCarry: r.SubSignal,
	}).GetResult()
}
