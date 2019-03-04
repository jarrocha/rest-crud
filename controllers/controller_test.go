package controllers_test

import (
	"fmt"
	"testing"

	"gitlab.com/codelittinc/golang-interview-project-jaime/rest-crud/controllers"
)

var test_string string = "%s(%s, %s) -> got %v want %v\n"

func TestStrFormValue(t *testing.T) {
	tests := []struct {
		input1 string
		input2 bool
		want1  string
		want2  error
	}{
		{"", false, "", nil},
		{"test", true, "test", nil},
	}

	for _, test := range tests {
		got1, got2 := controllers.StrFormValue(test.input1, test.input2)

		if got1 != test.want1 {
			t.Errorf(test_string, "strFormValue", test.input1, test.input2, got1, test.want1)
		}
		fmt.Println(got2)
		fmt.Println(test.want2)
		if got2 != test.want2 {
			t.Errorf(test_string, "strFormValue", test.input1, test.input2, got2, test.want2)
		}
	}
}

func TestIntFormValue(t *testing.T) {
	tests := []struct {
		input1 string
		input2 bool
		want1  int
		want2  error
	}{
		{"64", false, int(64), nil},
		{"", false, int(0), nil},
	}

	for _, test := range tests {
		got1, got2 := controllers.IntFormValue(test.input1, test.input2)

		if got1 != test.want1 {
			t.Errorf(test_string, "intFormValue", test.input1, test.input2, got1, test.want1)
		}
		if got2 != test.want2 {
			t.Errorf(test_string, "intFormValue", test.input1, test.input2, got2, test.want2)
		}
	}
}
