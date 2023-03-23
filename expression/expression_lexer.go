// Code generated from Expression.g4 by ANTLR 4.12.0. DO NOT EDIT.

package expression

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var expressionlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func expressionlexerLexerInit() {
	staticData := &expressionlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'+'", "'&&'", "'||'", "'b'", "','", "'.'", "'['", "']'", "'('",
		"')'", "'=='", "'!='",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "DOT", "LBRACK", "RBRACK", "LPAREN", "RPAREN",
		"EQUAL", "NOTEQUAL", "WhiteSpaces", "Identifier", "IntegerLiteral",
		"FloatingPointLiteral", "BooleanLiteral", "StringLiteral",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "DOT", "LBRACK", "RBRACK", "LPAREN",
		"RPAREN", "EQUAL", "NOTEQUAL", "WhiteSpaces", "Identifier", "IntegerLiteral",
		"FloatingPointLiteral", "BooleanLiteral", "StringLiteral", "StringCharacters",
		"StringCharacter", "EscapeSequence", "HexEscape", "DecimalIntegerLiteral",
		"HexIntegerLiteral", "OctalIntegerLiteral", "BinaryIntegerLiteral",
		"IntegerTypeSuffix", "DecimalNumeral", "Digits", "Digit", "NonZeroDigit",
		"DigitsAndUnderscores", "DigitOrUnderscore", "Underscores", "HexNumeral",
		"HexDigits", "HexDigit", "HexDigitsAndUnderscores", "HexDigitOrUnderscore",
		"OctalNumeral", "OctalDigits", "OctalDigit", "OctalDigitsAndUnderscores",
		"OctalDigitOrUnderscore", "BinaryNumeral", "BinaryDigits", "BinaryDigit",
		"BinaryDigitsAndUnderscores", "BinaryDigitOrUnderscore", "DecimalFloatingPointLiteral",
		"ExponentPart", "ExponentIndicator", "SignedInteger", "Sign", "FloatTypeSuffix",
		"HexadecimalFloatingPointLiteral", "HexSignificand", "BinaryExponent",
		"BinaryExponentIndicator",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 410, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1,
		3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 4, 12, 149, 8,
		12, 11, 12, 12, 12, 150, 1, 12, 1, 12, 1, 13, 4, 13, 156, 8, 13, 11, 13,
		12, 13, 157, 1, 13, 5, 13, 161, 8, 13, 10, 13, 12, 13, 164, 9, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 3, 14, 170, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 183, 8, 16, 1, 17,
		1, 17, 3, 17, 187, 8, 17, 1, 17, 1, 17, 1, 17, 3, 17, 192, 8, 17, 1, 17,
		3, 17, 195, 8, 17, 1, 18, 4, 18, 198, 8, 18, 11, 18, 12, 18, 199, 1, 19,
		1, 19, 3, 19, 204, 8, 19, 1, 20, 1, 20, 1, 20, 3, 20, 209, 8, 20, 1, 21,
		1, 21, 4, 21, 213, 8, 21, 11, 21, 12, 21, 214, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 3, 22, 222, 8, 22, 1, 23, 1, 23, 3, 23, 226, 8, 23, 1, 24, 1,
		24, 3, 24, 230, 8, 24, 1, 25, 1, 25, 3, 25, 234, 8, 25, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 3, 27, 241, 8, 27, 1, 27, 1, 27, 1, 27, 3, 27, 246, 8,
		27, 3, 27, 248, 8, 27, 1, 28, 1, 28, 3, 28, 252, 8, 28, 1, 28, 3, 28, 255,
		8, 28, 1, 29, 1, 29, 3, 29, 259, 8, 29, 1, 30, 1, 30, 1, 31, 4, 31, 264,
		8, 31, 11, 31, 12, 31, 265, 1, 32, 1, 32, 3, 32, 270, 8, 32, 1, 33, 4,
		33, 273, 8, 33, 11, 33, 12, 33, 274, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35,
		1, 35, 3, 35, 283, 8, 35, 1, 35, 3, 35, 286, 8, 35, 1, 36, 1, 36, 1, 37,
		4, 37, 291, 8, 37, 11, 37, 12, 37, 292, 1, 38, 1, 38, 3, 38, 297, 8, 38,
		1, 39, 1, 39, 3, 39, 301, 8, 39, 1, 39, 1, 39, 1, 40, 1, 40, 3, 40, 307,
		8, 40, 1, 40, 3, 40, 310, 8, 40, 1, 41, 1, 41, 1, 42, 4, 42, 315, 8, 42,
		11, 42, 12, 42, 316, 1, 43, 1, 43, 3, 43, 321, 8, 43, 1, 44, 1, 44, 1,
		44, 1, 44, 1, 45, 1, 45, 3, 45, 329, 8, 45, 1, 45, 3, 45, 332, 8, 45, 1,
		46, 1, 46, 1, 47, 4, 47, 337, 8, 47, 11, 47, 12, 47, 338, 1, 48, 1, 48,
		3, 48, 343, 8, 48, 1, 49, 1, 49, 1, 49, 3, 49, 348, 8, 49, 1, 49, 3, 49,
		351, 8, 49, 1, 49, 3, 49, 354, 8, 49, 1, 49, 1, 49, 1, 49, 3, 49, 359,
		8, 49, 1, 49, 3, 49, 362, 8, 49, 1, 49, 1, 49, 1, 49, 3, 49, 367, 8, 49,
		1, 49, 1, 49, 1, 49, 3, 49, 372, 8, 49, 1, 50, 1, 50, 1, 50, 1, 51, 1,
		51, 1, 52, 3, 52, 380, 8, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54,
		1, 55, 1, 55, 1, 55, 3, 55, 391, 8, 55, 1, 56, 1, 56, 3, 56, 395, 8, 56,
		1, 56, 1, 56, 1, 56, 3, 56, 400, 8, 56, 1, 56, 1, 56, 3, 56, 404, 8, 56,
		1, 57, 1, 57, 1, 57, 1, 58, 1, 58, 0, 0, 59, 1, 1, 3, 2, 5, 3, 7, 4, 9,
		5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14,
		29, 15, 31, 16, 33, 17, 35, 18, 37, 0, 39, 0, 41, 0, 43, 0, 45, 0, 47,
		0, 49, 0, 51, 0, 53, 0, 55, 0, 57, 0, 59, 0, 61, 0, 63, 0, 65, 0, 67, 0,
		69, 0, 71, 0, 73, 0, 75, 0, 77, 0, 79, 0, 81, 0, 83, 0, 85, 0, 87, 0, 89,
		0, 91, 0, 93, 0, 95, 0, 97, 0, 99, 0, 101, 0, 103, 0, 105, 0, 107, 0, 109,
		0, 111, 0, 113, 0, 115, 0, 117, 0, 1, 0, 16, 4, 0, 9, 9, 11, 12, 32, 32,
		160, 160, 4, 0, 36, 36, 65, 90, 95, 95, 97, 122, 5, 0, 36, 36, 48, 57,
		65, 90, 95, 95, 97, 122, 5, 0, 10, 10, 13, 13, 34, 34, 39, 39, 92, 92,
		8, 0, 34, 34, 39, 39, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116,
		116, 2, 0, 76, 76, 108, 108, 1, 0, 49, 57, 2, 0, 88, 88, 120, 120, 3, 0,
		48, 57, 65, 70, 97, 102, 1, 0, 48, 55, 2, 0, 66, 66, 98, 98, 1, 0, 48,
		49, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 4, 0, 68, 68, 70, 70,
		100, 100, 102, 102, 2, 0, 80, 80, 112, 112, 422, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 1, 119, 1, 0, 0, 0, 3, 121, 1, 0, 0, 0, 5, 124, 1,
		0, 0, 0, 7, 127, 1, 0, 0, 0, 9, 129, 1, 0, 0, 0, 11, 131, 1, 0, 0, 0, 13,
		133, 1, 0, 0, 0, 15, 135, 1, 0, 0, 0, 17, 137, 1, 0, 0, 0, 19, 139, 1,
		0, 0, 0, 21, 141, 1, 0, 0, 0, 23, 144, 1, 0, 0, 0, 25, 148, 1, 0, 0, 0,
		27, 155, 1, 0, 0, 0, 29, 169, 1, 0, 0, 0, 31, 171, 1, 0, 0, 0, 33, 182,
		1, 0, 0, 0, 35, 194, 1, 0, 0, 0, 37, 197, 1, 0, 0, 0, 39, 203, 1, 0, 0,
		0, 41, 208, 1, 0, 0, 0, 43, 210, 1, 0, 0, 0, 45, 219, 1, 0, 0, 0, 47, 223,
		1, 0, 0, 0, 49, 227, 1, 0, 0, 0, 51, 231, 1, 0, 0, 0, 53, 235, 1, 0, 0,
		0, 55, 247, 1, 0, 0, 0, 57, 249, 1, 0, 0, 0, 59, 258, 1, 0, 0, 0, 61, 260,
		1, 0, 0, 0, 63, 263, 1, 0, 0, 0, 65, 269, 1, 0, 0, 0, 67, 272, 1, 0, 0,
		0, 69, 276, 1, 0, 0, 0, 71, 280, 1, 0, 0, 0, 73, 287, 1, 0, 0, 0, 75, 290,
		1, 0, 0, 0, 77, 296, 1, 0, 0, 0, 79, 298, 1, 0, 0, 0, 81, 304, 1, 0, 0,
		0, 83, 311, 1, 0, 0, 0, 85, 314, 1, 0, 0, 0, 87, 320, 1, 0, 0, 0, 89, 322,
		1, 0, 0, 0, 91, 326, 1, 0, 0, 0, 93, 333, 1, 0, 0, 0, 95, 336, 1, 0, 0,
		0, 97, 342, 1, 0, 0, 0, 99, 371, 1, 0, 0, 0, 101, 373, 1, 0, 0, 0, 103,
		376, 1, 0, 0, 0, 105, 379, 1, 0, 0, 0, 107, 383, 1, 0, 0, 0, 109, 385,
		1, 0, 0, 0, 111, 387, 1, 0, 0, 0, 113, 403, 1, 0, 0, 0, 115, 405, 1, 0,
		0, 0, 117, 408, 1, 0, 0, 0, 119, 120, 5, 43, 0, 0, 120, 2, 1, 0, 0, 0,
		121, 122, 5, 38, 0, 0, 122, 123, 5, 38, 0, 0, 123, 4, 1, 0, 0, 0, 124,
		125, 5, 124, 0, 0, 125, 126, 5, 124, 0, 0, 126, 6, 1, 0, 0, 0, 127, 128,
		5, 98, 0, 0, 128, 8, 1, 0, 0, 0, 129, 130, 5, 44, 0, 0, 130, 10, 1, 0,
		0, 0, 131, 132, 5, 46, 0, 0, 132, 12, 1, 0, 0, 0, 133, 134, 5, 91, 0, 0,
		134, 14, 1, 0, 0, 0, 135, 136, 5, 93, 0, 0, 136, 16, 1, 0, 0, 0, 137, 138,
		5, 40, 0, 0, 138, 18, 1, 0, 0, 0, 139, 140, 5, 41, 0, 0, 140, 20, 1, 0,
		0, 0, 141, 142, 5, 61, 0, 0, 142, 143, 5, 61, 0, 0, 143, 22, 1, 0, 0, 0,
		144, 145, 5, 33, 0, 0, 145, 146, 5, 61, 0, 0, 146, 24, 1, 0, 0, 0, 147,
		149, 7, 0, 0, 0, 148, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0, 0, 150, 148,
		1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 6, 12,
		0, 0, 153, 26, 1, 0, 0, 0, 154, 156, 7, 1, 0, 0, 155, 154, 1, 0, 0, 0,
		156, 157, 1, 0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158,
		162, 1, 0, 0, 0, 159, 161, 7, 2, 0, 0, 160, 159, 1, 0, 0, 0, 161, 164,
		1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 28, 1, 0,
		0, 0, 164, 162, 1, 0, 0, 0, 165, 170, 3, 45, 22, 0, 166, 170, 3, 47, 23,
		0, 167, 170, 3, 49, 24, 0, 168, 170, 3, 51, 25, 0, 169, 165, 1, 0, 0, 0,
		169, 166, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 168, 1, 0, 0, 0, 170,
		30, 1, 0, 0, 0, 171, 172, 3, 99, 49, 0, 172, 32, 1, 0, 0, 0, 173, 174,
		5, 116, 0, 0, 174, 175, 5, 114, 0, 0, 175, 176, 5, 117, 0, 0, 176, 183,
		5, 101, 0, 0, 177, 178, 5, 102, 0, 0, 178, 179, 5, 97, 0, 0, 179, 180,
		5, 108, 0, 0, 180, 181, 5, 115, 0, 0, 181, 183, 5, 101, 0, 0, 182, 173,
		1, 0, 0, 0, 182, 177, 1, 0, 0, 0, 183, 34, 1, 0, 0, 0, 184, 186, 5, 34,
		0, 0, 185, 187, 3, 37, 18, 0, 186, 185, 1, 0, 0, 0, 186, 187, 1, 0, 0,
		0, 187, 188, 1, 0, 0, 0, 188, 195, 5, 34, 0, 0, 189, 191, 5, 39, 0, 0,
		190, 192, 3, 37, 18, 0, 191, 190, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192,
		193, 1, 0, 0, 0, 193, 195, 5, 39, 0, 0, 194, 184, 1, 0, 0, 0, 194, 189,
		1, 0, 0, 0, 195, 36, 1, 0, 0, 0, 196, 198, 3, 39, 19, 0, 197, 196, 1, 0,
		0, 0, 198, 199, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0,
		200, 38, 1, 0, 0, 0, 201, 204, 8, 3, 0, 0, 202, 204, 3, 41, 20, 0, 203,
		201, 1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 40, 1, 0, 0, 0, 205, 206, 5,
		92, 0, 0, 206, 209, 7, 4, 0, 0, 207, 209, 3, 43, 21, 0, 208, 205, 1, 0,
		0, 0, 208, 207, 1, 0, 0, 0, 209, 42, 1, 0, 0, 0, 210, 212, 5, 92, 0, 0,
		211, 213, 5, 120, 0, 0, 212, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214,
		212, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 1, 0, 0, 0, 216, 217,
		3, 73, 36, 0, 217, 218, 3, 73, 36, 0, 218, 44, 1, 0, 0, 0, 219, 221, 3,
		55, 27, 0, 220, 222, 3, 53, 26, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0,
		0, 0, 222, 46, 1, 0, 0, 0, 223, 225, 3, 69, 34, 0, 224, 226, 3, 53, 26,
		0, 225, 224, 1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 48, 1, 0, 0, 0, 227,
		229, 3, 79, 39, 0, 228, 230, 3, 53, 26, 0, 229, 228, 1, 0, 0, 0, 229, 230,
		1, 0, 0, 0, 230, 50, 1, 0, 0, 0, 231, 233, 3, 89, 44, 0, 232, 234, 3, 53,
		26, 0, 233, 232, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234, 52, 1, 0, 0, 0,
		235, 236, 7, 5, 0, 0, 236, 54, 1, 0, 0, 0, 237, 248, 5, 48, 0, 0, 238,
		245, 3, 61, 30, 0, 239, 241, 3, 57, 28, 0, 240, 239, 1, 0, 0, 0, 240, 241,
		1, 0, 0, 0, 241, 246, 1, 0, 0, 0, 242, 243, 3, 67, 33, 0, 243, 244, 3,
		57, 28, 0, 244, 246, 1, 0, 0, 0, 245, 240, 1, 0, 0, 0, 245, 242, 1, 0,
		0, 0, 246, 248, 1, 0, 0, 0, 247, 237, 1, 0, 0, 0, 247, 238, 1, 0, 0, 0,
		248, 56, 1, 0, 0, 0, 249, 254, 3, 59, 29, 0, 250, 252, 3, 63, 31, 0, 251,
		250, 1, 0, 0, 0, 251, 252, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 255,
		3, 59, 29, 0, 254, 251, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 58, 1, 0,
		0, 0, 256, 259, 5, 48, 0, 0, 257, 259, 3, 61, 30, 0, 258, 256, 1, 0, 0,
		0, 258, 257, 1, 0, 0, 0, 259, 60, 1, 0, 0, 0, 260, 261, 7, 6, 0, 0, 261,
		62, 1, 0, 0, 0, 262, 264, 3, 65, 32, 0, 263, 262, 1, 0, 0, 0, 264, 265,
		1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 64, 1, 0,
		0, 0, 267, 270, 3, 59, 29, 0, 268, 270, 5, 95, 0, 0, 269, 267, 1, 0, 0,
		0, 269, 268, 1, 0, 0, 0, 270, 66, 1, 0, 0, 0, 271, 273, 5, 95, 0, 0, 272,
		271, 1, 0, 0, 0, 273, 274, 1, 0, 0, 0, 274, 272, 1, 0, 0, 0, 274, 275,
		1, 0, 0, 0, 275, 68, 1, 0, 0, 0, 276, 277, 5, 48, 0, 0, 277, 278, 7, 7,
		0, 0, 278, 279, 3, 71, 35, 0, 279, 70, 1, 0, 0, 0, 280, 285, 3, 73, 36,
		0, 281, 283, 3, 75, 37, 0, 282, 281, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0,
		283, 284, 1, 0, 0, 0, 284, 286, 3, 73, 36, 0, 285, 282, 1, 0, 0, 0, 285,
		286, 1, 0, 0, 0, 286, 72, 1, 0, 0, 0, 287, 288, 7, 8, 0, 0, 288, 74, 1,
		0, 0, 0, 289, 291, 3, 77, 38, 0, 290, 289, 1, 0, 0, 0, 291, 292, 1, 0,
		0, 0, 292, 290, 1, 0, 0, 0, 292, 293, 1, 0, 0, 0, 293, 76, 1, 0, 0, 0,
		294, 297, 3, 73, 36, 0, 295, 297, 5, 95, 0, 0, 296, 294, 1, 0, 0, 0, 296,
		295, 1, 0, 0, 0, 297, 78, 1, 0, 0, 0, 298, 300, 5, 48, 0, 0, 299, 301,
		3, 67, 33, 0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302, 1,
		0, 0, 0, 302, 303, 3, 81, 40, 0, 303, 80, 1, 0, 0, 0, 304, 309, 3, 83,
		41, 0, 305, 307, 3, 85, 42, 0, 306, 305, 1, 0, 0, 0, 306, 307, 1, 0, 0,
		0, 307, 308, 1, 0, 0, 0, 308, 310, 3, 83, 41, 0, 309, 306, 1, 0, 0, 0,
		309, 310, 1, 0, 0, 0, 310, 82, 1, 0, 0, 0, 311, 312, 7, 9, 0, 0, 312, 84,
		1, 0, 0, 0, 313, 315, 3, 87, 43, 0, 314, 313, 1, 0, 0, 0, 315, 316, 1,
		0, 0, 0, 316, 314, 1, 0, 0, 0, 316, 317, 1, 0, 0, 0, 317, 86, 1, 0, 0,
		0, 318, 321, 3, 83, 41, 0, 319, 321, 5, 95, 0, 0, 320, 318, 1, 0, 0, 0,
		320, 319, 1, 0, 0, 0, 321, 88, 1, 0, 0, 0, 322, 323, 5, 48, 0, 0, 323,
		324, 7, 10, 0, 0, 324, 325, 3, 91, 45, 0, 325, 90, 1, 0, 0, 0, 326, 331,
		3, 93, 46, 0, 327, 329, 3, 95, 47, 0, 328, 327, 1, 0, 0, 0, 328, 329, 1,
		0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 332, 3, 93, 46, 0, 331, 328, 1, 0,
		0, 0, 331, 332, 1, 0, 0, 0, 332, 92, 1, 0, 0, 0, 333, 334, 7, 11, 0, 0,
		334, 94, 1, 0, 0, 0, 335, 337, 3, 97, 48, 0, 336, 335, 1, 0, 0, 0, 337,
		338, 1, 0, 0, 0, 338, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 96, 1,
		0, 0, 0, 340, 343, 3, 93, 46, 0, 341, 343, 5, 95, 0, 0, 342, 340, 1, 0,
		0, 0, 342, 341, 1, 0, 0, 0, 343, 98, 1, 0, 0, 0, 344, 345, 3, 57, 28, 0,
		345, 347, 5, 46, 0, 0, 346, 348, 3, 57, 28, 0, 347, 346, 1, 0, 0, 0, 347,
		348, 1, 0, 0, 0, 348, 350, 1, 0, 0, 0, 349, 351, 3, 101, 50, 0, 350, 349,
		1, 0, 0, 0, 350, 351, 1, 0, 0, 0, 351, 353, 1, 0, 0, 0, 352, 354, 3, 109,
		54, 0, 353, 352, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 372, 1, 0, 0, 0,
		355, 356, 5, 46, 0, 0, 356, 358, 3, 57, 28, 0, 357, 359, 3, 101, 50, 0,
		358, 357, 1, 0, 0, 0, 358, 359, 1, 0, 0, 0, 359, 361, 1, 0, 0, 0, 360,
		362, 3, 109, 54, 0, 361, 360, 1, 0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 372,
		1, 0, 0, 0, 363, 364, 3, 57, 28, 0, 364, 366, 3, 101, 50, 0, 365, 367,
		3, 109, 54, 0, 366, 365, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0, 367, 372, 1,
		0, 0, 0, 368, 369, 3, 57, 28, 0, 369, 370, 3, 109, 54, 0, 370, 372, 1,
		0, 0, 0, 371, 344, 1, 0, 0, 0, 371, 355, 1, 0, 0, 0, 371, 363, 1, 0, 0,
		0, 371, 368, 1, 0, 0, 0, 372, 100, 1, 0, 0, 0, 373, 374, 3, 103, 51, 0,
		374, 375, 3, 105, 52, 0, 375, 102, 1, 0, 0, 0, 376, 377, 7, 12, 0, 0, 377,
		104, 1, 0, 0, 0, 378, 380, 3, 107, 53, 0, 379, 378, 1, 0, 0, 0, 379, 380,
		1, 0, 0, 0, 380, 381, 1, 0, 0, 0, 381, 382, 3, 57, 28, 0, 382, 106, 1,
		0, 0, 0, 383, 384, 7, 13, 0, 0, 384, 108, 1, 0, 0, 0, 385, 386, 7, 14,
		0, 0, 386, 110, 1, 0, 0, 0, 387, 388, 3, 113, 56, 0, 388, 390, 3, 115,
		57, 0, 389, 391, 3, 109, 54, 0, 390, 389, 1, 0, 0, 0, 390, 391, 1, 0, 0,
		0, 391, 112, 1, 0, 0, 0, 392, 394, 3, 69, 34, 0, 393, 395, 5, 46, 0, 0,
		394, 393, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 404, 1, 0, 0, 0, 396,
		397, 5, 48, 0, 0, 397, 399, 7, 7, 0, 0, 398, 400, 3, 71, 35, 0, 399, 398,
		1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 402, 5, 46,
		0, 0, 402, 404, 3, 71, 35, 0, 403, 392, 1, 0, 0, 0, 403, 396, 1, 0, 0,
		0, 404, 114, 1, 0, 0, 0, 405, 406, 3, 117, 58, 0, 406, 407, 3, 105, 52,
		0, 407, 116, 1, 0, 0, 0, 408, 409, 7, 15, 0, 0, 409, 118, 1, 0, 0, 0, 51,
		0, 150, 157, 162, 169, 182, 186, 191, 194, 199, 203, 208, 214, 221, 225,
		229, 233, 240, 245, 247, 251, 254, 258, 265, 269, 274, 282, 285, 292, 296,
		300, 306, 309, 316, 320, 328, 331, 338, 342, 347, 350, 353, 358, 361, 366,
		371, 379, 390, 394, 399, 403, 1, 0, 1, 0,
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

// ExpressionLexerInit initializes any static state used to implement ExpressionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewExpressionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressionLexerInit() {
	staticData := &expressionlexerLexerStaticData
	staticData.once.Do(expressionlexerLexerInit)
}

// NewExpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewExpressionLexer(input antlr.CharStream) *ExpressionLexer {
	ExpressionLexerInit()
	l := new(ExpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &expressionlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Expression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionLexer tokens.
const (
	ExpressionLexerT__0                 = 1
	ExpressionLexerT__1                 = 2
	ExpressionLexerT__2                 = 3
	ExpressionLexerT__3                 = 4
	ExpressionLexerT__4                 = 5
	ExpressionLexerDOT                  = 6
	ExpressionLexerLBRACK               = 7
	ExpressionLexerRBRACK               = 8
	ExpressionLexerLPAREN               = 9
	ExpressionLexerRPAREN               = 10
	ExpressionLexerEQUAL                = 11
	ExpressionLexerNOTEQUAL             = 12
	ExpressionLexerWhiteSpaces          = 13
	ExpressionLexerIdentifier           = 14
	ExpressionLexerIntegerLiteral       = 15
	ExpressionLexerFloatingPointLiteral = 16
	ExpressionLexerBooleanLiteral       = 17
	ExpressionLexerStringLiteral        = 18
)
