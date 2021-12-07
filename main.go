package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Wait struct {
	time int
}

type Echo struct {
	time        int
	payloadPath string
}

func (wait *Wait) healthcheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("wait %d second", wait.time)
	time.Sleep(time.Duration(wait.time) * time.Second)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"up"}`))
	log.Println("returned")
}

func (echo *Echo) echo(w http.ResponseWriter, r *http.Request) {
	log.Printf("wait %d second", echo.time)
	log.Println("received")
	time.Sleep(time.Duration(echo.time) * time.Second)

	w.WriteHeader(http.StatusOK)
	jsonFile, err := os.Open(echo.payloadPath)
	if err != nil {
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	w.Write([]byte(byteValue))
	log.Println("returned")
}

func main() {
	wait := flag.Int("wait", 3, "")
	port := flag.Int("port", 8080, "")
	path := flag.String("path", "/api/echo", "")
	method := flag.String("method", "POST", "")
	payload := flag.String("payload", "", "path to payload json file")

	flag.Parse()

	w := &Wait{time: *wait}
	e := &Echo{time: *wait, payloadPath: *payload}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/healthcheck", w.healthcheck).Methods("GET")
	router.HandleFunc(*path, e.echo).Methods(*method)

	http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
}
