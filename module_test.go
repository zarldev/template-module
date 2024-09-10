package module_test

import (
	"testing"

	module "github.com/zarldev/template-module"
)

func TestModule_String(t *testing.T) {
	tests := []struct {
		name string
		m    *module.Module
		want string
	}{
		{
			name: "new module",
			m:    &module.Module{},
			want: "OFF",
		},
		{
			name: "module with toggled true",
			m:    module.New(true),
			want: "ON",
		},
		{
			name: "module with toggled false",
			m:    module.New(false),
			want: "OFF",
		},
		{
			name: "nil module",
			m:    nil,
			want: "OFF",
		},
	}
	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("module.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
