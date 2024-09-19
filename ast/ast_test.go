package ast

import (
	"example/monkey/token"
	"testing"
)

// In this test we construct the AST by hand. When writing tests for the parser we don’t, of
// course, but make assertions about the AST the parser produces. For demonstration purposes,
// this test shows us how we can add another easily readable layer of tests for our parser by just
// comparing the parser output with strings. That’s going to be especially handy when parsing
// expressions.
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Token: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				}, // page --> 52
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
