package expression

import (
	"context"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"regexp"
	"strconv"
)

type Environment struct {
	builtin map[string]EFunction
	vars    Vars
	Context context.Context
}

var builtin = map[string]EFunction{
	"bytes":     &eFunctionToBytes{},
	"string":    &eFunctionToString{},
	"randomInt": &eFunctionRandomInt{},
}

var expressionStringRegex = regexp.MustCompile(`\{\{.*?\}\}`)

func NewEnvironment(context context.Context, vars Vars) *Environment {
	if vars == nil {
		vars = MapVars{}
	}
	return &Environment{
		builtin: builtin,
		vars:    vars,
		Context: context,
	}
}

func (e *Environment) GetString(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	expressionStringRegex.ReplaceAllStringFunc(str, func(s string) string {
		value, err := e.Eval(str[2 : len(str)-2])
		if err != nil {
			return str
		}
		if value == nil {
			return ""
		}
		return string(value.ToString())
	})
	return str, nil
}

func (e *Environment) Verify(expression string) (err error) {
	_, err = e.parse(expression)
	return
}

func (e *Environment) Eval(expression string) (EValue, error) {
	expr, err := e.parse(expression)
	if err != nil {
		return nil, err
	}
	visitor := &evaluator{env: e}
	val := expr.Accept(visitor)
	if visitor.err != nil {
		return nil, visitor.err
	}
	return val.(EValue), nil
}

func (e *Environment) EvalWithVars(expression string, vars Vars) (EValue, error) {
	backup := e.vars
	e.vars = combinedVars{
		backup,
		vars,
	}
	value, err := e.Eval(expression)
	e.vars = backup

	return value, err
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

type EFunction interface {
	Call(args []EValue) (EValue, error)
	ToString() EString
}

type EValue interface {
	ToString() EString
}

type EInt int

func (v EInt) ToString() EString {
	return EString(strconv.Itoa(int(v)))
}

type EFloat float64

func (v EFloat) ToString() EString {
	return EString(strconv.FormatFloat(float64(v), 'f', -1, 64))
}

type EBool bool

func (v EBool) ToString() EString {
	return EString(strconv.FormatBool(bool(v)))
}

type EString string

func (v EString) ToString() EString {
	return v
}

type EBytes []byte

func (v EBytes) ToString() EString {
	return EString(v)
}

type EObject interface {
	EValue
}
