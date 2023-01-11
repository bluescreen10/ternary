package ternary_test

import (
	"testing"

	"github.com/bluescreen10/ternary"
)

func TrueCondition(t *testing.T) {
	ifTrue := "true"
	ifFalse := "false"

	result := ternary.T(true, ifTrue, ifFalse)
	if result != ifTrue {
		t.Errorf("expected %s got %s", ifTrue, result)
	}
}

func FalseCondition(t *testing.T) {
	ifTrue := 1.0
	ifFalse := 0.0

	result := ternary.T(false, ifTrue, ifFalse)
	if result != ifTrue {
		t.Errorf("expected %f got %f", ifTrue, result)
	}
}
