package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"encoding/json"

	"github.com/go-redis/redis"
)

type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	TimeStamp int64  `json:"time"`
}

var names = []string{
	"Hugo",
	"Paco",
	"Luis",
	"Carlos",
	"Alejandro",
	"Elliot",
}

var wg sync.WaitGroup

func main() {

	start := time.Now()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//fmt.Println(names)

	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func(ind int) {
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(names))
			nm := names[index]

			rand.Seed(time.Now().UnixNano())
			ag := rand.Intn(50)

			ts := time.Now().UnixNano()

			field := Author{
				Name:      nm,
				Age:       ag,
				TimeStamp: ts,
			}

			json, err := json.Marshal(field)
			if err != nil {
				fmt.Println(err)
			}

			id := "ID-"

			for {
				id += "0"
				if len(id)+len(string(strconv.Itoa(ind))) == 10 {
					id += strconv.Itoa(ind)
					break
				}
			}

			//fmt.Println(id)

			err = client.Set(id, json, 0).Err()
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(i)

	}
	wg.Wait()

	fmt.Printf("\nTiempo ejecución: %v ms\n", time.Since(start).Milliseconds())

	//val, err := client.Get("id1234").Result()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(val)

}
