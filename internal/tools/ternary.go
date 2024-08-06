package tools

func Ternary(cond bool, a interface{}, b interface{}) interface{} {
	if cond {
		return a
	}
	return b
}
