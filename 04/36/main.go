package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echo(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, "You are visitor number %d", counter)
	mutex.Unlock()
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	target, _ := url.Parse("http://baidu.com")
	proxy := httputil.NewSingleHostReverseProxy(target)
	r.URL.Host = target.Host
	r.URL.Scheme = target.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = target.Host

	proxy.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/proxy", handleProxy)
	http.ListenAndServe(":8080", nil)
}
