package main

import "fmt"

func main() {
	a := NewBitSet(16)
	a.Fill(253, 0, 8)
	fmt.Println(a.toString())

	b := NewBitSet(16)
	b.Fill(32, 0, 8)
	fmt.Println(b.toString())

	bs, cf := (&NBitComplementAddSubComputer{
		InputA:    a,
		InputB:    b,
		SubSignal: true,
	}).GetResult()

	fmt.Println(bs.Size(), bs.Convert(), cf)
}
