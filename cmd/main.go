package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Response struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
	Header string `json:"h"`
}

func main() {
	r, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		fmt.Printf("Ошибка при получении цитаты!\n%s", err)
		os.Exit(0)
	}
	var quote []byte

	quote, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа АПИ\n%s", err)
		os.Exit(0)
	}

	var result []Response

	err = json.Unmarshal(quote, &result)
	if err != nil {
		fmt.Printf("Ошибка при переводе из json\n%s\n", err)
	}
	fmt.Printf("Цитата: %s\nАвтор: %s", result[0].Quote, result[0].Author)
}
