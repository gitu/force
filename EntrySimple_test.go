package main

import (
	"testing"
)

func TestGetGroups(t *testing.T) {
	groups := GetGroups()

	if len(groups) != 28 {
		t.Log("Length was ", len(groups), " expected would be 28")
		t.FailNow()
	}
}

func TestGetEntries(t *testing.T) {
	type args struct {
		groupIds []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Single Selector - kader", args: args{groupIds: []string{"kader"}}, want: 4},
		{name: "Single Selector - elektro", args: args{groupIds: []string{"elektro"}}, want: 2},
		{name: "Multi Selector - kader & elektro", args: args{groupIds: []string{"elektro", "kader"}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEntries(tt.args.groupIds); len(got) != tt.want {
				t.Errorf("GetEntries() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
