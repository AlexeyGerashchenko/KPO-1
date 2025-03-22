package commands

import (
	"KPO1/application/facade"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// CreateOperationCommand представляет команду для создания операции
type CreateOperationCommand struct {
	CommandBase
	facade        *facade.OperationFacade
	operationType models.OperationType
	bankAccountID int
	categoryID    int
	amount        float64
	date          time.Time
	description   string
	resultCh      chan *models.Operation
	errorCh       chan error
}

// NewCreateOperationCommand создаёт новую команду для создания операции
func NewCreateOperationCommand(
	facade *facade.OperationFacade,
	operationType models.OperationType,
	bankAccountID int,
	categoryID int,
	amount float64,
	date time.Time,
	description string,
	resultCh chan *models.Operation,
	errorCh chan error,
) interfaces.Command {
	return &CreateOperationCommand{
		CommandBase:   NewCommandBase("CreateOperation"),
		facade:        facade,
		operationType: operationType,
		bankAccountID: bankAccountID,
		categoryID:    categoryID,
		amount:        amount,
		date:          date,
		description:   description,
		resultCh:      resultCh,
		errorCh:       errorCh,
	}
}

// Execute выполняет команду создания операции
func (c *CreateOperationCommand) Execute() error {
	operation, err := c.facade.CreateOperation(
		c.bankAccountID,
		c.categoryID,
		c.amount,
		c.date,
		c.description,
	)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- operation
	}
	return nil
}

// GetOperationCommand представляет команду для получения операции по ID
type GetOperationCommand struct {
	CommandBase
	facade   *facade.OperationFacade
	id       int
	resultCh chan *models.Operation
	errorCh  chan error
}

// NewGetOperationCommand создаёт новую команду для получения операции
func NewGetOperationCommand(
	facade *facade.OperationFacade,
	id int,
	resultCh chan *models.Operation,
	errorCh chan error,
) interfaces.Command {
	return &GetOperationCommand{
		CommandBase: NewCommandBase("GetOperation"),
		facade:      facade,
		id:          id,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения операции
func (c *GetOperationCommand) Execute() error {
	operation, err := c.facade.GetOperationDetails(c.id)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- operation
	}
	return nil
}

// ListOperationsCommand представляет команду для получения списка всех операций
type ListOperationsCommand struct {
	CommandBase
	facade   *facade.OperationFacade
	resultCh chan []*models.Operation
	errorCh  chan error
}

// NewListOperationsCommand создаёт новую команду для получения списка операций
func NewListOperationsCommand(
	facade *facade.OperationFacade,
	resultCh chan []*models.Operation,
	errorCh chan error,
) interfaces.Command {
	return &ListOperationsCommand{
		CommandBase: NewCommandBase("ListOperations"),
		facade:      facade,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения списка операций
func (c *ListOperationsCommand) Execute() error {
	operations, err := c.facade.GetAllOperations()
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- operations
	}
	return nil
}

// ListOperationsByAccountCommand представляет команду для получения списка операций по счету
type ListOperationsByAccountCommand struct {
	CommandBase
	facade        *facade.OperationFacade
	bankAccountID int
	resultCh      chan []*models.Operation
	errorCh       chan error
}

// NewListOperationsByAccountCommand создаёт новую команду для получения списка операций по счету
func NewListOperationsByAccountCommand(
	facade *facade.OperationFacade,
	bankAccountID int,
	resultCh chan []*models.Operation,
	errorCh chan error,
) interfaces.Command {
	return &ListOperationsByAccountCommand{
		CommandBase:   NewCommandBase("ListOperationsByAccount"),
		facade:        facade,
		bankAccountID: bankAccountID,
		resultCh:      resultCh,
		errorCh:       errorCh,
	}
}

// Execute выполняет команду получения списка операций по счету
func (c *ListOperationsByAccountCommand) Execute() error {
	operations, err := c.facade.GetOperationsByBankAccount(c.bankAccountID)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- operations
	}
	return nil
}

// ListOperationsByCategoryCommand представляет команду для получения списка операций по категории
type ListOperationsByCategoryCommand struct {
	CommandBase
	facade     *facade.OperationFacade
	categoryID int
	resultCh   chan []*models.Operation
	errorCh    chan error
}

// NewListOperationsByCategoryCommand создаёт новую команду для получения списка операций по категории
func NewListOperationsByCategoryCommand(
	facade *facade.OperationFacade,
	categoryID int,
	resultCh chan []*models.Operation,
	errorCh chan error,
) interfaces.Command {
	return &ListOperationsByCategoryCommand{
		CommandBase: NewCommandBase("ListOperationsByCategory"),
		facade:      facade,
		categoryID:  categoryID,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения списка операций по категории
func (c *ListOperationsByCategoryCommand) Execute() error {
	operations, err := c.facade.GetOperationsByCategory(c.categoryID)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- operations
	}
	return nil
}

// DeleteOperationCommand представляет команду для удаления операции
type DeleteOperationCommand struct {
	CommandBase
	facade  *facade.OperationFacade
	id      int
	errorCh chan error
}

// NewDeleteOperationCommand создаёт новую команду для удаления операции
func NewDeleteOperationCommand(
	facade *facade.OperationFacade,
	id int,
	errorCh chan error,
) interfaces.Command {
	return &DeleteOperationCommand{
		CommandBase: NewCommandBase("DeleteOperation"),
		facade:      facade,
		id:          id,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду удаления операции
func (c *DeleteOperationCommand) Execute() error {
	err := c.facade.DeleteOperation(c.id)
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}
