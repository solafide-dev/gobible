package gobible

import (
	_ "embed"
	"reflect"
	"testing"
)

func TestNewBible(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want *Bible
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBible(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBibleFromString(t *testing.T) {
	type args struct {
		bibleJSON string
	}
	tests := []struct {
		name string
		args args
		want *Bible
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBibleFromString(tt.args.bibleJSON); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBibleFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBibleKJV(t *testing.T) {
	tests := []struct {
		name string
		want *Bible
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBibleKJV(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBibleKJV() = %v, want %v", got, tt.want)
			}
		})
	}
}
