package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type resp struct {
	Results []int
}

func RandStringBytes(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	var newWord = strings.ToLower(RandStringBytes(5))
	fmt.Println(newWord)
	var tempWord = strings.Split(newWord, "")
	var count = make(map[string]int)
	for _, v := range tempWord {
		if count[v] == 0 {
			count[v] = 1
			continue
		}
		count[v] += 1
	}

	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		var wrd string = strings.ToLower(r.URL.Query().Get("word"))

		if wrd == strings.Join(tempWord, "") {
			var ret []int
			for i := 0; i < 5; i++ {
				ret = append(ret, 1)
			}
			var response, _ = json.Marshal(resp{Results: ret})
			w.Write(response)
			return
		}

		var retArr []int
		var currCount = make(map[string]int)

		for i, v := range strings.Split(wrd, "") {
			if string(v) == tempWord[i] {
				retArr = append(retArr, 1)
				currCount[string(v)] += 1
				continue
			}
			if count[v] != 0 && currCount[v] != count[v] {
				retArr = append(retArr, 2)
				currCount[v] += 1
				continue
			}
			retArr = append(retArr, 0)
		}
		var response, _ = json.Marshal(resp{Results: retArr})
		w.Write(response)
	})

	port := ":8080"

	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
