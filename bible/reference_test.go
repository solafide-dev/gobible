package bible

import (
	"reflect"
	"testing"

	"github.com/applehat/gobible"
)

func TestBible_ParseReference(t *testing.T) {
	type args struct {
		reference string
	}
	tests := []struct {
		name    string
		b       *Bible
		args    args
		want    []Reference
		wantErr bool
	}{
		{
			name: "John 3:16",
			b:    gobible.NewBibleKJV(),
			args: args{
				reference: "John 3:16",
			},
			want: []Reference{
				{Book: "John", Chapter: 3, VerseNumber: 16},
			},
			wantErr: false,
		},
		{
			name: "John 3:16-18",
			b:    NewBibleKJV(),
			args: args{
				reference: "John 3:16-18",
			},
			want: []Reference{
				{Book: "John", Chapter: 3, VerseNumber: 16},
				{Book: "John", Chapter: 3, VerseNumber: 17},
				{Book: "John", Chapter: 3, VerseNumber: 18},
			},
			wantErr: false,
		},
		{
			name: "John 3:35-4:2",
			b:    NewBibleKJV(),
			args: args{
				reference: "John 3:35-4:2",
			},
			want: []Reference{
				{Book: "John", Chapter: 3, VerseNumber: 35},
				{Book: "John", Chapter: 3, VerseNumber: 36},
				{Book: "John", Chapter: 4, VerseNumber: 1},
				{Book: "John", Chapter: 4, VerseNumber: 2},
			},
			wantErr: false,
		},
		{
			name: "Rev 22:1",
			b:    NewBibleKJV(),
			args: args{
				reference: "Rev 22:1",
			},
			want: []Reference{
				{Book: "Revelation", Chapter: 22, VerseNumber: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.ParseReference(tt.args.reference)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bible.ParseReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bible.ParseReference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReference_NormalizeBook(t *testing.T) {
	tests := []struct {
		name string
		r    *Reference
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.r.NormalizeBook()
		})
	}
}

func Test_NormalizeBookName(t *testing.T) {
	type args struct {
		bookIn string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeBookName(tt.args.bookIn); got != tt.want {
				t.Errorf("NormalizeBookName() = %v, want %v", got, tt.want)
			}
		})
	}
}
