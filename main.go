package main

import (
	"log"
	"net/http"
	"fmt"
	"bufio"
	"os"
	"bytes"
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
	"io/ioutil"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"FileInputCreate",
		"POST",
		"/fileInput",
		FileInputCreate,
	},
}

type FileInput struct {
	FileBytes []byte   `json:"fileBytes"` 
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the file parser service!\n")
}

func FileInputCreate(w http.ResponseWriter, r *http.Request) {
	var input FileInput
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &input); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := FindLongestLine(input.FileBytes)
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func FindLongestLine(s []byte) string { 

	maxLength := 0
	var lineLength int
	var longestLine string
	b := bytes.NewReader(s)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
		fmt.Println("length: " + strconv.Itoa(len(scanner.Text())))
		lineLength = len(scanner.Text())
		if lineLength > maxLength {
			maxLength = lineLength 
			longestLine = scanner.Text()
		}
		fmt.Println("max: " + strconv.Itoa(maxLength))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return longestLine
}
