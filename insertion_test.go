package gort

import (
	"testing"
	"time"
)

func TestInsertion(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "100 items",
			args: args{
				arr: CreateIntSlice(100),
			},
		},
		{
			name: "1000 items",
			args: args{
				arr: CreateIntSlice(1000),
			},
		},
		{
			name: "10000 items",
			args: args{
				arr: CreateIntSlice(10000),
			},
		},
		{
			name: "100000 items",
			args: args{
				arr: CreateIntSlice(100000),
			},
		},
		{
			name: "200000 items",
			args: args{
				arr: CreateIntSlice(200000),
			},
		},
		{
			name: "300000 items",
			args: args{
				arr: CreateIntSlice(300000),
			},
		},
	}
	timeResult := make([]int64, 0)
	var temp = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	Insertion(temp)
	t.Log(temp)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			Insertion(tt.args.arr)
			t.Log("took for", time.Duration(time.Since(now)).Milliseconds(), tt.name)
			timeResult = append(timeResult, time.Duration(time.Since(now)).Milliseconds())
			tt.args.arr = nil
			//t.Log(tt.args.arr)
		})
	}
	t.Log(timeResult)
}
