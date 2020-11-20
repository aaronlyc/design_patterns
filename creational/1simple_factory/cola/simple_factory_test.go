package cola

import (
	"reflect"
	"testing"
)

func TestNewCola(t *testing.T) {
	type args struct {
		t int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				t: 0,
			},
			want: "",
		},
		{
			name: "test2",
			args: args{
				t: 1,
			},
			want: "Drink the cacaCola",
		},
		{
			name: "test3",
			args: args{
				t: 2,
			},
			want: "Drink the pesiCola",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cola := NewCola(tt.args.t)
			if cola == nil {
				if !reflect.DeepEqual(cola, tt.want) {
					t.Errorf("NewCola() = %s, want %s", "", tt.want)
				}
			} else {
				if got := cola.Drink(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewCola() = %s, want %s", got, tt.want)
				}
			}
		})
	}
}
