package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EduPav/questions-app/db"
	"github.com/EduPav/questions-app/models"
	"github.com/EduPav/questions-app/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()

	//Add migrations to create tables in our db. Questions and Users
	db.DB.AutoMigrate(models.Question{})
	db.DB.AutoMigrate(models.User{}) // It executes the empty struct (to import it), and reads all its properties.

	router := mux.NewRouter()
	// router.HandleFunc("/", routes.HomeHandler).Methods("GET")
	// router.HandleFunc("/", routes.LoginHandler).Methods("POST")
	// router.HandleFunc("/register", routes.AddUserHandler)
	// router.HandleFunc("/add-question", routes.AddQuestionHandler).Methods("POST")

	router.HandleFunc("/api/users", routes.PostUsersHandler).Methods("POST")
	router.HandleFunc("/api/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/api/users/{id}", routes.GetOneUserHandler).Methods("GET") //id is in-between {} because it's a parameter.
	router.HandleFunc("/api/users/{id}", routes.PatchUsersHandler).Methods("PATCH")
	router.HandleFunc("/api/users/{id}", routes.DeleteUsersHandler).Methods("DELETE")

	//Question routes
	router.HandleFunc("/api/questions", routes.PostQuestionsHandler).Methods("POST")
	router.HandleFunc("/api/questions", routes.GetQuestionsHandler).Methods("GET")
	router.HandleFunc("/api/questions/{id}", routes.GetOneQuestionHandler).Methods("GET")
	router.HandleFunc("/api/rquestions/{id}", routes.GetRandomQuestionHandler).Methods("GET")
	router.HandleFunc("/api/questions/{id}", routes.PatchQuestionsHandler).Methods("PATCH")
	router.HandleFunc("/api/questions/{id}", routes.DeleteQuestionsHandler).Methods("DELETE")

	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
