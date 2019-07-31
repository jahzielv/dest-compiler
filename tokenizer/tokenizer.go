package tokenizer

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// TokenType is a token.
type TokenType struct {
	ttype string
	regex string
}

var tokenTypes = []TokenType{
	{"func_def", `\Adef`},
	{"func_end", `\Aend`},
	{"return_stat", `\Areturn`},
	{"identifier", `\A[a-zA-z]+`},
	{"integer", `\A[0-9]+`},
	{"oparen", `\A\(`},
	{"cparen", `\A\)`},
	{"comma", `\A,`},
}

// Tokenizer tokenizes.
type Tokenizer struct {
	Code string
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("oops", r)
	}
}

// Tokenize makes tokens!
func (t *Tokenizer) Tokenize() ([]Token, error) {
	// iterate and match against regexes
	tokens := make([]Token, 0)
	defer handlePanic()
	for len(t.Code) > 0 {
		tkn, err := t.tokenizeOneToken()
		t.Code = strings.TrimSpace(t.Code)
		if err != nil {
			return []Token{}, err
		}
		tokens = append(tokens, tkn)
	}
	return tokens, nil
}

func (t *Tokenizer) tokenizeOneToken() (Token, error) {
	for _, toktype := range tokenTypes {
		re, err := regexp.Compile(toktype.regex)
		if err != nil {
			return Token{}, err
		}
		if match := re.MatchString(t.Code); match {
			val := re.FindString(t.Code)
			t.Code = t.Code[len(val):]
			return Token{toktype.ttype, val}, nil
		}
	}
	return Token{}, errors.New("tokenize failed :(")
}

// Token is a token.
type Token struct {
	Type  string
	Value string
}

// ToString formats the token into a string
func (t *Token) ToString() string {
	return fmt.Sprintf("Token {type: %s, value: %s}", t.Type, t.Value)
}
