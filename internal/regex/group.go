package regex

type GroupInfo struct {
	Index int
	Name  string
}

func Groups(c *Compiled) []GroupInfo {
	out := []GroupInfo{}
	for i, n := range c.RE.SubexpNames() {
		if i > 0 {
			out = append(out, GroupInfo{i, n})
		}
	}
	return out
}
