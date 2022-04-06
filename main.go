package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type response struct {
	Payload map[string]interface{} `json:"payload"`
	Headers map[string][]string    `json:"headers"`
	EnvironmentVariables []string  `json:"environmentVariables"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	var response = response{
		Payload: data,
		Headers: r.Header,
		EnvironmentVariables: get_environment_variables(),
	}

	log.Println(response)


	marshalledData, _ := json.Marshal(response)
	log.Println(string(marshalledData))
	fmt.Println(string(marshalledData))

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(
		w,
		string(marshalledData),
	)
}

func get_environment_variables () []string {
	var env = []string{}
	for _, e := range os.Environ() {
		log.Println(e)
		if(strings.Contains(e, "PRACTICE")) {
			env = append(env, e)
		}
	}
	return env
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(handler)
	return r
}

// Initiate web server
func main() {
	router := router()
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server started")

	log.Fatal(srv.ListenAndServe())
}