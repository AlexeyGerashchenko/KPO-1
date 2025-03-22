package interfaces

import (
	"KPO1/domain/models"
	"time"
)

// Repository представляет общий интерфейс для репозиториев
type Repository[T any] interface {
	GetByID(id int) (*T, error)
	GetAll() ([]*T, error)
	Save(item *T) error
	Update(item *T) error
	Delete(id int) error
}

// BankAccountRepository представляет репозиторий для работы с банковскими счетами
type BankAccountRepository interface {
	Repository[models.BankAccount]
}

// CategoryRepository представляет репозиторий для работы с категориями
type CategoryRepository interface {
	Repository[models.Category]
	GetByType(opType models.OperationType) ([]*models.Category, error)
}

// OperationRepository представляет репозиторий для работы с операциями
type OperationRepository interface {
	Repository[models.Operation]
	GetByBankAccountID(bankAccountID int) ([]*models.Operation, error)
	GetByCategoryID(categoryID int) ([]*models.Operation, error)
	GetByDateRange(start, end time.Time) ([]*models.Operation, error)
	GetByTypeAndDateRange(opType models.OperationType, start, end time.Time) ([]*models.Operation, error)
}
