package expression

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type Environment struct {
}

func NewEnvironment() *Environment {
	return &Environment{}
}

func (e *Environment) GetString(str string) (string, error) {
	return str, nil
}

type TreeShapeListener struct {
	*BaseExpressionListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (l *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func (e *Environment) Verify(expression string) (err error) {
	_, err = e.parse(expression)
	return
}

func (e *Environment) Eval(expression string) (any, error) {
	expr, err := e.parse(expression)
	if err != nil {
		return nil, err
	}
	visitor := &evaluator{env: e}
	val := expr.Accept(visitor)
	if visitor.err != nil {
		return nil, err
	}
	return val, nil
}

func (e *Environment) parse(expression string) (IExpressionContext, error) {
	input := antlr.NewInputStream(expression)
	lexer := NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewExpressionParser(stream)
	var parserErrorListener = new(parserErrorListener)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrorListener)
	expr := parser.Expression()

	if parserErrorListener.Err != nil {
		return nil, parserErrorListener.Err
	}

	return expr, nil
}

type parserErrorListener struct {
	antlr.DefaultErrorListener
	Err error
}

func (l *parserErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if l.Err == nil {
		l.Err = errors.New(fmt.Sprintf("syntax error %d:%d - %s\n", line, column, msg))
	}
}
