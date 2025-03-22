package interfaces

// Command представляет команду для выполнения
type Command interface {
	Execute() error
	GetName() string
}

// CommandDecorator представляет декоратор для команды
type CommandDecorator interface {
	Decorate(Command) Command
}
