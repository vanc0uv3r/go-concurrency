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

type nameTest struct {
    arg1 string
	arg2 Command
	expected bool
}

type argsTest struct {
    arg1 []string
	arg2 Command
	expected bool
}

var ArgsCommandTests = []argsTest{{[]string{"SET", "key", "val"}, 
					&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}},
					true},
				
					{[]string{"GET", "key",}, 
					&Get{CommandDesc: CommandDesc{name: "GET", argsNum: 2, args: make([]string, 0),}},
					false},
				
					{[]string{"DEL", "key",}, 
					&Del{CommandDesc: CommandDesc{name: "DEL", argsNum: 1, args: make([]string, 0),}},
					true},
				
					{[]string{"SET"}, 
					&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 0, args: make([]string, 0),}},
					true},
				
					{[]string{"SET"}, 
					&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 0, args: make([]string, 0),}},
					true},
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

var CommandNameCheckTest = [...]nameTest{{"SET", 
			&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}}, true,},
			{"sET", 
			&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}}, false},
			{"DEL", 
			&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}}, false}}
	

var CommandArgsCheckTest = [...]nameTest{{"SET", 
			&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}}, true,},
			{"sET", 
			&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}}, false},
			{"DEL", 
			&Set{CommandDesc: CommandDesc{name: "SET", argsNum: 2, args: make([]string, 0),}}, false}}
	


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

func TestCheckCommandName(t *testing.T) {
	engine := NewEngine()
	for _, test := range CommandNameCheckTest{
		if res := engine.checkCommandName(test.arg1, test.arg2); res != test.expected {
    		t.Errorf("Output of test %q not equal to expected %t and return %t", test.arg1, 
																				test.expected,
																				res)
		}
	}
}


func TestCheckCommandArgs(t *testing.T) {
	engine := NewEngine()
	for _, test := range ArgsCommandTests{
		engine.SetLexemes(test.arg1)
		res := engine.checkCommandArgs(test.arg2)
		if res != test.expected {
			t.Errorf("Output of test %q not equal to expected %t and return %t", test.arg1, 
																				test.expected,
																				res)
		}
	}
}