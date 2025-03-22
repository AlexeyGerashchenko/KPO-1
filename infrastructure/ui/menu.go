package ui

import (
	"KPO1/application/commands"
	"KPO1/di"
	"KPO1/domain/interfaces"
	"KPO1/domain/models"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type MainMenu struct {
	console   *ConsoleUI
	container *di.Container
}

func NewMainMenu(console *ConsoleUI, container *di.Container) *MainMenu {
	return &MainMenu{
		console:   console,
		container: container,
	}
}

func (m *MainMenu) Display() {
	fmt.Println("1. Управление банковскими счетами")
	fmt.Println("2. Управление категориями")
	fmt.Println("3. Управление операциями")
	fmt.Println("4. Аналитика")
	fmt.Println("5. Импорт/Экспорт данных")
	fmt.Println("0. Выход")
}

func (m *MainMenu) HandleInput(input string, reader *bufio.Reader) error {
	switch input {
	case "1":
		return m.bankAccountsMenu(reader)
	case "2":
		return m.categoriesMenu(reader)
	case "3":
		return m.operationsMenu(reader)
	case "4":
		return m.analyticsMenu(reader)
	case "5":
		return m.importExportMenu(reader)
	default:
		fmt.Println("Неверный выбор. Повторите попытку.")
	}
	return nil
}

func (m *MainMenu) bankAccountsMenu(reader *bufio.Reader) error {
	fmt.Println("\n--- Управление банковскими счетами ---")
	fmt.Println("1. Создать счет")
	fmt.Println("2. Получить счет по ID")
	fmt.Println("3. Список счетов")
	fmt.Println("4. Обновить счет")
	fmt.Println("5. Удалить счет")
	fmt.Println("6. Пересчитать баланс")
	fmt.Println("0. Назад")
	fmt.Print("Выберите опцию: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		fmt.Print("Введите название счета: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		resultCh := make(chan *models.BankAccount, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewCreateBankAccountCommand(
			m.container.GetBankAccountFacade(),
			name,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			account := <-resultCh
			fmt.Printf("Создан счет: %+v\n", account)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "2":
		fmt.Print("Введите ID счета: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		resultCh := make(chan *models.BankAccount, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewGetBankAccountCommand(
			m.container.GetBankAccountFacade(),
			id,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			account := <-resultCh
			fmt.Printf("Счет: %+v\n", account)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "3":
		resultCh := make(chan []*models.BankAccount, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewListBankAccountsCommand(
			m.container.GetBankAccountFacade(),
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			accounts := <-resultCh
			fmt.Println("Список счетов:")
			for _, acc := range accounts {
				fmt.Printf("%+v\n", acc)
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "4":
		fmt.Print("Введите ID счета для обновления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		fmt.Print("Введите новое название счета: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		resultCh := make(chan *models.BankAccount, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewUpdateBankAccountCommand(
			m.container.GetBankAccountFacade(),
			id,
			name,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			account := <-resultCh
			fmt.Printf("Обновленный счет: %+v\n", account)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "5":
		fmt.Print("Введите ID счета для удаления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		errorCh := make(chan error, 1)
		cmd := commands.NewDeleteBankAccountCommand(
			m.container.GetBankAccountFacade(),
			id,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Счет успешно удален.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "6":
		fmt.Print("Введите ID счета для пересчета баланса: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		account, err := m.container.GetBankAccountFacade().RecalculateBalance(id)
		if err == nil {
			fmt.Printf("Пересчитанный счет: %+v\n", account)
		} else {
			fmt.Printf("Ошибка: %v\n", err)
		}
	case "0":
		return nil
	default:
		fmt.Println("Неверный выбор.")
	}
	return nil
}

func (m *MainMenu) categoriesMenu(reader *bufio.Reader) error {
	fmt.Println("\n--- Управление категориями ---")
	fmt.Println("1. Создать категорию")
	fmt.Println("2. Получить категорию по ID")
	fmt.Println("3. Список категорий")
	fmt.Println("4. Список категорий по типу")
	fmt.Println("5. Обновить категорию")
	fmt.Println("6. Удалить категорию")
	fmt.Println("0. Назад")
	fmt.Print("Выберите опцию: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		fmt.Print("Введите название категории: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		fmt.Print("Введите тип категории (1 - доход, 2 - расход): ")
		typeStr, _ := reader.ReadString('\n')
		typeStr = strings.TrimSpace(typeStr)
		var opType models.OperationType
		if typeStr == "1" {
			opType = models.Income
		} else {
			opType = models.Expense
		}
		resultCh := make(chan *models.Category, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewCreateCategoryCommand(
			m.container.GetCategoryFacade(),
			name,
			opType,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			category := <-resultCh
			fmt.Printf("Создана категория: %+v\n", category)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "2":
		fmt.Print("Введите ID категории: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		resultCh := make(chan *models.Category, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewGetCategoryCommand(
			m.container.GetCategoryFacade(),
			id,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			category := <-resultCh
			fmt.Printf("Категория: %+v\n", category)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "3":
		resultCh := make(chan []*models.Category, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewListCategoriesCommand(
			m.container.GetCategoryFacade(),
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			categories := <-resultCh
			fmt.Println("Список категорий:")
			for _, cat := range categories {
				fmt.Printf("%+v\n", cat)
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "4":
		fmt.Print("Введите тип категории (1 - доход, 2 - расход): ")
		typeStr, _ := reader.ReadString('\n')
		typeStr = strings.TrimSpace(typeStr)
		var opType models.OperationType
		if typeStr == "1" {
			opType = models.Income
		} else {
			opType = models.Expense
		}
		resultCh := make(chan []*models.Category, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewListCategoriesByTypeCommand(
			m.container.GetCategoryFacade(),
			opType,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			categories := <-resultCh
			fmt.Println("Список категорий:")
			for _, cat := range categories {
				fmt.Printf("%+v\n", cat)
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "5":
		fmt.Print("Введите ID категории для обновления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		fmt.Print("Введите новое название категории: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		fmt.Print("Введите тип категории (1 - доход, 2 - расход): ")
		typeStr, _ := reader.ReadString('\n')
		typeStr = strings.TrimSpace(typeStr)
		var opType models.OperationType
		if typeStr == "1" {
			opType = models.Income
		} else {
			opType = models.Expense
		}
		resultCh := make(chan *models.Category, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewUpdateCategoryCommand(
			m.container.GetCategoryFacade(),
			id,
			name,
			opType,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			category := <-resultCh
			fmt.Printf("Обновленная категория: %+v\n", category)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "6":
		fmt.Print("Введите ID категории для удаления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		errorCh := make(chan error, 1)
		cmd := commands.NewDeleteCategoryCommand(
			m.container.GetCategoryFacade(),
			id,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Категория успешно удалена.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "0":
		return nil
	default:
		fmt.Println("Неверный выбор.")
	}
	return nil
}

func (m *MainMenu) operationsMenu(reader *bufio.Reader) error {
	fmt.Println("\n--- Управление операциями ---")
	fmt.Println("1. Создать операцию")
	fmt.Println("2. Получить операцию по ID")
	fmt.Println("3. Список операций")
	fmt.Println("4. Список операций по счету")
	fmt.Println("5. Список операций по категории")
	fmt.Println("6. Обновить операцию")
	fmt.Println("7. Удалить операцию")
	fmt.Println("0. Назад")
	fmt.Print("Выберите опцию: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		fmt.Print("Введите ID счета: ")
		bankStr, _ := reader.ReadString('\n')
		bankStr = strings.TrimSpace(bankStr)
		bankID, _ := strconv.Atoi(bankStr)
		fmt.Print("Введите ID категории: ")
		catStr, _ := reader.ReadString('\n')
		catStr = strings.TrimSpace(catStr)
		categoryID, _ := strconv.Atoi(catStr)
		fmt.Print("Введите сумму операции: ")
		amountStr, _ := reader.ReadString('\n')
		amountStr = strings.TrimSpace(amountStr)
		amount, _ := strconv.ParseFloat(amountStr, 64)
		fmt.Print("Введите тип операции (1 - доход, 2 - расход): ")
		typeStr, _ := reader.ReadString('\n')
		typeStr = strings.TrimSpace(typeStr)
		var opType models.OperationType
		if typeStr == "1" {
			opType = models.Income
		} else {
			opType = models.Expense
		}
		fmt.Print("Введите дату операции (формат YYYY-MM-DD): ")
		dateStr, _ := reader.ReadString('\n')
		dateStr = strings.TrimSpace(dateStr)
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Неверный формат даты.")
			return nil
		}
		fmt.Print("Введите описание операции: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)
		resultCh := make(chan *models.Operation, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewCreateOperationCommand(
			m.container.GetOperationFacade(),
			opType,
			bankID,
			categoryID,
			amount,
			date,
			description,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			operation := <-resultCh
			fmt.Printf("Создана операция: %+v\n", operation)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "2":
		fmt.Print("Введите ID операции: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		resultCh := make(chan *models.Operation, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewGetOperationCommand(
			m.container.GetOperationFacade(),
			id,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			operation := <-resultCh
			fmt.Printf("Операция: %+v\n", operation)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "3":
		resultCh := make(chan []*models.Operation, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewListOperationsCommand(
			m.container.GetOperationFacade(),
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			operations := <-resultCh
			fmt.Println("Список операций:")
			for _, op := range operations {
				fmt.Printf("%+v\n", op)
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "4":
		fmt.Print("Введите ID счета: ")
		bankStr, _ := reader.ReadString('\n')
		bankStr = strings.TrimSpace(bankStr)
		bankID, _ := strconv.Atoi(bankStr)
		resultCh := make(chan []*models.Operation, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewListOperationsByAccountCommand(
			m.container.GetOperationFacade(),
			bankID,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			operations := <-resultCh
			fmt.Println("Операции по счету:")
			for _, op := range operations {
				fmt.Printf("%+v\n", op)
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "5":
		fmt.Print("Введите ID категории: ")
		catStr, _ := reader.ReadString('\n')
		catStr = strings.TrimSpace(catStr)
		categoryID, _ := strconv.Atoi(catStr)
		resultCh := make(chan []*models.Operation, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewListOperationsByCategoryCommand(
			m.container.GetOperationFacade(),
			categoryID,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			operations := <-resultCh
			fmt.Println("Операции по категории:")
			for _, op := range operations {
				fmt.Printf("%+v\n", op)
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "6":
		fmt.Print("Введите ID операции для обновления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		fmt.Print("Введите новый ID счета: ")
		bankStr, _ := reader.ReadString('\n')
		bankStr = strings.TrimSpace(bankStr)
		bankID, _ := strconv.Atoi(bankStr)
		fmt.Print("Введите новый ID категории: ")
		catStr, _ := reader.ReadString('\n')
		catStr = strings.TrimSpace(catStr)
		categoryID, _ := strconv.Atoi(catStr)
		fmt.Print("Введите новую сумму операции: ")
		amountStr, _ := reader.ReadString('\n')
		amountStr = strings.TrimSpace(amountStr)
		amount, _ := strconv.ParseFloat(amountStr, 64)
		fmt.Print("Введите новый тип операции (1 - доход, 2 - расход): ")
		typeStr, _ := reader.ReadString('\n')
		typeStr = strings.TrimSpace(typeStr)
		var opType models.OperationType
		if typeStr == "1" {
			opType = models.Income
		} else {
			opType = models.Expense
		}
		fmt.Print("Введите новую дату операции (YYYY-MM-DD): ")
		dateStr, _ := reader.ReadString('\n')
		dateStr = strings.TrimSpace(dateStr)
		date, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			fmt.Println("Неверный формат даты.")
			return nil
		}
		fmt.Print("Введите новое описание операции: ")
		description, _ := reader.ReadString('\n')
		description = strings.TrimSpace(description)
		resultCh := make(chan *models.Operation, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewUpdateOperationCommand(
			m.container.GetOperationFacade(),
			id,
			bankID,
			categoryID,
			amount,
			opType,
			date,
			description,
			resultCh,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			operation := <-resultCh
			fmt.Printf("Обновленная операция: %+v\n", operation)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "7":
		fmt.Print("Введите ID операции для удаления: ")
		idStr, _ := reader.ReadString('\n')
		idStr = strings.TrimSpace(idStr)
		id, _ := strconv.Atoi(idStr)
		errorCh := make(chan error, 1)
		cmd := commands.NewDeleteOperationCommand(
			m.container.GetOperationFacade(),
			id,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Операция успешно удалена.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "0":
		return nil
	default:
		fmt.Println("Неверный выбор.")
	}
	return nil
}

func (m *MainMenu) analyticsMenu(reader *bufio.Reader) error {
	fmt.Println("\nМеню аналитики:")
	fmt.Println("1. Разница доходов и расходов за период")
	fmt.Println("2. Группировка по категориям")
	fmt.Println("3. Месячная динамика")
	fmt.Println("0. Назад")
	fmt.Print("\nВыберите действие: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		start, end := readDateRange(reader)
		resultCh := make(chan float64, 1)
		errorCh := make(chan error, 1)

		cmd := commands.NewBalanceByPeriodCommand(
			m.container.GetAnalyticsFacade(),
			start,
			end,
			resultCh,
			errorCh,
		)

		// Оборачиваем команду в декоратор для измерения времени
		decoratedCmd := m.wrapWithTimeDecorator(cmd)

		if err := decoratedCmd.Execute(); err == nil {
			diff := <-resultCh
			fmt.Printf("Разница доходов и расходов за период: %.2f руб.\n", diff)
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "2":
		start, end := readDateRange(reader)
		resultCh := make(chan map[string]float64, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewExpensesByCategoryCommand(
			m.container.GetAnalyticsFacade(),
			start,
			end,
			resultCh,
			errorCh,
		)

		// Оборачиваем команду в декоратор для измерения времени
		decoratedCmd := m.wrapWithTimeDecorator(cmd)

		if err := decoratedCmd.Execute(); err == nil {
			expenses := <-resultCh
			fmt.Println("Расходы по категориям:")
			for cat, amt := range expenses {
				fmt.Printf("%s: %.2f\n", cat, amt)
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "3":
		fmt.Print("Введите год: ")
		yearStr, _ := reader.ReadString('\n')
		yearStr = strings.TrimSpace(yearStr)
		year, _ := strconv.Atoi(yearStr)
		resultCh := make(chan map[time.Month]map[models.OperationType]float64, 1)
		errorCh := make(chan error, 1)
		cmd := commands.NewMonthlyDynamicsCommand(
			m.container.GetAnalyticsFacade(),
			year,
			resultCh,
			errorCh,
		)

		// Оборачиваем команду в декоратор для измерения времени
		decoratedCmd := m.wrapWithTimeDecorator(cmd)

		if err := decoratedCmd.Execute(); err == nil {
			dynamics := <-resultCh
			fmt.Println("Месячная динамика:")
			for month, data := range dynamics {
				fmt.Printf("%s - Доход: %.2f, Расход: %.2f\n", month, data[models.Income], data[models.Expense])
			}
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "0":
		return nil
	default:
		fmt.Println("Неверный выбор.")
	}
	return nil
}

func (m *MainMenu) importExportMenu(reader *bufio.Reader) error {
	fmt.Println("\n--- Импорт/Экспорт данных ---")
	fmt.Println("1. Экспорт в CSV")
	fmt.Println("2. Экспорт в JSON")
	fmt.Println("3. Экспорт в YAML")
	fmt.Println("4. Импорт из CSV")
	fmt.Println("5. Импорт из JSON")
	fmt.Println("6. Импорт из YAML")
	fmt.Println("0. Назад")
	fmt.Print("Выберите опцию: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		fmt.Print("Введите путь для экспорта CSV: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		errorCh := make(chan error, 1)
		cmd := commands.NewExportCSVCommand(
			m.container.GetBankAccountRepository(),
			m.container.GetCategoryRepository(),
			m.container.GetOperationRepository(),
			path,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Экспорт CSV выполнен успешно.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "2":
		fmt.Print("Введите путь для экспорта JSON: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		errorCh := make(chan error, 1)
		cmd := commands.NewExportJSONCommand(
			m.container.GetBankAccountRepository(),
			m.container.GetCategoryRepository(),
			m.container.GetOperationRepository(),
			path,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Экспорт JSON выполнен успешно.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "3":
		fmt.Print("Введите путь для экспорта YAML: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		errorCh := make(chan error, 1)
		cmd := commands.NewExportYAMLCommand(
			m.container.GetBankAccountRepository(),
			m.container.GetCategoryRepository(),
			m.container.GetOperationRepository(),
			path,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Экспорт YAML выполнен успешно.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "4":
		fmt.Print("Введите путь для импорта CSV: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		errorCh := make(chan error, 1)
		cmd := commands.NewImportCSVCommand(
			m.container.GetBankAccountRepository(),
			m.container.GetCategoryRepository(),
			m.container.GetOperationRepository(),
			path,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Импорт CSV выполнен успешно.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "5":
		fmt.Print("Введите путь для импорта JSON: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		errorCh := make(chan error, 1)
		cmd := commands.NewImportJSONCommand(
			m.container.GetBankAccountRepository(),
			m.container.GetCategoryRepository(),
			m.container.GetOperationRepository(),
			path,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Импорт JSON выполнен успешно.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "6":
		fmt.Print("Введите путь для импорта YAML: ")
		path, _ := reader.ReadString('\n')
		path = strings.TrimSpace(path)
		errorCh := make(chan error, 1)
		cmd := commands.NewImportYAMLCommand(
			m.container.GetBankAccountRepository(),
			m.container.GetCategoryRepository(),
			m.container.GetOperationRepository(),
			path,
			errorCh,
		)
		if err := cmd.Execute(); err == nil {
			fmt.Println("Импорт YAML выполнен успешно.")
		} else {
			fmt.Printf("Ошибка: %v\n", <-errorCh)
		}
	case "0":
		return nil
	default:
		fmt.Println("Неверный выбор.")
	}
	return nil
}

func readDateRange(reader *bufio.Reader) (time.Time, time.Time) {
	fmt.Print("Введите дату начала (YYYY-MM-DD): ")
	startStr, _ := reader.ReadString('\n')
	startStr = strings.TrimSpace(startStr)
	start, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		fmt.Println("Неверный формат даты. Используется текущая дата.")
		start = time.Now()
	}
	fmt.Print("Введите дату окончания (YYYY-MM-DD): ")
	endStr, _ := reader.ReadString('\n')
	endStr = strings.TrimSpace(endStr)
	end, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		fmt.Println("Неверный формат даты. Используется текущая дата.")
		end = time.Now()
	}
	return start, end
}

// Обертываем команду в декоратор для измерения времени выполнения
func (m *MainMenu) wrapWithTimeDecorator(cmd interfaces.Command) interfaces.Command {
	return commands.NewTimeMeasurementDecorator(cmd)
}
