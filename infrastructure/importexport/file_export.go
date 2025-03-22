package importexport

import (
	"KPO1/domain/interfaces"
	"fmt"
	"os"
)

// FileExporter экспортирует данные в файлы
type FileExporter struct {
	visitor    interfaces.ExportVisitor
	repository interfaces.CompositeRepository
	exportPath string
}

// NewFileExporter создает новый экспортер файлов
func NewFileExporter(format FileFormat, path string, repository interfaces.CompositeRepository) *FileExporter {
	// Создаем директорию для экспорта, если не существует
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			panic(fmt.Sprintf("Не удалось создать директорию для экспорта: %v", err))
		}
	}

	return &FileExporter{
		visitor:    NewExportVisitor(format, path),
		repository: repository,
		exportPath: path,
	}
}

// ExportAll экспортирует все данные в файлы
func (e *FileExporter) ExportAll() error {
	if err := e.ExportBankAccounts(); err != nil {
		return err
	}

	if err := e.ExportCategories(); err != nil {
		return err
	}

	if err := e.ExportOperations(); err != nil {
		return err
	}

	return nil
}

// ExportBankAccounts экспортирует банковские счета
func (e *FileExporter) ExportBankAccounts() error {
	accounts, err := e.repository.GetBankAccounts()
	if err != nil {
		return fmt.Errorf("ошибка получения счетов: %w", err)
	}

	err = e.visitor.VisitBankAccounts(accounts)
	if err != nil {
		return fmt.Errorf("ошибка экспорта счетов: %w", err)
	}

	return nil
}

// ExportCategories экспортирует категории
func (e *FileExporter) ExportCategories() error {
	categories, err := e.repository.GetCategories()
	if err != nil {
		return fmt.Errorf("ошибка получения категорий: %w", err)
	}

	err = e.visitor.VisitCategories(categories)
	if err != nil {
		return fmt.Errorf("ошибка экспорта категорий: %w", err)
	}

	return nil
}

// ExportOperations экспортирует операции
func (e *FileExporter) ExportOperations() error {
	operations, err := e.repository.GetOperations()
	if err != nil {
		return fmt.Errorf("ошибка получения операций: %w", err)
	}

	err = e.visitor.VisitOperations(operations)
	if err != nil {
		return fmt.Errorf("ошибка экспорта операций: %w", err)
	}

	return nil
}
