package main

import (
	"fmt"
	"os"

	"KPO1/di"
	"KPO1/infrastructure/ui"
)

func main() {
	// Создаём консольный интерфейс
	console := ui.NewConsoleUI()

	// Подготавливаем директорию для данных
	dataDir := "./data"
	ensureDir(dataDir)

	// Создаём DI-контейнер
	container := di.NewContainer()

	// Создаём главное меню с доступом к DI-контейнеру
	menu := ui.NewMainMenu(console, container)
	console.SetMenu(menu)

	// Выводим приветствие и запускаем основной цикл
	fmt.Println("Добро пожаловать в систему учета финансов ВШЭ-банка!")
	fmt.Println("=================================================")
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
