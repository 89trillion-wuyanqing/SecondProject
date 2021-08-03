package service

import "testing"

func TestValStrService(t *testing.T) {
	result := ValStrService("333+2")

	t.Log(result)
}

func TestCalculate(t *testing.T) {
	result := Calculate("333+2")
	t.Log(result)
}
