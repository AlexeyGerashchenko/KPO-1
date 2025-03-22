package commands

import (
	"KPO1/application/facade"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
)

// CreateCategoryCommand представляет команду для создания категории
type CreateCategoryCommand struct {
	CommandBase
	facade       *facade.CategoryFacade
	name         string
	categoryType models.OperationType
	resultCh     chan *models.Category
	errorCh      chan error
}

// NewCreateCategoryCommand создаёт новую команду для создания категории
func NewCreateCategoryCommand(
	facade *facade.CategoryFacade,
	name string,
	categoryType models.OperationType,
	resultCh chan *models.Category,
	errorCh chan error,
) interfaces.Command {
	return &CreateCategoryCommand{
		CommandBase:  NewCommandBase("CreateCategory"),
		facade:       facade,
		name:         name,
		categoryType: categoryType,
		resultCh:     resultCh,
		errorCh:      errorCh,
	}
}

// Execute выполняет команду создания категории
func (c *CreateCategoryCommand) Execute() error {
	category, err := c.facade.CreateCategory(c.name, c.categoryType)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- category
	}
	return nil
}

// GetCategoryCommand представляет команду для получения категории по ID
type GetCategoryCommand struct {
	CommandBase
	facade   *facade.CategoryFacade
	id       int
	resultCh chan *models.Category
	errorCh  chan error
}

// NewGetCategoryCommand создаёт новую команду для получения категории
func NewGetCategoryCommand(
	facade *facade.CategoryFacade,
	id int,
	resultCh chan *models.Category,
	errorCh chan error,
) interfaces.Command {
	return &GetCategoryCommand{
		CommandBase: NewCommandBase("GetCategory"),
		facade:      facade,
		id:          id,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения категории
func (c *GetCategoryCommand) Execute() error {
	category, err := c.facade.GetCategory(c.id)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- category
	}
	return nil
}

// ListCategoriesCommand представляет команду для получения списка всех категорий
type ListCategoriesCommand struct {
	CommandBase
	facade   *facade.CategoryFacade
	resultCh chan []*models.Category
	errorCh  chan error
}

// NewListCategoriesCommand создаёт новую команду для получения списка категорий
func NewListCategoriesCommand(
	facade *facade.CategoryFacade,
	resultCh chan []*models.Category,
	errorCh chan error,
) interfaces.Command {
	return &ListCategoriesCommand{
		CommandBase: NewCommandBase("ListCategories"),
		facade:      facade,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения списка категорий
func (c *ListCategoriesCommand) Execute() error {
	categories, err := c.facade.GetAllCategories()
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- categories
	}
	return nil
}

// ListCategoriesByTypeCommand представляет команду для получения списка категорий по типу
type ListCategoriesByTypeCommand struct {
	CommandBase
	facade       *facade.CategoryFacade
	categoryType models.OperationType
	resultCh     chan []*models.Category
	errorCh      chan error
}

// NewListCategoriesByTypeCommand создаёт новую команду для получения списка категорий по типу
func NewListCategoriesByTypeCommand(
	facade *facade.CategoryFacade,
	categoryType models.OperationType,
	resultCh chan []*models.Category,
	errorCh chan error,
) interfaces.Command {
	return &ListCategoriesByTypeCommand{
		CommandBase:  NewCommandBase("ListCategoriesByType"),
		facade:       facade,
		categoryType: categoryType,
		resultCh:     resultCh,
		errorCh:      errorCh,
	}
}

// Execute выполняет команду получения списка категорий по типу
func (c *ListCategoriesByTypeCommand) Execute() error {
	categories, err := c.facade.GetCategoriesByType(c.categoryType)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- categories
	}
	return nil
}

// UpdateCategoryCommand представляет команду для обновления категории
type UpdateCategoryCommand struct {
	CommandBase
	facade       *facade.CategoryFacade
	id           int
	name         string
	categoryType models.OperationType
	resultCh     chan *models.Category
	errorCh      chan error
}

// NewUpdateCategoryCommand создаёт новую команду для обновления категории
func NewUpdateCategoryCommand(
	facade *facade.CategoryFacade,
	id int,
	name string,
	categoryType models.OperationType,
	resultCh chan *models.Category,
	errorCh chan error,
) interfaces.Command {
	return &UpdateCategoryCommand{
		CommandBase:  NewCommandBase("UpdateCategory"),
		facade:       facade,
		id:           id,
		name:         name,
		categoryType: categoryType,
		resultCh:     resultCh,
		errorCh:      errorCh,
	}
}

// Execute выполняет команду обновления категории
func (c *UpdateCategoryCommand) Execute() error {
	category, err := c.facade.UpdateCategory(c.id, c.name, c.categoryType)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- category
	}
	return nil
}

// DeleteCategoryCommand представляет команду для удаления категории
type DeleteCategoryCommand struct {
	CommandBase
	facade  *facade.CategoryFacade
	id      int
	errorCh chan error
}

// NewDeleteCategoryCommand создаёт новую команду для удаления категории
func NewDeleteCategoryCommand(
	facade *facade.CategoryFacade,
	id int,
	errorCh chan error,
) interfaces.Command {
	return &DeleteCategoryCommand{
		CommandBase: NewCommandBase("DeleteCategory"),
		facade:      facade,
		id:          id,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду удаления категории
func (c *DeleteCategoryCommand) Execute() error {
	err := c.facade.DeleteCategory(c.id)
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}
