package routes

import (
	"encoding/json"
	"net/http"

	"github.com/EduPav/questions-app/db"
	"github.com/EduPav/questions-app/models"
	"github.com/gorilla/mux"
)

// API Handlers
func PostQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	var questions models.Question
	json.NewDecoder(r.Body).Decode(&questions)
	createdQuestion := db.DB.Create(&questions)
	err := createdQuestion.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - " + err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&questions)
}

func GetQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	var questions []models.Question
	db.DB.Find(&questions)
	json.NewEncoder(w).Encode(&questions)
	//200 by default
}

func GetOneQuestionHandler(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	params := mux.Vars(r)

	db.DB.First(&question, params["id"])

	if question.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Question not found"))
		return
	}

	json.NewEncoder(w).Encode(&question)
}

func GetRandomQuestionHandler(w http.ResponseWriter, r *http.Request) {
	//Pick a random question from the id user
	var question models.Question
	params := mux.Vars(r)

	db.DB.Where("creator_id = ?", params["id"]).Order("RANDOM()").First(&question)
	if question.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Question not found"))
		return
	}

	json.NewEncoder(w).Encode(&question)
}

func PatchQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	//UPDATE. We use patch because we are only updating the requested fields.
	var question models.Question
	var updates map[string]interface{}

	params := mux.Vars(r)
	db.DB.First(&question, params["id"])
	if question.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Question not found"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Model(&question).Updates(updates)

	json.NewEncoder(w).Encode(&question)
}

func DeleteQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	var question models.Question
	params := mux.Vars(r)

	db.DB.First(&question, params["id"])

	if question.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Question not found"))
		return
	}

	db.DB.Unscoped().Delete(&question)
	w.WriteHeader(http.StatusNoContent) //Status 204. All went OK but there is no content to return.
}

//HTML Handlers. Commented temporally. Functionality not availabe. Only API.
// func AddQuestionHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println("There was an error", err)
// 	}
// 	//Search id in users database from the user that is logged in
// 	user := models.User{}
// 	db.DB.Where("username = ?", r.FormValue("username")).First(&user)

// 	question := models.Question{
// 		Description: r.FormValue("description"),
// 		CreatorID:   user.ID,
// 	}

// 	createdQuestion := db.DB.Create(&question)
// 	err = createdQuestion.Error //Si hay un error, lo guardo en err
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)    //Envío un status 400
// 		w.Write([]byte("400 - " + err.Error())) //Envío el error
// 		return
// 	}

// 	//Send back to user home page of succesful register
// 	http.ServeFile(w, r, "./static/userPage.html")
// }
