package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func httpFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.URL.Path)
	fmt.Fprintf(w, "1111\n")
	fmt.Fprintf(w, r.URL.Scheme)
	fmt.Fprintf(w, "2222\n")
	fmt.Fprintf(w, "这是测试httpFunc")
	fmt.Fprint(w, "333")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, err := template.ParseFiles("./login.html")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error parsing template:", err)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	} else {
		// 解析表单
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			log.Println("Error parsing form:", err)
			return
		}

		// 获取表单字段
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println("username:", username)
		fmt.Println("password:", password)
	}
}

func main() {
	http.HandleFunc("/", httpFunc)
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
