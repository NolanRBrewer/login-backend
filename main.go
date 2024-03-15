package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type login_input struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "You connected, but try /login!\n")
}
func loginRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var credentials login_input
		err := decoder.Decode(&credentials)
		if err != nil {
			panic(err)
		}
		// upon login request capture current time
		submission_time := time.Now().Format("1504")
		fmt.Printf(submission_time)

		/*TODO create a database for users and passwords, set up appropriate query for validating users*/

		/* Example query
		 */
		//check credentials
		if (credentials.User == "c137@onecause.com") && (credentials.Password == "#th@nH@rm#y#r!$100%D0p#") && (credentials.Token == submission_time) {
			//valid login
			validation, err := json.Marshal(true)
			if err != nil {
				fmt.Printf("Correct credentials, %v", err)
				return
			} else {
				io.WriteString(w, string(validation))
			}
		} else {
			//invalid login
			validation, err := json.Marshal(false)
			if err != nil {
				fmt.Printf("Incorrect credentials, %v", err)
				return
			} else {
				io.WriteString(w, string(validation))
			}
		}

	default:
		//Non-POST request handling
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
