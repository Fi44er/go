package tasks

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"root/typesstruct"
	"root/utils"
)

var EMOTIONS = [...]string{
	// Happy Emoticons
	"(✿◕‿◕✿)",
	"(❤‿❤)",
	"(✌◕◕✌)",
	"(‿◕✿◕‿)",
	"(◕‿◕✿)",
	// Sad Emoticons
	"(ಥ﹏ಥ)",
	"(╥﹏╥)",
	"(╭﹏╮)",
	"(‿﹏‿)",
	"(︶︹︺)",
	// Angry Emoticons
	"(凸ಠ益ಠ)凸",
	"(‡▼益▼)",
	"(╤﹏╤)",
	"(‿▀‿)",
	"(ಠ益ಠ)",
	// Surprised Emoticons
	"(‿⊙‿)",
	"(◕‿◕)",
	"(✿◕‿◕✿)",
	"(‿◕✿◕‿)",
	"(‿▃‿)",
}

type TodoWithEmotion struct {
	typesstruct.Todo
	Emotion string `json:"emotion"`
}

func TaskThre() {
	utilFunc2()
}

func utilFunc2() {
	url := "https://jsonplaceholder.typicode.com/comments?_limit=3"
	data := utils.GetMock[typesstruct.Todo](url)
	var todoWithEmotion []TodoWithEmotion

	for _, object := range data {
		emotion := EMOTIONS[rand.Intn(len(EMOTIONS))]
		test := TodoWithEmotion{Todo: object, Emotion: emotion}
		todoWithEmotion = append(todoWithEmotion, test)
	}

	jsonData, err := json.MarshalIndent(todoWithEmotion, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}
