package regex

type Group struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
	Start int    `json:"start"`
	End   int    `json:"end"`
	Text  string `json:"text"`
}
type Match struct {
	Index  int     `json:"index"`
	Start  int     `json:"start"`
	End    int     `json:"end"`
	Text   string  `json:"text"`
	Groups []Group `json:"groups"`
}

func runeOffset(s string, bytePos int) int { return len([]rune(s[:bytePos])) }
func (c *Compiled) one(text string, idx int, loc []int) Match {
	m := Match{Index: idx, Start: runeOffset(text, loc[0]), End: runeOffset(text, loc[1]), Text: text[loc[0]:loc[1]], Groups: groupBuffer()}
	names := c.RE.SubexpNames()
	for i := 0; i < len(loc)/2; i++ {
		a, b := loc[2*i], loc[2*i+1]
		g := Group{Index: i, Name: names[i]}
		if a >= 0 {
			g.Start = runeOffset(text, a)
			g.End = runeOffset(text, b)
			g.Text = text[a:b]
		} else {
			g.Start = -1
			g.End = -1
		}
		m.Groups = append(m.Groups, g)
	}
	saveGroupBuffer(m.Groups)
	return m
}
func (c *Compiled) FindAll(text string, first bool) []Match {
	locs := c.RE.FindAllStringSubmatchIndex(text, -1)
	out := make([]Match, 0, len(locs))
	for i, l := range locs {
		out = append(out, c.one(text, i, l))
		if first {
			break
		}
	}
	return out
}
func (c *Compiled) MatchString(text string) bool     { return c.RE.MatchString(text) }
func (c *Compiled) Replace(text, repl string) string { return c.RE.ReplaceAllString(text, repl) }
func (c *Compiled) Split(text string, n int) []string {
	if n == 0 {
		return []string{}
	}
	locs := c.RE.FindAllStringIndex(text, -1)
	if n > 0 && len(locs) >= n {
		locs = locs[:n-1]
	}
	out := make([]string, 0, len(locs)+1)
	last := 0
	for _, loc := range locs {
		out = append(out, text[last:loc[0]])
		last = loc[1]
	}
	return append(out, text[last:])
}
