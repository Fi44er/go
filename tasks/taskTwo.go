package tasks

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

type Todo struct {
	PostId int    `json:"postId"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func TaskTwo() {
	todos := make(chan Todo)
	pause := make(chan struct{})
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)

	var result []Todo

	go func(countUrls int) {
		defer wg.Done()
		for i := 0; i < countUrls; i++ {
			todos <- utilFunc()
			pause <- struct{}{}
			<-pause
		}
		close(todos)
	}(2)

	go func() {
		defer wg.Done()
		for data := range todos {
			<-pause
			mu.Lock()
			binarySearchAndInsert(&result, data)
			mu.Unlock()
			pause <- struct{}{}
		}
	}()

	wg.Wait()
	close(pause)
	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}

func utilFunc() Todo {
	randomId := rand.Intn(500)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?id=%d", randomId)
	// fmt.Print(url)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	var data []Todo
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	return data[0]
}

func binarySearchAndInsert(arr *[]Todo, target Todo) {
	low, high := 0, len(*arr)-1
	for low <= high {
		mid := (low + (high-low)/2)
		midValue := (*arr)[mid]
		if midValue.Id == target.Id {
			return
		} else if target.Id < midValue.Id {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	*arr = append((*arr)[:low], append([]Todo{target}, (*arr)[low:]...)...)
}
