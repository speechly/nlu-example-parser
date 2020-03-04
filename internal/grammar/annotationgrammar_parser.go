// Code generated from internal/grammar/AnnotationGrammar.g4 by ANTLR 4.8. DO NOT EDIT.

package grammar // AnnotationGrammar
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 74, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 6, 2, 20, 10, 2, 13, 2, 14, 2, 21, 3, 2, 3,
	2, 3, 3, 6, 3, 27, 10, 3, 13, 3, 14, 3, 28, 3, 3, 7, 3, 32, 10, 3, 12,
	3, 14, 3, 35, 11, 3, 3, 4, 7, 4, 38, 10, 4, 12, 4, 14, 4, 41, 11, 4, 3,
	4, 3, 4, 3, 4, 6, 4, 46, 10, 4, 13, 4, 14, 4, 47, 3, 5, 6, 5, 51, 10, 5,
	13, 5, 14, 5, 52, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 69, 10, 9, 12, 9, 14, 9, 72, 11, 9,
	3, 9, 2, 2, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 3, 4, 2, 8, 8, 10, 10, 2,
	73, 2, 19, 3, 2, 2, 2, 4, 26, 3, 2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 50, 3,
	2, 2, 2, 10, 54, 3, 2, 2, 2, 12, 57, 3, 2, 2, 2, 14, 61, 3, 2, 2, 2, 16,
	65, 3, 2, 2, 2, 18, 20, 5, 4, 3, 2, 19, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2,
	2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 24,
	7, 2, 2, 3, 24, 3, 3, 2, 2, 2, 25, 27, 5, 6, 4, 2, 26, 25, 3, 2, 2, 2,
	27, 28, 3, 2, 2, 2, 28, 26, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 33, 3,
	2, 2, 2, 30, 32, 7, 9, 2, 2, 31, 30, 3, 2, 2, 2, 32, 35, 3, 2, 2, 2, 33,
	31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 33, 3, 2, 2,
	2, 36, 38, 7, 10, 2, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2,
	42, 45, 5, 16, 9, 2, 43, 46, 5, 8, 5, 2, 44, 46, 5, 10, 6, 2, 45, 43, 3,
	2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47,
	48, 3, 2, 2, 2, 48, 7, 3, 2, 2, 2, 49, 51, 9, 2, 2, 2, 50, 49, 3, 2, 2,
	2, 51, 52, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 9, 3,
	2, 2, 2, 54, 55, 5, 12, 7, 2, 55, 56, 5, 14, 8, 2, 56, 11, 3, 2, 2, 2,
	57, 58, 7, 6, 2, 2, 58, 59, 5, 8, 5, 2, 59, 60, 7, 7, 2, 2, 60, 13, 3,
	2, 2, 2, 61, 62, 7, 4, 2, 2, 62, 63, 7, 8, 2, 2, 63, 64, 7, 5, 2, 2, 64,
	15, 3, 2, 2, 2, 65, 66, 7, 3, 2, 2, 66, 70, 7, 8, 2, 2, 67, 69, 7, 10,
	2, 2, 68, 67, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71,
	3, 2, 2, 2, 71, 17, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 10, 21, 28, 33, 39,
	45, 47, 52, 70,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "'('", "')'", "'['", "']'",
}
var symbolicNames = []string{
	"", "INTENT_NAME_START", "OPEN_PAREN", "CLOSE_PAREN", "OPEN_SB", "CLOSE_SB",
	"WORD", "END", "WHITESPACE",
}

