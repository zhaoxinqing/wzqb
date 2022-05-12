package wzqb

import "testing"

func TestFoo1(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"w"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Foo1()
		})
	}
}
