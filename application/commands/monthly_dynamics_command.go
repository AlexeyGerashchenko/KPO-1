package commands

import (
	"KPO1/application/facade"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// MonthlyDynamicsCommand представляет команду для получения месячной динамики
type MonthlyDynamicsCommand struct {
	CommandBase
	facade   *facade.AnalyticsFacade
	year     int
	resultCh chan map[time.Month]map[models.OperationType]float64
	errorCh  chan error
}

// NewMonthlyDynamicsCommand создает новую команду для получения месячной динамики
func NewMonthlyDynamicsCommand(
	facade *facade.AnalyticsFacade,
	year int,
	resultCh chan map[time.Month]map[models.OperationType]float64,
	errorCh chan error,
) interfaces.Command {
	return &MonthlyDynamicsCommand{
		CommandBase: NewCommandBase("MonthlyDynamics"),
		facade:      facade,
		year:        year,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения месячной динамики
func (c *MonthlyDynamicsCommand) Execute() error {
	dynamics, err := c.facade.GetMonthlyDynamics(c.year)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}
	if c.resultCh != nil {
		c.resultCh <- dynamics
	}
	return nil
}
