package ui

import (
	"KPO1/application/commands"
	"KPO1/domain/interfaces"
	"fmt"
)

// MainMenu представляет главное меню приложения
type MainMenu struct {
	ui               *ConsoleUI
	bankAccCmds      map[string]interfaces.Command
	categoryCmds     map[string]interfaces.Command
	operationCmds    map[string]interfaces.Command
	analyticsCmds    map[string]interfaces.Command
	importExportCmds map[string]interfaces.Command
}

// NewMainMenu создает новый экземпляр главного меню
func NewMainMenu(
	ui *ConsoleUI,
	bankAccCmds map[string]interfaces.Command,
	categoryCmds map[string]interfaces.Command,
	operationCmds map[string]interfaces.Command,
	analyticsCmds map[string]interfaces.Command,
	importExportCmds map[string]interfaces.Command,
) *MainMenu {
	return &MainMenu{
		ui:               ui,
		bankAccCmds:      bankAccCmds,
		categoryCmds:     categoryCmds,
		operationCmds:    operationCmds,
		analyticsCmds:    analyticsCmds,
		importExportCmds: importExportCmds,
	}
}

// Start запускает главное меню приложения
func (m *MainMenu) Start() {
	for {
		m.ui.ClearScreen()
		m.ui.PrintMessage("=== ВШЭ-Банк: Модуль учета финансов ===")
		m.ui.PrintMessage("")

		choice, err := m.ui.ReadChoice("Выберите раздел:", []string{
			"Управление счетами",
			"Управление категориями",
			"Управление операциями",
			"Аналитика",
			"Импорт/Экспорт данных",
			"Выход",
		})

		if err != nil {
			m.ui.PrintError(err)
			m.ui.PauseExecution()
			continue
		}

		switch choice {
		case 1:
			m.showBankAccountMenu()
		case 2:
			m.showCategoryMenu()
		case 3:
			m.showOperationMenu()
		case 4:
			m.showAnalyticsMenu()
		case 5:
			m.showImportExportMenu()
		case 6:
			m.ui.PrintMessage("Спасибо за использование ВШЭ-Банк!")
			return
		}
	}
}

// showBankAccountMenu отображает меню для работы со счетами
func (m *MainMenu) showBankAccountMenu() {
	for {
		m.ui.ClearScreen()
		m.ui.PrintMessage("=== Управление счетами ===")
		m.ui.PrintMessage("")

		choice, err := m.ui.ReadChoice("Выберите действие:", []string{
			"Создать новый счет",
			"Просмотреть список счетов",
			"Просмотреть детали счета",
			"Редактировать счет",
			"Удалить счет",
			"Вернуться в главное меню",
		})

		if err != nil {
			m.ui.PrintError(err)
			m.ui.PauseExecution()
			continue
		}

		switch choice {
		case 1:
			m.executeCommand(m.bankAccCmds["create"])
		case 2:
			m.executeCommand(m.bankAccCmds["list"])
		case 3:
			m.executeCommand(m.bankAccCmds["details"])
		case 4:
			m.executeCommand(m.bankAccCmds["edit"])
		case 5:
			m.executeCommand(m.bankAccCmds["delete"])
		case 6:
			return
		}

		m.ui.PauseExecution()
	}
}

// showCategoryMenu отображает меню для работы с категориями
func (m *MainMenu) showCategoryMenu() {
	for {
		m.ui.ClearScreen()
		m.ui.PrintMessage("=== Управление категориями ===")
		m.ui.PrintMessage("")

		choice, err := m.ui.ReadChoice("Выберите действие:", []string{
			"Создать новую категорию",
			"Просмотреть список категорий",
			"Просмотреть категории доходов",
			"Просмотреть категории расходов",
			"Редактировать категорию",
			"Удалить категорию",
			"Вернуться в главное меню",
		})

		if err != nil {
			m.ui.PrintError(err)
			m.ui.PauseExecution()
			continue
		}

		switch choice {
		case 1:
			m.executeCommand(m.categoryCmds["create"])
		case 2:
			m.executeCommand(m.categoryCmds["list"])
		case 3:
			m.executeCommand(m.categoryCmds["listIncome"])
		case 4:
			m.executeCommand(m.categoryCmds["listExpense"])
		case 5:
			m.executeCommand(m.categoryCmds["edit"])
		case 6:
			m.executeCommand(m.categoryCmds["delete"])
		case 7:
			return
		}

		m.ui.PauseExecution()
	}
}

