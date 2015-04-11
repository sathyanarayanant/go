package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var numRequests uint64 = 0
var numRequestsMutex sync.Mutex

var logHandle *log.Logger

func incrementAndGet() uint64 {
	numRequestsMutex.Lock()
	defer numRequestsMutex.Unlock()
	numRequests = numRequests + 1
	return numRequests
}

func serve(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	//requestNo := incrementAndGet()

	t1 := time.Now()
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		logMessage := fmt.Sprintf("Invalid n. err message: %v", err)
		fmt.Fprintln(w, logMessage)
		logHandle.Println(logMessage)
		return
	}

	output := getRandomString(n)
	logMessage := fmt.Sprintf("Received request. n: %v, Completed in %v", n, t1.Sub(t0))

	fmt.Fprintln(w, output)
	logHandle.Println(logMessage)
}

func getRandomString(n int) string {
	var buffer bytes.Buffer

	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(rand.Intn(10)))
	}

	return buffer.String()
}

func main() {
	//initialise logging
	file, err := os.OpenFile("http-test-server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("unable to open logging. exiting...")
		return
	}

	multi := io.MultiWriter(file, os.Stdout)

	logHandle = log.New(multi, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	logHandle.Println("http-test-server listening in 8080")

	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}
