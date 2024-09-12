package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
	"root/utils"
	"sync"
)

func TaskOne() {
	urls := []string{
		"https://jsonplaceholder.typicode.com/todos?_limit=3",
		"https://jsonplaceholder.typicode.com/posts?_limit=2",
		"https://jsonplaceholder.typicode.com/comments?_limit=1",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(urls))

	result := make(chan []map[string]interface{})

	go func() {
		wg.Wait()
		close(result)
	}()

	for url, _ := range utils.Countdown(urls) {
		go func(url string) {
			defer wg.Done()

			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			var data []map[string]interface{}
			err = json.NewDecoder(res.Body).Decode(&data)
			if err != nil {
				fmt.Println(err)
				return
			}

			mu.Lock()
			result <- data
			mu.Unlock()
		}(url)
	}

	for data := range result {
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonData))
	}
}
