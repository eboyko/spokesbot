package spokesbot

type Conversation struct {
	id                        int64
	deferredReactionCountdown uint
}

func NewConversation(id int64, behaviour *Behaviour) *Conversation {
	return &Conversation{
		id:                        id,
		deferredReactionCountdown: behaviour.DeferredReactionDelay,
	}
}
