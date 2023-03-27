package expression

import (
	"bytes"
	"context"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"hash"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

type Environment struct {
	builtin map[string]EFunction
	vars    Vars
	Context context.Context
}

var builtin = map[string]EFunction{
	"bytes":           &eFunctionToBytes{},
	"string":          &eFunctionToString{},
	"randomInt":       &eFunctionRandomInt{},
	"randomLowercase": &eFunctionRandomAlpha{upper: false},
	"randomUppercase": &eFunctionRandomAlpha{upper: true},
	"md5": &eFunctionHash{h: func() hash.Hash {
		return md5.New()
	}},
	"sha1": &eFunctionHash{h: func() hash.Hash {
		return sha1.New()
	}},
	"sha256": &eFunctionHash{h: func() hash.Hash {
		return sha256.New()
	}},
	"base64": &eFunctionCodec{fun: func(value EValue) (EValue, error) {
		switch value.(type) {
		case EBytes:
			return EString(base64.StdEncoding.EncodeToString(value.(EBytes))), nil
		case EString:
			return EString(base64.StdEncoding.EncodeToString([]byte(value.(EString)))), nil
		default:
			return nil, fmt.Errorf("expect EString,EBytes, got %s", reflect.TypeOf(value))
		}
	}},
	"urldecode": &eFunctionCodec{fun: func(value EValue) (EValue, error) {

		var v string
		switch value.(type) {
		case EBytes:
			v = string(value.(EBytes))
		case EString:
			v = string(value.(EString))
		default:
			return nil, fmt.Errorf("expect EString,EBytes, got %s", reflect.TypeOf(value))
		}
		decodedValue, err := url.QueryUnescape(v)
		if err != nil {
			return nil, err
		}
		return EString(decodedValue), err
	}},
}

var expressionStringRegex = regexp.MustCompile(`\{\{.*?\}\}`)

func NewEnvironment(context context.Context, vars ...Vars) *Environment {
	if vars == nil || len(vars) == 0 {
		vars = combinedVars{}
	}
	return &Environment{
		builtin: builtin,
		vars:    combinedVars(vars),
		Context: context,
	}
}

func (e *Environment) GetString(str string) (string, error) {
	if str == "" {
		return "", nil
	}
	return expressionStringRegex.ReplaceAllStringFunc(str, func(s string) string {
		value, err := e.Eval(s[2 : len(s)-2])
		if err == nil && value != nil {
			return string(value.ToString())
		}
		if value == nil {
			return ""
		}
		return string(value.ToString())
	}), nil
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

/* EString */

type EString string

func (v EString) ToString() EString {
	return v
}

func (v EString) Get(name string) (EValue, error) {
	switch name {
	case "bsubmatch":
		return &estringSubmatch{expr: string(v)}, nil
	case "submatch":
		return &estringSubmatch{expr: string(v)}, nil
	}
	panic("implement me")
}

type estringSubmatch struct {
	expr string
}

func (e estringSubmatch) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expect one argument, got %d", len(args))
	}
	regex, err := regexp.Compile(e.expr)
	if err != nil {
		return nil, err
	}
	arg := args[0]

	var str string
	switch arg.(type) {
	case EBytes:
		str = string(arg.(EBytes))
	case EString:
		str = string(arg.(EString))
	default:
		return nil, fmt.Errorf("expect EBytes,EString, got %s", reflect.TypeOf(arg))
	}
	obj := &MapObject{}
	match := regex.FindStringSubmatch(str)

	if match != nil {
		for i, name := range regex.SubexpNames() {
			if i != 0 && name != "" {
				obj.Set(name, EString(match[i]))
			}
		}
	}

	return obj, nil
}

func (e estringSubmatch) ToString() EString {
	panic("implement me")
}

/* EString */

/* EBytes */

type EBytes []byte

func (v EBytes) Get(name string) (EValue, error) {
	switch name {
	case "bcontains":
		return &ebytesBcontains{b: v}, nil
	}
	panic("implement me")
}

func (v EBytes) ToString() EString {
	return EString(v)
}

type ebytesBcontains struct {
	b EBytes
}

func (e ebytesBcontains) Call(args []EValue) (EValue, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("expect one argument, got %d", len(args))
	}
	arg := args[0]

	switch arg.(type) {
	case EBytes:
		return EBool(bytes.Contains(e.b, arg.(EBytes))), nil
	default:
		return nil, fmt.Errorf("expect EBytes, got %s", reflect.TypeOf(arg))
	}
}

func (e ebytesBcontains) ToString() EString {
	panic("implement me")
}

/* EBytes */

type EObject interface {
	EValue
	Get(name string) (EValue, error)
}

type MapObject map[string]EValue

func (m MapObject) ToString() EString {
	return "MapObject"
}

func (m MapObject) Set(name string, value EValue) {
	m[name] = value
}

func (m MapObject) Get(name string) (EValue, error) {
	if value, ok := m[name]; ok {
		return value, nil
	}
	return nil, fmt.Errorf("'%s' undefined", name)
}
