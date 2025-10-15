package models_test

import (
	"demo/models"
	"testing"
)

func TestValidateContact(t *testing.T) {
	contact1 := &models.Contact{Name: "jiten", Email: "jitenp@outlook.com", Mobile: "9618558500"}
	err := contact1.Validate()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
	contact2 := &models.Contact{Email: "jitenp@outlook.com", Mobile: "9618558500"}
	err = contact2.Validate()
	if err == nil {
		t.Fail()
	}
}

func TestStringReverse1(t *testing.T) {
	str := "abcdef"
	result := models.ReverseStr(str)
	expected := "fedcba"
	if result != expected {
		t.Fail()
	}
}
