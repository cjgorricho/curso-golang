package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

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

	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func(ind int) {
			rand.Seed(time.Now().UnixNano())
			index := rand.Intn(len(names))
			nm := names[index]

			rand.Seed(time.Now().UnixNano())
			//time.Sleep(5 * time.Millisecond)
			ag := rand.Intn(50)

			field := Author{
				Name: nm,
				Age:  ag,
			}

			id := "ID-H"
			for {
				id += "0"
				if len(id)+len(string(strconv.Itoa(ind))) == 10 {
					id += strconv.Itoa(ind)
					break
				}
			}

			//fmt.Println(id)

			err := client.HMSet(id, map[string]interface{}{"Name": field.Name, "Age": field.Age}).Err()
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
