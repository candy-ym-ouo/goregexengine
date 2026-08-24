package regex

func Execute(c *Compiled, text string) []Match { return c.FindAll(text, false) }
