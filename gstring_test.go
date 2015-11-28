package gstring

import (
	"bytes"
	"testing"
)

func TestSprintf(t *testing.T) {
	Printm("Test is: {test} and other value is: {other:%v} foobar\n",
		map[string]interface{}{"test": "baz", "other": []string{"a", "b"}})
	o := Sprintm("{just_string}\n", map[string]interface{}{"just_string": "just string!!!"})
	if o != "just string!!!\n" {
		t.Fatal("Expected 'just string!!!', got: ", o)
	}
}

func Test_Sprintf(t *testing.T) {
	if s := Sprintm("{test}", map[string]interface{}{"test": "hello gstring"}); s != "hello gstring" {
		t.Fatal("Expected 'hello gstring', got: ", s)
	}
}

func TestEscapeBracket(t *testing.T) {
	res := Sprintm("{{ some text }}", map[string]interface{}{})
	if res != "{ some text }" {
		t.Fatal("Got message: ", res)
	}
}

func TestErrorm(t *testing.T) {
	res := Errorm("{err}", map[string]interface{}{"err": "msg"})
	if res.Error() != "msg" {
		t.Fail()
	}
}

func TestFprintm(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	Fprintm(b, "{k}", map[string]interface{}{"k": "v"})
	if b.String() != "v" {
		t.Fatal("Expected 'k', got: ", b.String())
	}
}
