package facade

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
)

// CategoryFacade представляет фасад для работы с категориями
type CategoryFacade struct {
	categoryService interfaces.CategoryService
}

// NewCategoryFacade создаёт новый фасад для работы с категориями
func NewCategoryFacade(categoryService interfaces.CategoryService) *CategoryFacade {
	return &CategoryFacade{
		categoryService: categoryService,
	}
}

// CreateCategory создает новую категорию
func (f *CategoryFacade) CreateCategory(name string, opType models.OperationType) (*models.Category, error) {
	// Валидация входных данных
	if name == "" {
		return nil, &models.ValidationError{Message: "Название категории не может быть пустым"}
	}

	if opType != models.Income && opType != models.Expense {
		return nil, &models.ValidationError{Message: "Тип категории должен быть INCOME или EXPENSE"}
	}

	return f.categoryService.CreateCategory(name, opType)
}

// GetCategory получает категорию по ID
func (f *CategoryFacade) GetCategory(id int) (*models.Category, error) {
	if id <= 0 {
		return nil, &models.ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	return f.categoryService.GetCategory(id)
}

// GetAllCategories получает все категории
func (f *CategoryFacade) GetAllCategories() ([]*models.Category, error) {
	return f.categoryService.GetAllCategories()
}

// GetCategoriesByType получает категории по типу операции
func (f *CategoryFacade) GetCategoriesByType(opType models.OperationType) ([]*models.Category, error) {
	if opType != models.Income && opType != models.Expense {
		return nil, &models.ValidationError{Message: "Тип категории должен быть INCOME или EXPENSE"}
	}

	return f.categoryService.GetCategoriesByType(opType)
}

// UpdateCategory обновляет категорию
func (f *CategoryFacade) UpdateCategory(id int, name string, opType models.OperationType) (*models.Category, error) {
	if id <= 0 {
		return nil, &models.ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	if name == "" {
		return nil, &models.ValidationError{Message: "Название категории не может быть пустым"}
	}

	if opType != models.Income && opType != models.Expense {
		return nil, &models.ValidationError{Message: "Тип категории должен быть INCOME или EXPENSE"}
	}

	return f.categoryService.UpdateCategory(id, name, opType)
}

// DeleteCategory удаляет категорию
func (f *CategoryFacade) DeleteCategory(id int) error {
	if id <= 0 {
		return &models.ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	return f.categoryService.DeleteCategory(id)
}
