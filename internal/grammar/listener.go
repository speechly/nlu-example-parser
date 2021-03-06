package grammar

import (
	"sync"
)

type NluRuleListener struct {
	BaseAnnotationGrammarListener
	utterance Utterance
	output    chan Utterance
	lock      sync.Mutex
}

func NewNluRuleListener(bufSize uint16) *NluRuleListener {
	return &NluRuleListener{
		output:    make(chan Utterance, bufSize),
		utterance: NewUtterance(),
	}
}

func (l *NluRuleListener) ExitUtterance(c *UtteranceContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.output <- l.utterance
	l.utterance = NewUtterance()
}

func (l *NluRuleListener) EnterText(c *TextContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if c == nil {
		return
	}

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

func (l *NluRuleListener) EnterEntity(c *EntityContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if c == nil {
		return
	}

	e := c.Entity_name()
	entityName, ok := e.(*Entity_nameContext)
	if !ok || entityName == nil {
		return
	}

	l.utterance.Entity = entityName.WORD().GetText()
}

func (l *NluRuleListener) ExitEntity(c *EntityContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.utterance.Entity = ""
}

func (l *NluRuleListener) EnterIntent_name(c *Intent_nameContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if c == nil {
		return
	}

	l.utterance.Intent = c.WORD().GetText()
	l.utterance.IntentBIO = IntentTagBeginning
}

func (l *NluRuleListener) ExitReply(c *ReplyContext) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.utterance.Intent = ""
}

func (l *NluRuleListener) Utterances() <-chan Utterance {
	return l.output
}

func (l *NluRuleListener) Close() {
	close(l.output)
}
