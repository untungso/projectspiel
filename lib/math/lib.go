package math

// Abs is doing something
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Mod is doing something
func Mod(a, b int) int {
	m := a % b
	if m < 0 {
		m += Abs(b)
	}
	return m
}
