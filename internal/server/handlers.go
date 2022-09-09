package server

import (
	"encoding/json"
	"github.com/mr1hm/wildfire-golang-assessment/internal/api"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	counter = 1
)

func getMessage(w http.ResponseWriter, r *http.Request) {
	var req_err error

	// Since we already know the number of requests that need to be sent, let's use buffered channels
	len_urls := len(api.URLs)
	jobs := make(chan api.Request, len_urls)
	results := make(chan api.Response, len_urls)

	for _, req := range api.URLs {
		jobs <- req
	}
	// Close jobs channel to prevent blocking when full
	close(jobs)

	var wg sync.WaitGroup
	for job := range jobs {
		// Launch goroutines for each url
		wg.Add(1)
		// Send requests concurrently
		go func(j api.Request) {
			err := api.GetURL(j, results)
			if err != nil {
				req_err = err
			}
			wg.Done()
		}(job)
		wg.Wait()

		// If any requests errors have occurred, it would've been assigned to `req_err`
		if req_err != nil {
			log.Println("Request error:", req_err.Error())
		}
	}

	// Consume values from results channel
	wg.Add(1)
	go func() {
		for r := 1; r <= len_urls; r++ {
			res := <-results

			switch v := res.Body.(type) {
			case api.NameResponse:
				if set_err := api.RespData.Set(v, counter); set_err != nil {
					log.Println(set_err.Error())
				}
			case api.JokeResponse:
				if set_err := api.RespData.Set(v, counter); set_err != nil {
					log.Println(set_err.Error())
				}
			}
		}
		wg.Done()
	}()
	wg.Wait()

	log.Println("Random Full Name:", api.RespData.Data[counter].Name.FirstName+" "+api.RespData.Data[counter].Name.LastName)

	// Create our new joke by replacing "John Doe" with the random name we requested for
	full_msg_sl := []string{}
	joke_sl := strings.Split(api.RespData.Data[counter].JokeData.Value.Joke, " ")
	for _, word := range joke_sl {
		if strings.Contains(word, "John") {
			full_msg_sl = append(full_msg_sl, api.RespData.Data[counter].Name.FirstName)
		} else if strings.Contains(word, "Doe") {
			full_msg_sl = append(full_msg_sl, api.RespData.Data[counter].Name.LastName)
		} else if strings.Contains(word, "JohnDoe") {
			split_word := strings.Split(word, "")
			if split_word[2] != "" {
				full_msg_sl = append(full_msg_sl, api.RespData.Data[counter].Name.FirstName+api.RespData.Data[counter].Name.LastName+split_word[2])
			}
			full_msg_sl = append(full_msg_sl, api.RespData.Data[counter].Name.FirstName+api.RespData.Data[counter].Name.LastName)
		} else {
			full_msg_sl = append(full_msg_sl, word)
		}
	}

	// Once our string manipulation is done, we send our encoded response
	full_msg := strings.Join(full_msg_sl, " ")
	log.Println("Response:", full_msg)

	if err := json.NewEncoder(w).Encode(full_msg); err != nil {
		log.Println(err.Error())
	}

	// `counter` is used to keep track of map keys in `api.RespData.Data`
	counter++
	log.Printf("counter is now: %d", counter)

	// Feel free to uncomment this line if you want to see the ALL data received after sending this app multiple requests
	// log.Printf("api.RespData.Data: %+v", api.RespData.Data)
}
