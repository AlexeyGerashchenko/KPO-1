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

	// мьютексы для потокобезопасности
	repoMu    sync.Mutex
	factoryMu sync.Mutex
	serviceMu sync.Mutex
	facadeMu  sync.Mutex
}

// NewContainer создает новый контейнер для внедрения зависимостей
func NewContainer() *Container {
	return &Container{}
}

// GetMemoryRepository возвращает репозиторий в памяти
func (c *Container) GetMemoryRepository() *persistence.MemoryRepository {
	c.repoMu.Lock()
	defer c.repoMu.Unlock()

	if c.memoryRepository == nil {
		c.memoryRepository = persistence.NewMemoryRepository()
	}

	return c.memoryRepository
}

// GetBankAccountRepository возвращает репозиторий банковских счетов
func (c *Container) GetBankAccountRepository() interfaces.BankAccountRepository {
	c.repoMu.Lock()
	defer c.repoMu.Unlock()

	if c.bankAccountRepository == nil {
		// Используем напрямую memoryRepository вместо вызова GetMemoryRepository()
		if c.memoryRepository == nil {
			c.memoryRepository = persistence.NewMemoryRepository()
		}
		memRepo := c.memoryRepository

		c.bankAccountRepository = persistence.NewBankAccountRepository(memRepo)
	}

	return c.bankAccountRepository
}

// GetCategoryRepository возвращает репозиторий категорий
func (c *Container) GetCategoryRepository() interfaces.CategoryRepository {
	c.repoMu.Lock()
	defer c.repoMu.Unlock()

	if c.categoryRepository == nil {
		// Используем напрямую memoryRepository вместо вызова GetMemoryRepository()
		if c.memoryRepository == nil {
			c.memoryRepository = persistence.NewMemoryRepository()
		}
		memRepo := c.memoryRepository

		c.categoryRepository = persistence.NewCategoryRepository(memRepo)
	}

	return c.categoryRepository
}

// GetOperationRepository возвращает репозиторий операций
func (c *Container) GetOperationRepository() interfaces.OperationRepository {
	c.repoMu.Lock()
	defer c.repoMu.Unlock()

	if c.operationRepository == nil {
		// Используем напрямую memoryRepository вместо вызова GetMemoryRepository()
		if c.memoryRepository == nil {
			c.memoryRepository = persistence.NewMemoryRepository()
		}
		memRepo := c.memoryRepository

		c.operationRepository = persistence.NewOperationRepository(memRepo)
	}

	return c.operationRepository
}

// GetBankAccountFactory возвращает фабрику банковских счетов
func (c *Container) GetBankAccountFactory() *factory.BankAccountFactory {
	c.factoryMu.Lock()
	defer c.factoryMu.Unlock()

	if c.bankAccountFactory == nil {
		c.bankAccountFactory = factory.NewBankAccountFactory()
	}

	return c.bankAccountFactory
}

// GetCategoryFactory возвращает фабрику категорий
func (c *Container) GetCategoryFactory() *factory.CategoryFactory {
	c.factoryMu.Lock()
	defer c.factoryMu.Unlock()

	if c.categoryFactory == nil {
		c.categoryFactory = factory.NewCategoryFactory()
	}

	return c.categoryFactory
}

// GetOperationFactory возвращает фабрику операций
func (c *Container) GetOperationFactory() *factory.OperationFactory {
	c.factoryMu.Lock()
	defer c.factoryMu.Unlock()

	if c.operationFactory == nil {
		c.operationFactory = factory.NewOperationFactory()
	}

	return c.operationFactory
}

