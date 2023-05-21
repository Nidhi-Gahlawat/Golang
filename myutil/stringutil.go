package myutil

//  Concatenate returns concatenation of two strings.
func Concatenate(s1, s2 string) string {
	return s1 + " " + s2
}

//  Greet returns a greeting string.
func Greet(name string) string {
	return Concatenate("Hello", name) + "!"
}
