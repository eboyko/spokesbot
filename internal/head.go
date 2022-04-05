package spokesbot

import (
	"regexp"
	"time"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Head struct {
	memory             *Memory
	behaviour          *Behaviour
	exchange           *telegram.BotAPI
	conversations      []*Conversation
	initializationTime time.Time
}

func NewHead(token string, memory *Memory, behaviour *Behaviour) *Head {
	exchange, err := telegram.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	return &Head{
		memory:             memory,
		behaviour:          behaviour,
		exchange:           exchange,
		initializationTime: time.Now(),
	}
}

func (head *Head) Listen() {
	updatePolicy := telegram.NewUpdate(0)
	updatePolicy.Timeout = 60

	for update := range head.exchange.GetUpdatesChan(updatePolicy) {
		go head.processUpdate(&update)
	}
}

func (head *Head) processUpdate(update *telegram.Update) {
	if head.isUpdateProcessable(update) {
		conversation := head.getConversation(update.Message.Chat.ID)

		if head.isReactionRequired(conversation, update.Message) {
			head.answer(conversation)
		}
	}
}

func (head *Head) getConversation(id int64) *Conversation {
	for index := range head.conversations {
		if head.conversations[index].id == id {
			return head.conversations[index]
		}
	}

	head.conversations = append(head.conversations, NewConversation(id, head.behaviour))
	return head.conversations[len(head.conversations)-1]
}

func (head *Head) isUpdateProcessable(update *telegram.Update) bool {
	return update.Message != nil && update.Message.Time().After(head.initializationTime)
}

func (head *Head) isReactionRequired(conversation *Conversation, message *telegram.Message) bool {
	return head.isMandatoryReactionRequired(message) || head.isDeferredReactionRequired(conversation, message)
}

func (head *Head) isMandatoryReactionRequired(message *telegram.Message) bool {
	for _, pattern := range head.behaviour.MandatoryReactionPatterns {
		match, _ := regexp.Match(pattern, []byte(message.Text))

		if match {
			return true
		}
	}

	return false
}

func (head *Head) isDeferredReactionRequired(conversation *Conversation, message *telegram.Message) bool {
	for _, pattern := range head.behaviour.DeferredReactionPatterns {
		match, _ := regexp.Match(pattern, []byte(message.Text))

		if match {
			conversation.deferredReactionCountdown--

			if conversation.deferredReactionCountdown == 0 {
				conversation.deferredReactionCountdown = head.behaviour.DeferredReactionDelay
				return true
			} else {
				break
			}
		}
	}

	return false
}

func (head *Head) answer(conversation *Conversation) {
	head.exchange.Send(telegram.NewMessage(conversation.id, head.memory.GetRandomAnswer().Text))
}
