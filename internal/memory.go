package spokesbot

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

type Memory struct {
	answers    []Answer
	randomizer *rand.Rand
}

func NewMemory(filepath string) *Memory {
	var memory = Memory{
		randomizer: rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	answers, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(answers, &memory.answers)
	if err != nil {
		panic(err)
	}

	return &memory
}

func (memory *Memory) GetRandomAnswer() *Answer {
	return &memory.answers[memory.randomizer.Intn(len(memory.answers))]
}
