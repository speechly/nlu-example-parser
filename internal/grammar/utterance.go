package grammar

// TODO: make these custom types
const (
	IntentTagBeginning = "b"
	IntentTagInside    = "i"
	IntentTagOutside   = "o"
)

type Node struct {
	Text      string `json:"text"`
	Entity    string `json:"entity"`
	Intent    string `json:"intent"`
	IntentBIO string `json:"intent_bio"`
}

type Utterance struct {
	Nodes     []Node `json:"parsed"`
	Entity    string `json:"-"`
	Intent    string `json:"-"`
	IntentBIO string `json:"-"`
}

func NewUtterance() Utterance {
	return Utterance{
		IntentBIO: IntentTagOutside,
		Nodes:     make([]Node, 0, 1), // TODO: figure out how much we want to pre-allocate here
	}
}
