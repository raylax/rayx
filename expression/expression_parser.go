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
		"", "'+'", "'&&'", "'||'", "'b'", "','", "'.'", "'['", "']'", "'('",
		"')'", "'=='", "'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "DOT", "LBRACK", "RBRACK", "LPAREN", "RPAREN",
		"EQUAL", "NOTEQUAL", "WhiteSpaces", "BooleanLiteral", "Identifier",
		"IntegerLiteral", "FloatingPointLiteral", "StringLiteral", "Whitespace",
	}
	staticData.ruleNames = []string{
		"expression", "expressionSingle", "expressionMember", "expressionConst",
		"expressionArguments", "expressionArgument",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 88, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 3, 1, 23, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 48, 8, 1, 10, 1, 12, 1, 51, 9, 1, 1, 2, 1, 2, 3, 2, 55,
		8, 2, 1, 2, 3, 2, 58, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 63, 8, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 71, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 76,
		8, 4, 10, 4, 12, 4, 79, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 86, 8,
		5, 1, 5, 0, 1, 2, 6, 0, 2, 4, 6, 8, 10, 0, 1, 1, 0, 11, 12, 102, 0, 12,
		1, 0, 0, 0, 2, 22, 1, 0, 0, 0, 4, 62, 1, 0, 0, 0, 6, 70, 1, 0, 0, 0, 8,
		72, 1, 0, 0, 0, 10, 85, 1, 0, 0, 0, 12, 13, 3, 2, 1, 0, 13, 14, 5, 0, 0,
		1, 14, 1, 1, 0, 0, 0, 15, 16, 6, 1, -1, 0, 16, 23, 3, 6, 3, 0, 17, 23,
		5, 15, 0, 0, 18, 19, 5, 9, 0, 0, 19, 20, 3, 2, 1, 0, 20, 21, 5, 10, 0,
		0, 21, 23, 1, 0, 0, 0, 22, 15, 1, 0, 0, 0, 22, 17, 1, 0, 0, 0, 22, 18,
		1, 0, 0, 0, 23, 49, 1, 0, 0, 0, 24, 25, 10, 4, 0, 0, 25, 26, 5, 1, 0, 0,
		26, 48, 3, 2, 1, 5, 27, 28, 10, 3, 0, 0, 28, 29, 7, 0, 0, 0, 29, 48, 3,
		2, 1, 4, 30, 31, 10, 2, 0, 0, 31, 32, 5, 2, 0, 0, 32, 48, 3, 2, 1, 3, 33,
		34, 10, 1, 0, 0, 34, 35, 5, 3, 0, 0, 35, 48, 3, 2, 1, 2, 36, 37, 10, 7,
		0, 0, 37, 48, 3, 4, 2, 0, 38, 39, 10, 6, 0, 0, 39, 40, 5, 7, 0, 0, 40,
		41, 5, 16, 0, 0, 41, 48, 5, 8, 0, 0, 42, 43, 10, 5, 0, 0, 43, 44, 5, 9,
		0, 0, 44, 45, 3, 8, 4, 0, 45, 46, 5, 10, 0, 0, 46, 48, 1, 0, 0, 0, 47,
		24, 1, 0, 0, 0, 47, 27, 1, 0, 0, 0, 47, 30, 1, 0, 0, 0, 47, 33, 1, 0, 0,
		0, 47, 36, 1, 0, 0, 0, 47, 38, 1, 0, 0, 0, 47, 42, 1, 0, 0, 0, 48, 51,
		1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 3, 1, 0, 0, 0,
		51, 49, 1, 0, 0, 0, 52, 54, 5, 6, 0, 0, 53, 55, 5, 4, 0, 0, 54, 53, 1,
		0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0, 0, 56, 58, 5, 15, 0, 0, 57,
		56, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0, 58, 63, 1, 0, 0, 0, 59, 60, 5, 7, 0,
		0, 60, 61, 5, 18, 0, 0, 61, 63, 5, 8, 0, 0, 62, 52, 1, 0, 0, 0, 62, 59,
		1, 0, 0, 0, 63, 5, 1, 0, 0, 0, 64, 71, 5, 14, 0, 0, 65, 66, 5, 4, 0, 0,
		66, 71, 5, 18, 0, 0, 67, 71, 5, 18, 0, 0, 68, 71, 5, 16, 0, 0, 69, 71,
		5, 17, 0, 0, 70, 64, 1, 0, 0, 0, 70, 65, 1, 0, 0, 0, 70, 67, 1, 0, 0, 0,
		70, 68, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 7, 1, 0, 0, 0, 72, 77, 3, 10,
		5, 0, 73, 74, 5, 5, 0, 0, 74, 76, 3, 10, 5, 0, 75, 73, 1, 0, 0, 0, 76,
		79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 9, 1, 0, 0,
		0, 79, 77, 1, 0, 0, 0, 80, 86, 1, 0, 0, 0, 81, 86, 5, 15, 0, 0, 82, 86,
		5, 16, 0, 0, 83, 86, 5, 18, 0, 0, 84, 86, 3, 2, 1, 0, 85, 80, 1, 0, 0,
		0, 85, 81, 1, 0, 0, 0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84,
		1, 0, 0, 0, 86, 11, 1, 0, 0, 0, 9, 22, 47, 49, 54, 57, 62, 70, 77, 85,
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
	ExpressionParserT__3                 = 4
	ExpressionParserT__4                 = 5
	ExpressionParserDOT                  = 6
	ExpressionParserLBRACK               = 7
	ExpressionParserRBRACK               = 8
	ExpressionParserLPAREN               = 9
	ExpressionParserRPAREN               = 10
	ExpressionParserEQUAL                = 11
	ExpressionParserNOTEQUAL             = 12
	ExpressionParserWhiteSpaces          = 13
	ExpressionParserBooleanLiteral       = 14
	ExpressionParserIdentifier           = 15
	ExpressionParserIntegerLiteral       = 16
	ExpressionParserFloatingPointLiteral = 17
	ExpressionParserStringLiteral        = 18
	ExpressionParserWhitespace           = 19
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_expression          = 0
	ExpressionParserRULE_expressionSingle    = 1
	ExpressionParserRULE_expressionMember    = 2
	ExpressionParserRULE_expressionConst     = 3
	ExpressionParserRULE_expressionArguments = 4
	ExpressionParserRULE_expressionArgument  = 5
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
		p.SetState(12)
		p.expressionSingle(0)
	}
	{
		p.SetState(13)
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

type MemberAccessExpressionContext struct {
	*ExpressionSingleContext
}

func NewMemberAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberAccessExpressionContext {
	var p = new(MemberAccessExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *MemberAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessExpressionContext) ExpressionSingle() IExpressionSingleContext {
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

func (s *MemberAccessExpressionContext) ExpressionMember() IExpressionMemberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionMemberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionMemberContext)
}

func (s *MemberAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterMemberAccessExpression(s)
	}
}

