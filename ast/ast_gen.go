package ast

import (
	"github.com/regmicmahesh/crafting-interpreters/token"
)

type Visitor interface {
	VisitAssignExpr(expr *Assign) string
	VisitBinaryExpr(expr *Binary) string
	VisitLiteralExpr(expr *Literal) string
}

type Expr interface {
	Accept(visitor Visitor) string
}

type Assign struct {
	Name  *token.Token
	Value Expr
}

func (expr *Assign) Accept(visitor Visitor) string {
	return visitor.VisitAssignExpr(expr)
}

func NewAssign(name *token.Token, value Expr) *Assign {
	return &Assign{name, value}

}

type Binary struct {
	Left     Expr
	Operator *token.Token
	Right    Expr
}

func (expr *Binary) Accept(visitor Visitor) string {
	return visitor.VisitBinaryExpr(expr)
}

func NewBinary(left Expr, operator *token.Token, right Expr) *Binary {
	return &Binary{left, operator, right}

}

type Literal struct {
	Value any
}

func (expr *Literal) Accept(visitor Visitor) string {
	return visitor.VisitLiteralExpr(expr)
}

func NewLiteral(value any) *Literal {
	return &Literal{value}

}
