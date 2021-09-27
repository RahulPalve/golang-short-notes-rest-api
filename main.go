package main

import (
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	DB_DSN = "postgres://rahulpalve:th!s1s@db:5432/short_notes_db"
)
var db *sql.DB


func Form(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	form_data := r.PostForm
	insert_form := fmt.Sprintf("insert into forms (forms_data) values (%s)",form_data["data"])
	db.QueryRow(insert_form)
}

func AllResponses(w http.ResponseWriter, r *http.Request){

	
}

func Home(w http.ResponseWriter, r *http.Request){

	html_str := `<html>
	<body>
	<form action="/form" method="POST>
	<input type="text" id="data" name="data"><input type="submit" value="Submit">
	</form>
	</body>
	</html>`
	w.Write([]byte(html_str))
}


func main(){
	db, _ = sql.Open("postgres", DB_DSN)
	
	defer db.Close()


	http.HandleFunc("/result", AllResponses)
	http.HandleFunc("/form", Form)
	http.HandleFunc("/",Home)


	http.ListenAndServe(":8000", nil)
}