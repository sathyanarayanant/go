package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var numRequests uint64 = 0
var numRequestsMutex sync.Mutex

func incrementAndGet() uint64 {
	numRequestsMutex.Lock()
	defer numRequestsMutex.Unlock()
	numRequests = numRequests + 1
	return numRequests
}

func serve(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	requestNo := incrementAndGet()
	time.Sleep(5 * time.Second)

	t1 := time.Now()
	output := fmt.Sprintf("Hello world. Request no: %v, Completed in %v", requestNo, t1.Sub(t0))
	fmt.Fprintln(w, output)
}

func main() {
	http.HandleFunc("/", serve)
	http.ListenAndServe(":80", nil)
}
