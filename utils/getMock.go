package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetMock[T any](url string) []T {
	var data []T
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return data
}
