package facade

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
)

// BankAccountFacade представляет фасад для работы с банковскими счетами
type BankAccountFacade struct {
	bankAccountService interfaces.BankAccountService
}

// NewBankAccountFacade создаёт новый фасад для работы с банковскими счетами
func NewBankAccountFacade(bankAccountService interfaces.BankAccountService) *BankAccountFacade {
	return &BankAccountFacade{
		bankAccountService: bankAccountService,
	}
}

// CreateBankAccount создает новый банковский счёт
func (f *BankAccountFacade) CreateBankAccount(name string) (*models.BankAccount, error) {
	// Валидация входных данных
	if name == "" {
		return nil, &models.ValidationError{Message: "Название счета не может быть пустым"}
	}

	return f.bankAccountService.CreateBankAccount(name)
}

// GetBankAccount получает банковский счёт по ID
func (f *BankAccountFacade) GetBankAccount(id int) (*models.BankAccount, error) {
	if id <= 0 {
		return nil, &models.ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	return f.bankAccountService.GetBankAccount(id)
}

// GetAllBankAccounts получает все банковские счета
func (f *BankAccountFacade) GetAllBankAccounts() ([]*models.BankAccount, error) {
	return f.bankAccountService.GetAllBankAccounts()
}

// UpdateBankAccount обновляет банковский счёт
func (f *BankAccountFacade) UpdateBankAccount(id int, name string) (*models.BankAccount, error) {
	if id <= 0 {
		return nil, &models.ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	if name == "" {
		return nil, &models.ValidationError{Message: "Название счета не может быть пустым"}
	}

	return f.bankAccountService.UpdateBankAccount(id, name)
}

// DeleteBankAccount удаляет банковский счёт
func (f *BankAccountFacade) DeleteBankAccount(id int) error {
	if id <= 0 {
		return &models.ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	return f.bankAccountService.DeleteBankAccount(id)
}

// RecalculateBalance пересчитывает баланс счёта
func (f *BankAccountFacade) RecalculateBalance(id int) (*models.BankAccount, error) {
	if id <= 0 {
		return nil, &models.ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	return f.bankAccountService.RecalculateBalance(id)
}
