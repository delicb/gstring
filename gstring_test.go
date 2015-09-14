package gst

import (
	"testing"
)

func TestSprintf(t *testing.T) {
	Printm("Test is: {test} and other value is: {other:%v} foobar\n",
		map[string]interface{}{"test": "baz", "other": []string{"a", "b"}})
	o := Sprintm("{just_string}\n", map[string]interface{}{"just_string": "just string!!!"})
	if o != "just string!!!\n" {
		t.Fail()
	}
}

func Test_Sprintf(t *testing.T) {
	if s := Sprintm("{test}", map[string]interface{}{"test": "hello gstring"}); s != "hello gstring" {
		t.Fail()
	}
}
