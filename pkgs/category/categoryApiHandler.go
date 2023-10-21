package category

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Category struct {
	Name string
	Id   string
}

func ApiCreateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Create category
	result, err := createCategory(category)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			json.NewEncoder(w).Encode(map[string]string{"error": "Category already exists"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		return
	}

	// Return category as JSON
	json.NewEncoder(w).Encode(result)
}

func ApiGetCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := getIdParamFromUrl(r)

	result, err := getCategory(categoryId)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	if result.Id == "0" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Category not found"})
		return
	}

	// Return category as JSON
	json.NewEncoder(w).Encode(result)
}

func ApiGetCategories(w http.ResponseWriter, r *http.Request) {
	result, err := getCategories()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Return categories as JSON
	json.NewEncoder(w).Encode(result)
}

func ApiDeleteCategory(w http.ResponseWriter, r *http.Request) {

	categoryId := getIdParamFromUrl(r)

	err := deleteCategory(categoryId)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Return success
	json.NewEncoder(w).Encode(map[string]string{"success": "Category deleted"})
}

func ApiUpdateCategory(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	result, err := updateCategory(category)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			json.NewEncoder(w).Encode(map[string]string{"error": "Category already exists"})
		} else {
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		}
		return
	}

	// Return category as JSON
	json.NewEncoder(w).Encode(result)
}

func getIdParamFromUrl(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}
