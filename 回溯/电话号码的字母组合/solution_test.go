package main

import (
	"reflect"
	"testing"
)

func TestLetter(t *testing.T) {
	got := letterCombinations("23")
	want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("letterCombinations(\"23\") = %v, want %v", got, want)
	}
}
