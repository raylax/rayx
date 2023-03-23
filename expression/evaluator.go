package expression

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"reflect"
	"strconv"
)

type evaluator struct {
	env *Environment
	err error
}

func (e *evaluator) Visit(tree antlr.ParseTree) any {
	if e.err != nil {
		return nil
	}
	switch tree.(type) {
	case *PlusExpressionContext:
		return e.visitPlusExpression(tree.(*PlusExpressionContext))
	case *ConstExpressionContext:
		return e.Visit(tree.(*ConstExpressionContext).ExpressionConst())
	case *IntegerLiteralContext:
		return e.visitIntegerLiteralContext(tree.(*IntegerLiteralContext))
	case *StringLiteralContext:
		return e.visitStringLiteralContext(tree.(*StringLiteralContext))
	case *BinaryStringLiteralContext:
		return e.visitBinaryStringLiteralContext(tree.(*BinaryStringLiteralContext))
	}
	panic(reflect.TypeOf(tree))
}

func (e *evaluator) visitIntegerLiteralContext(c *IntegerLiteralContext) any {
	val, _ := strconv.Atoi(c.IntegerLiteral().GetText())
	return val
}

func (e *evaluator) visitBinaryStringLiteralContext(c *BinaryStringLiteralContext) any {
	text := c.StringLiteral().GetText()
	text, err := strconv.Unquote(`"` + text[1:len(text)-1] + `"`)
	if err != nil {
		e.Error(err)
		return nil
	}
	return []byte(text)
}

func (e *evaluator) visitStringLiteralContext(c *StringLiteralContext) any {
	text := c.StringLiteral().GetText()
	text, err := strconv.Unquote(`"` + text[1:len(text)-1] + `"`)
	if err != nil {
		e.Error(err)
		return nil
	}
	return text
}

func (e *evaluator) visitPlusExpression(c *PlusExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	right := e.Visit(c.ExpressionSingle(1))
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		e.Error(errors.New(fmt.Sprintf("%s + %s", reflect.TypeOf(left), reflect.TypeOf(left))))
		return nil
	}
	switch left.(type) {
	case int:
		return left.(int) + right.(int)
	case float64:
		return left.(float64) + right.(float64)
	case string:
		return left.(string) + right.(string)
	case []byte:
		return append(left.([]byte), right.([]byte)...)
	}
	panic("unreached code")
}

func (e *evaluator) Error(err error) {
	if e.err == nil {
		e.err = err
	}
}

func (e *evaluator) VisitChildren(node antlr.RuleNode) any {
	return e.Visit(node.GetChild(0).(antlr.ParseTree))
}

func (e *evaluator) VisitTerminal(node antlr.TerminalNode) any {
	//TODO implement me
	panic("implement me")
}

func (e *evaluator) VisitErrorNode(node antlr.ErrorNode) any {
	//TODO implement me
	panic("implement me")
}
