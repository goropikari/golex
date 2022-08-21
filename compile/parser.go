package compile

import (
	"errors"
	"io"
)

var (
	ErrParse = errors.New("parse error")
)

type Parser struct {
	tokens []Token
	pos    int
	length int
}

func NewParser(tokens []Token) *Parser {
	return &Parser{tokens: tokens, pos: 0, length: len(tokens)}
}

type NodeVisitor interface {
	VisitSumExpr(SumExpr)
	VisitConcatExpr(ConcatExpr)
	VisitStarExpr(StarExpr)
	VisitSymbolExpr(SymbolExpr)
}

type RegexExpr interface {
	Accept(v NodeVisitor)
}

func (p *Parser) Parse() (RegexExpr, error) {
	return p.sum()
}

func (p *Parser) read() (Token, error) {
	if p.pos >= p.length {
		return Token{}, io.EOF
	}
	p.pos++
	return p.tokens[p.pos-1], nil
}

func (p *Parser) peek() (Token, error) {
	if p.pos >= p.length {
		return Token{}, io.EOF
	}
	return p.tokens[p.pos], nil
}

func (p *Parser) next() (Token, error) {
	if p.pos+1 >= p.length {
		return Token{}, io.EOF
	}
	return p.tokens[p.pos+1], nil
}

func (p *Parser) sum() (RegexExpr, error) {
	lhs, err := p.concat()
	if err != nil {
		return nil, err
	}

	op, err := p.peek()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return lhs, nil
		}
		return nil, err
	}

	if op.GetType() == BarTokenType {
		p.read()
		rhs, err := p.sum()
		if err != nil {
			return nil, err
		}
		return SumExpr{lhs: lhs, rhs: rhs}, nil
	}

	return lhs, nil
}

func (p *Parser) concat() (RegexExpr, error) {
	lhs, err := p.star()
	if err != nil {
		return nil, err
	}

	b, err := p.peek()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return lhs, nil
		}
		return nil, err
	}
	// if b.GetType() == RParenTokenType {
	// 	return lhs, nil
	// }
	// if b.GetType() != BarTokenType {
	// 	rhs, err := p.sum()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return NewConcatExpr(lhs, rhs), nil
	// }
	if b.GetType() == SymbolTokenType || b.GetType() == LParenTokenType {
		rhs, err := p.concat()
		if err != nil {
			return nil, err
		}
		return NewConcatExpr(lhs, rhs), nil
	}
	return lhs, nil
}

func (p *Parser) star() (RegexExpr, error) {
	expr, err := p.primary()
	if err != nil {
		return nil, err
	}
	st, err := p.peek()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return expr, nil
		}
		return nil, err
	}
	if st.GetType() == StarTokenType {
		p.read()
		return NewStarExpr(expr), nil
	}
	return expr, nil
}

func (p *Parser) primary() (RegexExpr, error) {
	s, err := p.read()
	if err != nil {
		return nil, err
	}
	if s.GetType() == SymbolTokenType {
		return NewSymbolExpr(s.GetRune()), nil
	}
	if s.GetType() != LParenTokenType {
		return nil, ErrParse
	}

	sum, err := p.sum()
	if err != nil {
		return nil, err
	}
	r, err := p.read()
	if err != nil {
		return nil, err
	}
	if r.GetType() == RParenTokenType {
		return sum, nil
	}
	return nil, ErrParse
}

type SumExpr struct {
	lhs RegexExpr
	rhs RegexExpr
}

func NewSumExpr(lhs, rhs RegexExpr) SumExpr {
	return SumExpr{lhs: lhs, rhs: rhs}
}

func (expr SumExpr) Accept(v NodeVisitor) {
	v.VisitSumExpr(expr)
}

type ConcatExpr struct {
	lhs RegexExpr
	rhs RegexExpr
}

func NewConcatExpr(lhs, rhs RegexExpr) ConcatExpr {
	return ConcatExpr{lhs: lhs, rhs: rhs}
}

func (expr ConcatExpr) Accept(v NodeVisitor) {
	v.VisitConcatExpr(expr)
}

type StarExpr struct {
	expr RegexExpr
}

func NewStarExpr(expr RegexExpr) RegexExpr {
	return StarExpr{expr: expr}
}

func (expr StarExpr) Accept(v NodeVisitor) {
	v.VisitStarExpr(expr)
}

type SymbolExpr struct {
	sym rune
}

func NewSymbolExpr(sym rune) SymbolExpr {
	return SymbolExpr{sym: sym}
}

func (expr SymbolExpr) Accept(v NodeVisitor) {
	v.VisitSymbolExpr(expr)
}
