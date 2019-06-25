package ast

import (
	"testing"

	"../token"
)

// TestString :
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.VAR,
					Literal: "var",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if "var myVar := anotherVar" != program.String() {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}
