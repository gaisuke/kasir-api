package handlers

import (
	"encoding/json"
	"kasir-api/model"
	"kasir-api/services"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (handler *CategoryHandler) HandleCategories(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetAll(w, r)
	case http.MethodPost:
		handler.Create(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *CategoryHandler) HandleCategoriesById(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/categories/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		handler.GetByID(w, r, id)
	case http.MethodPut:
		handler.Update(w, r, id)
	case http.MethodDelete:
		handler.Delete(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) ([]model.Category, error) {
	categories, err := handler.service.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve categories", http.StatusInternalServerError)
		return nil, err
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
	return categories, nil
}

func (handler *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = handler.service.Create(&category)
	if err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

func (handler *CategoryHandler) GetByID(w http.ResponseWriter, r *http.Request, id int) {
	category, err := handler.service.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (handler *CategoryHandler) Update(w http.ResponseWriter, r *http.Request, id int) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	category.ID = id
	err = handler.service.Update(&category)
	if err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (handler *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request, id int) {
	err := handler.service.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
