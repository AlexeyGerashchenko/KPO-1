package commands

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"KPO1/infrastructure/importexport"
)

// ExportCSVCommand представляет команду для экспорта данных в CSV
type ExportCSVCommand struct {
	CommandBase
	exporter *importexport.FileExporter
	path     string
	errorCh  chan error
}

// NewExportCSVCommand создаёт новую команду для экспорта данных в CSV
func NewExportCSVCommand(
	bankAccountRepo interfaces.BankAccountRepository,
	categoryRepo interfaces.CategoryRepository,
	operationRepo interfaces.OperationRepository,
	path string,
	errorCh chan error,
) interfaces.Command {
	// Создаем композитный репозиторий для экспорта
	repository := &CompositeRepository{
		bankAccountRepo: bankAccountRepo,
		categoryRepo:    categoryRepo,
		operationRepo:   operationRepo,
	}
	return &ExportCSVCommand{
		CommandBase: NewCommandBase("ExportCSV"),
		exporter:    importexport.NewFileExporter(importexport.CSV, path, repository),
		path:        path,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду экспорта данных в CSV
func (c *ExportCSVCommand) Execute() error {
	err := c.exporter.ExportAll()
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}

// ExportJSONCommand представляет команду для экспорта данных в JSON
type ExportJSONCommand struct {
	CommandBase
	exporter *importexport.FileExporter
	path     string
	errorCh  chan error
}

// NewExportJSONCommand создаёт новую команду для экспорта данных в JSON
func NewExportJSONCommand(
	bankAccountRepo interfaces.BankAccountRepository,
	categoryRepo interfaces.CategoryRepository,
	operationRepo interfaces.OperationRepository,
	path string,
	errorCh chan error,
) interfaces.Command {
	// Создаем композитный репозиторий для экспорта
	repository := &CompositeRepository{
		bankAccountRepo: bankAccountRepo,
		categoryRepo:    categoryRepo,
		operationRepo:   operationRepo,
	}
	return &ExportJSONCommand{
		CommandBase: NewCommandBase("ExportJSON"),
		exporter:    importexport.NewFileExporter(importexport.JSON, path, repository),
		path:        path,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду экспорта данных в JSON
func (c *ExportJSONCommand) Execute() error {
	err := c.exporter.ExportAll()
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}

// ExportYAMLCommand представляет команду для экспорта данных в YAML
type ExportYAMLCommand struct {
	CommandBase
	exporter *importexport.FileExporter
	path     string
	errorCh  chan error
}

// NewExportYAMLCommand создаёт новую команду для экспорта данных в YAML
func NewExportYAMLCommand(
	bankAccountRepo interfaces.BankAccountRepository,
	categoryRepo interfaces.CategoryRepository,
	operationRepo interfaces.OperationRepository,
	path string,
	errorCh chan error,
) interfaces.Command {
	// Создаем композитный репозиторий для экспорта
	repository := &CompositeRepository{
		bankAccountRepo: bankAccountRepo,
		categoryRepo:    categoryRepo,
		operationRepo:   operationRepo,
	}
	return &ExportYAMLCommand{
		CommandBase: NewCommandBase("ExportYAML"),
		exporter:    importexport.NewFileExporter(importexport.YAML, path, repository),
		path:        path,
		errorCh:     errorCh,
	}
}

// Execute выполняет команду экспорта данных в YAML
func (c *ExportYAMLCommand) Execute() error {
	err := c.exporter.ExportAll()
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}

// CompositeRepository объединяет все репозитории для импорта/экспорта
type CompositeRepository struct {
	bankAccountRepo interfaces.BankAccountRepository
	categoryRepo    interfaces.CategoryRepository
	operationRepo   interfaces.OperationRepository
}

// GetBankAccounts возвращает все банковские счета
func (r *CompositeRepository) GetBankAccounts() ([]*models.BankAccount, error) {
	return r.bankAccountRepo.GetAll()
}

// GetCategories возвращает все категории
func (r *CompositeRepository) GetCategories() ([]*models.Category, error) {
	return r.categoryRepo.GetAll()
}

// GetOperations возвращает все операции
func (r *CompositeRepository) GetOperations() ([]*models.Operation, error) {
	return r.operationRepo.GetAll()
}

// ImportCSVCommand представляет команду для импорта данных из CSV
type ImportCSVCommand struct {
	CommandBase
	importer *importexport.FileImporter
	path     string
	errorCh  chan error
}

// NewImportCSVCommand создаёт новую команду для импорта данных из CSV
func NewImportCSVCommand(
	bankAccountRepo interfaces.BankAccountRepository,
	categoryRepo interfaces.CategoryRepository,
	operationRepo interfaces.OperationRepository,
	path string,
	errorCh chan error,
) interfaces.Command {
	return &ImportCSVCommand{
		CommandBase: NewCommandBase("ImportCSV"),
		importer: importexport.NewFileImporter(
			importexport.CSV,
			path,
			bankAccountRepo,
			categoryRepo,
			operationRepo,
		),
		path:    path,
		errorCh: errorCh,
	}
}

// Execute выполняет команду импорта данных из CSV
func (c *ImportCSVCommand) Execute() error {
	err := c.importer.ImportAll()
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}

// ImportJSONCommand представляет команду для импорта данных из JSON
type ImportJSONCommand struct {
	CommandBase
	importer *importexport.FileImporter
	path     string
	errorCh  chan error
}

// NewImportJSONCommand создаёт новую команду для импорта данных из JSON
func NewImportJSONCommand(
	bankAccountRepo interfaces.BankAccountRepository,
	categoryRepo interfaces.CategoryRepository,
	operationRepo interfaces.OperationRepository,
	path string,
	errorCh chan error,
) interfaces.Command {
	return &ImportJSONCommand{
		CommandBase: NewCommandBase("ImportJSON"),
		importer: importexport.NewFileImporter(
			importexport.JSON,
			path,
			bankAccountRepo,
			categoryRepo,
			operationRepo,
		),
		path:    path,
		errorCh: errorCh,
	}
}

// Execute выполняет команду импорта данных из JSON
func (c *ImportJSONCommand) Execute() error {
	err := c.importer.ImportAll()
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}

// ImportYAMLCommand представляет команду для импорта данных из YAML
type ImportYAMLCommand struct {
	CommandBase
	importer *importexport.FileImporter
	path     string
	errorCh  chan error
}

// NewImportYAMLCommand создаёт новую команду для импорта данных из YAML
func NewImportYAMLCommand(
	bankAccountRepo interfaces.BankAccountRepository,
	categoryRepo interfaces.CategoryRepository,
	operationRepo interfaces.OperationRepository,
	path string,
	errorCh chan error,
) interfaces.Command {
	return &ImportYAMLCommand{
		CommandBase: NewCommandBase("ImportYAML"),
		importer: importexport.NewFileImporter(
			importexport.YAML,
			path,
			bankAccountRepo,
			categoryRepo,
			operationRepo,
		),
		path:    path,
		errorCh: errorCh,
	}
}

// Execute выполняет команду импорта данных из YAML
func (c *ImportYAMLCommand) Execute() error {
	err := c.importer.ImportAll()
	if err != nil && c.errorCh != nil {
		c.errorCh <- err
	}
	return err
}
