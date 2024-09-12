package tasks

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"root/typesstruct"
	"root/utils"
	"sync"
)

func TaskTwo() {
	todos := make(chan typesstruct.Todo)
	pause := make(chan struct{})
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)

	var result []typesstruct.Todo

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

func utilFunc() typesstruct.Todo {
	randomId := rand.Intn(500)
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?id=%d", randomId)
	data := utils.GetMock[typesstruct.Todo](url)
	return data[0]
}

func binarySearchAndInsert(arr *[]typesstruct.Todo, target typesstruct.Todo) {
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
	*arr = append((*arr)[:low], append([]typesstruct.Todo{target}, (*arr)[low:]...)...)
}
