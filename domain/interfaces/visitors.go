package interfaces

import (
	"KPO1/domain/models"
)

// DataVisitor представляет интерфейс для посетителя данных
type DataVisitor interface {
	VisitBankAccount(account models.BankAccount) error
	VisitCategory(category models.Category) error
	VisitOperation(operation models.Operation) error
}

// Visitable представляет интерфейс для объектов, которые можно посещать
type Visitable interface {
	Accept(visitor DataVisitor) error
}
