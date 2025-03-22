package persistence

import (
	"KPO1/domain/models"
	"errors"
	"sync"
	"time"
)

// MemoryRepository реализация хранилища данных в памяти
type MemoryRepository struct {
	bankAccounts    map[int]*models.BankAccount
	categories      map[int]*models.Category
	operations      map[int]*models.Operation
	mu              sync.RWMutex
	nextBankAccID   int
	nextCategoryID  int
	nextOperationID int
}

// NewMemoryRepository создает новый экземпляр репозитория в памяти
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		bankAccounts:    make(map[int]*models.BankAccount),
		categories:      make(map[int]*models.Category),
		operations:      make(map[int]*models.Operation),
		nextBankAccID:   1,
		nextCategoryID:  1,
		nextOperationID: 1,
	}
}

// GetBankAccountByID возвращает банковский счет по его ID
func (r *MemoryRepository) GetBankAccountByID(id int) (*models.BankAccount, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	account, exists := r.bankAccounts[id]
	if !exists {
		return nil, errors.New("банковский счет не найден")
	}
	return account, nil
}

// GetAllBankAccounts возвращает все банковские счета
func (r *MemoryRepository) GetAllBankAccounts() ([]*models.BankAccount, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	accounts := make([]*models.BankAccount, 0, len(r.bankAccounts))
	for _, account := range r.bankAccounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// SaveBankAccount сохраняет банковский счет
func (r *MemoryRepository) SaveBankAccount(account *models.BankAccount) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if account.ID == 0 {
		account.ID = r.nextBankAccID
		r.nextBankAccID++
	} else if account.ID >= r.nextBankAccID {
		r.nextBankAccID = account.ID + 1
	}

	r.bankAccounts[account.ID] = account
	return nil
}

// UpdateBankAccount обновляет банковский счет
func (r *MemoryRepository) UpdateBankAccount(account *models.BankAccount) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bankAccounts[account.ID]; !exists {
		return errors.New("банковский счет не найден")
	}

	r.bankAccounts[account.ID] = account
	return nil
}

// DeleteBankAccount удаляет банковский счет
func (r *MemoryRepository) DeleteBankAccount(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.bankAccounts[id]; !exists {
		return errors.New("банковский счет не найден")
	}

	delete(r.bankAccounts, id)
	return nil
}

// GetCategoryByID возвращает категорию по её ID
func (r *MemoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	category, exists := r.categories[id]
	if !exists {
		return nil, errors.New("категория не найдена")
	}
	return category, nil
}

// GetAllCategories возвращает все категории
func (r *MemoryRepository) GetAllCategories() ([]*models.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categories := make([]*models.Category, 0, len(r.categories))
	for _, category := range r.categories {
		categories = append(categories, category)
	}
	return categories, nil
}

// GetCategoriesByType возвращает категории определенного типа
func (r *MemoryRepository) GetCategoriesByType(opType models.OperationType) ([]*models.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	categories := make([]*models.Category, 0)
	for _, category := range r.categories {
		if category.Type == opType {
			categories = append(categories, category)
		}
	}
	return categories, nil
}

// SaveCategory сохраняет категорию
func (r *MemoryRepository) SaveCategory(category *models.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if category.ID == 0 {
		category.ID = r.nextCategoryID
		r.nextCategoryID++
	} else if category.ID >= r.nextCategoryID {
		r.nextCategoryID = category.ID + 1
	}

	r.categories[category.ID] = category
	return nil
}

// UpdateCategory обновляет категорию
func (r *MemoryRepository) UpdateCategory(category *models.Category) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[category.ID]; !exists {
		return errors.New("категория не найдена")
	}

	r.categories[category.ID] = category
	return nil
}

// DeleteCategory удаляет категорию
func (r *MemoryRepository) DeleteCategory(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.categories[id]; !exists {
		return errors.New("категория не найдена")
	}

	delete(r.categories, id)
	return nil
}

// GetOperationByID возвращает операцию по её ID
func (r *MemoryRepository) GetOperationByID(id int) (*models.Operation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	operation, exists := r.operations[id]
	if !exists {
		return nil, errors.New("операция не найдена")
	}
	return operation, nil
}

// GetAllOperations возвращает все операции
func (r *MemoryRepository) GetAllOperations() ([]*models.Operation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	operations := make([]*models.Operation, 0, len(r.operations))
	for _, operation := range r.operations {
		operations = append(operations, operation)
	}
	return operations, nil
}

// GetOperationsByBankAccountID возвращает операции по ID банковского счета
func (r *MemoryRepository) GetOperationsByBankAccountID(bankAccountID int) ([]*models.Operation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	operations := make([]*models.Operation, 0)
	for _, operation := range r.operations {
		if operation.BankAccountID == bankAccountID {
			operations = append(operations, operation)
		}
	}
	return operations, nil
}

// GetOperationsByCategoryID возвращает операции по ID категории
func (r *MemoryRepository) GetOperationsByCategoryID(categoryID int) ([]*models.Operation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	operations := make([]*models.Operation, 0)
	for _, operation := range r.operations {
		if operation.CategoryID == categoryID {
			operations = append(operations, operation)
		}
	}
	return operations, nil
}

// GetOperationsByDateRange возвращает операции в указанном диапазоне дат
func (r *MemoryRepository) GetOperationsByDateRange(start, end time.Time) ([]*models.Operation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	operations := make([]*models.Operation, 0)
	for _, operation := range r.operations {
		if (operation.Date.Equal(start) || operation.Date.After(start)) &&
			(operation.Date.Equal(end) || operation.Date.Before(end)) {
			operations = append(operations, operation)
		}
	}
	return operations, nil
}

// GetOperationsByTypeAndDateRange возвращает операции определенного типа в указанном диапазоне дат
func (r *MemoryRepository) GetOperationsByTypeAndDateRange(opType models.OperationType, start, end time.Time) ([]*models.Operation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	operations := make([]*models.Operation, 0)
	for _, operation := range r.operations {
		if operation.Type == opType &&
			(operation.Date.Equal(start) || operation.Date.After(start)) &&
			(operation.Date.Equal(end) || operation.Date.Before(end)) {
			operations = append(operations, operation)
		}
	}
	return operations, nil
}

// SaveOperation сохраняет операцию
func (r *MemoryRepository) SaveOperation(operation *models.Operation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if operation.ID == 0 {
		operation.ID = r.nextOperationID
		r.nextOperationID++
	} else if operation.ID >= r.nextOperationID {
		r.nextOperationID = operation.ID + 1
	}

	r.operations[operation.ID] = operation
	return nil
}

// UpdateOperation обновляет операцию
func (r *MemoryRepository) UpdateOperation(operation *models.Operation) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.operations[operation.ID]; !exists {
		return errors.New("операция не найдена")
	}

	r.operations[operation.ID] = operation
	return nil
}

// DeleteOperation удаляет операцию
func (r *MemoryRepository) DeleteOperation(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.operations[id]; !exists {
		return errors.New("операция не найдена")
	}

	delete(r.operations, id)
	return nil
}
