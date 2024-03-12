package handlers

import (
	"MyModule/App/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
    "github.com/gorilla/mux"
	"gorm.io/gorm"
)

func RegisterHandlers(router *mux.Router, db *gorm.DB) {
	router.HandleFunc("/api/create", func(w http.ResponseWriter, r *http.Request) {
		newUserHandler(w, r, db)
	}).Methods("POST")

	router.HandleFunc("/api/getUser/{id}", func(w http.ResponseWriter, r *http.Request) {
		getUserByIdHandler(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/api/getUser", func(w http.ResponseWriter, r *http.Request) {
		getUserHandler(w, r, db)
	}).Methods("GET")

	router.HandleFunc("/api/update", func(w http.ResponseWriter, r *http.Request) {
		updateUserHandler(w, r, db)
	}).Methods("PUT")

	router.HandleFunc("/api/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		deleteUser(w, r, db)
	}).Methods("DELETE")

}

func newUserHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var people model.People
	_ = json.NewDecoder(r.Body).Decode(&people)
	newUser := model.People{
		ID:     people.ID,
		Name:   people.Name,
		Age:    people.Age,
		RoleID: people.RoleID,
	}

	if db == nil {
		http.Error(w, "Database connection is nil", http.StatusInternalServerError)
		return
	}

	db.Save(&newUser)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data saved to the database successfully!")
}
func getUserByIdHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var user model.People
	params := mux.Vars(r)
	userId := params["id"]
	userIDUint, _ := strconv.ParseUint(userId, 10, 64)
	db.First(&user, userIDUint)
	json.NewEncoder(w).Encode(&user)

}
func getUserHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var users []model.People
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var user model.People
	var role uint
	json.NewDecoder(r.Body).Decode(&user)
	var peoples []model.People
	db.Find(&peoples)
	fmt.Print(peoples)
	for _, v := range peoples {
		if v.ID == user.ID {
			role = v.RoleID
		}
	}
	updateUser := model.People{
		ID:     user.ID,
		Name:   user.Name,
		Age:    user.Age,
		RoleID: role,
	}
	db.Save(&updateUser)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated successfully!")
}

func deleteUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var user model.People
	params := mux.Vars(r)
	userId := params["id"]
	userIDUint, _ := strconv.ParseUint(userId, 10, 64)
	result := db.Delete(&user, userIDUint)
	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User deleted successfully")
}
