package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo-app/internal/models"
	"todo-app/internal/services"
	"todo-app/pkg/auth"

	"github.com/gorilla/mux"
)

var mockService = services.NewMockService()

func Login(jwtService *auth.JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		user, err := mockService.GetUserByUsername(req.Username)
		if err != nil || user == nil || user.Password != req.Password {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, err := jwtService.GenerateToken(user)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(models.LoginResponse{Token: token})
	}
}

func GetTodoLists(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	isAdmin := r.Context().Value("role").(string) == "admin"

	lists, err := mockService.GetTodoLists(userID, isAdmin)
	if err != nil {
		http.Error(w, "Error getting todo lists", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(lists)
}

func CreateTodoList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	var list models.TodoList
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	list.UserID = userID
	if err := mockService.CreateTodoList(&list); err != nil {
		http.Error(w, "Error creating todo list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(list)
}

func GetTodoList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	isAdmin := r.Context().Value("role").(string) == "admin"

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	list, err := mockService.GetTodoList(id, userID, isAdmin)
	if err != nil {
		http.Error(w, "Error getting todo list", http.StatusInternalServerError)
		return
	}

	if list == nil {
		http.Error(w, "Todo list not found", http.StatusNotFound)
		return
	}

	// Get items for the list
	items, err := mockService.GetTodoItems(id)
	if err != nil {
		http.Error(w, "Error getting todo items", http.StatusInternalServerError)
		return
	}

	// Create a response that includes both list and items
	response := struct {
		*models.TodoList
		Items []models.TodoItem `json:"items"`
	}{
		TodoList: list,
		Items:    items,
	}

	json.NewEncoder(w).Encode(response)
}

func UpdateTodoList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	isAdmin := r.Context().Value("role").(string) == "admin"

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	list, err := mockService.GetTodoList(id, userID, isAdmin)
	if err != nil {
		http.Error(w, "Error getting todo list", http.StatusInternalServerError)
		return
	}

	if list == nil {
		http.Error(w, "Todo list not found", http.StatusNotFound)
		return
	}

	var updatedList models.TodoList
	if err := json.NewDecoder(r.Body).Decode(&updatedList); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedList.ID = id
	updatedList.UserID = userID
	if err := mockService.UpdateTodoList(&updatedList); err != nil {
		http.Error(w, "Error updating todo list", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedList)
}

func DeleteTodoList(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	isAdmin := r.Context().Value("role").(string) == "admin"

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	list, err := mockService.GetTodoList(id, userID, isAdmin)
	if err != nil {
		http.Error(w, "Error getting todo list", http.StatusInternalServerError)
		return
	}

	if list == nil {
		http.Error(w, "Todo list not found", http.StatusNotFound)
		return
	}

	if err := mockService.DeleteTodoList(id); err != nil {
		http.Error(w, "Error deleting todo list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetTodoItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["listId"])
	if err != nil {
		http.Error(w, "Invalid list ID", http.StatusBadRequest)
		return
	}

	items, err := mockService.GetTodoItems(listID)
	if err != nil {
		http.Error(w, "Error getting todo items", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

func CreateTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	listID, err := strconv.Atoi(vars["listId"])
	if err != nil {
		http.Error(w, "Invalid list ID", http.StatusBadRequest)
		return
	}

	var item models.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Debug logging
	log.Printf("Creating todo item for list %d: %+v", listID, item)

	item.ListID = listID
	if err := mockService.CreateTodoItem(&item); err != nil {
		http.Error(w, "Error creating todo item", http.StatusInternalServerError)
		return
	}

	// Debug logging
	log.Printf("Created todo item: %+v", item)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

func UpdateTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["itemId"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.TodoItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	item.ID = itemID
	if err := mockService.UpdateTodoItem(&item); err != nil {
		http.Error(w, "Error updating todo item", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

func DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID, err := strconv.Atoi(vars["itemId"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	if err := mockService.DeleteTodoItem(itemID); err != nil {
		http.Error(w, "Error deleting todo item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
