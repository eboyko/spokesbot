package spokesbot

import (
	"encoding/json"
	"io/ioutil"
)

type Behaviour struct {
	MandatoryReactionPatterns []string `json:"mandatory_reaction_patterns"`
	DeferredReactionPatterns  []string `json:"deferred_reaction_patterns"`
	DeferredReactionDelay     uint     `json:"deferred_reaction_delay"`
}

func NewBehaviour(filepath string) *Behaviour {
	var behaviour Behaviour

	preferences, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(preferences, &behaviour)
	if err != nil {
		panic(err)
	}

	return &behaviour
}
