package importexport

import (
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// FileImporter импортирует данные из файлов
type FileImporter struct {
	format      FileFormat
	importPath  string
	bankAccRepo interfaces.BankAccountRepository
	catRepo     interfaces.CategoryRepository
	opRepo      interfaces.OperationRepository
}

// NewFileImporter создает новый импортер файлов
func NewFileImporter(
	format FileFormat,
	path string,
	bankAccRepo interfaces.BankAccountRepository,
	catRepo interfaces.CategoryRepository,
	opRepo interfaces.OperationRepository,
) *FileImporter {
	return &FileImporter{
		format:      format,
		importPath:  path,
		bankAccRepo: bankAccRepo,
		catRepo:     catRepo,
		opRepo:      opRepo,
	}
}

// ImportAll импортирует все данные из файлов
func (i *FileImporter) ImportAll() error {
	if err := i.ImportBankAccounts(); err != nil {
		return err
	}

	if err := i.ImportCategories(); err != nil {
		return err
	}

	if err := i.ImportOperations(); err != nil {
		return err
	}

	return nil
}

// ImportBankAccounts импортирует банковские счета
func (i *FileImporter) ImportBankAccounts() error {
	switch i.format {
	case CSV:
		return i.importBankAccountsFromCSV()
	case JSON:
		return i.importBankAccountsFromJSON()
	case YAML:
		return i.importBankAccountsFromYAML()
	default:
		return fmt.Errorf("неподдерживаемый формат: %s", i.format)
	}
}

// ImportCategories импортирует категории
func (i *FileImporter) ImportCategories() error {
	switch i.format {
	case CSV:
		return i.importCategoriesFromCSV()
	case JSON:
		return i.importCategoriesFromJSON()
	case YAML:
		return i.importCategoriesFromYAML()
	default:
		return fmt.Errorf("неподдерживаемый формат: %s", i.format)
	}
}

// ImportOperations импортирует операции
func (i *FileImporter) ImportOperations() error {
	switch i.format {
	case CSV:
		return i.importOperationsFromCSV()
	case JSON:
		return i.importOperationsFromJSON()
	case YAML:
		return i.importOperationsFromYAML()
	default:
		return fmt.Errorf("неподдерживаемый формат: %s", i.format)
	}
}

// importBankAccountsFromCSV импортирует банковские счета из CSV
func (i *FileImporter) importBankAccountsFromCSV() error {
	file, err := os.Open(fmt.Sprintf("%s/accounts.csv", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Пропускаем заголовок
	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if len(record) < 3 {
			return fmt.Errorf("неверный формат записи: %v", record)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return fmt.Errorf("ошибка преобразования ID: %w", err)
		}

		name := record[1]

		balance, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return fmt.Errorf("ошибка преобразования баланса: %w", err)
		}

		account := &models.BankAccount{
			ID:      id,
			Name:    name,
			Balance: balance,
		}

		err = i.bankAccRepo.Save(account)
		if err != nil {
			return fmt.Errorf("ошибка создания счета: %w", err)
		}
	}

	return nil
}

// importBankAccountsFromJSON импортирует банковские счета из JSON
func (i *FileImporter) importBankAccountsFromJSON() error {
	file, err := os.Open(fmt.Sprintf("%s/accounts.json", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	var accounts []*models.BankAccount
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&accounts); err != nil {
		return err
	}

	for _, account := range accounts {
		err := i.bankAccRepo.Save(account)
		if err != nil {
			return fmt.Errorf("ошибка создания счета: %w", err)
		}
	}

	return nil
}

// importBankAccountsFromYAML импортирует банковские счета из YAML
func (i *FileImporter) importBankAccountsFromYAML() error {
	file, err := os.Open(fmt.Sprintf("%s/accounts.yaml", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	var accounts []*models.BankAccount
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&accounts); err != nil {
		return err
	}

	for _, account := range accounts {
		err := i.bankAccRepo.Save(account)
		if err != nil {
			return fmt.Errorf("ошибка создания счета: %w", err)
		}
	}

	return nil
}

// importCategoriesFromCSV импортирует категории из CSV
func (i *FileImporter) importCategoriesFromCSV() error {
	file, err := os.Open(fmt.Sprintf("%s/categories.csv", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Пропускаем заголовок
	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if len(record) < 3 {
			return fmt.Errorf("неверный формат записи: %v", record)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return fmt.Errorf("ошибка преобразования ID: %w", err)
		}

		categoryType := models.OperationType(record[1])
		if categoryType != models.Income && categoryType != models.Expense {
			return fmt.Errorf("неверный тип категории: %s", categoryType)
		}

		name := record[2]

		category := &models.Category{
			ID:   id,
			Type: categoryType,
			Name: name,
		}

		err = i.catRepo.Save(category)
		if err != nil {
			return fmt.Errorf("ошибка создания категории: %w", err)
		}
	}

	return nil
}

// importCategoriesFromJSON импортирует категории из JSON
func (i *FileImporter) importCategoriesFromJSON() error {
	file, err := os.Open(fmt.Sprintf("%s/categories.json", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	var categories []*models.Category
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&categories); err != nil {
		return err
	}

	for _, category := range categories {
		err := i.catRepo.Save(category)
		if err != nil {
			return fmt.Errorf("ошибка создания категории: %w", err)
		}
	}

	return nil
}

// importCategoriesFromYAML импортирует категории из YAML
func (i *FileImporter) importCategoriesFromYAML() error {
	file, err := os.Open(fmt.Sprintf("%s/categories.yaml", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	var categories []*models.Category
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&categories); err != nil {
		return err
	}

	for _, category := range categories {
		err := i.catRepo.Save(category)
		if err != nil {
			return fmt.Errorf("ошибка создания категории: %w", err)
		}
	}

	return nil
}

// importOperationsFromCSV импортирует операции из CSV
func (i *FileImporter) importOperationsFromCSV() error {
	file, err := os.Open(fmt.Sprintf("%s/operations.csv", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Пропускаем заголовок
	_, err = reader.Read()
	if err != nil {
		return err
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if len(record) < 7 {
			return fmt.Errorf("неверный формат записи: %v", record)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return fmt.Errorf("ошибка преобразования ID: %w", err)
		}

		opType := models.OperationType(record[1])
		if opType != models.Income && opType != models.Expense {
			return fmt.Errorf("неверный тип операции: %s", opType)
		}

		bankAccountID, err := strconv.Atoi(record[2])
		if err != nil {
			return fmt.Errorf("ошибка преобразования ID счета: %w", err)
		}

		amount, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return fmt.Errorf("ошибка преобразования суммы: %w", err)
		}

		date, err := time.Parse(time.RFC3339, record[4])
		if err != nil {
			return fmt.Errorf("ошибка преобразования даты: %w", err)
		}

		description := record[5]

		categoryID, err := strconv.Atoi(record[6])
		if err != nil {
			return fmt.Errorf("ошибка преобразования ID категории: %w", err)
		}

		operation := &models.Operation{
			ID:            id,
			Type:          opType,
			BankAccountID: bankAccountID,
			Amount:        amount,
			Date:          date,
			Description:   description,
			CategoryID:    categoryID,
		}

		err = i.opRepo.Save(operation)
		if err != nil {
			return fmt.Errorf("ошибка создания операции: %w", err)
		}
	}

	return nil
}

// importOperationsFromJSON импортирует операции из JSON
func (i *FileImporter) importOperationsFromJSON() error {
	file, err := os.Open(fmt.Sprintf("%s/operations.json", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	var operations []*models.Operation
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&operations); err != nil {
		return err
	}

	for _, operation := range operations {
		err := i.opRepo.Save(operation)
		if err != nil {
			return fmt.Errorf("ошибка создания операции: %w", err)
		}
	}

	return nil
}

// importOperationsFromYAML импортирует операции из YAML
func (i *FileImporter) importOperationsFromYAML() error {
	file, err := os.Open(fmt.Sprintf("%s/operations.yaml", i.importPath))
	if err != nil {
		return err
	}
	defer file.Close()

	var operations []*models.Operation
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&operations); err != nil {
		return err
	}

	for _, operation := range operations {
		err := i.opRepo.Save(operation)
		if err != nil {
			return fmt.Errorf("ошибка создания операции: %w", err)
		}
	}

	return nil
}
