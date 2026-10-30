package helpers

import (
	"reflect"
	"testing"
)

func TestMergeStringMaps(t *testing.T) {
	tests := []struct {
		name string
		maps []map[string]string
		want map[string]string
	}{
		{
			name: "no input maps",
			maps: nil,
			want: map[string]string{},
		},
		{
			name: "empty variadic maps slice",
			maps: []map[string]string{},
			want: map[string]string{},
		},
		{
			name: "single map",
			maps: []map[string]string{
				{"a": "1", "b": "2"},
			},
			want: map[string]string{"a": "1", "b": "2"},
		},
		{
			name: "multiple maps with disjoint keys",
			maps: []map[string]string{
				{"a": "1"},
				{"b": "2"},
				{"c": "3"},
			},
			want: map[string]string{"a": "1", "b": "2", "c": "3"},
		},
		{
			name: "multiple maps with overlapping keys (latter overwrites former)",
			maps: []map[string]string{
				{"a": "1", "b": "2"},
				{"b": "20", "c": "30"},
				{"a": "100"},
			},
			want: map[string]string{"a": "100", "b": "20", "c": "30"},
		},
		{
			name: "empty map inputs",
			maps: []map[string]string{
				{},
				{},
			},
			want: map[string]string{},
		},
		{
			name: "nil map inputs",
			maps: []map[string]string{
				nil,
				nil,
			},
			want: map[string]string{},
		},
		{
			name: "combination of nil, empty, and populated maps",
			maps: []map[string]string{
				nil,
				{"a": "1"},
				{},
				{"b": "2", "a": "override"},
				nil,
			},
			want: map[string]string{"a": "override", "b": "2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MergeStringMaps(tt.maps...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeStringMaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
