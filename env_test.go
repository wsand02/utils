package utils

import (
	"testing"
)

var tBool1 bool = false
var tBool2 bool = true

var tInt64_1 int64 = 100
var tInt64_2 int64 = 55

var tUint64_1 uint64 = 77
var tUint64_2 uint64 = 69

var tString1 string = "owo"
var tString2 string = "uwu"

func Test_GetEnvAsBool_1(t *testing.T) {
	// Will test if the default value works.
	v := GetEnvAsBool("utils_bool_aboow", tBool1)
	if v != tBool1 {
		t.Fail()
	}
}

func Test_GetEnvAsBool_2(t *testing.T) {
	// Will test if the regular value works.
	v := GetEnvAsBool("utils_bool_test", tBool1)
	if v != tBool2 {
		t.Fail()
	}
}

func Test_GetEnvAsInt64_1(t *testing.T) {
	v := GetEnvAsInt64("utils_int64_aboow", tInt64_1)
	if v != tInt64_1 {
		t.Fail()
	}
}

func Test_GetEnvAsInt64_2(t *testing.T) {
	v := GetEnvAsInt64("utils_int64_test", tInt64_1)
	if v != tInt64_2 {
		t.Fail()
	}
}

func Test_GetEnvAsUint64_1(t *testing.T) {
	v := GetEnvAsUint64("utils_uint64_aboow", tUint64_1)
	if v != tUint64_1 {
		t.Fail()
	}
}

func Test_GetEnvAsUint64_2(t *testing.T) {
	v := GetEnvAsUint64("utils_uint64_test", tUint64_1)
	if v != tUint64_2 {
		t.Fail()
	}
}

func Test_GetEnvWDefault_1(t *testing.T) {
	v := GetEnvWDefault("utils_str_aboow", tString1)
	if v != tString1 {
		t.Fail()
	}
}

func Test_GetEnvWDefault_2(t *testing.T) {
	v := GetEnvWDefault("utils_str_test", tString1)
	if v != tString2 {
		t.Fail()
	}
}
