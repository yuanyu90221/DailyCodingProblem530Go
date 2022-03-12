package diff_string

import "testing"

func TestGetStringDistance(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "kitten, sitting",
			args: args{str1: "kitten", str2: "sitting"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetStringDistance(tt.args.str1, tt.args.str2); got != tt.want {
				t.Errorf("GetStringDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
