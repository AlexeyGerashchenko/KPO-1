package models

import (
	"fmt"
	"time"
)

// BankAccount представляет банковский счёт пользователя
type BankAccount struct {
	ID        int
	Name      string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate проверяет валидность банковского счёта
func (b *BankAccount) Validate() error {
	if b.Name == "" {
		return &ValidationError{Message: "Название счета не может быть пустым"}
	}

	if b.ID <= 0 {
		return &ValidationError{Message: "ID счета должен быть положительным числом"}
	}

	return nil
}

// String возвращает строковое представление банковского счёта
func (b *BankAccount) String() string {
	return fmt.Sprintf("Счет #%d: %s (Баланс: %.2f руб.)", b.ID, b.Name, b.Balance)
}
