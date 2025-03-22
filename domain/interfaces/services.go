package interfaces

import (
	"KPO1/domain/models"
	"time"
)

// BankAccountService представляет сервис для управления банковскими счетами
type BankAccountService interface {
	CreateBankAccount(name string) (*models.BankAccount, error)
	GetBankAccount(id int) (*models.BankAccount, error)
	GetAllBankAccounts() ([]*models.BankAccount, error)
	UpdateBankAccount(id int, name string) (*models.BankAccount, error)
	DeleteBankAccount(id int) error
	RecalculateBalance(id int) (*models.BankAccount, error)
}

// CategoryService представляет сервис для управления категориями
type CategoryService interface {
	CreateCategory(name string, opType models.OperationType) (*models.Category, error)
	GetCategory(id int) (*models.Category, error)
	GetAllCategories() ([]*models.Category, error)
	GetCategoriesByType(opType models.OperationType) ([]*models.Category, error)
	UpdateCategory(id int, name string, opType models.OperationType) (*models.Category, error)
	DeleteCategory(id int) error
}

// OperationService представляет сервис для управления операциями
type OperationService interface {
	CreateOperation(bankAccountID, categoryID int, amount float64, opType models.OperationType, date time.Time, description string) (*models.Operation, error)
	GetOperation(id int) (*models.Operation, error)
	GetAllOperations() ([]*models.Operation, error)
	GetOperationsByBankAccount(bankAccountID int) ([]*models.Operation, error)
	GetOperationsByCategory(categoryID int) ([]*models.Operation, error)
	GetOperationsByDateRange(start, end time.Time) ([]*models.Operation, error)
	UpdateOperation(id, bankAccountID, categoryID int, amount float64, opType models.OperationType, date time.Time, description string) (*models.Operation, error)
	DeleteOperation(id int) error
}

// AnalyticsService представляет сервис для аналитики финансов
type AnalyticsService interface {
	GetIncomeExpenseDifference(start, end time.Time) (float64, error)
	GetCategorySummary(start, end time.Time) (map[*models.Category]float64, error)
	GetMonthlyDynamics(year int) (map[time.Month]map[models.OperationType]float64, error)
}
