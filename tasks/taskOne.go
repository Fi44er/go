package tasks

import (
	"encoding/json"
	"fmt"
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
			data := utils.GetMock[map[string]interface{}](url)
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