var ruleNames = []string{
	"annotation", "utterance", "reply", "text", "entity", "entity_value", "entity_name",
	"intent_name",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AnnotationGrammarParser struct {
	*antlr.BaseParser
}

func NewAnnotationGrammarParser(input antlr.TokenStream) *AnnotationGrammarParser {
	this := new(AnnotationGrammarParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AnnotationGrammar.g4"

	return this
}

// AnnotationGrammarParser tokens.
const (
	AnnotationGrammarParserEOF               = antlr.TokenEOF
	AnnotationGrammarParserINTENT_NAME_START = 1
	AnnotationGrammarParserOPEN_PAREN        = 2
	AnnotationGrammarParserCLOSE_PAREN       = 3
	AnnotationGrammarParserOPEN_SB           = 4
	AnnotationGrammarParserCLOSE_SB          = 5
	AnnotationGrammarParserWORD              = 6
	AnnotationGrammarParserEND               = 7
	AnnotationGrammarParserWHITESPACE        = 8
)

// AnnotationGrammarParser rules.
const (
	AnnotationGrammarParserRULE_annotation   = 0
	AnnotationGrammarParserRULE_utterance    = 1
	AnnotationGrammarParserRULE_reply        = 2
	AnnotationGrammarParserRULE_text         = 3
	AnnotationGrammarParserRULE_entity       = 4
	AnnotationGrammarParserRULE_entity_value = 5
	AnnotationGrammarParserRULE_entity_name  = 6
	AnnotationGrammarParserRULE_intent_name  = 7
)

// IAnnotationContext is an interface to support dynamic dispatch.
type IAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotationContext differentiates from other interfaces.
	IsAnnotationContext()
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *AnnotationContext) EOF() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserEOF, 0)
}

func (s *AnnotationContext) AllUtterance() []IUtteranceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUtteranceContext)(nil)).Elem())
	var tst = make([]IUtteranceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUtteranceContext)
		}
	}

	return tst
}

func (s *AnnotationContext) Utterance(i int) IUtteranceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUtteranceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUtteranceContext)
}

func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (p *AnnotationGrammarParser) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AnnotationGrammarParserRULE_annotation)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == AnnotationGrammarParserINTENT_NAME_START || _la == AnnotationGrammarParserWHITESPACE {
		{
			p.SetState(16)
			p.Utterance()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(21)
		p.Match(AnnotationGrammarParserEOF)
	}

	return localctx
}

// IUtteranceContext is an interface to support dynamic dispatch.
type IUtteranceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUtteranceContext differentiates from other interfaces.
	IsUtteranceContext()
}

type UtteranceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUtteranceContext() *UtteranceContext {
	var p = new(UtteranceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_utterance
	return p
}

func (*UtteranceContext) IsUtteranceContext() {}

func NewUtteranceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UtteranceContext {
	var p = new(UtteranceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_utterance

	return p
}

func (s *UtteranceContext) GetParser() antlr.Parser { return s.parser }

func (s *UtteranceContext) AllReply() []IReplyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReplyContext)(nil)).Elem())
	var tst = make([]IReplyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReplyContext)
		}
	}

	return tst
}

func (s *UtteranceContext) Reply(i int) IReplyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReplyContext)
}

func (s *UtteranceContext) AllEND() []antlr.TerminalNode {
	return s.GetTokens(AnnotationGrammarParserEND)
}

func (s *UtteranceContext) END(i int) antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserEND, i)
}

func (s *UtteranceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UtteranceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UtteranceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterUtterance(s)
	}
}

func (s *UtteranceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitUtterance(s)
	}
}

func (p *AnnotationGrammarParser) Utterance() (localctx IUtteranceContext) {
	localctx = NewUtteranceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AnnotationGrammarParserRULE_utterance)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(23)
				p.Reply()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnnotationGrammarParserEND {
		{
			p.SetState(28)
			p.Match(AnnotationGrammarParserEND)
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReplyContext is an interface to support dynamic dispatch.
type IReplyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplyContext differentiates from other interfaces.
	IsReplyContext()
}

type ReplyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplyContext() *ReplyContext {
	var p = new(ReplyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_reply
	return p
}

func (*ReplyContext) IsReplyContext() {}

func NewReplyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplyContext {
	var p = new(ReplyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_reply

	return p
}

func (s *ReplyContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplyContext) Intent_name() IIntent_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntent_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntent_nameContext)
}

func (s *ReplyContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(AnnotationGrammarParserWHITESPACE)
}

func (s *ReplyContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserWHITESPACE, i)
}

