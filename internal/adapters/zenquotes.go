package adapters

import (
	"log"
	"net/http"
	"time"
)

type ZenQuote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

type Quotes struct {
	Quotes []ZenQuote
}

func NewQuotes() *Quotes {
	return &Quotes{}
}

func (Quotes) Get() error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest("GET", "https://zenquotes.io/api/random", nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
		return err
	}

}
