package p886

import "testing"

func Test_possibleBipartition(t *testing.T) {
	type args struct {
		n        int
		dislikes [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				n:        4,
				dislikes: [][]int{{1, 2}, {1, 3}, {2, 4}},
			},
			want: true,
		},
		{
			args: args{
				n:        3,
				dislikes: [][]int{{1, 2}, {1, 3}, {2, 3}},
			},
			want: false,
		},
		{
			args: args{
				n: 10,
				dislikes: [][]int{{4, 7}, {4, 8}, {5, 6}, {1, 6}, {3, 7},
					{2, 5}, {5, 8}, {1, 2}, {4, 9}, {6, 10}, {8, 10}, {3, 6},
					{2, 10}, {9, 10}, {3, 9}, {2, 3}, {1, 9}, {4, 6}, {5, 7},
					{3, 8}, {1, 8}, {1, 7}, {2, 4}},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := possibleBipartition(tt.args.n,
				tt.args.dislikes); got != tt.want {
				t.Errorf("possibleBipartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