// showOperationMenu отображает меню для работы с операциями
func (m *MainMenu) showOperationMenu() {
	for {
		m.ui.ClearScreen()
		m.ui.PrintMessage("=== Управление операциями ===")
		m.ui.PrintMessage("")

		choice, err := m.ui.ReadChoice("Выберите действие:", []string{
			"Создать новую операцию дохода",
			"Создать новую операцию расхода",
			"Просмотреть список всех операций",
			"Просмотреть операции по счету",
			"Просмотреть операции по категории",
			"Просмотреть детали операции",
			"Удалить операцию",
			"Вернуться в главное меню",
		})

		if err != nil {
			m.ui.PrintError(err)
			m.ui.PauseExecution()
			continue
		}

		switch choice {
		case 1:
			m.executeCommand(m.operationCmds["createIncome"])
		case 2:
			m.executeCommand(m.operationCmds["createExpense"])
		case 3:
			m.executeCommand(m.operationCmds["list"])
		case 4:
			m.executeCommand(m.operationCmds["listByAccount"])
		case 5:
			m.executeCommand(m.operationCmds["listByCategory"])
		case 6:
			m.executeCommand(m.operationCmds["details"])
		case 7:
			m.executeCommand(m.operationCmds["delete"])
		case 8:
			return
		}

		m.ui.PauseExecution()
	}
}

// showAnalyticsMenu отображает меню для аналитики
func (m *MainMenu) showAnalyticsMenu() {
	for {
		m.ui.ClearScreen()
		m.ui.PrintMessage("=== Аналитика ===")
		m.ui.PrintMessage("")

		choice, err := m.ui.ReadChoice("Выберите действие:", []string{
			"Баланс доходов и расходов за период",
			"Группировка расходов по категориям",
			"Группировка доходов по категориям",
			"Общая статистика",
			"Вернуться в главное меню",
		})

		if err != nil {
			m.ui.PrintError(err)
			m.ui.PauseExecution()
			continue
		}

		switch choice {
		case 1:
			m.executeCommand(m.analyticsCmds["balance"])
		case 2:
			m.executeCommand(m.analyticsCmds["expensesByCategory"])
		case 3:
			m.executeCommand(m.analyticsCmds["incomesByCategory"])
		case 4:
			m.executeCommand(m.analyticsCmds["statistics"])
		case 5:
			return
		}

		m.ui.PauseExecution()
	}
}

// showImportExportMenu отображает меню для импорта/экспорта данных
func (m *MainMenu) showImportExportMenu() {
	for {
		m.ui.ClearScreen()
		m.ui.PrintMessage("=== Импорт/Экспорт данных ===")
		m.ui.PrintMessage("")

		choice, err := m.ui.ReadChoice("Выберите действие:", []string{
			"Экспорт данных в CSV",
			"Экспорт данных в JSON",
			"Экспорт данных в YAML",
			"Импорт данных из CSV",
			"Импорт данных из JSON",
			"Импорт данных из YAML",
			"Вернуться в главное меню",
		})

		if err != nil {
			m.ui.PrintError(err)
			m.ui.PauseExecution()
			continue
		}

		switch choice {
		case 1:
			m.executeCommand(m.importExportCmds["exportCSV"])
		case 2:
			m.executeCommand(m.importExportCmds["exportJSON"])
		case 3:
			m.executeCommand(m.importExportCmds["exportYAML"])
		case 4:
			m.executeCommand(m.importExportCmds["importCSV"])
		case 5:
			m.executeCommand(m.importExportCmds["importJSON"])
		case 6:
			m.executeCommand(m.importExportCmds["importYAML"])
		case 7:
			return
		}

		m.ui.PauseExecution()
	}
}

// executeCommand выполняет команду и обрабатывает ошибки
func (m *MainMenu) executeCommand(cmd interfaces.Command) {
	if cmd == nil {
		m.ui.PrintError(fmt.Errorf("команда не реализована"))
		return
	}

	// Оборачиваем команду в декоратор для измерения времени
	timedCmd := commands.NewTimeMeasurementDecorator(cmd)

	err := timedCmd.Execute()
	if err != nil {
		m.ui.PrintError(err)
	}
}
