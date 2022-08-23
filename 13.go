package main

import (
	"encoding/json"
	"fmt"
)

type response struct {
	PageNumber int      `json:"page"`
	Fruits     []string `json:"fruits"`
}

func main() {
	decode()
	encode()
}

func decode() {
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))
}

func encode() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.PageNumber)
}
