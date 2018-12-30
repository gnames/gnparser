package grammar

func (p *Engine) ParsedName() string {
	if p.tokens32.tree == nil {
		return ""
	}
	for i := len(p.tokens32.tree) - 1; i >= 0; i-- {
		t := p.tokens32.tree[i]
		if t.pegRule == ruleSciName1 {
			return string(p.buffer[t.begin:t.end])
		}
	}
	return ""
}
