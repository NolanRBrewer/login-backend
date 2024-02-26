package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type login_input struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func loginRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var t login_input
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		// log.Print(t)
		// log.Printf("Username: %s", t.User)
		// log.Printf("Password: %s", t.Password)
		//check credentials
		if (t.User == "c137@onecause.com") && (t.Password == "#th@nH@rm#y#r!$100%D0p#") {
			validation, err := json.Marshal(true)
			if err != nil {
				fmt.Printf("Correct credentials, %v", err)
				return
			} else {
				io.WriteString(w, string(validation))
			}
		} else {
			validation, err := json.Marshal(false)
			if err != nil {
				fmt.Printf("Incorrect credentials, %v", err)
				return
			} else {
				io.WriteString(w, string(validation))
			}
		}
		// error example:
		// if err := r.ParseForm(); err != nil {
		// 	fmt.Fprintf(w, "ParseForm() err: %v", err)
		// 	return
		// }

	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
	fmt.Printf("got /login request\n")
}
func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/login", loginRequest)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
