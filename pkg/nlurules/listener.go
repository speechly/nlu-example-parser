package nlurules

import (
	"fmt"

	"speechly/nlu-rules-parser/pkg/parser"
)

type Utterance struct {
	Parsed           []map[string]string
	CurrentIntent    string
	CurrentEntity    string
	CurrentIntentBio string
}

const (
	// TODO: better names
	IntentBioO = "o"
	IntentBioB = "b"
	IntentBioI = "i"

	TextEmptySpace = " "
)

func NewUtterance() Utterance {
	return Utterance{
		CurrentIntentBio: IntentBioO,
	}
}

type NluRuleListener struct {
	parser.BaseAnnotationGrammarListener
	ParsedRules []Utterance
	utterance   Utterance
	debug       bool
	ch          chan Utterance
}

func NewNluRuleListener(bufSize uint64, debug bool) *NluRuleListener {
	return &NluRuleListener{
		ch: make(chan Utterance, bufSize),
	}
}

func (l *NluRuleListener) ExitUtterance(c *parser.UtteranceContext) {
	// TODO: make listener a per-utterance setup?
	l.ch <- l.utterance

	// TODO: this will eat a lot of memory
	l.ParsedRules = append(l.ParsedRules, l.utterance)

	l.utterance = NewUtterance()
}

func (l *NluRuleListener) EnterText(c *parser.TextContext) {
	newText := map[string]string{
		"text": c.GetText(),
	}

	if l.utterance.CurrentEntity != "" {
		newText["entity"] = l.utterance.CurrentEntity
	}

	if l.utterance.CurrentIntent != "" {
		newText["intent"] = l.utterance.CurrentIntent
		newText["intent_bio"] = l.utterance.CurrentIntentBio

		if l.utterance.CurrentIntentBio == IntentBioB {
			l.utterance.CurrentIntentBio = IntentBioI
		}
	}

	l.utterance.Parsed = append(l.utterance.Parsed, newText)
}

func (l *NluRuleListener) EnterIndent(c *parser.IndentContext) {
	l.utterance.Parsed = append(
		l.utterance.Parsed,
		map[string]string{
			"text": TextEmptySpace,
		},
	)
}

func (l *NluRuleListener) EnterEntity(c *parser.EntityContext) {
	e := c.Entity_name()
	entityName, ok := e.(*parser.Entity_nameContext)
	if (!ok || entityName == nil) && l.debug {
		fmt.Printf("failed to parse entity name: %+v\n", e)
	}

	l.utterance.CurrentEntity = entityName.WORD().GetText()
}

func (l *NluRuleListener) ExitEntity(c *parser.EntityContext) {
	l.utterance.CurrentEntity = ""
}

func (l *NluRuleListener) EnterIntent_name(c *parser.Intent_nameContext) {
	l.utterance.CurrentIntent = c.WORD().GetText()
	l.utterance.CurrentIntentBio = IntentBioB
}

func (l *NluRuleListener) ExitReply(c *parser.ReplyContext) {
	l.utterance.CurrentIntent = ""
}

func (l *NluRuleListener) Utterances() <-chan Utterance {
	return l.ch
}

func (l *NluRuleListener) Close() {
	close(l.ch)
}
