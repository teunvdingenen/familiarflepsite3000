package resolvers

type Token struct {
	Token string
	ExpiredAt int32
}

type TokenResolver struct {
	t *Token
}

func (t *TokenResolver) Token() string {
	return t.t.Token
}

func (t *TokenResolver) ExpiredAt() int32 {
	return t.t.ExpiredAt
}