func (s *ReplyContext) AllText() []ITextContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITextContext)(nil)).Elem())
	var tst = make([]ITextContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITextContext)
		}
	}

	return tst
}

func (s *ReplyContext) Text(i int) ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *ReplyContext) AllEntity() []IEntityContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEntityContext)(nil)).Elem())
	var tst = make([]IEntityContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEntityContext)
		}
	}

	return tst
}

func (s *ReplyContext) Entity(i int) IEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEntityContext)
}

func (s *ReplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterReply(s)
	}
}

func (s *ReplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitReply(s)
	}
}

func (p *AnnotationGrammarParser) Reply() (localctx IReplyContext) {
	localctx = NewReplyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AnnotationGrammarParserRULE_reply)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AnnotationGrammarParserWHITESPACE {
		{
			p.SetState(34)
			p.Match(AnnotationGrammarParserWHITESPACE)
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Intent_name()
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(43)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case AnnotationGrammarParserWORD, AnnotationGrammarParserWHITESPACE:
				{
					p.SetState(41)
					p.Text()
				}

			case AnnotationGrammarParserOPEN_SB:
				{
					p.SetState(42)
					p.Entity()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ITextContext is an interface to support dynamic dispatch.
type ITextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextContext differentiates from other interfaces.
	IsTextContext()
}

type TextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextContext() *TextContext {
	var p = new(TextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_text
	return p
}

func (*TextContext) IsTextContext() {}

func NewTextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextContext {
	var p = new(TextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_text

	return p
}

func (s *TextContext) GetParser() antlr.Parser { return s.parser }

func (s *TextContext) AllWORD() []antlr.TerminalNode {
	return s.GetTokens(AnnotationGrammarParserWORD)
}

func (s *TextContext) WORD(i int) antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserWORD, i)
}

func (s *TextContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(AnnotationGrammarParserWHITESPACE)
}

func (s *TextContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserWHITESPACE, i)
}

func (s *TextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterText(s)
	}
}

func (s *TextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitText(s)
	}
}

func (p *AnnotationGrammarParser) Text() (localctx ITextContext) {
	localctx = NewTextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AnnotationGrammarParserRULE_text)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(47)
				_la = p.GetTokenStream().LA(1)

				if !(_la == AnnotationGrammarParserWORD || _la == AnnotationGrammarParserWHITESPACE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IEntityContext is an interface to support dynamic dispatch.
type IEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityContext differentiates from other interfaces.
	IsEntityContext()
}

type EntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityContext() *EntityContext {
	var p = new(EntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_entity
	return p
}

func (*EntityContext) IsEntityContext() {}

func NewEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityContext {
	var p = new(EntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_entity

	return p
}

func (s *EntityContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityContext) Entity_value() IEntity_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntity_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntity_valueContext)
}

func (s *EntityContext) Entity_name() IEntity_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntity_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntity_nameContext)
}

func (s *EntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterEntity(s)
	}
}

func (s *EntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitEntity(s)
	}
}

func (p *AnnotationGrammarParser) Entity() (localctx IEntityContext) {
	localctx = NewEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AnnotationGrammarParserRULE_entity)

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
		p.SetState(52)
		p.Entity_value()
	}
	{
		p.SetState(53)
		p.Entity_name()
	}

	return localctx
}

// IEntity_valueContext is an interface to support dynamic dispatch.
type IEntity_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntity_valueContext differentiates from other interfaces.
	IsEntity_valueContext()
}

type Entity_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntity_valueContext() *Entity_valueContext {
	var p = new(Entity_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_entity_value
	return p
}

func (*Entity_valueContext) IsEntity_valueContext() {}

func NewEntity_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Entity_valueContext {
	var p = new(Entity_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_entity_value

	return p
}

func (s *Entity_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Entity_valueContext) OPEN_SB() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserOPEN_SB, 0)
}

