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

	// EnterPlusExpression is called when entering the PlusExpression production.
	EnterPlusExpression(c *PlusExpressionContext)

	// EnterLogicalAndExpression is called when entering the LogicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterConstExpression is called when entering the ConstExpression production.
	EnterConstExpression(c *ConstExpressionContext)

	// EnterLogicalOrExpression is called when entering the LogicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterParenExpression is called when entering the ParenExpression production.
	EnterParenExpression(c *ParenExpressionContext)

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

	// EnterBooleanLiteral is called when entering the BooleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterBinaryStringLiteral is called when entering the BinaryStringLiteral production.
	EnterBinaryStringLiteral(c *BinaryStringLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterIntegerLiteral is called when entering the IntegerLiteral production.
	EnterIntegerLiteral(c *IntegerLiteralContext)

	// EnterFloatingPointLiteral is called when entering the FloatingPointLiteral production.
	EnterFloatingPointLiteral(c *FloatingPointLiteralContext)

	// EnterExpressionArguments is called when entering the expressionArguments production.
	EnterExpressionArguments(c *ExpressionArgumentsContext)

	// EnterExpressionArgument is called when entering the expressionArgument production.
	EnterExpressionArgument(c *ExpressionArgumentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitChainExpression is called when exiting the ChainExpression production.
	ExitChainExpression(c *ChainExpressionContext)

	// ExitPlusExpression is called when exiting the PlusExpression production.
	ExitPlusExpression(c *PlusExpressionContext)

	// ExitLogicalAndExpression is called when exiting the LogicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitConstExpression is called when exiting the ConstExpression production.
	ExitConstExpression(c *ConstExpressionContext)

	// ExitLogicalOrExpression is called when exiting the LogicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitParenExpression is called when exiting the ParenExpression production.
	ExitParenExpression(c *ParenExpressionContext)

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

	// ExitBooleanLiteral is called when exiting the BooleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitBinaryStringLiteral is called when exiting the BinaryStringLiteral production.
	ExitBinaryStringLiteral(c *BinaryStringLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitIntegerLiteral is called when exiting the IntegerLiteral production.
	ExitIntegerLiteral(c *IntegerLiteralContext)

	// ExitFloatingPointLiteral is called when exiting the FloatingPointLiteral production.
	ExitFloatingPointLiteral(c *FloatingPointLiteralContext)

	// ExitExpressionArguments is called when exiting the expressionArguments production.
	ExitExpressionArguments(c *ExpressionArgumentsContext)

	// ExitExpressionArgument is called when exiting the expressionArgument production.
	ExitExpressionArgument(c *ExpressionArgumentContext)
}
