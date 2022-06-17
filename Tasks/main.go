package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"

	"log"
	"net/http"
)

var database *sql.DB

func init() {
	defer database.Close()
	database, err = sql.Open("sqlite3", "./tasks.db")

	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	// http.HandleFunc("/complete/", CompleteTaskFunc)
	// http.HandleFunc("/delete/", DeleteTaskFunc)
	// http.HandleFunc("/deleted/", ShowTrashTaskFunc)
	// http.HandleFunc("/trash/", TrashTaskFunc)
	// http.HandleFunc("/edit/", EditTaskFunc)
	// http.HandleFunc("/completed/", ShowCompleteTasksFunc)
	// http.HandleFunc("/restore/", RestoreTaskFunc)
	// http.HandleFunc("/add/", AddTaskFunc)
	// http.HandleFunc("/update/", UpdateTaskFunc)
	// http.HandleFunc("/search/", SearchTaskFunc)
	// http.HandleFunc("/login", GetLogin)
	// http.HandleFunc("/register", PostRegister)
	// http.HandleFunc("/admin", HandleAdmin)
	// http.HandleFunc("/add_user", PostAddUser)
	// http.HandleFunc("/change", PostChange)
	// http.HandleFunc("/logout", HandleLogout)
	http.HandleFunc("/", ShowAllTasksFunc)

	http.Handle("/static/", http.FileServer(http.Dir("public")))
	log.Print("running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	var task []Task
	var context Context
	var TaskID int
	var TaskTitle string
	var TaskContent string
	var TaskCreated time.Time
	var getTasksql string

	getTasksql = "select id, title, content, created_date from task;"

	rows, err := database.Query(getTaskSQL)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&TaskID, &TaskTitle, &TaskContent, &TaskCreated)
		TaskContent = strings.Replace(TaskContent, "\n", "<br>", -1)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(TaskID, TaskTitle, TaskContent, TaskCreated)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "all pending tasks GET"
	} else {
		message = "all pending tasks POST"
	}
	w.Write([]byte(message))
}
