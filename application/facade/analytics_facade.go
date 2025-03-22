package facade

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"fmt"
	"time"
)

// AnalyticsFacade представляет фасад для аналитики финансов
type AnalyticsFacade struct {
	analyticsService interfaces.AnalyticsService
}

// NewAnalyticsFacade создаёт новый фасад для аналитики финансов
func NewAnalyticsFacade(analyticsService interfaces.AnalyticsService) *AnalyticsFacade {
	return &AnalyticsFacade{
		analyticsService: analyticsService,
	}
}

// GetIncomeExpenseDifference получает разницу между доходами и расходами за период
func (f *AnalyticsFacade) GetIncomeExpenseDifference(start, end time.Time) (float64, error) {
	if start.After(end) {
		return 0, fmt.Errorf("дата начала не может быть позже даты окончания")
	}

	return f.analyticsService.GetIncomeExpenseDifference(start, end)
}

// GetCategorySummary получает суммарные доходы/расходы по категориям за период
func (f *AnalyticsFacade) GetCategorySummary(start, end time.Time) (map[*models.Category]float64, error) {
	if start.After(end) {
		return nil, fmt.Errorf("дата начала не может быть позже даты окончания")
	}

	return f.analyticsService.GetCategorySummary(start, end)
}

// GetMonthlyDynamics получает месячную динамику доходов и расходов за год
func (f *AnalyticsFacade) GetMonthlyDynamics(year int) (map[time.Month]map[models.OperationType]float64, error) {
	currentYear := time.Now().Year()
	if year < 2000 || year > currentYear+1 {
		return nil, fmt.Errorf("некорректный год (должен быть в диапазоне от 2000 до %d)", currentYear+1)
	}

	return f.analyticsService.GetMonthlyDynamics(year)
}
