package factory

import (
	"KPO1/domain/models"
	"time"
)

// BankAccountFactory представляет фабрику для создания банковских счетов
type BankAccountFactory struct {
	nextID int
}

// NewBankAccountFactory создаёт новую фабрику счетов
func NewBankAccountFactory() *BankAccountFactory {
	return &BankAccountFactory{
		nextID: 1,
	}
}

// CreateBankAccount создаёт новый банковский счёт
func (f *BankAccountFactory) CreateBankAccount(name string) (*models.BankAccount, error) {
	now := time.Now()
	account := &models.BankAccount{
		ID:        f.nextID,
		Name:      name,
		Balance:   0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Валидация счёта
	if err := account.Validate(); err != nil {
		return nil, err
	}

	f.nextID++
	return account, nil
}

// SetNextID устанавливает следующий ID для фабрики
func (f *BankAccountFactory) SetNextID(id int) {
	if id > f.nextID {
		f.nextID = id
	}
}
