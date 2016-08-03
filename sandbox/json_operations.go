package sandbox

import (
	"fmt"
	"encoding/json"
)

// https://blog.golang.org/json-and-go
type Message struct {
	Name string
	Body string
	Time int64
}

func Encode() {
	m := Message{
		Name:"Alice",
		Body:"Hello",
		Time:1294706395881547000,
	}
	b, err := json.Marshal(m)
	if b != nil && err == nil {
		fmt.Printf("Encode OK: %s\n", string(b))
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Encode Error")
	}
}

func Decode() {
	// body := "{\"Name\":\"Alice\",\"Body\":\"Hello\",\"Time\":1294706395881547000}"
	b := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	var m Message
	err := json.Unmarshal(b, &m)
	if err == nil {
		c, err2 := json.Marshal(m)
		if c != nil && err2 == nil {
			fmt.Printf("Decode OK: %s\n", string(c))
		} else if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Decode Error")
		}
	} else {
		fmt.Println(err)
	}
}