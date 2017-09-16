package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GETのみ
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "GET Method")
}

// POSTのみ
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "POST Method")
}

// jsonを返す
func handler_json(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name: "anraku",
		Age:  25,
	}
	// jsonオブジェクトをバイト列に変換
	res, _ := json.Marshal(user)
	// レスポンスデータに書き込み
	w.Write(res)

	// ログ出力
	buf := new(bytes.Buffer)
	buf.Write(res)

	// jsonの出力にインデントとかつける
	json.Indent(buf, res, "", "     ")
	fmt.Println(buf.String())
}

func GET(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		fn(w, r)
	}
}

func POST(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/get", GET(handler1))
	http.HandleFunc("/post", POST(handler2))
	http.HandleFunc("/json", handler_json)
	http.ListenAndServe(":8080", nil)
}
