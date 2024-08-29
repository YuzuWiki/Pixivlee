package types

import "fmt"

type IResponse[T any] interface {
	Assert() error
	Result() (*T, error)
}

type TPixivResponse[T interface{}] struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Body    T      `json:"body"`
}

func (r *TPixivResponse[T]) Assert() error {
	if r.Error || len(r.Message) > 0 {
		return fmt.Errorf(r.Message)
	}
	return nil
}

func (r *TPixivResponse[T]) Result() (*T, error) {
	if err := r.Assert(); err != nil {
		return nil, err
	}
	return &r.Body, nil
}

type TFanboxResponse[T interface{}] struct {
	Error string `json:"error"`
	Body  T      `json:"body"`
}

func (r *TFanboxResponse[T]) Assert() error {
	if len(r.Error) > 0 {
		return fmt.Errorf(r.Error)
	}
	return nil
}

func (r *TFanboxResponse[T]) Result() (*T, error) {
	if err := r.Assert(); err != nil {
		return nil, err
	}
	return &r.Body, nil
}
