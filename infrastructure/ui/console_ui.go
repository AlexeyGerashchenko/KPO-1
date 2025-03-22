package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ConsoleUI представляет консольный интерфейс
type ConsoleUI struct {
	reader  *bufio.Reader
	scanner *bufio.Scanner
	menu    *MainMenu
}

// NewConsoleUI создаёт новый консольный интерфейс
func NewConsoleUI() *ConsoleUI {
	consoleUI := &ConsoleUI{
		reader:  bufio.NewReader(os.Stdin),
		scanner: bufio.NewScanner(os.Stdin),
	}
	return consoleUI
}

// SetMenu устанавливает меню для консольного интерфейса
func (ui *ConsoleUI) SetMenu(menu *MainMenu) {
	ui.menu = menu
}

// Run запускает консольный интерфейс
func (ui *ConsoleUI) Run() error {
	// Запускаем меню
	ui.menu.Start()
	return nil
}

// ReadLine читает строку из консоли
func (ui *ConsoleUI) ReadLine() string {
	line, err := ui.reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(line)
}

// PrintMessage выводит сообщение
func (ui *ConsoleUI) PrintMessage(message string) {
	fmt.Println(message)
}

// PrintError выводит ошибку
func (ui *ConsoleUI) PrintError(err error) {
	fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
}

// ClearScreen очищает экран
func (ui *ConsoleUI) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// PauseExecution приостанавливает выполнение до нажатия Enter
func (ui *ConsoleUI) PauseExecution() {
	fmt.Print("Нажмите Enter для продолжения...")
	ui.ReadLine()
}

// ReadInt читает целое число
func (ui *ConsoleUI) ReadInt(prompt string) (int, error) {
	fmt.Print(prompt)
	line := ui.ReadLine()

	value, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return 0, fmt.Errorf("неверный формат числа: %v", err)
	}

	return value, nil
}

// ReadFloat читает число с плавающей точкой
func (ui *ConsoleUI) ReadFloat(prompt string) (float64, error) {
	fmt.Print(prompt)
	line := ui.ReadLine()

	value, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
	if err != nil {
		return 0, fmt.Errorf("неверный формат числа: %v", err)
	}

	return value, nil
}

// ReadChoice читает выбор пользователя из списка вариантов
func (ui *ConsoleUI) ReadChoice(prompt string, options []string) (int, error) {
	fmt.Println(prompt)
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}
	fmt.Print("Выберите вариант (1-" + strconv.Itoa(len(options)) + "): ")

	line := ui.ReadLine()
	value, err := strconv.Atoi(line)
	if err != nil {
		return 0, fmt.Errorf("неверный формат числа: %v", err)
	}

	if value < 1 || value > len(options) {
		return 0, fmt.Errorf("неверный выбор: должно быть число от 1 до %d", len(options))
	}

	return value - 1, nil
}

// ReadConfirmation читает подтверждение (да/нет)
func (ui *ConsoleUI) ReadConfirmation(prompt string) bool {
	fmt.Print(prompt + " (д/н): ")

	line := ui.ReadLine()
	line = strings.ToLower(strings.TrimSpace(line))
	return line == "д" || line == "да" || line == "y" || line == "yes"
}

// ReadDate читает дату в формате YYYY-MM-DD
func (ui *ConsoleUI) ReadDate(prompt string) string {
	fmt.Print(prompt + " (формат YYYY-MM-DD): ")
	line := ui.ReadLine()

	// Здесь можно добавить валидацию формата даты
	return strings.TrimSpace(line)
}

// PrintSuccess выводит сообщение об успехе на консоль
func (ui *ConsoleUI) PrintSuccess(message string) {
	fmt.Printf("УСПЕШНО: %s\n", message)
}

// ReadLineWithPrompt выводит подсказку и считывает строку
func (ui *ConsoleUI) ReadLineWithPrompt(prompt string) string {
	fmt.Print(prompt)
	return ui.ReadLine()
}

// ReadPositiveFloat считывает положительное дробное число от пользователя
func (ui *ConsoleUI) ReadPositiveFloat(prompt string) (float64, error) {
	for {
		num, err := ui.ReadFloat(prompt)
		if err != nil {
			return 0, err
		}

		if num <= 0 {
			ui.PrintError(fmt.Errorf("число должно быть положительным"))
			continue
		}

		return num, nil
	}
}
