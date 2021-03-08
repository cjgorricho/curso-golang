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
	Name string `json:"name"`
	Age  int    `json:"age"`
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

	for j := 0; j < 10; j++ {
		for i := 0; i < 5000; i++ {
			wg.Add(1)
			go func(ind int) {
				rand.Seed(time.Now().UnixNano())
				index := rand.Intn(len(names))
				nm := names[index]

				rand.Seed(time.Now().UnixNano())
				time.Sleep(5 * time.Millisecond)
				ag := rand.Intn(50)

				field := Author{
					Name: nm,
					Age:  ag,
				}

				json, err := json.Marshal(field)
				if err != nil {
					fmt.Println(err)
				}

				id := "ID-List"
				id += strconv.Itoa(ind)

				//fmt.Println(id)

				err = client.LPush(id, json).Err()
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}(j)

		}
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
