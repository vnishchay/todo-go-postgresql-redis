package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"todo/config"
	"todo/models"

	"github.com/gorilla/mux"
)

var ctx = context.Background()

func CreateTodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	// append the todo table
	if result := config.DB.Create(&todo); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&todo)

}

// response :  all todos
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	result := config.DB.Find(&todos)

	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.Rows())
	json.NewEncoder(w).Encode(todos)
}

// get todo by id
// look into cache if there
// else get from db
// push into cache

func GetTodoByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	// val, err := config.Cache.Get(ctx, params["id"]).Result()

	// if err == nil {
	// 	json.NewEncoder(w).Encode(&val)
	// 	return
	// }else {
	// 	panic(err)
	// }

	var todo models.Todo
	result := config.DB.First(&todo, params["id"])
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	// err = config.Cache.Set(ctx, params["id"], result, 0).Err()
	// if err != nil {
	// 	fmt.Println("Error in Redis")
	// }

	json.NewEncoder(w).Encode(&todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	if config.DB.Model(&todo).Where("id= ?", params["id"]).Updates(&todo).RowsAffected == 0 {
		config.DB.Create(&todo)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(&todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := config.DB.Where("id = ?", params["id"]).Delete(&models.Todo{})
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode("Deleted")
}

func DeleteAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos models.Todo
	result := config.DB.Where("Todo = ?", "").Delete(&todos)
	json.NewEncoder(w).Encode(&result)
}
