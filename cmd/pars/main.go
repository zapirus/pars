package main

import (
	"flag"
	"fmt"
	"github.com/zapirus/pars/services"
	"log"
	"net/http"
	"net/http/cookiejar"
	"sync"
	"time"
)

func main() {

	n := flag.Int("n", 1, "")
	flag.Parse()

	count := *n
	fmt.Printf("Будет атака %d горутинами.\n", count)
	time.Sleep(1 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ct int) {
		defer wg.Done()
		for i := 0; i < ct; i++ {

			jar, err := cookiejar.New(nil)
			if err != nil {
				log.Fatalln(err)
			}

			client := &http.Client{
				Jar: jar,
			}
			services.Entry(client)
			var data = services.Question1get(client)

			var step uint = 1
			for {
				var ok bool
				if ok, data = services.QuestionPost(client, step, data); ok {
					break
				}
				step++
			}

			log.Printf("Для прохождения теста понадобилось %d шагов", step)
		}

	}(count)
	wg.Wait()
	fmt.Println("The end")
}
