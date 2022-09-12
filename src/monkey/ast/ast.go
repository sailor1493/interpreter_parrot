package ast

import "monkey/token"

// Generic Node and Subclass Section

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expresssionNode()
}

// Statements Section

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Expressions Section

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expresssionNode()     {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Program and Program builder Section

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