func (s *MemberAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitMemberAccessExpression(s)
	}
}

type PlusExpressionContext struct {
	*ExpressionSingleContext
}

func NewPlusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusExpressionContext {
	var p = new(PlusExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *PlusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusExpressionContext) AllExpressionSingle() []IExpressionSingleContext {
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

func (s *PlusExpressionContext) ExpressionSingle(i int) IExpressionSingleContext {
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

func (s *PlusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterPlusExpression(s)
	}
}

func (s *PlusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitPlusExpression(s)
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

type ParenExpressionContext struct {
	*ExpressionSingleContext
}

func NewParenExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpressionContext {
	var p = new(ParenExpressionContext)

	p.ExpressionSingleContext = NewEmptyExpressionSingleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionSingleContext))

	return p
}

func (s *ParenExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *ParenExpressionContext) ExpressionSingle() IExpressionSingleContext {
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

func (s *ParenExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *ParenExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterParenExpression(s)
	}
}

func (s *ParenExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitParenExpression(s)
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
	p.SetState(22)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserT__3, ExpressionParserBooleanLiteral, ExpressionParserIntegerLiteral, ExpressionParserFloatingPointLiteral, ExpressionParserStringLiteral:
		localctx = NewConstExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(16)
			p.ExpressionConst()
		}

	case ExpressionParserIdentifier:
		localctx = NewIdentifierAccessExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(17)
			p.Match(ExpressionParserIdentifier)
		}

	case ExpressionParserLPAREN:
		localctx = NewParenExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(18)
			p.Match(ExpressionParserLPAREN)
		}
		{
			p.SetState(19)
			p.expressionSingle(0)
		}
		{
			p.SetState(20)
			p.Match(ExpressionParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPlusExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(24)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(25)
					p.Match(ExpressionParserT__0)
				}
				{
					p.SetState(26)
					p.expressionSingle(5)
				}

			case 2:
				localctx = NewEqualityExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(27)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(28)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserEQUAL || _la == ExpressionParserNOTEQUAL) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(29)
					p.expressionSingle(4)
				}

			case 3:
				localctx = NewLogicalAndExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(30)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(31)
					p.Match(ExpressionParserT__1)
				}
				{
					p.SetState(32)
					p.expressionSingle(3)
				}

			case 4:
				localctx = NewLogicalOrExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(34)
					p.Match(ExpressionParserT__2)
				}
				{
					p.SetState(35)
					p.expressionSingle(2)
				}

			case 5:
				localctx = NewMemberAccessExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(37)
					p.ExpressionMember()
				}

			case 6:
				localctx = NewArrayAccessExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(38)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(39)
					p.Match(ExpressionParserLBRACK)
				}
				{
					p.SetState(40)
					p.Match(ExpressionParserIntegerLiteral)
				}
				{
					p.SetState(41)
					p.Match(ExpressionParserRBRACK)
				}

			case 7:
				localctx = NewFunctionCallExpressionContext(p, NewExpressionSingleContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expressionSingle)
				p.SetState(42)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(43)
					p.Match(ExpressionParserLPAREN)
				}
				{
					p.SetState(44)
					p.ExpressionArguments()
				}
				{
					p.SetState(45)
					p.Match(ExpressionParserRPAREN)
				}

			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionMemberContext is an interface to support dynamic dispatch.
type IExpressionMemberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOT() antlr.TerminalNode
	Identifier() antlr.TerminalNode
	LBRACK() antlr.TerminalNode
	StringLiteral() antlr.TerminalNode
	RBRACK() antlr.TerminalNode

	// IsExpressionMemberContext differentiates from other interfaces.
	IsExpressionMemberContext()
}

type ExpressionMemberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionMemberContext() *ExpressionMemberContext {
	var p = new(ExpressionMemberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ExpressionParserRULE_expressionMember
	return p
}

func (*ExpressionMemberContext) IsExpressionMemberContext() {}

func NewExpressionMemberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionMemberContext {
	var p = new(ExpressionMemberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expressionMember

	return p
}

func (s *ExpressionMemberContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionMemberContext) DOT() antlr.TerminalNode {
	return s.GetToken(ExpressionParserDOT, 0)
}

func (s *ExpressionMemberContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIdentifier, 0)
}

func (s *ExpressionMemberContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLBRACK, 0)
}

func (s *ExpressionMemberContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *ExpressionMemberContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRBRACK, 0)
}

func (s *ExpressionMemberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionMemberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionMemberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterExpressionMember(s)
	}
}

func (s *ExpressionMemberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitExpressionMember(s)
	}
}

func (p *ExpressionParser) ExpressionMember() (localctx IExpressionMemberContext) {
	this := p
	_ = this

	localctx = NewExpressionMemberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExpressionParserRULE_expressionMember)

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

	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserDOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(ExpressionParserDOT)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(53)
				p.Match(ExpressionParserT__3)
			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(56)
				p.Match(ExpressionParserIdentifier)
			}

		}

	case ExpressionParserLBRACK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(ExpressionParserLBRACK)
		}
		{
			p.SetState(60)
			p.Match(ExpressionParserStringLiteral)
		}
		{
			p.SetState(61)
			p.Match(ExpressionParserRBRACK)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionConstContext is an interface to support dynamic dispatch.
type IExpressionConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *ExpressionConstContext) CopyFrom(ctx *ExpressionConstContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BinaryStringLiteralContext struct {
	*ExpressionConstContext
}

func NewBinaryStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryStringLiteralContext {
	var p = new(BinaryStringLiteralContext)

	p.ExpressionConstContext = NewEmptyExpressionConstContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionConstContext))

	return p
}

