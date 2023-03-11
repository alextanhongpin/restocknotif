package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alextanhongpin/restocknotif/rest/types"
)

func Success[T any](w http.ResponseWriter, data T, statusCode int) {
	err := json.NewEncoder(w).Encode(types.Result[T]{
		Data: data,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Failure(w http.ResponseWriter, failure error, statusCode int) {
	if failure == nil {
		panic("rest: Failure error not provided")
	}

	err := json.NewEncoder(w).Encode(types.Result[any]{
		Error: &types.Error{
			// TODO: Replace with domain error.
			Code:    fmt.Sprint(statusCode),
			Message: failure.Error(),
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ParseBody[T any](r *http.Request) (T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		return t, err
	}

	return t, nil
}
