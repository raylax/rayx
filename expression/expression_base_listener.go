// Code generated from Expression.g4 by ANTLR 4.12.0. DO NOT EDIT.

package expression // Expression
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type BaseExpressionListener struct{}

var _ ExpressionListener = &BaseExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExpressionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExpressionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterChainExpression is called when production ChainExpression is entered.
func (s *BaseExpressionListener) EnterChainExpression(ctx *ChainExpressionContext) {}

// ExitChainExpression is called when production ChainExpression is exited.
func (s *BaseExpressionListener) ExitChainExpression(ctx *ChainExpressionContext) {}

// EnterLogicalAndExpression is called when production LogicalAndExpression is entered.
func (s *BaseExpressionListener) EnterLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// ExitLogicalAndExpression is called when production LogicalAndExpression is exited.
func (s *BaseExpressionListener) ExitLogicalAndExpression(ctx *LogicalAndExpressionContext) {}

// EnterConstExpression is called when production ConstExpression is entered.
func (s *BaseExpressionListener) EnterConstExpression(ctx *ConstExpressionContext) {}

// ExitConstExpression is called when production ConstExpression is exited.
func (s *BaseExpressionListener) ExitConstExpression(ctx *ConstExpressionContext) {}

// EnterLogicalOrExpression is called when production LogicalOrExpression is entered.
func (s *BaseExpressionListener) EnterLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// ExitLogicalOrExpression is called when production LogicalOrExpression is exited.
func (s *BaseExpressionListener) ExitLogicalOrExpression(ctx *LogicalOrExpressionContext) {}

// EnterObjectAccessExpression is called when production ObjectAccessExpression is entered.
func (s *BaseExpressionListener) EnterObjectAccessExpression(ctx *ObjectAccessExpressionContext) {}

// ExitObjectAccessExpression is called when production ObjectAccessExpression is exited.
func (s *BaseExpressionListener) ExitObjectAccessExpression(ctx *ObjectAccessExpressionContext) {}

// EnterEqualityExpression is called when production EqualityExpression is entered.
func (s *BaseExpressionListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production EqualityExpression is exited.
func (s *BaseExpressionListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterArrayAccessExpression is called when production ArrayAccessExpression is entered.
func (s *BaseExpressionListener) EnterArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// ExitArrayAccessExpression is called when production ArrayAccessExpression is exited.
func (s *BaseExpressionListener) ExitArrayAccessExpression(ctx *ArrayAccessExpressionContext) {}

// EnterFunctionCallExpression is called when production FunctionCallExpression is entered.
func (s *BaseExpressionListener) EnterFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// ExitFunctionCallExpression is called when production FunctionCallExpression is exited.
func (s *BaseExpressionListener) ExitFunctionCallExpression(ctx *FunctionCallExpressionContext) {}

// EnterIdentifierAccessExpression is called when production IdentifierAccessExpression is entered.
func (s *BaseExpressionListener) EnterIdentifierAccessExpression(ctx *IdentifierAccessExpressionContext) {
}

// ExitIdentifierAccessExpression is called when production IdentifierAccessExpression is exited.
func (s *BaseExpressionListener) ExitIdentifierAccessExpression(ctx *IdentifierAccessExpressionContext) {
}

// EnterExpressionConst is called when production expressionConst is entered.
func (s *BaseExpressionListener) EnterExpressionConst(ctx *ExpressionConstContext) {}

// ExitExpressionConst is called when production expressionConst is exited.
func (s *BaseExpressionListener) ExitExpressionConst(ctx *ExpressionConstContext) {}

// EnterExpressionArguments is called when production expressionArguments is entered.
func (s *BaseExpressionListener) EnterExpressionArguments(ctx *ExpressionArgumentsContext) {}

// ExitExpressionArguments is called when production expressionArguments is exited.
func (s *BaseExpressionListener) ExitExpressionArguments(ctx *ExpressionArgumentsContext) {}

// EnterExpressionArgument is called when production expressionArgument is entered.
func (s *BaseExpressionListener) EnterExpressionArgument(ctx *ExpressionArgumentContext) {}

// ExitExpressionArgument is called when production expressionArgument is exited.
func (s *BaseExpressionListener) ExitExpressionArgument(ctx *ExpressionArgumentContext) {}
