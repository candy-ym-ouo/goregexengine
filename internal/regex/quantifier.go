package regex

type Quantifier struct {
	Min, Max int
	Greedy   bool
}

func NewQuantifier(min, max int, greedy bool) Quantifier { return Quantifier{min, max, greedy} }
