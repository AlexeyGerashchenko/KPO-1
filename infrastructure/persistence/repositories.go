package persistence

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// BankAccountRepositoryAdapter адаптер репозитория для банковских счетов
type BankAccountRepositoryAdapter struct {
	repo *MemoryRepository
}

// NewBankAccountRepository создает новый репозиторий для банковских счетов
func NewBankAccountRepository(repo *MemoryRepository) interfaces.BankAccountRepository {
	return &BankAccountRepositoryAdapter{repo: repo}
}

// GetByID получает банковский счет по ID
func (a *BankAccountRepositoryAdapter) GetByID(id int) (*models.BankAccount, error) {
	return a.repo.GetBankAccountByID(id)
}

// GetAll получает все банковские счета
func (a *BankAccountRepositoryAdapter) GetAll() ([]*models.BankAccount, error) {
	return a.repo.GetAllBankAccounts()
}

// Save сохраняет банковский счет
func (a *BankAccountRepositoryAdapter) Save(account *models.BankAccount) error {
	return a.repo.SaveBankAccount(account)
}

// Update обновляет банковский счет
func (a *BankAccountRepositoryAdapter) Update(account *models.BankAccount) error {
	return a.repo.UpdateBankAccount(account)
}

// Delete удаляет банковский счет
func (a *BankAccountRepositoryAdapter) Delete(id int) error {
	return a.repo.DeleteBankAccount(id)
}

// CategoryRepositoryAdapter адаптер репозитория для категорий
type CategoryRepositoryAdapter struct {
	repo *MemoryRepository
}

// NewCategoryRepository создает новый репозиторий для категорий
func NewCategoryRepository(repo *MemoryRepository) interfaces.CategoryRepository {
	return &CategoryRepositoryAdapter{repo: repo}
}

// GetByID получает категорию по ID
func (a *CategoryRepositoryAdapter) GetByID(id int) (*models.Category, error) {
	return a.repo.GetCategoryByID(id)
}

// GetAll получает все категории
func (a *CategoryRepositoryAdapter) GetAll() ([]*models.Category, error) {
	return a.repo.GetAllCategories()
}

// Save сохраняет категорию
func (a *CategoryRepositoryAdapter) Save(category *models.Category) error {
	return a.repo.SaveCategory(category)
}

// Update обновляет категорию
func (a *CategoryRepositoryAdapter) Update(category *models.Category) error {
	return a.repo.UpdateCategory(category)
}

// Delete удаляет категорию
func (a *CategoryRepositoryAdapter) Delete(id int) error {
	return a.repo.DeleteCategory(id)
}

// GetByType получает категории по типу операции
func (a *CategoryRepositoryAdapter) GetByType(opType models.OperationType) ([]*models.Category, error) {
	return a.repo.GetCategoriesByType(opType)
}

// OperationRepositoryAdapter адаптер репозитория для операций
type OperationRepositoryAdapter struct {
	repo *MemoryRepository
}

// NewOperationRepository создает новый репозиторий для операций
func NewOperationRepository(repo *MemoryRepository) interfaces.OperationRepository {
	return &OperationRepositoryAdapter{repo: repo}
}

// GetByID получает операцию по ID
func (a *OperationRepositoryAdapter) GetByID(id int) (*models.Operation, error) {
	return a.repo.GetOperationByID(id)
}

// GetAll получает все операции
func (a *OperationRepositoryAdapter) GetAll() ([]*models.Operation, error) {
	return a.repo.GetAllOperations()
}

// Save сохраняет операцию
func (a *OperationRepositoryAdapter) Save(operation *models.Operation) error {
	return a.repo.SaveOperation(operation)
}

// Update обновляет операцию
func (a *OperationRepositoryAdapter) Update(operation *models.Operation) error {
	return a.repo.UpdateOperation(operation)
}

// Delete удаляет операцию
func (a *OperationRepositoryAdapter) Delete(id int) error {
	return a.repo.DeleteOperation(id)
}

// GetByBankAccountID получает операции по ID банковского счета
func (a *OperationRepositoryAdapter) GetByBankAccountID(bankAccountID int) ([]*models.Operation, error) {
	return a.repo.GetOperationsByBankAccountID(bankAccountID)
}

// GetByCategoryID получает операции по ID категории
func (a *OperationRepositoryAdapter) GetByCategoryID(categoryID int) ([]*models.Operation, error) {
	return a.repo.GetOperationsByCategoryID(categoryID)
}

// GetByDateRange получает операции в указанном диапазоне дат
func (a *OperationRepositoryAdapter) GetByDateRange(start, end time.Time) ([]*models.Operation, error) {
	return a.repo.GetOperationsByDateRange(start, end)
}

// GetByTypeAndDateRange получает операции определенного типа в указанном диапазоне дат
func (a *OperationRepositoryAdapter) GetByTypeAndDateRange(opType models.OperationType, start, end time.Time) ([]*models.Operation, error) {
	return a.repo.GetOperationsByTypeAndDateRange(opType, start, end)
}
