package main

import (
	"fmt"
	"net/http"
	"flag"
	"time"
	"log"

	"github.com/gorilla/mux"
)

type Wait struct {
    time int
}

func (wait *Wait) healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("wait %d second", wait.time)
	log.Println("received")
	time.Sleep(time.Duration(wait.time) * time.Second)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
	log.Println("returned")
}

func main() {
	wait := flag.Int("wait", 3, "")
	port := flag.Int("port", 8080, "")
	
	flag.Parse()

	w := &Wait{time: *wait}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", w.healthcheck).Methods("GET")
	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}