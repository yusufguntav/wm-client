package models

type APIResponse[T any] struct {
	Data   T   `json:"data"`
	Status int `json:"status"`
}
