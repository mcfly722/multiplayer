package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("my_supper_pupper_secret_phase")
var lastUserID float64

func isAuthorized(endpoint func(float64, http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint((token.Claims.(jwt.MapClaims))["Id"].(float64), w, r)
				/*
					id, err := strconv.Atoi((token.Claims.(jwt.MapClaims))["Id"].(string))
					if err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						fmt.Fprintf(w, err.Error())
					} else {
					}
				*/
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func main() {
	var bindingAddress = "localhost:8080"

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/login", login)
	http.Handle("/api/movement", isAuthorized(movement))

	log.Printf("starting server at %v", bindingAddress)
	log.Fatal(http.ListenAndServe(bindingAddress, nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Id"] = lastUserID
	claims["Exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Printf("new user entered id=%v", lastUserID)
		lastUserID++
		w.Write([]byte(tokenString))
	}
}

// Movement - pressed buttons on client
type Movement struct {
	ArrowUp    bool
	ArrowDown  bool
	ArrowLeft  bool
	ArrowRight bool
	Space      bool
}

func movement(id float64, w http.ResponseWriter, r *http.Request) {
	var movement Movement

	err := json.NewDecoder(r.Body).Decode(&movement)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("movement id=%v %+v", id, movement)
}
