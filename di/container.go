package di

import (
	"KPO1/application/analytics"
	"KPO1/application/facade"
	"KPO1/application/services"
	"KPO1/domain/factory"
	"KPO1/domain/interfaces"
	"KPO1/infrastructure/persistence"
	"sync"
)

// Container представляет контейнер для внедрения зависимостей
type Container struct {
	// Репозитории
	memoryRepository      *persistence.MemoryRepository
	bankAccountRepository interfaces.BankAccountRepository
	categoryRepository    interfaces.CategoryRepository
	operationRepository   interfaces.OperationRepository

	// Фабрики
	bankAccountFactory *factory.BankAccountFactory
	categoryFactory    *factory.CategoryFactory
	operationFactory   *factory.OperationFactory

	// Сервисы
	bankAccountService interfaces.BankAccountService
	categoryService    interfaces.CategoryService
	operationService   interfaces.OperationService
	analyticsService   interfaces.AnalyticsService

	// Фасады
	bankAccountFacade *facade.BankAccountFacade
	categoryFacade    *facade.CategoryFacade
	operationFacade   *facade.OperationFacade
	analyticsFacade   *facade.AnalyticsFacade

	// мьютекс для потокобезопасности
	mu sync.Mutex
}

// NewContainer создает новый контейнер для внедрения зависимостей
func NewContainer() *Container {
	return &Container{}
}

// GetMemoryRepository возвращает репозиторий в памяти
func (c *Container) GetMemoryRepository() *persistence.MemoryRepository {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.memoryRepository == nil {
		c.memoryRepository = persistence.NewMemoryRepository()
	}

	return c.memoryRepository
}

// GetBankAccountRepository возвращает репозиторий банковских счетов
func (c *Container) GetBankAccountRepository() interfaces.BankAccountRepository {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.bankAccountRepository == nil {
		c.bankAccountRepository = persistence.NewBankAccountRepository(c.GetMemoryRepository())
	}

	return c.bankAccountRepository
}

// GetCategoryRepository возвращает репозиторий категорий
func (c *Container) GetCategoryRepository() interfaces.CategoryRepository {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.categoryRepository == nil {
		c.categoryRepository = persistence.NewCategoryRepository(c.GetMemoryRepository())
	}

	return c.categoryRepository
}

// GetOperationRepository возвращает репозиторий операций
func (c *Container) GetOperationRepository() interfaces.OperationRepository {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.operationRepository == nil {
		c.operationRepository = persistence.NewOperationRepository(c.GetMemoryRepository())
	}

	return c.operationRepository
}

// GetBankAccountFactory возвращает фабрику банковских счетов
func (c *Container) GetBankAccountFactory() *factory.BankAccountFactory {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.bankAccountFactory == nil {
		c.bankAccountFactory = factory.NewBankAccountFactory()
	}

	return c.bankAccountFactory
}

// GetCategoryFactory возвращает фабрику категорий
func (c *Container) GetCategoryFactory() *factory.CategoryFactory {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.categoryFactory == nil {
		c.categoryFactory = factory.NewCategoryFactory()
	}

	return c.categoryFactory
}

// GetOperationFactory возвращает фабрику операций
func (c *Container) GetOperationFactory() *factory.OperationFactory {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.operationFactory == nil {
		c.operationFactory = factory.NewOperationFactory()
	}

	return c.operationFactory
}

// GetBankAccountService возвращает сервис для управления банковскими счетами
func (c *Container) GetBankAccountService() interfaces.BankAccountService {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.bankAccountService == nil {
		c.bankAccountService = services.NewBankAccountService(
			c.GetBankAccountRepository(),
			c.GetOperationRepository(),
			c.GetBankAccountFactory(),
		)
	}

	return c.bankAccountService
}

// GetCategoryService возвращает сервис для управления категориями
func (c *Container) GetCategoryService() interfaces.CategoryService {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.categoryService == nil {
		c.categoryService = services.NewCategoryService(
			c.GetCategoryRepository(),
			c.GetOperationRepository(),
			c.GetCategoryFactory(),
		)
	}

	return c.categoryService
}

// GetOperationService возвращает сервис для управления операциями
func (c *Container) GetOperationService() interfaces.OperationService {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.operationService == nil {
		c.operationService = services.NewOperationService(
			c.GetOperationRepository(),
			c.GetBankAccountRepository(),
			c.GetCategoryRepository(),
			c.GetOperationFactory(),
		)
	}

	return c.operationService
}

// GetAnalyticsService возвращает сервис для аналитики финансов
func (c *Container) GetAnalyticsService() interfaces.AnalyticsService {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.analyticsService == nil {
		c.analyticsService = analytics.NewAnalyticsService(
			c.GetOperationRepository(),
			c.GetCategoryRepository(),
		)
	}

	return c.analyticsService
}

// GetBankAccountFacade возвращает фасад для работы с банковскими счетами
func (c *Container) GetBankAccountFacade() *facade.BankAccountFacade {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.bankAccountFacade == nil {
		c.bankAccountFacade = facade.NewBankAccountFacade(
			c.GetBankAccountService(),
		)
	}

	return c.bankAccountFacade
}

// GetCategoryFacade возвращает фасад для работы с категориями
func (c *Container) GetCategoryFacade() *facade.CategoryFacade {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.categoryFacade == nil {
		c.categoryFacade = facade.NewCategoryFacade(
			c.GetCategoryService(),
		)
	}

	return c.categoryFacade
}

// GetOperationFacade возвращает фасад для работы с операциями
func (c *Container) GetOperationFacade() *facade.OperationFacade {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.operationFacade == nil {
		c.operationFacade = facade.NewOperationFacade(
			c.GetOperationService(),
			c.GetCategoryService(),
		)
	}

	return c.operationFacade
}

// GetAnalyticsFacade возвращает фасад для аналитики
func (c *Container) GetAnalyticsFacade() *facade.AnalyticsFacade {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.analyticsFacade == nil {
		c.analyticsFacade = facade.NewAnalyticsFacade(
			c.GetAnalyticsService(),
		)
	}

	return c.analyticsFacade
}
