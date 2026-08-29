package ast

import (
	"bytes"
	"github.com/SegniAT/monkey-language-interpreter/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
	Start() token.Position
	End() token.Position
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) Start() token.Position {
	if len(p.Statements) > 0 {
		return p.Statements[0].Start()
	}

	return token.Position{Line: 1, Character: 1}
}

func (p *Program) End() token.Position {
	if len(p.Statements) > 0 {
		return p.Statements[len(p.Statements)-1].Start()
	}

	return token.Position{Line: 1, Character: 1}
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) Start() token.Position { return ls.Token.Range.Start }
func (ls *LetStatement) End() token.Position   { return ls.Value.End() }
func (ls *LetStatement) statementNode()        {}
func (ls *LetStatement) TokenLiteral() string  { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral())
	out.WriteString(" ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) Start() token.Position { return i.Token.Range.Start }
func (i *Identifier) End() token.Position   { return i.Token.Range.End }
func (i *Identifier) expressionNode()       {}
func (i *Identifier) TokenLiteral() string  { return i.Token.Literal }
func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) Start() token.Position { return rs.Token.Range.Start }
func (rs *ReturnStatement) End() token.Position   { return rs.ReturnValue.End() }
func (rs *ReturnStatement) statementNode()        {}
func (rs *ReturnStatement) TokenLiteral() string  { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())
	out.WriteString(" ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) Start() token.Position { return es.Token.Range.Start }
func (es *ExpressionStatement) End() token.Position   { return es.Expression.End() }
func (es *ExpressionStatement) statementNode()        {}
func (es *ExpressionStatement) TokenLiteral() string  { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) Start() token.Position { return il.Token.Range.Start }
func (il *IntegerLiteral) End() token.Position   { return il.Token.Range.End }
func (il *IntegerLiteral) expressionNode()       {}
func (il *IntegerLiteral) TokenLiteral() string  { return il.Token.Literal }
func (il *IntegerLiteral) String() string        { return il.Token.Literal }

type PrefixExpression struct {
	Token    token.Token // The prefix token, eg. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) Start() token.Position { return pe.Token.Range.Start }
func (pe *PrefixExpression) End() token.Position   { return pe.Right.End() }
func (pe *PrefixExpression) expressionNode()       {}
func (pe *PrefixExpression) TokenLiteral() string  { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) Start() token.Position { return oe.Token.Range.Start }
func (oe *InfixExpression) End() token.Position   { return oe.Right.End() }
func (oe *InfixExpression) expressionNode()       {}
func (oe *InfixExpression) TokenLiteral() string  { return oe.Token.Literal }
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" ")
	out.WriteString(oe.Operator)
	out.WriteString(" ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) Start() token.Position { return b.Token.Range.Start }
func (b *Boolean) End() token.Position   { return b.Token.Range.End }
func (b *Boolean) expressionNode()       {}
func (b *Boolean) TokenLiteral() string  { return b.Token.Literal }
func (b *Boolean) String() string        { return b.Token.Literal }

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) Start() token.Position { return ie.Token.Range.Start }
func (ie *IfExpression) End() token.Position {
	if ie.Alternative != nil {
		return ie.Alternative.StartToken.Range.End
	}
	return ie.Consequence.EndToken.Range.End
}
func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	StartToken token.Token // the { token
	Statements []Statement
	EndToken   token.Token // the } token
}

func (bs *BlockStatement) Start() token.Position { return bs.StartToken.Range.Start }
func (bs *BlockStatement) End() token.Position   { return bs.EndToken.Range.End }
func (bs *BlockStatement) statementNode()        {}
func (bs *BlockStatement) TokenLiteral() string  { return bs.StartToken.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token // the 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) Start() token.Position { return fl.Token.Range.Start }
func (fl *FunctionLiteral) End() token.Position   { return fl.Body.End() }
func (fl *FunctionLiteral) expressionNode()       {}
func (fl *FunctionLiteral) TokenLiteral() string  { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, param := range fl.Parameters {
		params = append(params, param.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
	EndToken  token.Token // the ) token
}

func (ce *CallExpression) Start() token.Position { return ce.Function.Start() }
func (ce *CallExpression) End() token.Position   { return ce.EndToken.Range.End }
func (ce *CallExpression) expressionNode()       {}
func (ce *CallExpression) TokenLiteral() string  { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) Start() token.Position { return sl.Token.Range.Start }
func (sl *StringLiteral) End() token.Position   { return sl.Token.Range.End }
func (sl *StringLiteral) expressionNode()       {}
func (sl *StringLiteral) TokenLiteral() string  { return sl.Token.Literal }
func (sl *StringLiteral) String() string        { return sl.Token.Literal }

type ArrayLiteral struct {
	StartToken token.Token
	Elements   []Expression
	EndToken   token.Token // the ] literal
}

func (al *ArrayLiteral) Start() token.Position { return al.StartToken.Range.Start }
func (al *ArrayLiteral) End() token.Position   { return al.EndToken.Range.End }
func (al *ArrayLiteral) expressionNode()       {}
func (al *ArrayLiteral) TokenLiteral() string  { return al.StartToken.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	StartToken token.Token // The [ token
	Left       Expression
	Index      Expression
	EndToken   token.Token // The ] token
}

func (ie *IndexExpression) Start() token.Position { return ie.StartToken.Range.Start }
func (ie *IndexExpression) End() token.Position   { return ie.EndToken.Range.End }
func (ie *IndexExpression) expressionNode()       {}
func (ie *IndexExpression) TokenLiteral() string  { return ie.StartToken.Literal }
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("]")
	out.WriteString(")")

	return out.String()
}

type HashLiteral struct {
	StartToken token.Token // The { token
	Pairs      map[Expression]Expression
	EndToken   token.Token // The } token
}

func (hl *HashLiteral) Start() token.Position { return hl.StartToken.Range.Start }
func (hl *HashLiteral) End() token.Position   { return hl.EndToken.Range.End }
func (hl *HashLiteral) expressionNode()       {}
func (hl *HashLiteral) TokenLiteral() string  { return hl.StartToken.Literal }
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	pairs := []string{}
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
