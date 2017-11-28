package lexer

type Token struct {
	Kind      TokenKind
	fullStart int
	start     int
	length    int
}

type TokenShortForm struct {
	Kind       string `json:"kind"`
	TextLength int    `json:"textLength"`
}

func (r Token) getText(text []rune) string {
	s := r.start
	end := s + r.length - (r.start - r.fullStart)
	return string(text[s:end])
}

func (r Token) getShortForm() TokenShortForm {
	return TokenShortForm{r.Kind.String(), r.length - (r.start - r.fullStart)}
}
