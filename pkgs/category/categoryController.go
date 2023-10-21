package category

import (
	"strconv"

	"github.com/Ctere1/golang-api/pkgs/storage"
)

func createCategory(category Category) (Category, error) {
	// Create category
	categoryID, err := storage.CreateCategory(category.Name)
	if err != nil {
		return Category{}, err
	}

	// Convert categoryID to int
	id, err := strconv.Atoi(categoryID)
	if err != nil {
		return Category{}, err
	}

	// Create category object
	category = Category{
		Name: category.Name,
		Id:   strconv.Itoa(id),
	}

	return category, nil
}

func getCategory(categoryId string) (Category, error) {
	// Get category
	category, err := storage.GetCategory(categoryId)
	if err != nil {
		return Category{}, err
	}

	// Convert storage.Category to Category
	result := Category{
		Name: category.Name,
		Id:   strconv.Itoa(int(category.ID)),
	}

	return result, nil
}

func getCategories() ([]Category, error) {
	// Get categories
	result, err := storage.GetAllCategories()
	if err != nil {
		return []Category{}, err
	}

	// Convert []storage.Category to []Category
	categories := make([]Category, len(result))
	for i, category := range result {
		categories[i] = Category{
			Name: category.Name,
			Id:   strconv.Itoa(int(category.ID)),
		}
	}

	return categories, nil
}

func deleteCategory(categoryId string) error {
	// Delete category
	err := storage.DeleteCategory(categoryId)
	if err != nil {
		return err
	}

	return nil
}

func updateCategory(category Category) (Category, error) {
	// Update category
	err := storage.UpdateCategory(category.Id, category.Name)
	if err != nil {
		return Category{}, err
	}

	return category, nil

}
