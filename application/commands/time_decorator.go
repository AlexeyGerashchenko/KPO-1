package commands

import (
	"KPO1/domain/interfaces"
	"fmt"
	"time"
)

// TimeMeasurementDecorator реализует паттерн Декоратор для измерения времени выполнения команд
type TimeMeasurementDecorator struct {
	wrappedCommand interfaces.Command
}

// NewTimeMeasurementDecorator создает новый декоратор для измерения времени
func NewTimeMeasurementDecorator(cmd interfaces.Command) *TimeMeasurementDecorator {
	return &TimeMeasurementDecorator{
		wrappedCommand: cmd,
	}
}

// Execute выполняет команду и измеряет время выполнения
func (d *TimeMeasurementDecorator) Execute() error {
	start := time.Now()

	// Выполняем оригинальную команду
	err := d.wrappedCommand.Execute()

	// Вычисляем затраченное время
	duration := time.Since(start)

	// Выводим информацию о времени выполнения
	fmt.Printf("\nВремя выполнения: %v\n", duration)

	return err
}

// GetName возвращает имя команды
func (d *TimeMeasurementDecorator) GetName() string {
	return "Измерение времени: " + d.wrappedCommand.GetName()
}
