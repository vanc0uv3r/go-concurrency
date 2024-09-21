package engine

import (
	"errors"
	"testing"
)

type addTest struct {
    arg1 []string
	expected Command
	expected_err error
}

var FindCommandTests = []addTest{{[]string{"SET", "key", "val"}, 
					&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}},
					nil},
				
					{[]string{"GET", "key",}, 
					&Get{CommandDesc: CommandDesc{name: "GET", argsNum: 1, args: make([]string, 0),}},
					nil},
				
					{[]string{"DEL", "key",}, 
					&Del{CommandDesc: CommandDesc{name: "DEL", argsNum: 1, args: make([]string, 0),}},
					nil},
				
					{[]string{"DEL", "key", "jsadf12"}, 
					nil,
					errors.New(invalidArgsNumErr)},
				
					{[]string{"DEL"}, 
					nil,
					errors.New(invalidArgsNumErr)},
				
					{[]string{"DeL", "key",}, 
					nil,
					errors.New(commandNotFoundErr)},
		}
    

func TestFindCommand(t *testing.T) {
	engine := NewEngine()
	for _, test := range FindCommandTests{
		engine.SetLexemes(test.arg1)
		cmd, err := engine.findCommand()
        
		if !(cmd == test.expected) && !(err == test.expected_err) {
            t.Errorf("Output %q of test %q not equal to expected %q and error %q", cmd, test.arg1, test.expected, test.expected_err)
        }

	}
}

	