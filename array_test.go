package go_

import "testing"

func TestJointInt64(t *testing.T) {
	type args struct {
		data []int64
		sep  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success 1",
			args: args{
				data: []int64{1, 2, 3, 4, 5, 6, 7, 165, 789, 123456789},
				sep:  ",",
			},
			want: "1,2,3,4,5,6,7,165,789,123456789",
		},
		{
			name: "Succ 2",
			args: args{
				data: []int64{1, 2, 10, 23, 34, 4567, 1, 239, 0},
				sep:  ", ",
			},
			want: "1, 2, 10, 23, 34, 4567, 1, 239, 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JointInt64(tt.args.data, tt.args.sep); got != tt.want {
				t.Errorf("JointInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
