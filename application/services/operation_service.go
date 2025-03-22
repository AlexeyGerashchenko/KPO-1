package services

import (
	"KPO1/domain/factory"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// OperationServiceImpl реализация сервиса для управления операциями
type OperationServiceImpl struct {
	operationRepo   interfaces.OperationRepository
	bankAccountRepo interfaces.BankAccountRepository
	categoryRepo    interfaces.CategoryRepository
	factory         *factory.OperationFactory
}

// NewOperationService создаёт новый сервис для управления операциями
func NewOperationService(
	operationRepo interfaces.OperationRepository,
	bankAccountRepo interfaces.BankAccountRepository,
	categoryRepo interfaces.CategoryRepository,
	factory *factory.OperationFactory,
) interfaces.OperationService {
	return &OperationServiceImpl{
		operationRepo:   operationRepo,
		bankAccountRepo: bankAccountRepo,
		categoryRepo:    categoryRepo,
		factory:         factory,
	}
}

// CreateOperation создает новую операцию
func (s *OperationServiceImpl) CreateOperation(
	bankAccountID, categoryID int,
	amount float64,
	opType models.OperationType,
	date time.Time,
	description string,
) (*models.Operation, error) {
	// Проверяем наличие счета
	account, err := s.bankAccountRepo.GetByID(bankAccountID)
	if err != nil {
		return nil, err
	}

	// Проверяем наличие категории
	category, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, err
	}

	// Проверяем соответствие типа категории и типа операции
	if category.Type != opType {
		return nil, &models.ValidationError{Message: "Тип категории не соответствует типу операции"}
	}

	// Создаем операцию
	operation, err := s.factory.CreateOperation(bankAccountID, categoryID, amount, opType, date, description)
	if err != nil {
		return nil, err
	}

	// Сохраняем операцию
	err = s.operationRepo.Save(operation)
	if err != nil {
		return nil, err
	}

	// Обновляем баланс счета
	if opType == models.Income {
		account.Balance += amount
	} else {
		account.Balance -= amount
	}
	account.UpdatedAt = time.Now()

	err = s.bankAccountRepo.Update(account)
	if err != nil {
		return nil, err
	}

	return operation, nil
}

// GetOperation получает операцию по ID
func (s *OperationServiceImpl) GetOperation(id int) (*models.Operation, error) {
	operation, err := s.operationRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return operation, nil
}

// GetAllOperations получает все операции
func (s *OperationServiceImpl) GetAllOperations() ([]*models.Operation, error) {
	return s.operationRepo.GetAll()
}

// GetOperationsByBankAccount получает операции по счету
func (s *OperationServiceImpl) GetOperationsByBankAccount(bankAccountID int) ([]*models.Operation, error) {
	return s.operationRepo.GetByBankAccountID(bankAccountID)
}

// GetOperationsByCategory получает операции по категории
func (s *OperationServiceImpl) GetOperationsByCategory(categoryID int) ([]*models.Operation, error) {
	return s.operationRepo.GetByCategoryID(categoryID)
}

// GetOperationsByDateRange получает операции в указанном диапазоне дат
func (s *OperationServiceImpl) GetOperationsByDateRange(start, end time.Time) ([]*models.Operation, error) {
	return s.operationRepo.GetByDateRange(start, end)
}

// UpdateOperation обновляет операцию
func (s *OperationServiceImpl) UpdateOperation(
	id, bankAccountID, categoryID int,
	amount float64,
	opType models.OperationType,
	date time.Time,
	description string,
) (*models.Operation, error) {
	// Получаем старую операцию
	oldOperation, err := s.operationRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Получаем старый банковский счет и откатываем баланс
	oldAccount, err := s.bankAccountRepo.GetByID(oldOperation.BankAccountID)
	if err != nil {
		return nil, err
	}

	if oldOperation.Type == models.Income {
		oldAccount.Balance -= oldOperation.Amount
	} else {
		oldAccount.Balance += oldOperation.Amount
	}

	// Проверяем новый счет
	var newAccount *models.BankAccount
	if oldOperation.BankAccountID != bankAccountID {
		newAccount, err = s.bankAccountRepo.GetByID(bankAccountID)
		if err != nil {
			return nil, err
		}
	} else {
		newAccount = oldAccount
	}

	// Проверяем новую категорию
	category, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return nil, err
	}

	// Проверяем соответствие типа категории и типа операции
	if category.Type != opType {
		return nil, &models.ValidationError{Message: "Тип категории не соответствует типу операции"}
	}

	// Обновляем операцию
	oldOperation.BankAccountID = bankAccountID
	oldOperation.CategoryID = categoryID
	oldOperation.Amount = amount
	oldOperation.Type = opType
	oldOperation.Date = date
	oldOperation.Description = description

	// Обновляем новый баланс счета
	if opType == models.Income {
		newAccount.Balance += amount
	} else {
		newAccount.Balance -= amount
	}
	newAccount.UpdatedAt = time.Now()

	// Сохраняем изменения
	err = s.operationRepo.Update(oldOperation)
	if err != nil {
		return nil, err
	}

	// Обновляем старый счет, если нужно
	if oldOperation.BankAccountID != bankAccountID {
		err = s.bankAccountRepo.Update(oldAccount)
		if err != nil {
			return nil, err
		}
	}

	// Обновляем новый счет
	err = s.bankAccountRepo.Update(newAccount)
	if err != nil {
		return nil, err
	}

	return oldOperation, nil
}

// DeleteOperation удаляет операцию
func (s *OperationServiceImpl) DeleteOperation(id int) error {
	// Получаем операцию
	operation, err := s.operationRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Получаем банковский счет
	account, err := s.bankAccountRepo.GetByID(operation.BankAccountID)
	if err != nil {
		return err
	}

	// Обновляем баланс счета
	if operation.Type == models.Income {
		account.Balance -= operation.Amount
	} else {
		account.Balance += operation.Amount
	}
	account.UpdatedAt = time.Now()

	// Удаляем операцию
	err = s.operationRepo.Delete(id)
	if err != nil {
		return err
	}

	// Обновляем счет
	return s.bankAccountRepo.Update(account)
}
