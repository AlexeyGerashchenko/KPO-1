package commands

import (
	"KPO1/application/facade"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
)

// CreateBankAccountCommand представляет команду для создания банковского счёта
type CreateBankAccountCommand struct {
	CommandBase
	facade   *facade.BankAccountFacade
	name     string
	resultCh chan *models.BankAccount
	errorCh  chan error
}

// NewCreateBankAccountCommand создаёт новую команду для создания банковского счёта
func NewCreateBankAccountCommand(
	facade *facade.BankAccountFacade,
	name string,
	resultCh chan *models.BankAccount,
	errorCh chan error,
) interfaces.Command {
	return &CreateBankAccountCommand{
		CommandBase: NewCommandBase("CreateBankAccount"),
		facade:      facade,
		name:        name,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду
func (c *CreateBankAccountCommand) Execute() error {
	account, err := c.facade.CreateBankAccount(c.name)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- account
	}

	return nil
}

// GetBankAccountCommand представляет команду для получения банковского счёта
type GetBankAccountCommand struct {
	CommandBase
	facade   *facade.BankAccountFacade
	id       int
	resultCh chan *models.BankAccount
	errorCh  chan error
}

// NewGetBankAccountCommand создаёт новую команду для получения банковского счёта
func NewGetBankAccountCommand(
	facade *facade.BankAccountFacade,
	id int,
	resultCh chan *models.BankAccount,
	errorCh chan error,
) interfaces.Command {
	return &GetBankAccountCommand{
		CommandBase: NewCommandBase("GetBankAccount"),
		facade:      facade,
		id:          id,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду
func (c *GetBankAccountCommand) Execute() error {
	account, err := c.facade.GetBankAccount(c.id)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- account
	}

	return nil
}

// ListBankAccountsCommand представляет команду для получения списка банковских счетов
type ListBankAccountsCommand struct {
	CommandBase
	facade   *facade.BankAccountFacade
	resultCh chan []*models.BankAccount
	errorCh  chan error
}

// NewListBankAccountsCommand создаёт новую команду для получения списка банковских счетов
func NewListBankAccountsCommand(
	facade *facade.BankAccountFacade,
	resultCh chan []*models.BankAccount,
	errorCh chan error,
) interfaces.Command {
	return &ListBankAccountsCommand{
		CommandBase: NewCommandBase("ListBankAccounts"),
		facade:      facade,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду
func (c *ListBankAccountsCommand) Execute() error {
	accounts, err := c.facade.GetAllBankAccounts()
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- accounts
	}

	return nil
}

// UpdateBankAccountCommand представляет команду для обновления банковского счёта
type UpdateBankAccountCommand struct {
	CommandBase
	facade   *facade.BankAccountFacade
	id       int
	name     string
	resultCh chan *models.BankAccount
	errorCh  chan error
}

// NewUpdateBankAccountCommand создаёт новую команду для обновления банковского счёта
func NewUpdateBankAccountCommand(
	facade *facade.BankAccountFacade,
	id int,
	name string,
	resultCh chan *models.BankAccount,
	errorCh chan error,
) interfaces.Command {
	return &UpdateBankAccountCommand{
		CommandBase: NewCommandBase("UpdateBankAccount"),
		facade:      facade,
		id:          id,
		name:        name,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду
func (c *UpdateBankAccountCommand) Execute() error {
	account, err := c.facade.UpdateBankAccount(c.id, c.name)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- account
	}

	return nil
}

// DeleteBankAccountCommand представляет команду для удаления банковского счёта
type DeleteBankAccountCommand struct {
	CommandBase
	facade  *facade.BankAccountFacade
	id      int
	errorCh chan error
}

// NewDeleteBankAccountCommand создаёт новую команду для удаления банковского счёта
func NewDeleteBankAccountCommand(
	facade *facade.BankAccountFacade,
	id int,
	errorCh chan error,
) interfaces.Command {
	return &DeleteBankAccountCommand{
		CommandBase: NewCommandBase("DeleteBankAccount"),
		facade:      facade,
		id:          id,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду
func (c *DeleteBankAccountCommand) Execute() error {
	err := c.facade.DeleteBankAccount(c.id)
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}

	return err
}
