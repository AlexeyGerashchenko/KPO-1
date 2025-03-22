package main

import (
	"KPO1/domain/interfaces"
	"KPO1/infrastructure/ui"
	"fmt"
	"os"
)

// Интерактивная команда представляет команду с интерактивным вводом
type InteractiveCommand struct {
	name    string
	execute func() error
}

// Execute выполняет команду
func (c *InteractiveCommand) Execute() error {
	return c.execute()
}

// GetName возвращает имя команды
func (c *InteractiveCommand) GetName() string {
	return c.name
}

func main() {
	// Создаем консольный интерфейс
	console := ui.NewConsoleUI()

	// Подготавливаем директорию для данных
	dataDir := "./data"
	ensureDir(dataDir)

	// Создаем пустые карты команд, которые будут заполнены позже
	bankAccountCommands := make(map[string]interfaces.Command)
	categoryCommands := make(map[string]interfaces.Command)
	operationCommands := make(map[string]interfaces.Command)
	analyticsCommands := make(map[string]interfaces.Command)
	importExportCommands := make(map[string]interfaces.Command)

	// Создаем главное меню
	menu := ui.NewMainMenu(
		console,
		bankAccountCommands,
		categoryCommands,
		operationCommands,
		analyticsCommands,
		importExportCommands,
	)

	// Устанавливаем меню для консольного интерфейса
	console.SetMenu(menu)

	// Запускаем приложение
	fmt.Println("Добро пожаловать в систему учета финансов ВШЭ-банка!")
	fmt.Println("=================================================")

	// Запускаем основной цикл обработки команд через консольный интерфейс
	if err := console.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("До свидания! Спасибо за использование системы учета финансов ВШЭ-банка!")
}

// ensureDir создает директорию, если она не существует
func ensureDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, 0755)
		if err != nil {
			panic(fmt.Sprintf("Не удалось создать директорию: %v", err))
		}
	}
}
