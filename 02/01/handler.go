package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Sum int `json:"sum"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	// Tell the client that the API version is 1.3
	w.Header().Add("API-VERSION", "1.3")

	// Get the two parameters from the query
	param1 := r.URL.Query().Get("param1")
	param2 := r.URL.Query().Get("param2")

	// Convert the query parameters to integers
	intParam1, _ := strconv.Atoi(param1)
	intParam2, _ := strconv.Atoi(param2)

	// Calculate the sum of the parameters
	sum := intParam1 + intParam2

	// Create a response object with the sum
	response := Response{Sum: sum}

	// Encode the response object as JSON
	responseJSON, _ := json.Marshal(response)

	w.Write(responseJSON)
}
