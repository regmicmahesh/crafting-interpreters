package ast

import (
	"fmt"

	"github.com/regmicmahesh/crafting-interpreters/token"
)

type AstPrinter struct{}

func (printer *AstPrinter) Print(expr Expr) string {
	return expr.Accept(printer)
}

func (printer *AstPrinter) VisitAssignExpr(expr *Assign) string {
	return fmt.Sprintf("%v = %v", expr.Name.Lexeme, expr.Value.Accept(printer))
}

func (printer *AstPrinter) VisitBinaryExpr(expr *Binary) string {
	return fmt.Sprintf("(%v %v %v)", expr.Left.Accept(printer), expr.Operator.Lexeme, expr.Right.Accept(printer))
}

func (printer *AstPrinter) VisitLiteralExpr(expr *Literal) string {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)

}

func main() {

	x := &Assign{
		Name:  token.NewToken(token.IDENTIFER, "x", nil, 1),
		Value: &Literal{Value: 1},
	}

	printer := &AstPrinter{}
	y := printer.VisitAssignExpr(x)
	fmt.Println(y)

}
