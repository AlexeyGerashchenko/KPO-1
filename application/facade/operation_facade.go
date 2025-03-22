package facade

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// OperationFacade представляет фасад для работы с операциями
type OperationFacade struct {
	operationService interfaces.OperationService
	categoryService  interfaces.CategoryService
}

// NewOperationFacade создаёт новый фасад для работы с операциями
func NewOperationFacade(
	operationService interfaces.OperationService,
	categoryService interfaces.CategoryService,
) *OperationFacade {
	return &OperationFacade{
		operationService: operationService,
		categoryService:  categoryService,
	}
}

// CreateOperation создает новую операцию
func (f *OperationFacade) CreateOperation(
	bankAccountID, categoryID int,
	amount float64,
	date time.Time,
	description string,
) (*models.Operation, error) {
	// Валидация входных данных
	if bankAccountID <= 0 {
		return nil, &models.ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	if categoryID <= 0 {
		return nil, &models.ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	if amount <= 0 {
		return nil, &models.ValidationError{Message: "Сумма операции должна быть положительным числом"}
	}

	// Определяем тип операции на основе категории
	category, err := f.categoryService.GetCategory(categoryID)
	if err != nil {
		return nil, err
	}

	// Создаем операцию с типом, соответствующим категории
	return f.operationService.CreateOperation(
		bankAccountID,
		categoryID,
		amount,
		category.Type,
		date,
		description,
	)
}

// GetOperationDetails получает детальную информацию об операции
func (f *OperationFacade) GetOperationDetails(id int) (*models.Operation, error) {
	if id <= 0 {
		return nil, &models.ValidationError{Message: "ID операции должен быть положительным числом"}
	}

	return f.operationService.GetOperation(id)
}

// GetAllOperations получает список всех операций
func (f *OperationFacade) GetAllOperations() ([]*models.Operation, error) {
	return f.operationService.GetAllOperations()
}

// GetOperationsByBankAccount получает список операций по банковскому счёту
func (f *OperationFacade) GetOperationsByBankAccount(bankAccountID int) ([]*models.Operation, error) {
	if bankAccountID <= 0 {
		return nil, &models.ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	return f.operationService.GetOperationsByBankAccount(bankAccountID)
}

// GetOperationsByCategory получает список операций по категории
func (f *OperationFacade) GetOperationsByCategory(categoryID int) ([]*models.Operation, error) {
	if categoryID <= 0 {
		return nil, &models.ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	return f.operationService.GetOperationsByCategory(categoryID)
}

// GetOperationsByDateRange получает список операций за период
func (f *OperationFacade) GetOperationsByDateRange(start, end time.Time) ([]*models.Operation, error) {
	if start.After(end) {
		return nil, &models.ValidationError{Message: "Дата начала не может быть позже даты окончания"}
	}

	return f.operationService.GetOperationsByDateRange(start, end)
}

// UpdateOperation обновляет информацию об операции
func (f *OperationFacade) UpdateOperation(
	id, bankAccountID, categoryID int,
	amount float64,
	date time.Time,
	description string,
) (*models.Operation, error) {
	// Валидация входных данных
	if id <= 0 {
		return nil, &models.ValidationError{Message: "ID операции должен быть положительным числом"}
	}

	if bankAccountID <= 0 {
		return nil, &models.ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	if categoryID <= 0 {
		return nil, &models.ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	if amount <= 0 {
		return nil, &models.ValidationError{Message: "Сумма операции должна быть положительным числом"}
	}

	// Определяем тип операции на основе категории
	category, err := f.categoryService.GetCategory(categoryID)
	if err != nil {
		return nil, err
	}

	// Обновляем операцию с типом, соответствующим категории
	return f.operationService.UpdateOperation(
		id,
		bankAccountID,
		categoryID,
		amount,
		category.Type,
		date,
		description,
	)
}

// DeleteOperation удаляет операцию
func (f *OperationFacade) DeleteOperation(id int) error {
	if id <= 0 {
		return &models.ValidationError{Message: "ID операции должен быть положительным числом"}
	}

	return f.operationService.DeleteOperation(id)
}
