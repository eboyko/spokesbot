package main

import (
	"os"

	spokesbot "github.com/eboyko/spokesbot/internal"
)

func main() {
	head := spokesbot.NewHead(
		os.Getenv("TOKEN"),
		spokesbot.NewMemory("configs/memory.json"),
		spokesbot.NewBehaviour("configs/behaviour.json"),
	)

	head.Listen()
}
