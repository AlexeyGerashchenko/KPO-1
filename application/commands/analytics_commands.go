package commands

import (
	"KPO1/application/facade"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// BalanceByPeriodCommand представляет команду для получения баланса за период
type BalanceByPeriodCommand struct {
	CommandBase
	facade    *facade.AnalyticsFacade
	startDate time.Time
	endDate   time.Time
	resultCh  chan float64
	errorCh   chan error
}

// NewBalanceByPeriodCommand создаёт новую команду для получения баланса за период
func NewBalanceByPeriodCommand(
	facade *facade.AnalyticsFacade,
	startDate time.Time,
	endDate time.Time,
	resultCh chan float64,
	errorCh chan error,
) interfaces.Command {
	return &BalanceByPeriodCommand{
		CommandBase: NewCommandBase("BalanceByPeriod"),
		facade:      facade,
		startDate:   startDate,
		endDate:     endDate,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения баланса за период
func (c *BalanceByPeriodCommand) Execute() error {
	balance, err := c.facade.GetIncomeExpenseDifference(c.startDate, c.endDate)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	if c.resultCh != nil {
		c.resultCh <- balance
	}
	return nil
}

// ExpensesByCategoryCommand представляет команду для получения расходов по категориям
type ExpensesByCategoryCommand struct {
	CommandBase
	facade    *facade.AnalyticsFacade
	startDate time.Time
	endDate   time.Time
	resultCh  chan map[string]float64
	errorCh   chan error
}

// NewExpensesByCategoryCommand создаёт новую команду для получения расходов по категориям
func NewExpensesByCategoryCommand(
	facade *facade.AnalyticsFacade,
	startDate time.Time,
	endDate time.Time,
	resultCh chan map[string]float64,
	errorCh chan error,
) interfaces.Command {
	return &ExpensesByCategoryCommand{
		CommandBase: NewCommandBase("ExpensesByCategory"),
		facade:      facade,
		startDate:   startDate,
		endDate:     endDate,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения расходов по категориям
func (c *ExpensesByCategoryCommand) Execute() error {
	categoryMap, err := c.facade.GetCategorySummary(c.startDate, c.endDate)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	// Преобразуем результаты для расходов
	expenses := make(map[string]float64)
	for category, amount := range categoryMap {
		if category.Type == models.Expense {
			expenses[category.Name] = amount
		}
	}

	if c.resultCh != nil {
		c.resultCh <- expenses
	}
	return nil
}

// IncomesByCategoryCommand представляет команду для получения доходов по категориям
type IncomesByCategoryCommand struct {
	CommandBase
	facade    *facade.AnalyticsFacade
	startDate time.Time
	endDate   time.Time
	resultCh  chan map[string]float64
	errorCh   chan error
}

// NewIncomesByCategoryCommand создаёт новую команду для получения доходов по категориям
func NewIncomesByCategoryCommand(
	facade *facade.AnalyticsFacade,
	startDate time.Time,
	endDate time.Time,
	resultCh chan map[string]float64,
	errorCh chan error,
) interfaces.Command {
	return &IncomesByCategoryCommand{
		CommandBase: NewCommandBase("IncomesByCategory"),
		facade:      facade,
		startDate:   startDate,
		endDate:     endDate,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения доходов по категориям
func (c *IncomesByCategoryCommand) Execute() error {
	categoryMap, err := c.facade.GetCategorySummary(c.startDate, c.endDate)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	// Преобразуем результаты для доходов
	incomes := make(map[string]float64)
	for category, amount := range categoryMap {
		if category.Type == models.Income {
			incomes[category.Name] = amount
		}
	}

	if c.resultCh != nil {
		c.resultCh <- incomes
	}
	return nil
}

// StatisticsCommand представляет команду для получения общей статистики
type StatisticsCommand struct {
	CommandBase
	facade    *facade.AnalyticsFacade
	startDate time.Time
	endDate   time.Time
	resultCh  chan map[string]interface{}
	errorCh   chan error
}

// NewStatisticsCommand создаёт новую команду для получения общей статистики
func NewStatisticsCommand(
	facade *facade.AnalyticsFacade,
	startDate time.Time,
	endDate time.Time,
	resultCh chan map[string]interface{},
	errorCh chan error,
) interfaces.Command {
	return &StatisticsCommand{
		CommandBase: NewCommandBase("Statistics"),
		facade:      facade,
		startDate:   startDate,
		endDate:     endDate,
		resultCh:    resultCh,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду получения общей статистики
func (c *StatisticsCommand) Execute() error {
	// Получаем разницу между доходами и расходами
	difference, err := c.facade.GetIncomeExpenseDifference(c.startDate, c.endDate)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	// Получаем суммарные доходы и расходы по категориям
	categorySummary, err := c.facade.GetCategorySummary(c.startDate, c.endDate)
	if err != nil {
		if c.errorCh != nil {
			c.errorCh <- err
		}
		return err
	}

	// Преобразуем данные о категориях
	incomes := make(map[string]float64)
	expenses := make(map[string]float64)
	for category, amount := range categorySummary {
		if category.Type == models.Income {
			incomes[category.Name] = amount
		} else {
			expenses[category.Name] = amount
		}
	}

	// Подготавливаем результат
	result := map[string]interface{}{
		"difference": difference,
		"incomes":    incomes,
		"expenses":   expenses,
	}

	if c.resultCh != nil {
		c.resultCh <- result
	}
	return nil
}
