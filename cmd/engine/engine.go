package engine

import (
	"errors"
)

type Command interface {
	execute() (string, error)
	getName() string
	getArgsNumber() int
	getArgs() []string
	setArgs([]string)
}

type CommandDesc struct {
	name    string
	argsNum int
	args    []string
}

type Set struct{ CommandDesc }

type Get struct{ CommandDesc }

type Del struct{ CommandDesc }

func (s *Set) execute() (string, error) {
	args := s.getArgs()
	storage[args[0]] = args[1]
	return args[1], nil
}
func (s *Set) getName() string       { return s.name }
func (s *Set) getArgsNumber() int    { return s.argsNum }
func (s *Set) getArgs() []string     { return s.args }
func (s *Set) setArgs(args []string) { s.args = args }

func (g *Get) execute() (string, error) {
	key := g.getArgs()[0]
	value, exists := storage[key]
	if exists {
		return value, nil
	}

	return "", errors.New(keyNotExistsErr)
}
func (g *Get) getName() string       { return g.name }
func (g *Get) getArgsNumber() int    { return g.argsNum }
func (g *Get) getArgs() []string     { return g.args }
func (g *Get) setArgs(args []string) { g.args = args }

func (d *Del) execute() (string, error) {
	key := d.getArgs()[0]
	value, exists := storage[key]
	if exists {
		delete(storage, key)
		return value, nil
	}

	return "", errors.New(keyNotExistsErr)
}
func (d *Del) getName() string       { return d.name }
func (d *Del) getArgsNumber() int    { return d.argsNum }
func (d *Del) getArgs() []string     { return d.args }
func (d *Del) setArgs(args []string) { d.args = args }

var (
	commands []Command
	storage  map[string]string
)

func init() {
	commands = []Command{
		&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0)}},
		&Get{CommandDesc: CommandDesc{name: "GET", argsNum: 1, args: make([]string, 0)}},
		&Del{CommandDesc: CommandDesc{name: "DEL", argsNum: 1, args: make([]string, 0)}},
	}
	storage = make(map[string]string)
}

type Engine struct {
	lexemes        []string
	current_lexeme string
	command        Command
}

func NewEngine() *Engine {
	return &Engine{
		lexemes:        make([]string, 0),
		current_lexeme: "",
		command:        nil,
	}
}

func (e *Engine) SetLexemes(lexemes []string) { e.lexemes = lexemes }
func (e *Engine) GetCommandName() string      { return e.command.getName() }

func (e *Engine) Execute() (string, error) {
	if len(e.lexemes) == 0 {
		return "", errors.New(emptyLineErr)
	}
	command, err := e.findCommand()
	if err != nil {
		return "", err
	}

	if !e.checkCommandArgs(command) {
		return "", errors.New(invalidArgsNumErr)
	}
	command.setArgs(e.lexemes[1:])

	e.command = command
	return command.execute()
}

func (e *Engine) findCommand() (Command, error) {
	cmdToFound := e.lexemes[0]
	for _, command := range commands {
		if e.checkCommandName(cmdToFound, command) {
			return command, nil
		}
	}

	return nil, errors.New(commandNotFoundErr)
}

func (e *Engine) checkCommandName(cmdToFound string, command Command) bool {
	name := command.getName()
	return cmdToFound == name
}

func (e *Engine) checkCommandArgs(command Command) bool {
	argsNum := command.getArgsNumber()
	return argsNum == len(e.lexemes)-1
}
