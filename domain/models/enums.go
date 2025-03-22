package models

// OperationType представляет тип операции (доход или расход)
type OperationType string

const (
	Income  OperationType = "INCOME"
	Expense OperationType = "EXPENSE"
)

// ValidationError представляет ошибку валидации
type ValidationError struct {
	Message string
}

func (e ValidationError) Error() string {
	return e.Message
}
