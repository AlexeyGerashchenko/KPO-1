package factory

import (
	"KPO1/domain/models"
	"time"
)

// CategoryFactory представляет фабрику для создания категорий
type CategoryFactory struct {
	nextID int
}

// NewCategoryFactory создаёт новую фабрику категорий
func NewCategoryFactory() *CategoryFactory {
	return &CategoryFactory{
		nextID: 1,
	}
}

// CreateCategory создаёт новую категорию
func (f *CategoryFactory) CreateCategory(name string, opType models.OperationType) (*models.Category, error) {
	now := time.Now()
	category := &models.Category{
		ID:        f.nextID,
		Type:      opType,
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Валидация категории
	if err := category.Validate(); err != nil {
		return nil, err
	}

	f.nextID++
	return category, nil
}

// SetNextID устанавливает следующий ID для фабрики
func (f *CategoryFactory) SetNextID(id int) {
	if id > f.nextID {
		f.nextID = id
	}
}
