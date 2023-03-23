// Code generated from Expression.g4 by ANTLR 4.12.0. DO NOT EDIT.

package expression // Expression
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ExpressionParser struct {
	*antlr.BaseParser
}

var expressionParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expressionParserInit() {
	staticData := &expressionParserStaticData
	staticData.literalNames = []string{
		"", "'&&'", "'||'", "','", "'.'", "'['", "']'", "'('", "')'", "'=='",
		"'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "DOT", "LBRACK", "RBRACK", "LPAREN", "RPAREN", "EQUAL",
		"NOTEQUAL", "WhiteSpaces", "Identifier", "IntegerLiteral", "FloatingPointLiteral",
		"BooleanLiteral", "StringLiteral",
	}
	staticData.ruleNames = []string{
		"expression", "expressionSingle", "expressionConst", "expressionArguments",
		"expressionArgument",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 16, 66, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 3, 1, 17, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 44, 8,
		1, 10, 1, 12, 1, 47, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 5, 3, 54, 8, 3,
		10, 3, 12, 3, 57, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 64, 8, 4, 1,
		4, 0, 1, 2, 5, 0, 2, 4, 6, 8, 0, 2, 1, 0, 9, 10, 2, 0, 13, 14, 16, 16,
		73, 0, 10, 1, 0, 0, 0, 2, 16, 1, 0, 0, 0, 4, 48, 1, 0, 0, 0, 6, 50, 1,
		0, 0, 0, 8, 63, 1, 0, 0, 0, 10, 11, 3, 2, 1, 0, 11, 12, 5, 0, 0, 1, 12,
		1, 1, 0, 0, 0, 13, 14, 6, 1, -1, 0, 14, 17, 5, 12, 0, 0, 15, 17, 3, 4,
		2, 0, 16, 13, 1, 0, 0, 0, 16, 15, 1, 0, 0, 0, 17, 45, 1, 0, 0, 0, 18, 19,
		10, 7, 0, 0, 19, 20, 5, 4, 0, 0, 20, 44, 3, 2, 1, 8, 21, 22, 10, 3, 0,
		0, 22, 23, 7, 0, 0, 0, 23, 44, 3, 2, 1, 4, 24, 25, 10, 2, 0, 0, 25, 26,
		5, 1, 0, 0, 26, 44, 3, 2, 1, 3, 27, 28, 10, 1, 0, 0, 28, 29, 5, 2, 0, 0,
		29, 44, 3, 2, 1, 2, 30, 31, 10, 6, 0, 0, 31, 32, 5, 5, 0, 0, 32, 33, 5,
		16, 0, 0, 33, 44, 5, 6, 0, 0, 34, 35, 10, 5, 0, 0, 35, 36, 5, 5, 0, 0,
		36, 37, 5, 13, 0, 0, 37, 44, 5, 6, 0, 0, 38, 39, 10, 4, 0, 0, 39, 40, 5,
		7, 0, 0, 40, 41, 3, 6, 3, 0, 41, 42, 5, 8, 0, 0, 42, 44, 1, 0, 0, 0, 43,
		18, 1, 0, 0, 0, 43, 21, 1, 0, 0, 0, 43, 24, 1, 0, 0, 0, 43, 27, 1, 0, 0,
		0, 43, 30, 1, 0, 0, 0, 43, 34, 1, 0, 0, 0, 43, 38, 1, 0, 0, 0, 44, 47,
		1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 3, 1, 0, 0, 0,
		47, 45, 1, 0, 0, 0, 48, 49, 7, 1, 0, 0, 49, 5, 1, 0, 0, 0, 50, 55, 3, 8,
		4, 0, 51, 52, 5, 3, 0, 0, 52, 54, 3, 8, 4, 0, 53, 51, 1, 0, 0, 0, 54, 57,
		1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0,
		57, 55, 1, 0, 0, 0, 58, 64, 1, 0, 0, 0, 59, 64, 5, 12, 0, 0, 60, 64, 5,
		13, 0, 0, 61, 64, 5, 16, 0, 0, 62, 64, 3, 2, 1, 0, 63, 58, 1, 0, 0, 0,
		63, 59, 1, 0, 0, 0, 63, 60, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 62, 1,
		0, 0, 0, 64, 9, 1, 0, 0, 0, 5, 16, 43, 45, 55, 63,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ExpressionParserInit initializes any static state used to implement ExpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressionParserInit() {
	staticData := &expressionParserStaticData
	staticData.once.Do(expressionParserInit)
}

// NewExpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExpressionParser(input antlr.TokenStream) *ExpressionParser {
	ExpressionParserInit()
	this := new(ExpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &expressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Expression.g4"

	return this
}

// ExpressionParser tokens.
const (
	ExpressionParserEOF                  = antlr.TokenEOF
	ExpressionParserT__0                 = 1
	ExpressionParserT__1                 = 2
	ExpressionParserT__2                 = 3
	ExpressionParserDOT                  = 4
	ExpressionParserLBRACK               = 5
	ExpressionParserRBRACK               = 6
	ExpressionParserLPAREN               = 7
	ExpressionParserRPAREN               = 8
	ExpressionParserEQUAL                = 9
	ExpressionParserNOTEQUAL             = 10
	ExpressionParserWhiteSpaces          = 11
	ExpressionParserIdentifier           = 12
	ExpressionParserIntegerLiteral       = 13
	ExpressionParserFloatingPointLiteral = 14
	ExpressionParserBooleanLiteral       = 15
	ExpressionParserStringLiteral        = 16
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_expression          = 0
	ExpressionParserRULE_expressionSingle    = 1
	ExpressionParserRULE_expressionConst     = 2
	ExpressionParserRULE_expressionArguments = 3
	ExpressionParserRULE_expressionArgument  = 4
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionSingle() IExpressionSingleContext
	EOF() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEOF, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *ExpressionParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ExpressionParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.expressionSingle(0)
	}
	{
		p.SetState(11)
		p.Match(ExpressionParserEOF)
	}

	return localctx
}

// IExpressionSingleContext is an interface to support dynamic dispatch.
type IExpressionSingleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionSingleContext differentiates from other interfaces.
	IsExpressionSingleContext()
}

type ExpressionSingleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSingleContext() *ExpressionSingleContext {
	var p = new(ExpressionSingleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionSingle
	return p
}

func (*ExpressionSingleContext) IsExpressionSingleContext() {}

func NewExpressionSingleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSingleContext {
	var p = new(ExpressionSingleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionSingle

	return p
}

func (s *ExpressionSingleContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSingleContext) CopyFrom(ctx *ExpressionSingleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSingleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ChainExpressionContext struct {
	*ExpressionSingleContext
}

func NewChainExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChainExpressionContext {
	var p = new(ChainExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ChainExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *ChainExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ChainExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExpressionParserDOT, 0)
}

func (s *ChainExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterChainExpression(s)
	}
}

func (s *ChainExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitChainExpression(s)
	}
}

type LogicalAndExpressionContext struct {
	*ExpressionSingleContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

type ConstExpressionContext struct {
	*ExpressionSingleContext
}

func NewConstExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstExpressionContext {
	var p = new(ConstExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ConstExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstExpressionContext) ExpressionConst() IExpressionConstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionConstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionConstContext)
}

func (s *ConstExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterConstExpression(s)
	}
}

func (s *ConstExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitConstExpression(s)
	}
}

type LogicalOrExpressionContext struct {
	*ExpressionSingleContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

type ObjectAccessExpressionContext struct {
	*ExpressionSingleContext
}

func NewObjectAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectAccessExpressionContext {
	var p = new(ObjectAccessExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ObjectAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectAccessExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ObjectAccessExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLBRACK, 0)
}

func (s *ObjectAccessExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *ObjectAccessExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRBRACK, 0)
}

func (s *ObjectAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterObjectAccessExpression(s)
	}
}

func (s *ObjectAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitObjectAccessExpression(s)
	}
}

type EqualityExpressionContext struct {
	*ExpressionSingleContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			len++
		}
	}

	tst := make([]IExpressionSingleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionSingleContext); ok {
			tst[i] = t.(IExpressionSingleContext)
			i++
		}
	}

	return tst
}

