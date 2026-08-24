package regex

func IsPredefinedClass(name rune) bool {
	return name == 'd' || name == 'D' || name == 'w' || name == 'W' || name == 's' || name == 'S'
}
