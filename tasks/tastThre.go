package tasks

import (
	"math/rand"
	"root/typesstruct"
	"root/utils"
)

var EMOTIONS = [...]string{
	// Happy Emoticons
	"(✿◕‿◕✿)",
	"(❤‿❤)",
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
	"(ಠ益ಠ)",
}

type TodoWithEmotion struct {
	typesstruct.Todo
	Emotion string `json:"emotion"`
}

func TaskThre() {
  data := make(chan []TodoWithEmotion)
  go func() {
    data <- utilFunc2()
    close(data)
  }()

  go func() {
    for res := range data {
      <- data
    }
  }()

}

func utilFunc2() []TodoWithEmotion{
	url := "https://jsonplaceholder.typicode.com/comments?_limit=3"
	data := utils.GetMock[typesstruct.Todo](url)
	var todoWithEmotion []TodoWithEmotion

	for _, object := range data {
		emotion := EMOTIONS[rand.Intn(len(EMOTIONS))]
		test := TodoWithEmotion{Todo: object, Emotion: emotion}
		todoWithEmotion = append(todoWithEmotion, test)
	}
  return todoWithEmotion
}

func sortData(data *[]TodoWithEmotion) {

}
