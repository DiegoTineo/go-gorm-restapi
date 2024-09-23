package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DiegoTineo/go-gorm-restapi/db"
	"github.com/DiegoTineo/go-gorm-restapi/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user []models.User
	db.DB.Find(&user)
	json.NewEncoder(w).Encode(&user)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	// esto es para obtener el parametro de la url
	params := mux.Vars(r)

	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)
}

func CreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User 
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}


	json.NewEncoder(w).Encode(&user)

}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user models.User
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	// si quieres eliminar de manera suave
	// db.DB.Delete(&user)

	// si quieres eliminar de manera definitiva
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
