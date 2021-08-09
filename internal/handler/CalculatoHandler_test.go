package handler

import (
	"fmt"
	"testing"
)

func TestCalculatoStr(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{"test1", "1+1-2+3/2-3", -2},
		{"test2", "1 +2 *3-2", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculatoStr(tt.args)

			if got.Code != "200" {
				t.Errorf("RegisterUser() = %v, want %v", got, tt.want)
			}
			fmt.Println(got.Data)
		})
	}
}
