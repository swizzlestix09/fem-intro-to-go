package utils

import (
	"fem-intro-to-go/05_toolkit/code/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 18
	actual := utils.Add(5, 6, 7)

	if actual != expected {
		t.Errorf("Add was incorrect! Expected %d, Actual was %d", expected, actual)
	}
}
