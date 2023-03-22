package gobible

import (
	"reflect"
	"testing"
)

func TestBible_GetBook(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		b    *Bible
		args args
		want *Book
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetBook(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bible.GetBook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBible_GetName(t *testing.T) {
	tests := []struct {
		name string
		b    *Bible
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetName(); got != tt.want {
				t.Errorf("Bible.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBible_GetAbbrev(t *testing.T) {
	tests := []struct {
		name string
		b    *Bible
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetAbbrev(); got != tt.want {
				t.Errorf("Bible.GetAbbrev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_GetName(t *testing.T) {
	tests := []struct {
		name string
		b    *Book
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetName(); got != tt.want {
				t.Errorf("Book.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_GetEnglishName(t *testing.T) {
	tests := []struct {
		name string
		b    *Book
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetEnglishName(); got != tt.want {
				t.Errorf("Book.GetEnglishName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_GetChapter(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		b    *Book
		args args
		want *Chapter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetChapter(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Book.GetChapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_GetVerseCount(t *testing.T) {
	tests := []struct {
		name string
		c    *Chapter
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetVerseCount(); got != tt.want {
				t.Errorf("Chapter.GetVerseCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_GetVerse(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		c    *Chapter
		args args
		want *Verse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetVerse(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chapter.GetVerse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapter_GetVerses(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		c    *Chapter
		args args
		want []Verse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetVerses(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chapter.GetVerses() = %v, want %v", got, tt.want)
			}
		})
	}
}
