package ast

import (
    "github.com/komlow/writing-an-interpreter-in-go/token"
    "testing"
)

func TestString(t *testing.T) {
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "myVar"},
                    Value: "myVar",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "another-Var"},
                    Value: "anotherVar",
                },
            },
        },
    }

    if program.String() != "let myVar = anotherVar;" {
        t.Errorf("program.String() wrong. got=%q", program.String())
    }
}
