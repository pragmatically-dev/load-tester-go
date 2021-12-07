package main

import (
	"fmt"
	"net/http"

	_ "go.uber.org/automaxprocs"
)

var (
	URL             = "http://localhost:3000/users/1"
	WORKERS         = 100000
	REQ_PER_WORKERS = 5000
)

func main() {
	var d int
	for workerID := 0; workerID < WORKERS; workerID++ {

		go makeRequest(workerID)
	}

	fmt.Scanf("waiting", &d)
}

func makeRequest(workerID int) {
	var client http.Client = *http.DefaultClient

	for i := 0; i < REQ_PER_WORKERS; i++ {

		res, err := client.Get(URL)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("WORKER-%d    - request n-%d  - status -> %d"+"\n", workerID, i, res.StatusCode)

		}

		//		res.Body.Close()
	}
}
