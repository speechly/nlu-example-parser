// Code generated from internal/grammar/AnnotationGrammar.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // AnnotationGrammar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// AnnotationGrammarListener is a complete listener for a parse tree produced by AnnotationGrammarParser.
type AnnotationGrammarListener interface {
	antlr.ParseTreeListener

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterUtterance is called when entering the utterance production.
	EnterUtterance(c *UtteranceContext)

	// EnterReply is called when entering the reply production.
	EnterReply(c *ReplyContext)

	// EnterText is called when entering the text production.
	EnterText(c *TextContext)

	// EnterIndent is called when entering the indent production.
	EnterIndent(c *IndentContext)

	// EnterEntity is called when entering the entity production.
	EnterEntity(c *EntityContext)

	// EnterEntity_value is called when entering the entity_value production.
	EnterEntity_value(c *Entity_valueContext)

	// EnterEntity_name is called when entering the entity_name production.
	EnterEntity_name(c *Entity_nameContext)

	// EnterIntent_name is called when entering the intent_name production.
	EnterIntent_name(c *Intent_nameContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitUtterance is called when exiting the utterance production.
	ExitUtterance(c *UtteranceContext)

	// ExitReply is called when exiting the reply production.
	ExitReply(c *ReplyContext)

	// ExitText is called when exiting the text production.
	ExitText(c *TextContext)

	// ExitIndent is called when exiting the indent production.
	ExitIndent(c *IndentContext)

	// ExitEntity is called when exiting the entity production.
	ExitEntity(c *EntityContext)

	// ExitEntity_value is called when exiting the entity_value production.
	ExitEntity_value(c *Entity_valueContext)

	// ExitEntity_name is called when exiting the entity_name production.
	ExitEntity_name(c *Entity_nameContext)

	// ExitIntent_name is called when exiting the intent_name production.
	ExitIntent_name(c *Intent_nameContext)
}
