package utils

// Function initial letter has to be capital to export in pkg
// otherwise it won't be available.

func SomeCustomFuncOfAddition(a int, b int) int {
	var result = a + b
	return result
}
