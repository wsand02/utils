package utils

import (
	"testing"
)

var tPassIn string = "password"
var tPassHash string = "$2a$14$hZGs0u4IW0MmwSrTZVxO.OYtjMIi2Diutl20GhuJS4CU4GMn7ajt2"

func Test_HashPassword(t *testing.T) {
	pass, err := HashPassword(tPassIn)
	if err != nil {
		t.Fail()
	}
	if pass == tPassIn {
		t.Fail()
	}
}

func Test_CheckPassword(t *testing.T) {
	if CheckPasswordHash(tPassIn, tPassHash) == false {
		t.Fail()
	}
}
