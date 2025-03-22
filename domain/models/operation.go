package models

import (
	"fmt"
	"time"
)

// Operation представляет финансовую операцию (доход или расход)
type Operation struct {
	ID            int
	Type          OperationType
	BankAccountID int
	CategoryID    int
	Amount        float64
	Date          time.Time
	Description   string
	CreatedAt     time.Time
}

// Validate проверяет валидность операции
func (o *Operation) Validate() error {
	if o.ID <= 0 {
		return &ValidationError{Message: "ID операции должен быть положительным числом"}
	}

	if o.Type != Income && o.Type != Expense {
		return &ValidationError{Message: "Тип операции должен быть INCOME или EXPENSE"}
	}

	if o.BankAccountID <= 0 {
		return &ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	if o.CategoryID <= 0 {
		return &ValidationError{Message: "ID категории должен быть положительным числом"}
	}

	if o.Amount <= 0 {
		return &ValidationError{Message: "Сумма операции должна быть положительным числом"}
	}

	return nil
}

// String возвращает строковое представление операции
func (o *Operation) String() string {
	typeStr := "Расход"
	if o.Type == Income {
		typeStr = "Доход"
	}
	return fmt.Sprintf("Операция #%d: %.2f руб. (Тип: %s, Счет: #%d, Категория: #%d, Дата: %s, Описание: %s)",
		o.ID, o.Amount, typeStr, o.BankAccountID, o.CategoryID, o.Date.Format("02.01.2006"), o.Description)
}