// GetBankAccountService возвращает сервис для управления банковскими счетами
func (c *Container) GetBankAccountService() interfaces.BankAccountService {
	c.serviceMu.Lock()
	defer c.serviceMu.Unlock()

	if c.bankAccountService == nil {
		// Получаем все зависимости вне мьютекса serviceMu
		c.repoMu.Lock()

		if c.memoryRepository == nil {
			c.memoryRepository = persistence.NewMemoryRepository()
		}

		if c.bankAccountRepository == nil {
			c.bankAccountRepository = persistence.NewBankAccountRepository(c.memoryRepository)
		}
		bankRepo := c.bankAccountRepository

		if c.operationRepository == nil {
			c.operationRepository = persistence.NewOperationRepository(c.memoryRepository)
		}
		opRepo := c.operationRepository

		c.repoMu.Unlock()

		// Инициализируем фабрику напрямую
		c.factoryMu.Lock()
		if c.bankAccountFactory == nil {
			c.bankAccountFactory = factory.NewBankAccountFactory()
		}
		bankFactory := c.bankAccountFactory
		c.factoryMu.Unlock()

		c.bankAccountService = services.NewBankAccountService(
			bankRepo,
			opRepo,
			bankFactory,
		)
	}

	return c.bankAccountService
}

// GetCategoryService возвращает сервис для управления категориями
func (c *Container) GetCategoryService() interfaces.CategoryService {
	c.serviceMu.Lock()
	defer c.serviceMu.Unlock()

	if c.categoryService == nil {
		// Получаем все зависимости до инициализации сервиса
		catRepo := c.GetCategoryRepository()
		opRepo := c.GetOperationRepository()
		factory := c.GetCategoryFactory()

		c.categoryService = services.NewCategoryService(
			catRepo,
			opRepo,
			factory,
		)
	}

	return c.categoryService
}

// GetOperationService возвращает сервис для управления операциями
func (c *Container) GetOperationService() interfaces.OperationService {
	c.serviceMu.Lock()
	defer c.serviceMu.Unlock()

	if c.operationService == nil {
		// Получаем все зависимости до инициализации сервиса
		opRepo := c.GetOperationRepository()
		bankRepo := c.GetBankAccountRepository()
		catRepo := c.GetCategoryRepository()
		factory := c.GetOperationFactory()

		c.operationService = services.NewOperationService(
			opRepo,
			bankRepo,
			catRepo,
			factory,
		)
	}

	return c.operationService
}

// GetAnalyticsService возвращает сервис для аналитики финансов
func (c *Container) GetAnalyticsService() interfaces.AnalyticsService {
	c.serviceMu.Lock()
	defer c.serviceMu.Unlock()

	if c.analyticsService == nil {
		// Получаем все зависимости до инициализации сервиса
		opRepo := c.GetOperationRepository()
		catRepo := c.GetCategoryRepository()

		c.analyticsService = analytics.NewAnalyticsService(
			opRepo,
			catRepo,
		)
	}

	return c.analyticsService
}

// GetBankAccountFacade возвращает фасад для управления банковскими счетами
func (c *Container) GetBankAccountFacade() *facade.BankAccountFacade {
	c.facadeMu.Lock()
	defer c.facadeMu.Unlock()

	if c.bankAccountFacade == nil {
		// Инициализируем сервис напрямую
		c.serviceMu.Lock()
		if c.bankAccountService == nil {
			// Инициализируем зависимости вне serviceMu
			c.serviceMu.Unlock()

			c.repoMu.Lock()

			if c.memoryRepository == nil {
				c.memoryRepository = persistence.NewMemoryRepository()
			}

			if c.bankAccountRepository == nil {
				c.bankAccountRepository = persistence.NewBankAccountRepository(c.memoryRepository)
			}
			bankRepo := c.bankAccountRepository

			if c.operationRepository == nil {
				c.operationRepository = persistence.NewOperationRepository(c.memoryRepository)
			}
			opRepo := c.operationRepository

			c.repoMu.Unlock()

			// Инициализируем фабрику напрямую
			c.factoryMu.Lock()
			if c.bankAccountFactory == nil {
				c.bankAccountFactory = factory.NewBankAccountFactory()
			}
			bankFactory := c.bankAccountFactory
			c.factoryMu.Unlock()

			// Теперь снова захватываем serviceMu
			c.serviceMu.Lock()

			// Проверяем еще раз, не был ли сервис инициализирован другой горутиной
			if c.bankAccountService == nil {
				c.bankAccountService = services.NewBankAccountService(
					bankRepo,
					opRepo,
					bankFactory,
				)
			}
		}
		service := c.bankAccountService
		c.serviceMu.Unlock()

		c.bankAccountFacade = facade.NewBankAccountFacade(service)
	}

	return c.bankAccountFacade
}

