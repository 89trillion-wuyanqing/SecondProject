package service

import "testing"

func TestValStrService(t *testing.T) {
	result, err := ValStrService("333+2")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
