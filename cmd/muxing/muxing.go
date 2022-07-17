package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

func getMessageParam(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r) // Gets param
	_, err := fmt.Fprintf(w, "Hello, %s!", param["PARAM"])
	if err != nil {
		log.Println(err)
	}
}

func getBadStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

}

func postDataWithParam(w http.ResponseWriter, r *http.Request) {
	//param := mux.Vars(r)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	_, err = fmt.Fprintf(w, "I got message:\n%s", string(body))
	if err != nil {
		log.Println(err)
	}
}

func postHeaders(w http.ResponseWriter, r *http.Request) {

	num1 := r.Header.Get("a")
	num2 := r.Header.Get("b")
	a, err := strconv.Atoi(num1)
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(num2)
	if err != nil {
		panic(err)
	}
	out := a + b
	res := strconv.Itoa(out)
	w.Header().Add("a+b", res)

}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", getMessageParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBadStatus).Methods(http.MethodGet)
	router.HandleFunc("/data", postDataWithParam).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