func (s *EqualityExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *EqualityExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEQUAL, 0)
}

func (s *EqualityExpressionContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserNOTEQUAL, 0)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

type ArrayAccessExpressionContext struct {
	*ExpressionSingleContext
}

func NewArrayAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayAccessExpressionContext {
	var p = new(ArrayAccessExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ArrayAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccessExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ArrayAccessExpressionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLBRACK, 0)
}

func (s *ArrayAccessExpressionContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIntegerLiteral, 0)
}

func (s *ArrayAccessExpressionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRBRACK, 0)
}

func (s *ArrayAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterArrayAccessExpression(s)
	}
}

func (s *ArrayAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitArrayAccessExpression(s)
	}
}

type FunctionCallExpressionContext struct {
	*ExpressionSingleContext
}

func NewFunctionCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionContext {
	var p = new(FunctionCallExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *FunctionCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *FunctionCallExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *FunctionCallExpressionContext) ExpressionArguments() IExpressionArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionArgumentsContext)
}

func (s *FunctionCallExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *FunctionCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterFunctionCallExpression(s)
	}
}

func (s *FunctionCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitFunctionCallExpression(s)
	}
}

type IdentifierAccessExpressionContext struct {
	*ExpressionSingleContext
}

func NewIdentifierAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierAccessExpressionContext {
	var p = new(IdentifierAccessExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *IdentifierAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierAccessExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIdentifier, 0)
}

func (s *IdentifierAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterIdentifierAccessExpression(s)
	}
}

func (s *IdentifierAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitIdentifierAccessExpression(s)
	}
}

func (p *ExpressionParser) ExpressionSingle() (localctx IExpressionSingleContext) {
	return p.expressionSingle(0)
}

func (p *ExpressionParser) expressionSingle(_p int) (localctx IExpressionSingleContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionSingleContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionSingleContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ExpressionParserRULE_expressionSingle, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(16)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserIdentifier:
		localctx = NewIdentifierAccessExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(14)
			p.Match(ExpressionParserIdentifier)
		}

	case ExpressionParserIntegerLiteral, ExpressionParserFloatingPointLiteral, ExpressionParserStringLiteral:
		localctx = NewConstExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(15)
			p.ExpressionConst()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(43)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewChainExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(18)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(19)
					p.Match(ExpressionParserDOT)
				}
				{
					p.SetState(20)
					p.expressionSingle(8)
				}

			case 2:
				localctx = NewEqualityExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(21)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(22)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserEQUAL || _la == ExpressionParserNOTEQUAL) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(23)
					p.expressionSingle(4)
				}

			case 3:
				localctx = NewLogicalAndExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(25)
					p.Match(ExpressionParserT__0)
				}
				{
					p.SetState(26)
					p.expressionSingle(3)
				}

			case 4:
				localctx = NewLogicalOrExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(28)
					p.Match(ExpressionParserT__1)
				}
				{
					p.SetState(29)
					p.expressionSingle(2)
				}

			case 5:
				localctx = NewObjectAccessExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(31)
					p.Match(ExpressionParserLBRACK)
				}
				{
					p.SetState(32)
					p.Match(ExpressionParserStringLiteral)
				}
				{
					p.SetState(33)
					p.Match(ExpressionParserRBRACK)
				}

			case 6:
				localctx = NewArrayAccessExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(34)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(35)
					p.Match(ExpressionParserLBRACK)
				}
				{
					p.SetState(36)
					p.Match(ExpressionParserIntegerLiteral)
				}
				{
					p.SetState(37)
					p.Match(ExpressionParserRBRACK)
				}

			case 7:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(39)
					p.Match(ExpressionParserLPAREN)
				}
				{
					p.SetState(40)
					p.ExpressionArguments()
				}
				{
					p.SetState(41)
					p.Match(ExpressionParserRPAREN)
				}

			}

		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionConstContext is an interface to support dynamic dispatch.
type IExpressionConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	StringLiteral() antlr.TerminalNode
	IntegerLiteral() antlr.TerminalNode
	FloatingPointLiteral() antlr.TerminalNode

	// IsExpressionConstContext differentiates from other interfaces.
	IsExpressionConstContext()
}

type ExpressionConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionConstContext() *ExpressionConstContext {
	var p = new(ExpressionConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionConst
	return p
}

func (*ExpressionConstContext) IsExpressionConstContext() {}

func NewExpressionConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionConstContext {
	var p = new(ExpressionConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionConst

	return p
}

func (s *ExpressionConstContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionConstContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *ExpressionConstContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIntegerLiteral, 0)
}

func (s *ExpressionConstContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserFloatingPointLiteral, 0)
}

func (s *ExpressionConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpressionConst(s)
	}
}

func (s *ExpressionConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpressionConst(s)
	}
}

func (p *ExpressionParser) ExpressionConst() (localctx IExpressionConstContext) {
	this := p
	_ = this

	localctx = NewExpressionConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExpressionParserRULE_expressionConst)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&90112) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionArgumentsContext is an interface to support dynamic dispatch.
type IExpressionArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpressionArgument() []IExpressionArgumentContext
	ExpressionArgument(i int) IExpressionArgumentContext

	// IsExpressionArgumentsContext differentiates from other interfaces.
	IsExpressionArgumentsContext()
}

type ExpressionArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionArgumentsContext() *ExpressionArgumentsContext {
	var p = new(ExpressionArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionArguments
	return p
}

func (*ExpressionArgumentsContext) IsExpressionArgumentsContext() {}

func NewExpressionArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionArgumentsContext {
	var p = new(ExpressionArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionArguments

	return p
}

func (s *ExpressionArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionArgumentsContext) AllExpressionArgument() []IExpressionArgumentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionArgumentContext); ok {
			len++
		}
	}

	tst := make([]IExpressionArgumentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionArgumentContext); ok {
			tst[i] = t.(IExpressionArgumentContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionArgumentsContext) ExpressionArgument(i int) IExpressionArgumentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionArgumentContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionArgumentContext)
}

func (s *ExpressionArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpressionArguments(s)
	}
}

func (s *ExpressionArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpressionArguments(s)
	}
}

func (p *ExpressionParser) ExpressionArguments() (localctx IExpressionArgumentsContext) {
	this := p
	_ = this

	localctx = NewExpressionArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExpressionParserRULE_expressionArguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.ExpressionArgument()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExpressionParserT__2 {
		{
			p.SetState(51)
			p.Match(ExpressionParserT__2)
		}
		{
			p.SetState(52)
			p.ExpressionArgument()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionArgumentContext is an interface to support dynamic dispatch.
type IExpressionArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	IntegerLiteral() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	ExpressionSingle() IExpressionSingleContext

	// IsExpressionArgumentContext differentiates from other interfaces.
	IsExpressionArgumentContext()
}

type ExpressionArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionArgumentContext() *ExpressionArgumentContext {
	var p = new(ExpressionArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionArgument
	return p
}

func (*ExpressionArgumentContext) IsExpressionArgumentContext() {}

func NewExpressionArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionArgumentContext {
	var p = new(ExpressionArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionArgument

	return p
}

func (s *ExpressionArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionArgumentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIdentifier, 0)
}

func (s *ExpressionArgumentContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIntegerLiteral, 0)
}

func (s *ExpressionArgumentContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *ExpressionArgumentContext) ExpressionSingle() IExpressionSingleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionSingleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionSingleContext)
}

func (s *ExpressionArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpressionArgument(s)
	}
}

func (s *ExpressionArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpressionArgument(s)
	}
}

func (p *ExpressionParser) ExpressionArgument() (localctx IExpressionArgumentContext) {
	this := p
	_ = this

	localctx = NewExpressionArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExpressionParserRULE_expressionArgument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(ExpressionParserIdentifier)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.Match(ExpressionParserIntegerLiteral)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.Match(ExpressionParserStringLiteral)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(62)
			p.expressionSingle(0)
		}

	}

	return localctx
}

func (p *ExpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpressionSingleContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionSingleContext)
		}
		return p.ExpressionSingle_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExpressionParser) ExpressionSingle_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
