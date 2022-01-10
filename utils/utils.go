package utils

func Pluralize(n int, s string) string {
	if n > 1 {
		return s + "s"
	}
	return s
}
