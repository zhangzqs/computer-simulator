package main

func And(a, b bool) bool {
	if a == true && b == true {
		return true
	}
	return false
}
func Not(v bool) bool {
	if v {
		return false
	}
	return true
}
func Or(a, b bool) bool {
	return Not(And(
		Not(a),
		Not(b),
	))
}

func Xor(a, b bool) bool {
	return Or(
		And(a, Not(b)),
		And(Not(a), b),
	)
}
