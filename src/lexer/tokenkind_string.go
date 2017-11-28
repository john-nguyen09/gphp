// Code generated by "stringer -type=TokenKind"; DO NOT EDIT.

package lexer

import "strconv"

const _TokenKind_name = "UnknownEndOfFileTokenNameVariableNameSkippedTokenMissingTokenQualifiedNameAbstractKeywordAndKeywordArrayKeywordAsKeywordBreakKeywordCallableKeywordCaseKeywordCatchKeywordClassKeywordCloneKeywordConstKeywordContinueKeywordDeclareKeywordDefaultKeywordDieKeywordDoKeywordEchoKeywordElseKeywordElseIfKeywordEmptyKeywordEndDeclareKeywordEndForKeywordEndForEachKeywordEndIfKeywordEndSwitchKeywordEndWhileKeywordEvalKeywordExitKeywordExtendsKeywordFinalKeywordFinallyKeywordForKeywordForeachKeywordFunctionKeywordGlobalKeywordGotoKeywordIfKeywordImplementsKeywordIncludeKeywordIncludeOnceKeywordInstanceOfKeywordInsteadOfKeywordInterfaceKeywordIsSetKeywordListKeywordNamespaceKeywordNewKeywordOrKeywordPrintKeywordPrivateKeywordProtectedKeywordPublicKeywordRequireKeywordRequireOnceKeywordReturnKeywordStaticKeywordSwitchKeywordThrowKeywordTraitKeywordTryKeywordUnsetKeywordUseKeywordVarKeywordWhileKeywordXorKeywordYieldKeywordYieldFromKeywordOpenBracketTokenCloseBracketTokenOpenParenTokenCloseParenTokenOpenBraceTokenCloseBraceTokenDotTokenArrowTokenPlusPlusTokenMinusMinusTokenAsteriskAsteriskTokenAsteriskTokenPlusTokenMinusTokenTildeTokenExclamationTokenDollarTokenSlashTokenPercentTokenLessThanLessThanTokenGreaterThanGreaterThanTokenLessThanTokenGreaterThanTokenLessThanEqualsTokenGreaterThanEqualsTokenEqualsEqualsTokenEqualsEqualsEqualsTokenExclamationEqualsTokenExclamationEqualsEqualsTokenCaretTokenBarTokenAmpersandTokenAmpersandAmpersandTokenBarBarTokenColonTokenSemicolonTokenEqualsTokenAsteriskAsteriskEqualsTokenAsteriskEqualsTokenSlashEqualsTokenPercentEqualsTokenPlusEqualsTokenMinusEqualsTokenDotEqualsTokenLessThanLessThanEqualsTokenGreaterThanGreaterThanEqualsTokenAmpersandEqualsTokenCaretEqualsTokenBarEqualsTokenCommaTokenQuestionQuestionTokenLessThanEqualsGreaterThanTokenDotDotDotTokenBackslashTokenColonColonTokenDoubleArrowTokenLessThanGreaterThanTokenAtSymbolTokenBacktickTokenQuestionTokenIntegerLiteralTokenOctalLiteralTokenHexadecimalLiteralTokenBinaryLiteralTokenFloatingLiteralTokenInvalidOctalLiteralTokenInvalidHexadecimalLiteralInvalidBinaryLiteralStringLiteralTokenIntReservedWordFloatReservedWordTrueReservedWordStringReservedWordBoolReservedWordNullReservedWordScriptSectionStartTagScriptSectionEndTagScriptSectionPrependedTextVoidReservedWordFalseReservedWordMemberNameExpressionBinaryReservedWordBooleanReservedWordDoubleReservedWordIntegerReservedWordObjectReservedWordRealReservedWordReturnTypeInlineHtml"

var _TokenKind_index = [...]uint16{0, 7, 21, 25, 37, 49, 61, 74, 89, 99, 111, 120, 132, 147, 158, 170, 182, 194, 206, 221, 235, 249, 259, 268, 279, 290, 303, 315, 332, 345, 362, 374, 390, 405, 416, 427, 441, 453, 467, 477, 491, 506, 519, 530, 539, 556, 570, 588, 605, 621, 637, 649, 660, 676, 686, 695, 707, 721, 737, 750, 764, 782, 795, 808, 821, 833, 845, 855, 867, 877, 887, 899, 909, 921, 937, 953, 970, 984, 999, 1013, 1028, 1036, 1046, 1059, 1074, 1095, 1108, 1117, 1127, 1137, 1153, 1164, 1174, 1186, 1207, 1234, 1247, 1263, 1282, 1304, 1321, 1344, 1366, 1394, 1404, 1412, 1426, 1449, 1460, 1470, 1484, 1495, 1522, 1541, 1557, 1575, 1590, 1606, 1620, 1647, 1680, 1700, 1716, 1730, 1740, 1761, 1791, 1805, 1819, 1834, 1850, 1874, 1887, 1900, 1913, 1932, 1949, 1972, 1990, 2010, 2034, 2059, 2079, 2097, 2112, 2129, 2145, 2163, 2179, 2195, 2216, 2235, 2261, 2277, 2294, 2304, 2314, 2332, 2351, 2369, 2388, 2406, 2422, 2432, 2442}

func (i TokenKind) String() string {
	if i < 0 || i >= TokenKind(len(_TokenKind_index)-1) {
		return "TokenKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenKind_name[_TokenKind_index[i]:_TokenKind_index[i+1]]
}
