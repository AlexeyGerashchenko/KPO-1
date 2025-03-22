package commands

import (
	"KPO1/application/facade"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// UpdateOperationCommand представляет команду для обновления операции
type UpdateOperationCommand struct {
	CommandBase
	facade        *facade.OperationFacade
	id            int
	bankAccountID int
	categoryID    int
	amount        float64
	opType        models.OperationType
	date          time.Time
	description   string
	resultCh      chan *models.Operation
	errorCh       chan error
}

// NewUpdateOperationCommand создает новую команду для обновления операции
func NewUpdateOperationCommand(
	facade *facade.OperationFacade,
	id int,
	bankAccountID int,
	categoryID int,
	amount float64,
	opType models.OperationType,
	date time.Time,
	description string,
	resultCh chan *models.Operation,
	errorCh chan error,
) interfaces.Command {
	return &UpdateOperationCommand{
		CommandBase:   NewCommandBase("UpdateOperation"),
		facade:        facade,
		id:            id,
		bankAccountID: bankAccountID,
		categoryID:    categoryID,
		amount:        amount,
		opType:        opType,
		date:          date,
		description:   description,
		resultCh:      resultCh,
		errorCh:       errorCh,
	}
}

// Execute выполняет команду обновления операции
func (c *UpdateOperationCommand) Execute() error {
	operation, err := c.facade.UpdateOperation(
		c.id,
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
