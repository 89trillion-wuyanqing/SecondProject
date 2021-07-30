package handler

import (
	"testing"
)

func TestCalculatoStr(t *testing.T) {
	result, err := CalculatoStr("1+1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
