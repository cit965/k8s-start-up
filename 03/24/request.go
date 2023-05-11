package request

import (
	"net/http"
	"time"
)

type Request struct {
	s        *Server
	err      error
	verb     string
	path     string
	body     interface{}
	selector labels.Selector
	timeout  time.Duration
	headers  map[string]string // Add this field to store headers
}

// Add this method to allow the user to add headers to the request
func (r *Request) Header(name string, value string) *Request {
	if r.err != nil {
		return r
	}
	if r.headers == nil {
		r.headers = make(map[string]string)
	}
	r.headers[name] = value
	return r
}

// Modify this method to include the headers when creating the request
func (r *Request) Do() (interface{}, error) {
	// The rest of the method is the same until creating the request...
	req, err := http.NewRequest(r.verb, finalUrl, body)
	if err != nil {
		return nil, err
	}
	// Add the headers to the request
	for name, value := range r.headers {
		req.Header.Add(name, value)
	}
	str, err := doRequest(req, r.s.auth)
	if err != nil {
		return nil, err
	}
	return api.Decode([]byte(str))
}
