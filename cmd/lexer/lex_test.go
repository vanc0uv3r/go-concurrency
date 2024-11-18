package lexer

import (
	"testing"
)

type addTest struct {
    arg1 string 
	expected []string
}

func compareSlices(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

var AnalyzeTests = []addTest{
    {"SET key val", []string{"SET", "key", "val"}},
    {"GET key", []string{"GET", "key"}},
    {"", []string{}},
    {"SET 31ke-as_asd_AS12 //unxi/sock.qas1", []string{"SET", "31ke-as_asd_AS12", "//unxi/sock.qas1"}},
}
    

func TestAnalyze(t *testing.T) {
	lex := NewLex()

	for _, test := range AnalyzeTests{
		lex.Analyze([]byte(test.arg1))
		output := lex.GetLexemes()
        if !compareSlices(output, test.expected) {
            t.Errorf("Output %q not equal to expected %q", output, test.expected)
        }
		lex.ClearLexer()
    }
	
}


func TestCleanLexer(t *testing.T) {
    lex := NewLex()

    for _, test := range AnalyzeTests{
		lex.Analyze([]byte(test.arg1))
        lex.ClearLexer()
        if !checkLexEmpty(lex) {
            t.Errorf("Expected lex struct to be clean %+v", lex)   
        }
    }
}

func checkLexEmpty(lex *Lex) bool {
    return  lex.buffer.Len() == 0 && 
            lex.current_state == none && 
            lex.new_state == none && 
            len(lex.lexeme_list) == 0
}