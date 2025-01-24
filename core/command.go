package core

type Command struct {
	Name        string
	Description string
	Execute     func(args []string) error
	Help        func()
}

var Commands = make(map[string]Command)

func RegisterCommand(name string, description string,
	execute func(args []string) error, help func()) {
	Commands[name] = Command{
		Name:        name,
		Description: description,
		Execute:     execute,
		Help:        help,
	}
}
