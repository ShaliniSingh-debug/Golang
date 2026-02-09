package interfacesgeneric



func Max[T int](a, b T) T {
	if a > b {
		return a
	}
	return b
}
