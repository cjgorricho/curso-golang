package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

//var mu sync.Mutex

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

			var id strings.Builder
			id.WriteString("ID-Hash")

			for {
				id.WriteString("0")
				if id.Len()+len(string(strconv.Itoa(ind))) == 10 {
					id.WriteString(strconv.Itoa(ind))
					break
				}
			}

			//fmt.Println(id)

			//mu.Lock()
			err := client.HMSet(id.String(), map[string]interface{}{"Name": field.Name, "Age": field.Age}).Err()
			//mu.Unlock()
			if err != nil {
				fmt.Println("Error escribiendo a la BD: ", err)
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
