// Code generated from Expression.g4 by ANTLR 4.12.0. DO NOT EDIT.

package expression // Expression
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// ExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type ExpressionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterChainExpression is called when entering the ChainExpression production.
	EnterChainExpression(c *ChainExpressionContext)

	// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterConstExpression is called when entering the ConstExpression production.
	EnterConstExpression(c *ConstExpressionContext)

	// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterObjectAccessExpression is called when entering the ObjectAccessExpression production.
	EnterObjectAccessExpression(c *ObjectAccessExpressionContext)

	// EnterEqualityExpression is called when entering the EqualityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterArrayAccessExpression is called when entering the ArrayAccessExpression production.
	EnterArrayAccessExpression(c *ArrayAccessExpressionContext)

	// EnterFunctionCallExpression is called when entering the FunctionCallExpression production.
	EnterFunctionCallExpression(c *FunctionCallExpressionContext)

	// EnterIdentifierAccessExpression is called when entering the IdentifierAccessExpression production.
	EnterIdentifierAccessExpression(c *IdentifierAccessExpressionContext)

	// EnterExpressionConst is called when entering the expressionConst production.
	EnterExpressionConst(c *ExpressionConstContext)

	// EnterExpressionArguments is called when entering the expressionArguments production.
	EnterExpressionArguments(c *ExpressionArgumentsContext)

	// EnterExpressionArgument is called when entering the expressionArgument production.
	EnterExpressionArgument(c *ExpressionArgumentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitChainExpression is called when exiting the ChainExpression production.
	ExitChainExpression(c *ChainExpressionContext)

	// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitConstExpression is called when exiting the ConstExpression production.
	ExitConstExpression(c *ConstExpressionContext)

	// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitObjectAccessExpression is called when exiting the ObjectAccessExpression production.
	ExitObjectAccessExpression(c *ObjectAccessExpressionContext)

	// ExitEqualityExpression is called when exiting the EqualityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitArrayAccessExpression is called when exiting the ArrayAccessExpression production.
	ExitArrayAccessExpression(c *ArrayAccessExpressionContext)

	// ExitFunctionCallExpression is called when exiting the FunctionCallExpression production.
	ExitFunctionCallExpression(c *FunctionCallExpressionContext)

	// ExitIdentifierAccessExpression is called when exiting the IdentifierAccessExpression production.
	ExitIdentifierAccessExpression(c *IdentifierAccessExpressionContext)

	// ExitExpressionConst is called when exiting the expressionConst production.
	ExitExpressionConst(c *ExpressionConstContext)

	// ExitExpressionArguments is called when exiting the expressionArguments production.
	ExitExpressionArguments(c *ExpressionArgumentsContext)

	// ExitExpressionArgument is called when exiting the expressionArgument production.
	ExitExpressionArgument(c *ExpressionArgumentContext)
}
