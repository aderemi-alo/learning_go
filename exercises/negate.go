package exercises

// Negate function takes the pointer of a boolean value and changes it to true/false at the address itself.
func Negate(myBoolean *bool) {
	*myBoolean = !*myBoolean
}
