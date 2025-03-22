package factory

import (
	"KPO1/domain/models"
	"time"
)

// OperationFactory представляет фабрику для создания операций
type OperationFactory struct {
	nextID int
}

// NewOperationFactory создаёт новую фабрику операций
func NewOperationFactory() *OperationFactory {
	return &OperationFactory{
		nextID: 1,
	}
}

// CreateOperation создаёт новую операцию
func (f *OperationFactory) CreateOperation(
	bankAccountID, categoryID int,
	amount float64,
	opType models.OperationType,
	date time.Time,
	description string,
) (*models.Operation, error) {
	now := time.Now()

	operation := &models.Operation{
		ID:            f.nextID,
		Type:          opType,
		BankAccountID: bankAccountID,
		CategoryID:    categoryID,
		Amount:        amount,
		Date:          date,
		Description:   description,
		CreatedAt:     now,
	}

	// Валидация операции
	if err := operation.Validate(); err != nil {
		return nil, err
	}

	f.nextID++
	return operation, nil
}

// SetNextID устанавливает следующий ID для фабрики
func (f *OperationFactory) SetNextID(id int) {
	if id > f.nextID {
		f.nextID = id
	}
}
