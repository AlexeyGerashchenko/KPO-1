package importexport

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Убедимся что ExportVisitor реализует интерфейс
var _ interfaces.ExportVisitor = (*ExportVisitor)(nil)

// ExportVisitor реализует паттерн Посетитель для экспорта данных
type ExportVisitor struct {
	format FileFormat
	path   string
}

// NewExportVisitor создает нового посетителя для экспорта
func NewExportVisitor(format FileFormat, path string) *ExportVisitor {
	return &ExportVisitor{
		format: format,
		path:   path,
	}
}

// VisitBankAccounts экспортирует банковские счета
func (v *ExportVisitor) VisitBankAccounts(accounts []*models.BankAccount) error {
	switch v.format {
	case CSV:
		return v.exportBankAccountsToCSV(accounts)
	case JSON:
		return v.exportBankAccountsToJSON(accounts)
	case YAML:
		return v.exportBankAccountsToYAML(accounts)
	default:
		return fmt.Errorf("неподдерживаемый формат: %s", v.format)
	}
}

// VisitCategories экспортирует категории
func (v *ExportVisitor) VisitCategories(categories []*models.Category) error {
	switch v.format {
	case CSV:
		return v.exportCategoriesToCSV(categories)
	case JSON:
		return v.exportCategoriesToJSON(categories)
	case YAML:
		return v.exportCategoriesToYAML(categories)
	default:
		return fmt.Errorf("неподдерживаемый формат: %s", v.format)
	}
}

// VisitOperations экспортирует операции
func (v *ExportVisitor) VisitOperations(operations []*models.Operation) error {
	switch v.format {
	case CSV:
		return v.exportOperationsToCSV(operations)
	case JSON:
		return v.exportOperationsToJSON(operations)
	case YAML:
		return v.exportOperationsToYAML(operations)
	default:
		return fmt.Errorf("неподдерживаемый формат: %s", v.format)
	}
}

// exportBankAccountsToCSV экспортирует банковские счета в CSV
func (v *ExportVisitor) exportBankAccountsToCSV(accounts []*models.BankAccount) error {
	file, err := os.Create(fmt.Sprintf("%s/accounts.csv", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись заголовка
	err = writer.Write([]string{"ID", "Name", "Balance"})
	if err != nil {
		return err
	}

	// Запись данных
	for _, account := range accounts {
		err = writer.Write([]string{
			fmt.Sprintf("%d", account.ID),
			account.Name,
			fmt.Sprintf("%.2f", account.Balance),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// exportBankAccountsToJSON экспортирует банковские счета в JSON
func (v *ExportVisitor) exportBankAccountsToJSON(accounts []*models.BankAccount) error {
	file, err := os.Create(fmt.Sprintf("%s/accounts.json", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(accounts)
}

// exportBankAccountsToYAML экспортирует банковские счета в YAML
func (v *ExportVisitor) exportBankAccountsToYAML(accounts []*models.BankAccount) error {
	file, err := os.Create(fmt.Sprintf("%s/accounts.yaml", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	return encoder.Encode(accounts)
}

// exportCategoriesToCSV экспортирует категории в CSV
func (v *ExportVisitor) exportCategoriesToCSV(categories []*models.Category) error {
	file, err := os.Create(fmt.Sprintf("%s/categories.csv", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись заголовка
	err = writer.Write([]string{"ID", "Type", "Name"})
	if err != nil {
		return err
	}

	// Запись данных
	for _, category := range categories {
		err = writer.Write([]string{
			fmt.Sprintf("%d", category.ID),
			string(category.Type),
			category.Name,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// exportCategoriesToJSON экспортирует категории в JSON
func (v *ExportVisitor) exportCategoriesToJSON(categories []*models.Category) error {
	file, err := os.Create(fmt.Sprintf("%s/categories.json", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(categories)
}

// exportCategoriesToYAML экспортирует категории в YAML
func (v *ExportVisitor) exportCategoriesToYAML(categories []*models.Category) error {
	file, err := os.Create(fmt.Sprintf("%s/categories.yaml", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	return encoder.Encode(categories)
}

// exportOperationsToCSV экспортирует операции в CSV
func (v *ExportVisitor) exportOperationsToCSV(operations []*models.Operation) error {
	file, err := os.Create(fmt.Sprintf("%s/operations.csv", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Запись заголовка
	err = writer.Write([]string{"ID", "Type", "BankAccountID", "Amount", "Date", "Description", "CategoryID"})
	if err != nil {
		return err
	}

	// Запись данных
	for _, op := range operations {
		err = writer.Write([]string{
			fmt.Sprintf("%d", op.ID),
			string(op.Type),
			fmt.Sprintf("%d", op.BankAccountID),
			fmt.Sprintf("%.2f", op.Amount),
			op.Date.Format(time.RFC3339),
			op.Description,
			fmt.Sprintf("%d", op.CategoryID),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// exportOperationsToJSON экспортирует операции в JSON
func (v *ExportVisitor) exportOperationsToJSON(operations []*models.Operation) error {
	file, err := os.Create(fmt.Sprintf("%s/operations.json", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(operations)
}

// exportOperationsToYAML экспортирует операции в YAML
func (v *ExportVisitor) exportOperationsToYAML(operations []*models.Operation) error {
	file, err := os.Create(fmt.Sprintf("%s/operations.yaml", v.path))
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	return encoder.Encode(operations)
}
