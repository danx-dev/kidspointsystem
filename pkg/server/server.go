package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/danx-dev/kidspointsystem/pkg/db"
)

type Kid struct {
	Id     int
	Name   string
	Points int
}

type Parrent struct {
	Id   int
	Name string
}

func NewServer() {
	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./static/dist"))
	http.Handle("/", fs)
	http.HandleFunc("/kids", kidsHandler)
	http.HandleFunc("/kidsinfo", kidInfoHandler)
	http.HandleFunc("/kidsassignments", kidAssignmentsHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	fmt.Println("Server started at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there\n")
}

func kidsHandler(w http.ResponseWriter, r *http.Request) {
	var kids = db.GetAllKids()
	json.NewEncoder(w).Encode(kids)
}

func kidInfoHandler(w http.ResponseWriter, r *http.Request) {
	kidIdString := r.URL.Query().Get("kidid")
	kidId, _ := strconv.ParseInt(kidIdString, 0, 64)
	var kids = db.GetAllKids()

	json.NewEncoder(w).Encode(db.FindKidInArray(kidId, kids))
}

func kidAssignmentsHandler(w http.ResponseWriter, r *http.Request) {
	kidIdString := r.URL.Query().Get("kidid")
	kidId, _ := strconv.ParseInt(kidIdString, 0, 64)

	var assignments = db.GetKidAssignments(kidId)
	json.NewEncoder(w).Encode(assignments)
}
