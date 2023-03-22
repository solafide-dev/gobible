package gobible

import (
	"reflect"
	"testing"
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
			b:    NewBibleKJV(),
			args: args{
				reference: "John 3:16",
			},
			want: []Reference{
				{
					Book:        "John",
					Chapter:     3,
					VerseNumber: 16,
				},
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
