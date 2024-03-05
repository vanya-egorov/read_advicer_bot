package main

import (
	"fmt"
	"log"
)

func main() {
	t:= mustToken()
	//tgClient = telegram.New(tocen)

	// fetcher = fetcher.New()

	// processor = processor.New()

	//consumer.Start(fetcher, processor)
}

func mustToken()string{
	token := flag.String(
		name: "token-bot-token", 
		value: "",
		usage: "token for access to telegram bot",
	)

	flag.Parse()

	if *token == ""{
		log.Fatal(v...:"token is specified")
	}

	return *token
} 