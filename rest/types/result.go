package types

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Result[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error *Error `json:"error,omitempty"`
}