func (s *Entity_valueContext) Text() ITextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextContext)
}

func (s *Entity_valueContext) CLOSE_SB() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserCLOSE_SB, 0)
}

func (s *Entity_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Entity_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Entity_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterEntity_value(s)
	}
}

func (s *Entity_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitEntity_value(s)
	}
}

func (p *AnnotationGrammarParser) Entity_value() (localctx IEntity_valueContext) {
	localctx = NewEntity_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AnnotationGrammarParserRULE_entity_value)

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
		p.SetState(55)
		p.Match(AnnotationGrammarParserOPEN_SB)
	}
	{
		p.SetState(56)
		p.Text()
	}
	{
		p.SetState(57)
		p.Match(AnnotationGrammarParserCLOSE_SB)
	}

	return localctx
}

// IEntity_nameContext is an interface to support dynamic dispatch.
type IEntity_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntity_nameContext differentiates from other interfaces.
	IsEntity_nameContext()
}

type Entity_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntity_nameContext() *Entity_nameContext {
	var p = new(Entity_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_entity_name
	return p
}

func (*Entity_nameContext) IsEntity_nameContext() {}

func NewEntity_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Entity_nameContext {
	var p = new(Entity_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_entity_name

	return p
}

func (s *Entity_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Entity_nameContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserOPEN_PAREN, 0)
}

func (s *Entity_nameContext) WORD() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserWORD, 0)
}

func (s *Entity_nameContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserCLOSE_PAREN, 0)
}

func (s *Entity_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Entity_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Entity_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterEntity_name(s)
	}
}

func (s *Entity_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitEntity_name(s)
	}
}

func (p *AnnotationGrammarParser) Entity_name() (localctx IEntity_nameContext) {
	localctx = NewEntity_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AnnotationGrammarParserRULE_entity_name)

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
		p.SetState(59)
		p.Match(AnnotationGrammarParserOPEN_PAREN)
	}
	{
		p.SetState(60)
		p.Match(AnnotationGrammarParserWORD)
	}
	{
		p.SetState(61)
		p.Match(AnnotationGrammarParserCLOSE_PAREN)
	}

	return localctx
}

// IIntent_nameContext is an interface to support dynamic dispatch.
type IIntent_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntent_nameContext differentiates from other interfaces.
	IsIntent_nameContext()
}

type Intent_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntent_nameContext() *Intent_nameContext {
	var p = new(Intent_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AnnotationGrammarParserRULE_intent_name
	return p
}

func (*Intent_nameContext) IsIntent_nameContext() {}

func NewIntent_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Intent_nameContext {
	var p = new(Intent_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AnnotationGrammarParserRULE_intent_name

	return p
}

func (s *Intent_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Intent_nameContext) INTENT_NAME_START() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserINTENT_NAME_START, 0)
}

func (s *Intent_nameContext) WORD() antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserWORD, 0)
}

func (s *Intent_nameContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(AnnotationGrammarParserWHITESPACE)
}

func (s *Intent_nameContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(AnnotationGrammarParserWHITESPACE, i)
}

func (s *Intent_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Intent_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Intent_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.EnterIntent_name(s)
	}
}

func (s *Intent_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationGrammarListener); ok {
		listenerT.ExitIntent_name(s)
	}
}

func (p *AnnotationGrammarParser) Intent_name() (localctx IIntent_nameContext) {
	localctx = NewIntent_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AnnotationGrammarParserRULE_intent_name)

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(AnnotationGrammarParserINTENT_NAME_START)
	}
	{
		p.SetState(64)
		p.Match(AnnotationGrammarParserWORD)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(65)
				p.Match(AnnotationGrammarParserWHITESPACE)
			}

		}
		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}
