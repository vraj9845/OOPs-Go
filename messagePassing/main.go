package main

import (
	"OOPs-Go/messagePassing/messagePassingOne"
	"time"
)

func main() {
	ingredientsChannel := make(chan string)
	messagePassingOne.New(ingredientsChannel)
	ingredientsChannel <- "Focaccia bread"
	ingredientsChannel <- "Mozzarella cheese"
	ingredientsChannel <- "Tomato sauce"
	time.Sleep(3 * time.Second)
}