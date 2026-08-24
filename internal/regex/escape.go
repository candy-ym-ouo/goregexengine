package regex

func DecodeEscape(s string) string {
	switch s {
	case "n":
		return "\n"
	case "t":
		return "\t"
	case "r":
		return "\r"
	case "f":
		return "\f"
	case "v":
		return "\v"
	}
	return s
}