func (s *BinaryStringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryStringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *BinaryStringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterBinaryStringLiteral(s)
	}
}

func (s *BinaryStringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitBinaryStringLiteral(s)
	}
}

type StringLiteralContext struct {
	*ExpressionConstContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.ExpressionConstContext = NewEmptyExpressionConstContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionConstContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserStringLiteral, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type FloatingPointLiteralContext struct {
	*ExpressionConstContext
}

func NewFloatingPointLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatingPointLiteralContext {
	var p = new(FloatingPointLiteralContext)

	p.ExpressionConstContext = NewEmptyExpressionConstContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionConstContext))

	return p
}

func (s *FloatingPointLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingPointLiteralContext) FloatingPointLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserFloatingPointLiteral, 0)
}

func (s *FloatingPointLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterFloatingPointLiteral(s)
	}
}

func (s *FloatingPointLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitFloatingPointLiteral(s)
	}
}

type BooleanLiteralContext struct {
	*ExpressionConstContext
}

func NewBooleanLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.ExpressionConstContext = NewEmptyExpressionConstContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionConstContext))

	return p
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserBooleanLiteral, 0)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

type IntegerLiteralContext struct {
	*ExpressionConstContext
}

func NewIntegerLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.ExpressionConstContext = NewEmptyExpressionConstContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionConstContext))

	return p
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIntegerLiteral, 0)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ExpressionListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *ExpressionParser) ExpressionConst() (localctx IExpressionConstContext) {
	this := p
	_ = this

	localctx = NewExpressionConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExpressionParserRULE_expressionConst)

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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserBooleanLiteral:
		localctx = NewBooleanLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(ExpressionParserBooleanLiteral)
		}

	case ExpressionParserT__3:
		localctx = NewBinaryStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Match(ExpressionParserT__3)
		}
		{
			p.SetState(66)
			p.Match(ExpressionParserStringLiteral)
		}

	case ExpressionParserStringLiteral:
		localctx = NewStringLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(67)
			p.Match(ExpressionParserStringLiteral)
		}

	case ExpressionParserIntegerLiteral:
		localctx = NewIntegerLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.Match(ExpressionParserIntegerLiteral)
		}

	case ExpressionParserFloatingPointLiteral:
		localctx = NewFloatingPointLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.Match(ExpressionParserFloatingPointLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.EnterRule(localctx, 8, ExpressionParserRULE_expressionArguments)
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
		p.SetState(72)
		p.ExpressionArgument()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ExpressionParserT__4 {
		{
			p.SetState(73)
			p.Match(ExpressionParserT__4)
		}
		{
			p.SetState(74)
			p.ExpressionArgument()
		}

		p.SetState(79)
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
	p.EnterRule(localctx, 10, ExpressionParserRULE_expressionArgument)

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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Match(ExpressionParserIdentifier)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.Match(ExpressionParserIntegerLiteral)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(83)
			p.Match(ExpressionParserStringLiteral)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
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
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
