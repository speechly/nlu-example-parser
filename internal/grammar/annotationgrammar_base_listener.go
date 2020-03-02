// Code generated from internal/grammar/AnnotationGrammar.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // AnnotationGrammar
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAnnotationGrammarListener is a complete listener for a parse tree produced by AnnotationGrammarParser.
type BaseAnnotationGrammarListener struct{}

var _ AnnotationGrammarListener = &BaseAnnotationGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAnnotationGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAnnotationGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAnnotationGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAnnotationGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseAnnotationGrammarListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseAnnotationGrammarListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterUtterance is called when production utterance is entered.
func (s *BaseAnnotationGrammarListener) EnterUtterance(ctx *UtteranceContext) {}

// ExitUtterance is called when production utterance is exited.
func (s *BaseAnnotationGrammarListener) ExitUtterance(ctx *UtteranceContext) {}

// EnterReply is called when production reply is entered.
func (s *BaseAnnotationGrammarListener) EnterReply(ctx *ReplyContext) {}

// ExitReply is called when production reply is exited.
func (s *BaseAnnotationGrammarListener) ExitReply(ctx *ReplyContext) {}

// EnterText is called when production text is entered.
func (s *BaseAnnotationGrammarListener) EnterText(ctx *TextContext) {}

// ExitText is called when production text is exited.
func (s *BaseAnnotationGrammarListener) ExitText(ctx *TextContext) {}

// EnterIndent is called when production indent is entered.
func (s *BaseAnnotationGrammarListener) EnterIndent(ctx *IndentContext) {}

// ExitIndent is called when production indent is exited.
func (s *BaseAnnotationGrammarListener) ExitIndent(ctx *IndentContext) {}

// EnterEntity is called when production entity is entered.
func (s *BaseAnnotationGrammarListener) EnterEntity(ctx *EntityContext) {}

// ExitEntity is called when production entity is exited.
func (s *BaseAnnotationGrammarListener) ExitEntity(ctx *EntityContext) {}

// EnterEntity_value is called when production entity_value is entered.
func (s *BaseAnnotationGrammarListener) EnterEntity_value(ctx *Entity_valueContext) {}

// ExitEntity_value is called when production entity_value is exited.
func (s *BaseAnnotationGrammarListener) ExitEntity_value(ctx *Entity_valueContext) {}

// EnterEntity_name is called when production entity_name is entered.
func (s *BaseAnnotationGrammarListener) EnterEntity_name(ctx *Entity_nameContext) {}

// ExitEntity_name is called when production entity_name is exited.
func (s *BaseAnnotationGrammarListener) ExitEntity_name(ctx *Entity_nameContext) {}

// EnterIntent_name is called when production intent_name is entered.
func (s *BaseAnnotationGrammarListener) EnterIntent_name(ctx *Intent_nameContext) {}

// ExitIntent_name is called when production intent_name is exited.
func (s *BaseAnnotationGrammarListener) ExitIntent_name(ctx *Intent_nameContext) {}
