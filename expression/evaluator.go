package expression

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/samber/lo"
	"reflect"
	"strconv"
)

type evaluator struct {
	env *Environment
	err error
}

func (e *evaluator) Visits(trees []antlr.ParseTree) []EValue {
	var values []EValue
	for _, vc := range trees {
		value := e.Visit(vc)
		if e.err != nil {
			return nil
		}
		values = append(values, value.(EValue))
	}
	return values
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
	case *FunctionCallExpressionContext:
		return e.visitFunctionCallExpressionContext(tree.(*FunctionCallExpressionContext))
	case *IdentifierAccessExpressionContext:
		return e.visitIdentifierAccessExpressionContext(tree.(*IdentifierAccessExpressionContext))
	case *FloatingPointLiteralContext:
		return e.visitFloatingPointLiteralContext(tree.(*FloatingPointLiteralContext))
	case *BooleanLiteralContext:
		return e.visitBooleanLiteralContext(tree.(*BooleanLiteralContext))
	}
	panic(reflect.TypeOf(tree))
}

func (e *evaluator) visitIdentifierAccessExpressionContext(c *IdentifierAccessExpressionContext) any {
	return e.visitIdentifier(c.Identifier())
}

func (e *evaluator) visitIdentifier(n antlr.TerminalNode) any {
	variable := n.GetText()
	if value, ok := e.env.vars[variable]; ok {
		return value
	}
	if value, ok := e.env.builtin[variable]; ok {
		return value
	}
	return fmt.Errorf("'%s' not defined", variable)
}

func (e *evaluator) visitFunctionCallExpressionContext(c *FunctionCallExpressionContext) any {

	calee, ok := e.Visit(c.ExpressionSingle()).(EFunction)
	if !ok {
		return e.Error(fmt.Errorf("%s is not function", c.ExpressionSingle().GetText()))
	}

	args := lo.Map(c.ExpressionArguments().AllExpressionArgument(), func(item IExpressionArgumentContext, index int) EValue {
		if v := item.StringLiteral(); v != nil {
			return e.visitStringLiteral(v).(EValue)
		}
		if v := item.IntegerLiteral(); v != nil {
			return e.visitIntegerLiteral(v).(EValue)
		}
		if v := item.Identifier(); v != nil {
			return e.visitIdentifier(v).(EValue)
		}
		if v := item.ExpressionSingle(); v != nil {
			return e.Visit(v).(EValue)
		}
		panic("unreached code")
	})

	if e.err != nil {
		return nil
	}

	value, err := calee.Call(args)
	if err != nil {
		return e.Error(fmt.Errorf("call '%s' error - %s", c.ExpressionSingle(), err.Error()))
	}
	return value
}

func (e *evaluator) visitIntegerLiteralContext(c *IntegerLiteralContext) any {
	return e.visitIntegerLiteral(c.IntegerLiteral())
}

func (e *evaluator) visitIntegerLiteral(n antlr.TerminalNode) any {
	val, _ := strconv.Atoi(n.GetText())
	return EInt(val)
}

func (e *evaluator) visitFloatingPointLiteralContext(c *FloatingPointLiteralContext) any {
	return e.visitFloatLiteral(c.FloatingPointLiteral())
}

func (e *evaluator) visitFloatLiteral(n antlr.TerminalNode) any {
	val, _ := strconv.ParseFloat(n.GetText(), 64)
	return EFloat(val)
}

func (e *evaluator) visitBinaryStringLiteralContext(c *BinaryStringLiteralContext) any {
	return e.visitBinaryStringLiteral(c.StringLiteral())
}

func (e *evaluator) visitBinaryStringLiteral(n antlr.TerminalNode) any {
	text := n.GetText()
	text, err := strconv.Unquote(`"` + text[1:len(text)-1] + `"`)
	if err != nil {
		return e.Error(err)
	}
	return EBytes(text)
}

func (e *evaluator) visitStringLiteralContext(c *StringLiteralContext) any {
	return e.visitStringLiteral(c.StringLiteral())
}

func (e *evaluator) visitStringLiteral(n antlr.TerminalNode) any {
	text := n.GetText()
	text, err := strconv.Unquote(`"` + text[1:len(text)-1] + `"`)
	if err != nil {
		return e.Error(err)
	}
	return EString(text)
}

func (e *evaluator) visitBooleanLiteralContext(c *BooleanLiteralContext) any {
	return e.visitBooleanLiteral(c.BooleanLiteral())
}

func (e *evaluator) visitBooleanLiteral(n antlr.TerminalNode) any {
	if n.GetText() == "true" {
		return EBool(true)
	}
	return EBool(false)
}

func (e *evaluator) visitPlusExpression(c *PlusExpressionContext) any {
	left := e.Visit(c.ExpressionSingle(0))
	right := e.Visit(c.ExpressionSingle(1))
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return e.Error(errors.New(fmt.Sprintf("%s + %s", reflect.TypeOf(left), reflect.TypeOf(right))))
	}
	switch left.(type) {
	case EInt:
		return left.(EInt) + right.(EInt)
	case EFloat:
		return left.(EFloat) + right.(EFloat)
	case EString:
		return left.(EString) + right.(EString)
	case EBytes:
		return append(left.(EBytes), right.(EBytes)...)
	}
	panic("unreached code")
}

func (e *evaluator) Error(err error) any {
	if e.err == nil {
		e.err = err
	}
	return nil
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