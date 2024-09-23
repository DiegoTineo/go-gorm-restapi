package main

import (
	"net/http"

	"github.com/DiegoTineo/go-gorm-restapi/db"
	"github.com/DiegoTineo/go-gorm-restapi/models"
	"github.com/DiegoTineo/go-gorm-restapi/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConection()
	db.DB.AutoMigrate(models.Task{})	
	db.DB.AutoMigrate(models.User{})	

	// w: response writer (to write the response)
	// r: request (to read the request)

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	// Users
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.CreateUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	// Tasks
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	
	http.ListenAndServe(":8080", r)
}
