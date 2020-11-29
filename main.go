package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func list(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user := User{"Cleitinho", "43 96718583", "cleito@h.com"}

	userJSON, err := json.Marshal(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Teste"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(userJSON)
	/* log.Print(w, "Teste") */
}

func main() {
	fmt.Println("oi")

	r := mux.NewRouter()

	r.HandleFunc("/users", list).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
