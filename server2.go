package main

import (
	"net/http"
	"text/template"
)

type Page struct { // テンプレート展開用のデータ構造
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", 1}                                       // テンプレート用のデータ
	tmpl, err := template.New("new").Parse("{{.Title}} {{.Count}} count") // テンプレート文字列
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page) // テンプレートをジェネレート
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler) // hello
	http.ListenAndServe(":8080", nil)
}
