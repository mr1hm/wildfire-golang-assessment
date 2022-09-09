package api

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type HandlerError struct {
	Err error
}

func (he *HandlerError) Error() string {
	return fmt.Sprintf("[ API ] Error: %v", he.Err)
}

type ResponseData struct {
	Data map[int]Responses
	Mtx  *sync.RWMutex
}

var (
	RespData ResponseData
)

func (rd *ResponseData) NewData() ResponseData {
	RespData = ResponseData{
		Data: make(map[int]Responses),
		Mtx:  new(sync.RWMutex),
	}
	return RespData
}

func (rd *ResponseData) Set(value interface{}, counter int) error {
	// Process based on type
	full_resp := RespData.Data[counter]
	switch v := value.(type) {
	case NameResponse:
		RespData.Mtx.Lock()
		full_resp.Name = NameResponse{
			FirstName: v.FirstName,
			LastName:  v.LastName,
		}
		RespData.Data[counter] = full_resp
		RespData.Mtx.Unlock()
	case JokeResponse:
		RespData.Mtx.Lock()
		full_resp.JokeData = JokeResponse{
			Type:  v.Type,
			Value: v.Value,
		}
		RespData.Data[counter] = full_resp
		RespData.Mtx.Unlock()
	default:
		return &HandlerError{
			Err: errors.New(fmt.Sprintf("Set() - Unknown type received: %v", reflect.TypeOf(v))),
		}
	}

	return nil
}
