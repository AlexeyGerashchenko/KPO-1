package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConsoleUI struct {
	menu *MainMenu
}

func NewConsoleUI() *ConsoleUI {
	return &ConsoleUI{}
}

func (c *ConsoleUI) SetMenu(menu *MainMenu) {
	c.menu = menu
}

func (c *ConsoleUI) Run() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Главное меню ---")
		c.menu.Display()
		fmt.Print("Выберите опцию: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)
		if input == "0" {
			fmt.Println("Выход из приложения.")
			break
		}
		err = c.menu.HandleInput(input, reader)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}
	}
	return nil
}
