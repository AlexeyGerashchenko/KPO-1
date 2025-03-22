package models

import (
	"fmt"
	"time"
)

// Category представляет категорию доходов или расходов
type Category struct {
	ID        int
	Type      OperationType
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate проверяет валидность категории
func (c *Category) Validate() error {
	if c.Name == "" {
		return &ValidationError{Message: "Название категории не может быть пустым"}
	}

	if c.ID <= 0 {
		return &ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	if c.Type != Income && c.Type != Expense {
		return &ValidationError{Message: "Тип категории должен быть INCOME или EXPENSE"}
	}

	return nil
}

// String возвращает строковое представление категории
func (c *Category) String() string {
	typeStr := "Расход"
	if c.Type == Income {
		typeStr = "Доход"
	}
	return fmt.Sprintf("Категория #%d: %s (Тип: %s)", c.ID, c.Name, typeStr)
}
