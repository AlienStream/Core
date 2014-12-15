package artist

import (
	"fmt"
	"net/http"

    //"AlienStream/Models"
)


// CRUD
func Info(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}

func Create(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}

func Edit(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}

func Delete(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}

func Flag(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}


// Info
func Tracks(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}

func Communities(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}


// Aggregation
func Trending(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}

func Popular(w http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(w,"here's what's trending\n")
}