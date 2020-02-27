// parser/nlu_rule_listener.go
package parser


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
	BaseAnnotationGrammarListener
	ParsedRules []Utterance
	utterance Utterance

}

func (l *NluRuleListener) ExitUtterance(c *UtteranceContext) {

	l.ParsedRules = append(l.ParsedRules, l.utterance)
	l.utterance = NewUtterance()
}

func (l *NluRuleListener) EnterText(c *TextContext) {
	text := c.GetText()

	new_text := make(map[string]string)
	new_text["text"] = text

	if l.utterance.current_entity != ""{
		new_text["entity"] = l.utterance.current_entity
	}
	if l.utterance.current_intent != ""{
		new_text["intent"] = l.utterance.current_intent
	}

	l.utterance.Parsed = append(l.utterance.Parsed, new_text)
}

func (l *NluRuleListener) EnterIndent(c *IndentContext) {
	new_text := make(map[string]string)
	new_text["text"] = " "
	l.utterance.Parsed = append(l.utterance.Parsed, new_text)
}

func (l *NluRuleListener) EnterEntity(c *EntityContext) {
	entity_name, ok := c.Entity_name().(*Entity_nameContext)
		if !ok || entity_name == nil {
			panic("failed to cast type") // TODO: graceful error handling
		}
	l.utterance.current_entity = entity_name.WORD().GetText()
}

func (l *NluRuleListener) ExitEntity(c *EntityContext) {
	l.utterance.current_entity = ""
}

func (l *NluRuleListener) EnterIntent_name(c *Intent_nameContext) {
	l.utterance.current_intent = c.WORD().GetText()
}

func (l *NluRuleListener) ExitReply(c *ReplyContext) {
	l.utterance.current_intent = ""
}

func NewNluRuleListener() *NluRuleListener {
	return new(NluRuleListener)
}
