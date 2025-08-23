package utils

import (
	"os"
	"testing"
)

const tArg string = "aboow.jxb"
const tArgDef string = "idktbh"

func Test_GetArgWithDefault_1(t *testing.T) {
	// save previous args
	prevArgs := os.Args
	defer func() { os.Args = prevArgs }()

	os.Args = []string{"", tArg}
	v := GetArgWDefault(1, tArgDef)
	if v != tArg {
		t.Fail()
	}
}

func Test_GetArgWithDefault_2(t *testing.T) {
	// save previous args
	prevArgs := os.Args
	defer func() { os.Args = prevArgs }()

	os.Args = []string{""}
	v := GetArgWDefault(1, tArgDef)
	if v != tArgDef {
		t.Fail()
	}
}
