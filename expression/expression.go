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

func (e *Environment) Eval(expression string) error {

	input := antlr.NewInputStream(expression)
	lexer := NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	parser := NewExpressionParser(stream)
	var parserErrorListener = new(parserErrorListener)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(parserErrorListener)
	expr := parser.Expression()

	if parserErrorListener.Err != nil {
		return parserErrorListener.Err
	}

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), expr)

	return nil
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
