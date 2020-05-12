package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	. "./models"
	util "./utils"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "usersh"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM usersh ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	res := []Users{}
	for selDB.Next() {
		var id int
		var name, surname string
		err = selDB.Scan(&id, &name, &surname)
		if err != nil {
			panic(err.Error())
		}
		usr.Id = id
		usr.Name = name
		usr.Surname = surname
		res = append(res, usr)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM usersh WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	for selDB.Next() {
		var id int
		var name, surname string
		err = selDB.Scan(&id, &name, &surname)
		if err != nil {
			panic(err.Error())
		}
		usr.Id = id
		usr.Name = name
		usr.Surname = surname
	}
	tmpl.ExecuteTemplate(w, "Show", usr)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM usersh WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	usr := Users{}
	for selDB.Next() {
		var id int
		var name, surname string
		err = selDB.Scan(&id, &name, &surname)
		if err != nil {
			panic(err.Error())
		}
		usr.Id = id
		usr.Name = name
		usr.Surname = surname
	}
	tmpl.ExecuteTemplate(w, "Edit", usr)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		insForm, err := db.Prepare("INSERT INTO usersh(name, surname) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, surname)
		log.Println("INSERT: Name: " + name + " | Surname: " + surname)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE usersh SET name=?, surname=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, surname, id)
		log.Println("UPDATE: Name: " + name + " | Surname: " + surname)
	}
	defer db.Close()
	http.Redirect(w, r, "/admin", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	usr := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM usersh WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(usr)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/admin", 301)
}

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Home", nil)

}

func User(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "User", nil)

}

func Administrator(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Admin", nil)

}

func Result(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "Resulttitle", nil)

	bytes, _ := util.ReadFile("./conf.json")
	var data []Conf
	json.Unmarshal([]byte(bytes), &data)

	load := LoadConf{}

	for i := range data {

		load.ID = data[i].ID
		load.Question = data[i].Question
		load.TrueAnswer = data[i].TrueAnswer

		tmpl.ExecuteTemplate(w, "Result", load)

	}

	tmpl.ExecuteTemplate(w, "Goback", nil)

}

func Questions(w http.ResponseWriter, r *http.Request) {

	bytes, _ := util.ReadFile("./conf.json")
	var data []Conf
	json.Unmarshal([]byte(bytes), &data)

	load := LoadConf{}

	for i := range data {

		load.Question = data[i].Question

		tmpl.ExecuteTemplate(w, "Questions", load)

	}

	tmpl.ExecuteTemplate(w, "Goback", nil)

}

func Start(w http.ResponseWriter, r *http.Request) {

	bytes, _ := util.ReadFile("./conf.json")
	var data []Conf
	json.Unmarshal([]byte(bytes), &data)

	load := LoadConf{}

	for i := range data {

		load.ID = data[i].ID
		load.Question = data[i].Question
		load.TrueAnswer = data[i].TrueAnswer
		load.FalseAnswer1 = data[i].FalseAnswer1
		load.FalseAnswer2 = data[i].FalseAnswer2

		tmpl.ExecuteTemplate(w, "Start", load)

	}

	tmpl.ExecuteTemplate(w, "And", nil)

}

func Adminlogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		adminname := r.FormValue("adminname")
		adminsurname := r.FormValue("adminsurname")

		aname := "root"
		asurname := "root"

		if adminname == aname && adminsurname == asurname {

			log.Println("Hi Admin!")
			http.Redirect(w, r, "/admin", 301)

		} else {

			log.Println("False name and surname!")
		}

	}

	http.Redirect(w, r, "/", 301)

}

func Userlogin(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	if r.Method == "POST" {
		username := r.FormValue("userrname")
		usersurname := r.FormValue("userrsurname")

		if username != "" && usersurname != "" {

			insForm, err := db.Prepare("INSERT INTO usersh(name, surname) VALUES(?,?)")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(username, usersurname)
			log.Println("User: Name: " + username + " | Surname: " + usersurname)
			log.Println("Start!")

			http.Redirect(w, r, "/start", 301)
		}

		http.Redirect(w, r, "/user", 301)

	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)

}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/admin", Index)
	http.HandleFunc("/adminlogin", Adminlogin)
	http.HandleFunc("/administrator", Administrator)
	http.HandleFunc("/login", Userlogin)
	http.HandleFunc("/user", User)
	http.HandleFunc("/start", Start)
	http.HandleFunc("/questions", Questions)
	http.HandleFunc("/", Home)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/result", Result)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
