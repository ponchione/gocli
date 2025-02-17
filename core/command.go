package core

type Command interface {
	Execute(args []string) error
	Help() string
}

var Commands = make(map[string]Command)
