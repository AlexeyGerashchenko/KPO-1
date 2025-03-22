package services

import (
	"KPO1/domain/factory"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"errors"
	"time"
)

// BankAccountServiceImpl реализация сервиса для управления банковскими счетами
type BankAccountServiceImpl struct {
	bankAccountRepo interfaces.BankAccountRepository
	operationRepo   interfaces.OperationRepository
	factory         *factory.BankAccountFactory
}

// NewBankAccountService создаёт новый сервис для управления банковскими счетами
func NewBankAccountService(
	bankAccountRepo interfaces.BankAccountRepository,
	operationRepo interfaces.OperationRepository,
	factory *factory.BankAccountFactory,
) interfaces.BankAccountService {
	return &BankAccountServiceImpl{
		bankAccountRepo: bankAccountRepo,
		operationRepo:   operationRepo,
		factory:         factory,
	}
}

// CreateBankAccount создает новый банковский счёт
func (s *BankAccountServiceImpl) CreateBankAccount(name string) (*models.BankAccount, error) {
	account, err := s.factory.CreateBankAccount(name)
	if err != nil {
		return nil, err
	}

	err = s.bankAccountRepo.Save(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// GetBankAccount получает банковский счёт по ID
func (s *BankAccountServiceImpl) GetBankAccount(id int) (*models.BankAccount, error) {
	account, err := s.bankAccountRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// GetAllBankAccounts получает все банковские счета
func (s *BankAccountServiceImpl) GetAllBankAccounts() ([]*models.BankAccount, error) {
	return s.bankAccountRepo.GetAll()
}

// UpdateBankAccount обновляет банковский счёт
func (s *BankAccountServiceImpl) UpdateBankAccount(id int, name string) (*models.BankAccount, error) {
	account, err := s.bankAccountRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	account.Name = name
	account.UpdatedAt = time.Now()

	if err := account.Validate(); err != nil {
		return nil, err
	}

	err = s.bankAccountRepo.Update(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// DeleteBankAccount удаляет банковский счёт
func (s *BankAccountServiceImpl) DeleteBankAccount(id int) error {
	// Проверяем наличие операций по этому счету
	operations, err := s.operationRepo.GetByBankAccountID(id)
	if err != nil {
		return err
	}

	if len(operations) > 0 {
		return errors.New("нельзя удалить счет, по которому есть операции")
	}

	return s.bankAccountRepo.Delete(id)
}

// RecalculateBalance пересчитывает баланс счёта
func (s *BankAccountServiceImpl) RecalculateBalance(id int) (*models.BankAccount, error) {
	account, err := s.bankAccountRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	operations, err := s.operationRepo.GetByBankAccountID(id)
	if err != nil {
		return nil, err
	}

	// Сбрасываем баланс и пересчитываем
	account.Balance = 0

	for _, op := range operations {
		if op.Type == models.Income {
			account.Balance += op.Amount
		} else {
			account.Balance -= op.Amount
		}
	}

	account.UpdatedAt = time.Now()

	err = s.bankAccountRepo.Update(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
