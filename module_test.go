package module_test

import (
	"testing"

	module "github.com/zarldev/template-module"
)

func TestModule_String(t *testing.T) {
	t.Parallel()
	stringTests := []struct {
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
	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.m.String(); got != tt.want {
				t.Errorf("module.String() mismatch, got '%v', wanted '%v'", got, tt.want)
			}
		})
	}

	stateTests := []struct {
		name string
		m    *module.Module
		want bool
	}{
		{
			name: "new module",
			m:    &module.Module{},
			want: false,
		},
		{
			name: "module with toggled true",
			m:    module.New(true),
			want: true,
		},
		{
			name: "module with toggled false",
			m:    module.New(false),
			want: false,
		},
		{
			name: "nil module",
			m:    nil,
			want: false,
		},
	}
	for _, tt := range stateTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.m.State(); got != tt.want {
				t.Errorf("module.State() mismatch, got '%v', wanted '%v'", got, tt.want)
			}
		})
	}

	toggleTests := []struct {
		name    string
		initial *module.Module
		want    bool
	}{
		{
			name:    "new module",
			initial: &module.Module{},
			want:    true,
		},
		{
			name:    "module with toggled true",
			initial: module.New(true),
			want:    false,
		},
		{
			name:    "module with toggled false",
			initial: module.New(false),
			want:    true,
		},
		{
			name:    "nil module",
			initial: nil,
			want:    false,
		},
	}
	for _, tt := range toggleTests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tt.initial.Toggle()
			if got := tt.initial.State(); got != tt.want {
				t.Errorf("module.Toggle() mismatch, got '%v', wanted '%v'", got, tt.want)
			}
		})
	}
}
