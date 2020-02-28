package nlurules

import "speechly/nlu-rules-parser/pkg/parser"

type Utterance struct {
	Parsed []map[string]string

	current_intent string
	current_entity string
}

func NewUtterance() Utterance {

	var l Utterance
	l.current_intent = ""
	l.current_entity = ""
	return l

}

type NluRuleListener struct {
	parser.BaseAnnotationGrammarListener
	ParsedRules []Utterance
	utterance   Utterance
	ch          chan Utterance
}

func (l *NluRuleListener) ExitUtterance(c *parser.UtteranceContext) {

	l.ch <- l.utterance
	l.ParsedRules = append(l.ParsedRules, l.utterance)
	l.utterance = NewUtterance()
}

func (l *NluRuleListener) EnterText(c *parser.TextContext) {
	text := c.GetText()

	new_text := make(map[string]string)
	new_text["text"] = text

	if l.utterance.current_entity != "" {
		new_text["entity"] = l.utterance.current_entity
	}
	if l.utterance.current_intent != "" {
		new_text["intent"] = l.utterance.current_intent
	}

	l.utterance.Parsed = append(l.utterance.Parsed, new_text)
}

func (l *NluRuleListener) EnterIndent(c *parser.IndentContext) {
	new_text := make(map[string]string)
	new_text["text"] = " "
	l.utterance.Parsed = append(l.utterance.Parsed, new_text)
}

func (l *NluRuleListener) EnterEntity(c *parser.EntityContext) {
	entity_name, ok := c.Entity_name().(*parser.Entity_nameContext)
	if !ok || entity_name == nil {
		panic("failed to cast type") // TODO: graceful error handling
	}
	l.utterance.current_entity = entity_name.WORD().GetText()
}

func (l *NluRuleListener) ExitEntity(c *parser.EntityContext) {
	l.utterance.current_entity = ""
}

func (l *NluRuleListener) EnterIntent_name(c *parser.Intent_nameContext) {
	l.utterance.current_intent = c.WORD().GetText()
}

func (l *NluRuleListener) ExitReply(c *parser.ReplyContext) {
	l.utterance.current_intent = ""
}

func (l *NluRuleListener) Utterances() <-chan Utterance {
	return l.ch
}

func (l *NluRuleListener) Close() {
	close(l.ch)
}

func NewNluRuleListener() *NluRuleListener {
	return &NluRuleListener{
		ch: make(chan Utterance),
	}
}
