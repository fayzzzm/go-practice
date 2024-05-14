package pointers

func ChangeValue(x *int) *int {
	*x = 1

	return x
}
