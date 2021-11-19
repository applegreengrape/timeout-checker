package main

import (
	"fmt"
	"net/http"
	"flag"
	"time"

	"github.com/gorilla/mux"
)

type Wait struct {
    time int
}

func (wait *Wait) healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recieved")
	time.Sleep(time.Duration(wait.time) * time.Second)

	fmt.Println("returned")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
}

func main() {
	wait := flag.Int("wait", 3, "")
	flag.Parse()

	w := &Wait{time: *wait}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", w.healthcheck).Methods("GET")
	http.ListenAndServe(":8080", router)
}