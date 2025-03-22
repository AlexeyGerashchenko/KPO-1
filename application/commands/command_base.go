package commands

import (
	"KPO1/domain/interfaces"
	"time"
)

// Command представляет команду для выполнения
type CommandBase struct {
	name string
}

// NewCommandBase создаёт новую базовую команду
func NewCommandBase(name string) CommandBase {
	return CommandBase{
		name: name,
	}
}

// GetName возвращает имя команды
func (c CommandBase) GetName() string {
	return c.name
}

// TimeLoggerDecorator представляет декоратор для логирования времени выполнения команды
type TimeLoggerDecorator struct {
	command interfaces.Command
}

// NewTimeLoggerDecorator создаёт новый декоратор для логирования времени
func NewTimeLoggerDecorator(command interfaces.Command) interfaces.Command {
	return &TimeLoggerDecorator{
		command: command,
	}
}

// Execute выполняет команду и логирует время выполнения
func (d *TimeLoggerDecorator) Execute() error {
	start := time.Now()
	err := d.command.Execute()
	elapsed := time.Since(start)

	// Здесь можно добавить логирование времени выполнения
	// log.Printf("Command '%s' executed in %s", d.GetName(), elapsed)

	_ = elapsed // чтобы избежать ошибки компиляции

	return err
}

// GetName возвращает имя команды
func (d *TimeLoggerDecorator) GetName() string {
	return d.command.GetName()
}
