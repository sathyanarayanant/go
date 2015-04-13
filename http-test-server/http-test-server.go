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

	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		logMessage := fmt.Sprintf("Invalid n. err message: %v", err)
		fmt.Fprintln(w, logMessage)
		logHandle.Println(logMessage)
		return
	}

	output := getRandomString(n)
	t1 := time.Now()
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
		fmt.Println("unable to open log file. exiting...")
		return
	}

	multi := io.MultiWriter(file, os.Stdout)
	logHandle = log.New(multi, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)

	if len(os.Args) < 2 {
		logHandle.Println("Incorrect usage. First command argument should be port. Exiting...")
		os.Exit(-1)
	}

	port := os.Args[1]

	if _, err := strconv.Atoi(port); err != nil {
		logHandle.Printf("Port %v is not an int. Exiting...", port)
		os.Exit(-1)
	}

	logHandle.Printf("http-test-server listening in %v", port)

	http.HandleFunc("/content", serve)
	http.ListenAndServe(":"+port, nil)
}
