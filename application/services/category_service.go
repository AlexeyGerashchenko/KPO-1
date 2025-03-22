package services

import (
	"KPO1/domain/factory"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"errors"
	"time"
)

// CategoryServiceImpl реализация сервиса для управления категориями
type CategoryServiceImpl struct {
	categoryRepo  interfaces.CategoryRepository
	operationRepo interfaces.OperationRepository
	factory       *factory.CategoryFactory
}

// NewCategoryService создаёт новый сервис для управления категориями
func NewCategoryService(
	categoryRepo interfaces.CategoryRepository,
	operationRepo interfaces.OperationRepository,
	factory *factory.CategoryFactory,
) interfaces.CategoryService {
	return &CategoryServiceImpl{
		categoryRepo:  categoryRepo,
		operationRepo: operationRepo,
		factory:       factory,
	}
}

// CreateCategory создает новую категорию
func (s *CategoryServiceImpl) CreateCategory(name string, opType models.OperationType) (*models.Category, error) {
	category, err := s.factory.CreateCategory(name, opType)
	if err != nil {
		return nil, err
	}

	err = s.categoryRepo.Save(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// GetCategory получает категорию по ID
func (s *CategoryServiceImpl) GetCategory(id int) (*models.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

// GetAllCategories получает все категории
func (s *CategoryServiceImpl) GetAllCategories() ([]*models.Category, error) {
	return s.categoryRepo.GetAll()
}

// GetCategoriesByType получает категории по типу операции
func (s *CategoryServiceImpl) GetCategoriesByType(opType models.OperationType) ([]*models.Category, error) {
	return s.categoryRepo.GetByType(opType)
}

// UpdateCategory обновляет категорию
func (s *CategoryServiceImpl) UpdateCategory(id int, name string, opType models.OperationType) (*models.Category, error) {
	category, err := s.categoryRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	category.Name = name
	category.Type = opType
	category.UpdatedAt = time.Now()

	if err := category.Validate(); err != nil {
		return nil, err
	}

	err = s.categoryRepo.Update(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory удаляет категорию
func (s *CategoryServiceImpl) DeleteCategory(id int) error {
	// Проверяем наличие операций с этой категорией
	operations, err := s.operationRepo.GetByCategoryID(id)
	if err != nil {
		return err
	}

	if len(operations) > 0 {
		return errors.New("нельзя удалить категорию, по которой есть операции")
	}

	return s.categoryRepo.Delete(id)
}
