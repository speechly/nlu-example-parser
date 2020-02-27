grammar AnnotationGrammar;

/*
 * Parser Rules
 */

annotation          : utterance+ EOF ;
utterance			: (indent? reply)+ END*;
reply				: intent_name (text | entity)+ ;
text				: (WORD | WHITESPACE)+ ;
indent				: INDENT ;
entity				: entity_value entity_name ;
entity_value		: OPEN_SB text CLOSE_SB ;
entity_name			: OPEN_PAREN WORD CLOSE_PAREN ;
intent_name			: INTENT_NAME_START WORD WHITESPACE* ;

/*
 * Lexer Rules
 */

INTENT_NAME_START	: '*' ;
OPEN_PAREN			: '(' ;
CLOSE_PAREN			: ')' ;
OPEN_SB				: '[' ;
CLOSE_SB			: ']' ;
WORD				: ~('['|']'|'('|')'|' '|'*'|'\n'|'\t')+ ;
INDENT				: END ('  ' | '\t' | '    ') WHITESPACE* ;
END             	: WHITESPACE* '\n'+ ;
WHITESPACE          : (' ' | '\t')+ ;
