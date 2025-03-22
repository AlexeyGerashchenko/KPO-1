package analytics

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"time"
)

// AnalyticsServiceImpl реализация сервиса для аналитики финансов
type AnalyticsServiceImpl struct {
	operationRepo interfaces.OperationRepository
	categoryRepo  interfaces.CategoryRepository
}

// NewAnalyticsService создаёт новый сервис для аналитики финансов
func NewAnalyticsService(
	operationRepo interfaces.OperationRepository,
	categoryRepo interfaces.CategoryRepository,
) interfaces.AnalyticsService {
	return &AnalyticsServiceImpl{
		operationRepo: operationRepo,
		categoryRepo:  categoryRepo,
	}
}

// GetIncomeExpenseDifference рассчитывает разницу между доходами и расходами за период
func (s *AnalyticsServiceImpl) GetIncomeExpenseDifference(start, end time.Time) (float64, error) {
	operations, err := s.operationRepo.GetByDateRange(start, end)
	if err != nil {
		return 0, err
	}

	var income, expense float64
	for _, op := range operations {
		if op.Type == models.Income {
			income += op.Amount
		} else {
			expense += op.Amount
		}
	}

	return income - expense, nil
}

// GetCategorySummary получает сумму операций по каждой категории за период
func (s *AnalyticsServiceImpl) GetCategorySummary(start, end time.Time) (map[*models.Category]float64, error) {
	operations, err := s.operationRepo.GetByDateRange(start, end)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[int]*models.Category)
	categories, err := s.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, cat := range categories {
		categoryMap[cat.ID] = cat
	}

	result := make(map[*models.Category]float64)
	for _, op := range operations {
		category, exists := categoryMap[op.CategoryID]
		if !exists {
			continue
		}

		if _, ok := result[category]; !ok {
			result[category] = 0
		}
		result[category] += op.Amount
	}

	return result, nil
}

// GetMonthlyDynamics получает месячную динамику доходов и расходов за год
func (s *AnalyticsServiceImpl) GetMonthlyDynamics(year int) (map[time.Month]map[models.OperationType]float64, error) {
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	end := time.Date(year, 12, 31, 23, 59, 59, 999999999, time.Local)

	operations, err := s.operationRepo.GetByDateRange(start, end)
	if err != nil {
		return nil, err
	}

	result := make(map[time.Month]map[models.OperationType]float64)
	for i := time.January; i <= time.December; i++ {
		result[i] = map[models.OperationType]float64{
			models.Income:  0,
			models.Expense: 0,
		}
	}

	for _, op := range operations {
		month := op.Date.Month()
		result[month][op.Type] += op.Amount
	}

	return result, nil
}
