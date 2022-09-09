package api

import (
	"encoding/json"
	"errors"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Service string
	URL     string
}
type Response struct {
	Service string
	Body    interface{}
}

var (
	URLs = []Request{
		{
			Service: "name",
			URL:     "https://names.mcquay.me/api/v0/",
		},
		{
			Service: "joke",
			URL:     "http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=nerdy",
		},
	}
)

type APIError struct {
	Err error
}

func (ae *APIError) Error() string {
	return fmt.Sprintf("[ API ] Error: %v", ae.Err)
}

func GetURL(job Request, results chan<- Response) error {
	// Send request and populate results channel with responses
	var api_err error

	resp, err := http.Get(job.URL)
	if err != nil {
		log.Println(err.Error())
		api_err = &APIError{
			Err: errors.New(fmt.Sprintf("%v", err)),
		}
	}
	defer resp.Body.Close()

	switch job.Service {
	case "name":
		var resp_payload NameResponse
		if err := json.NewDecoder(resp.Body).Decode(&resp_payload); err != nil {
			// b, io_err := ioutil.ReadAll(resp.Body)
			// if io_err != nil {
			// 	log.Println(io_err.Error())
			// }
			// if e, ok := err.(*json.SyntaxError); ok {
			// 	log.Printf("syntax error at byte offset %d", e.Offset)
			// }
			// log.Printf("resp.Body: %q", b)

			log.Println("name decode error", err.Error())
			api_err = &APIError{
				Err: errors.New(fmt.Sprintf("Decode error: %v", err)),
			}
		}
		results <- Response{
			Service: "name",
			Body: NameResponse{
				FirstName: resp_payload.FirstName,
				LastName:  resp_payload.LastName,
			},
		}
	case "joke":
		var resp_payload JokeResponse
		if err := json.NewDecoder(resp.Body).Decode(&resp_payload); err != nil {
			log.Println(err.Error())
			api_err = &APIError{
				Err: errors.New(fmt.Sprintf("Decode error: %v", err)),
			}
		}
		results <- Response{
			Service: "joke",
			Body: JokeResponse{
				Type:  resp_payload.Type,
				Value: resp_payload.Value,
			},
		}
	}

	if api_err != nil {
		return api_err
	}

	return nil
}
