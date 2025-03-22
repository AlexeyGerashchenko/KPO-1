package interfaces

import (
	"KPO1/domain/models"
)

// ExportVisitor интерфейс для экспорта данных с использованием паттерна Посетитель
type ExportVisitor interface {
	VisitBankAccounts(accounts []*models.BankAccount) error
	VisitCategories(categories []*models.Category) error
	VisitOperations(operations []*models.Operation) error
}

// CompositeRepository интерфейс композитного репозитория для экспорта/импорта
type CompositeRepository interface {
	GetBankAccounts() ([]*models.BankAccount, error)
	GetCategories() ([]*models.Category, error)
	GetOperations() ([]*models.Operation, error)
}
