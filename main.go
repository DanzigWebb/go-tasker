package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"task-app/models"
)

func main() {
	http.HandleFunc("/task/create", handleCreateTask)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCreateTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(body))
		var t models.Task
		err = json.Unmarshal(body, &t)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(200)
		w.Write(body)
	}
}
