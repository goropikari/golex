package compile

import (
	"github.com/google/uuid"
	"github.com/goropikari/golex/collection"
)

type CodeGenerator struct {
	nfa NFA
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

func (gen *CodeGenerator) GetNFA() NFA {
	return gen.nfa
}

func (gen *CodeGenerator) VisitSumExpr(expr SumExpr) {
	expr.lhs.Accept(gen)
	lhs := gen.nfa
	expr.rhs.Accept(gen)
	rhs := gen.nfa

	gen.nfa = lhs.Sum(rhs)
}

func (gen *CodeGenerator) VisitConcatExpr(expr ConcatExpr) {
	expr.lhs.Accept(gen)
	lhs := gen.nfa
	expr.rhs.Accept(gen)
	rhs := gen.nfa

	gen.nfa = lhs.Concat(rhs)
}

func (gen *CodeGenerator) VisitStarExpr(expr StarExpr) {
	expr.expr.Accept(gen)
	gen.nfa = gen.nfa.Star()
}

func (gen *CodeGenerator) VisitSymbolExpr(expr SymbolExpr) {
	from := NewState(uuid.New().String())
	to := NewState(uuid.New().String())

	gen.nfa = NewNFA(
		collection.NewSet[State]().Insert(from).Insert(to),
		Transition{
			NewTuple[State, rune](from, expr.sym): collection.NewSet[State]().Insert(to),
		},
		collection.NewSet[State]().Insert(from),
		collection.NewSet[State]().Insert(to),
	)
}
