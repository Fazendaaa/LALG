package ast

import (
	"testing"

	"../token"
)

// TestString :
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VarStatement{
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
				Type: token.Token{
					Type:    token.INTEGER_KEYWORD,
					Literal: "integer",
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

	if "var myVar: integer := anotherVar;" != program.String() {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}