// GetCategoryFacade возвращает фасад для работы с категориями
func (c *Container) GetCategoryFacade() *facade.CategoryFacade {
	c.facadeMu.Lock()
	defer c.facadeMu.Unlock()

	if c.categoryFacade == nil {
		// Получаем сервис до инициализации фасада
		service := c.GetCategoryService()

		c.categoryFacade = facade.NewCategoryFacade(
			service,
		)
	}

	return c.categoryFacade
}

// GetOperationFacade возвращает фасад для управления операциями
func (c *Container) GetOperationFacade() *facade.OperationFacade {
	c.facadeMu.Lock()
	defer c.facadeMu.Unlock()

	if c.operationFacade == nil {
		// Инициализируем сервисы напрямую
		c.serviceMu.Lock()
		if c.operationService == nil {
			// Получаем все зависимости для сервиса операций
			c.repoMu.Lock()

			if c.memoryRepository == nil {
				c.memoryRepository = persistence.NewMemoryRepository()
			}

			if c.operationRepository == nil {
				c.operationRepository = persistence.NewOperationRepository(c.memoryRepository)
			}
			opRepo := c.operationRepository

			if c.bankAccountRepository == nil {
				c.bankAccountRepository = persistence.NewBankAccountRepository(c.memoryRepository)
			}
			bankRepo := c.bankAccountRepository

			if c.categoryRepository == nil {
				c.categoryRepository = persistence.NewCategoryRepository(c.memoryRepository)
			}
			catRepo := c.categoryRepository

			c.repoMu.Unlock()

			// Инициализируем фабрику напрямую
			c.factoryMu.Lock()
			if c.operationFactory == nil {
				c.operationFactory = factory.NewOperationFactory()
			}
			opFactory := c.operationFactory
			c.factoryMu.Unlock()

			c.operationService = services.NewOperationService(
				opRepo,
				bankRepo,
				catRepo,
				opFactory,
			)
		}
		opService := c.operationService

		// Инициализируем сервис категорий, если он еще не создан
		if c.categoryService == nil {
			// Получаем все зависимости для сервиса категорий
			c.repoMu.Lock()

			if c.memoryRepository == nil {
				c.memoryRepository = persistence.NewMemoryRepository()
			}

			if c.categoryRepository == nil {
				c.categoryRepository = persistence.NewCategoryRepository(c.memoryRepository)
			}
			catRepo := c.categoryRepository

			if c.operationRepository == nil {
				c.operationRepository = persistence.NewOperationRepository(c.memoryRepository)
			}
			opRepo := c.operationRepository

			c.repoMu.Unlock()

			// Инициализируем фабрику напрямую
			c.factoryMu.Lock()
			if c.categoryFactory == nil {
				c.categoryFactory = factory.NewCategoryFactory()
			}
			catFactory := c.categoryFactory
			c.factoryMu.Unlock()

			c.categoryService = services.NewCategoryService(
				catRepo,
				opRepo,
				catFactory,
			)
		}
		catService := c.categoryService
		c.serviceMu.Unlock()

		c.operationFacade = facade.NewOperationFacade(opService, catService)
	}

	return c.operationFacade
}

// GetAnalyticsFacade возвращает фасад для аналитики
func (c *Container) GetAnalyticsFacade() *facade.AnalyticsFacade {
	c.facadeMu.Lock()
	defer c.facadeMu.Unlock()

	if c.analyticsFacade == nil {
		// Получаем сервис до инициализации фасада
		service := c.GetAnalyticsService()

		c.analyticsFacade = facade.NewAnalyticsFacade(
			service,
		)
	}

	return c.analyticsFacade
}
