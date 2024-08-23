package apis

import "fmt"

type ResponseData[T interface{}] struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    T      `json:"body"`
}

func (r *ResponseData[T]) Assert() error {
	if r.Error || len(r.Message) > 0 {
		return fmt.Errorf(r.Message)
	}
	return nil
}

func (r *ResponseData[T]) Result() (*T, error) {
	if err := r.Assert(); err != nil {
		return nil, err
	}
	return &r.Body, nil
}
