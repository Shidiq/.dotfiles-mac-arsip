package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"E001", "ethan", 21},
	{"E002", "kun", 22},
	{"E003", "jhon", 24},
	{"E004", "wich", 25},
	{"E005", "tuti", 22},
	{"E006", "nur", 20},
	{"E007", "bond", 21},
	{"E008", "james", 21},
	{"E009", "kina", 27},
	{"E010", "doni", 24},
	{"E011", "tata", 22},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var results, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(results)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var results []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				results, err = json.Marshal(each)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(results)
				return
			}
		}
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)
	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
