package routes

import (
	"encoding/json"
	"net/http"

	"github.com/EduPav/questions-app/db"
	"github.com/EduPav/questions-app/models"
	"github.com/gorilla/mux"
)

// API Handlers
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	//CREATE
	var user models.User
	json.NewDecoder(r.Body).Decode(&user) //We use & because decode receives a pointer to an object.

	createdUser := db.DB.Create(&user) //Create user in the database. And also save the new data in createdUser
	err := createdUser.Error           //In case there is an error. Otherwise it will be nil.
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)    //400
		w.Write([]byte("400 - " + err.Error())) //Send error details
		return
	}

	json.NewEncoder(w).Encode(&user) //Returns the created data in json format
}

func GetOneUserHandler(w http.ResponseWriter, r *http.Request) {
	//READ
	var user models.User
	params := mux.Vars(r) //Extract URL parameters.

	db.DB.First(&user, params["id"]) //Search for the first user that matches the id that I pass as a parameter in the DB and stores it in user.
	if user.ID == 0 {                //We get ID=0 by default in go when the user doesn't exist.
		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("404 - User not found"))
		return //Return before sending back the default user (empty)
	}
	//Table association
	db.DB.Model(&user).Association("Questions").Find(&user.Questions) //Search for all questions made by the id user.

	json.NewEncoder(w).Encode(&user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users) //Get all users(as I didn't specify any condition) in the db and saves them in users variable.
	json.NewEncoder(w).Encode(&users)
	//Status 200 by default
}

func PatchUsersHandler(w http.ResponseWriter, r *http.Request) {
	//UPDATE. We use patch because we are only updating the requested fields.
	var user models.User
	var updates map[string]interface{} //We use an interface because it can store any type of data. It's like a generic type.
	//In this case we use an interface because we don't know what type we will get from decoding the JSON data.
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - User not found"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.DB.Model(&user).Updates(updates)

	json.NewEncoder(w).Encode(&user)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	//DELETE
	user := models.User{}
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - User not found"))
		return
	}

	//Real delete
	db.DB.Unscoped().Delete(&user)

	//Logic Delete
	// db.DB.Delete(&user) //Deletes the user from the db table.
	//gorm doesn't really deletes it. It changes the deleted field to the date, a field to false and now it doesn't bring it when you use find.

	w.WriteHeader(http.StatusOK) //200. Not neccessary. It is by default.
}

//HTML Handlers. Commented temporally. Functionality not availabe. Only API.
// func AddUserHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		http.ServeFile(w, r, "./static/register.html")
// 	}

// 	if r.Method == "POST" {
// 		err := r.ParseForm()
// 		if err != nil {
// 			log.Println("There was an error", err)
// 		}
// 		//We can get the data from the form in the registration html page
// 		//Create a User struct with my read values
// 		user := models.User{
// 			Username: r.FormValue("username"),
// 			Password: r.FormValue("password"),
// 		}

// 		createdUser := db.DB.Create(&user) //Crea el usuario en la base de datos. Y además lo guardo al dato nuevo en createdUser
// 		err = createdUser.Error            //Si hay un error, lo guardo en err
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)    //Envío un status 400
// 			w.Write([]byte("400 - " + err.Error())) //Envío el error
// 			return
// 		}
// 		//Send to a page of succesful register
// 		http.ServeFile(w, r, "./static/afterRegister.html")
// 	}
// }

// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Println("There was an error", err)
// 	}

// 	user := models.User{}
// 	//Search if there is such username and password in my db
// 	db.DB.Where("username = ? AND password = ?", r.FormValue("username"), r.FormValue("password")).First(&user)
// 	//If it doesn't find any, return an error 404
// 	if user.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte("404 - User not found"))
// 		return
// 	}
// 	//If it finds it, send him to the user page
// 	tmpl, err := template.ParseFiles("./static/userPage.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, user)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// }
