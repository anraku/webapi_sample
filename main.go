package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GETのみ
func handlerGET(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "GET Method")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// POSTのみ
func handlerPOST(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Fprintln(w, "POST Method")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// jsonを返す
func handler_user(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		user := User{
			Name: "anraku",
			Age:  25,
		}
		res, _ := json.Marshal(user)
		w.Write(res)

		// ログ出力
		data := new(User)
		json.Unmarshal(res, data)
		fmt.Println(data)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/get", handlerGET)
	http.HandleFunc("/post", handlerPOST)
	http.HandleFunc("/user", handler_user)
	http.ListenAndServe(":8080", nil)
}
