package nlurules

import (
	"fmt"
	"sync"

	"speechly/nlu-rules-parser/pkg/parser"
)

const (
	TextEmptySpace = " "
)

type NluRuleListener struct {
	parser.BaseAnnotationGrammarListener
	utterance Utterance
	debug     bool
	output    chan Utterance
	lock      sync.Mutex
}

func NewNluRuleListener(bufSize uint64, debug bool) *NluRuleListener {
	return &NluRuleListener{
		output:    make(chan Utterance, bufSize),
		utterance: NewUtterance(),
		debug:     debug,
	}
}

func (l *NluRuleListener) ExitUtterance(c *parser.UtteranceContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.output <- l.utterance
	l.utterance = NewUtterance()
}

func (l *NluRuleListener) EnterText(c *parser.TextContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	node := Node{
		Text: c.GetText(),
	}

	if l.utterance.Entity != "" {
		node.Entity = l.utterance.Entity
	}

	if l.utterance.Intent != "" {
		node.Intent = l.utterance.Intent
		node.IntentBIO = l.utterance.IntentBIO

		if l.utterance.IntentBIO == IntentTagBeginning {
			l.utterance.IntentBIO = IntentTagInside
		}
	}

	l.utterance.Nodes = append(l.utterance.Nodes, node)
}

func (l *NluRuleListener) EnterIndent(c *parser.IndentContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.utterance.Nodes = append(
		l.utterance.Nodes,
		Node{
			Text: TextEmptySpace,
		},
	)
}

func (l *NluRuleListener) EnterEntity(c *parser.EntityContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	e := c.Entity_name()
	entityName, ok := e.(*parser.Entity_nameContext)

	if (!ok || entityName == nil) && l.debug {
		fmt.Printf("failed to parse entity name: %+v\n", e)
		return
	}

	l.utterance.Entity = entityName.WORD().GetText()
}

func (l *NluRuleListener) ExitEntity(c *parser.EntityContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.utterance.Entity = ""
}

func (l *NluRuleListener) EnterIntent_name(c *parser.Intent_nameContext) {
	l.utterance.Intent = c.WORD().GetText()
	l.utterance.IntentBIO = IntentTagBeginning
}

func (l *NluRuleListener) ExitReply(c *parser.ReplyContext) {
	l.utterance.Intent = ""
}

func (l *NluRuleListener) Utterances() <-chan Utterance {
	return l.output
}

func (l *NluRuleListener) Close() {
	close(l.output)
}
