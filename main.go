package multiplayer

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey []byte

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
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mySigningKey = []byte(randomString(32))

	var bindingAddress = "localhost:8080"

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/api/login", login)
	http.Handle("/api/movement", isAuthorized(movement))
	http.Handle("/api/state", isAuthorized(state))

	log.Printf("starting server at %v (session key:%s)", bindingAddress, mySigningKey)
	go PlayGame()
	log.Fatal(http.ListenAndServe(bindingAddress, nil))
}

func login(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	playerID := NewPlayer()
	log.Printf("new player created id=%v", playerID)

	claims["Id"] = playerID
	claims["Exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write([]byte(tokenString))
	}
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func movement(id float64, w http.ResponseWriter, r *http.Request) {
	var movement Movement

	err := json.NewDecoder(r.Body).Decode(&movement)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ApplyPlayerMovement(int(id), movement)
	//	log.Printf("movement id=%v %+v", id, movement)
}

func state(id float64, w http.ResponseWriter, r *http.Request) {
	jsonString, err := json.Marshal(Players)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write([]byte(jsonString))
	}
}
