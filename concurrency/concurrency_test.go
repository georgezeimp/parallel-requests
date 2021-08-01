package concurrency

import (
	"sort"
	"sync"
	"testing"
)

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		cs       *ConcurrentSlice
		newItems []string
		want     *ConcurrentSlice
	}{
		{
			name: "Test append single",
			cs: &ConcurrentSlice{
				items: []string{
					"Go",
					"PHP",
					"Python",
				},
			},
			newItems: []string{
				"Java",
			},
			want: &ConcurrentSlice{
				items: []string{
					"Go",
					"PHP",
					"Python",
					"Java",
				},
			},
		},
		{
			name: "Test append multiple",
			cs: &ConcurrentSlice{
				items: []string{
					"Go",
				},
			},
			newItems: []string{
				"PHP",
				"Python",
				"Java",
				"Kotlin",
				"Scala",
			},
			want: &ConcurrentSlice{
				items: []string{
					"Go",
					"PHP",
					"Python",
					"Java",
					"Kotlin",
					"Scala",
				},
			},
		},
		{
			name: "Test append multiple to empty slice",
			cs: &ConcurrentSlice{
				items: []string{},
			},
			newItems: []string{
				"Go",
				"PHP",
				"Python",
				"Java",
				"Kotlin",
				"Scala",
			},
			want: &ConcurrentSlice{
				items: []string{
					"Go",
					"PHP",
					"Python",
					"Java",
					"Kotlin",
					"Scala",
				},
			},
		},
	}

	for _, tt := range tests {
		var waitGroup sync.WaitGroup
		waitGroup.Add(len(tt.newItems))

		for _, item := range tt.newItems {
			go func(i string, wg *sync.WaitGroup) {
				tt.cs.Append(i)
				waitGroup.Done()
			}(item, &waitGroup)
		}

		waitGroup.Wait()

		sort.Strings(tt.want.items)
		sort.Strings(tt.cs.items)
		wantItems := tt.want.items
		gotItems := tt.cs.items
		equal := true
		if len(wantItems) != len(gotItems) {
			equal = false
		}

		for i := range wantItems {
			if wantItems[i] != gotItems[i] {
				equal = false
			}
		}

		if !equal {
			t.Errorf("Test \"%s\" failed for Input: %+v, Expected Output: %+v, Actual Output: %+v", tt.name, tt.cs, tt.want, tt.cs)
		}

	}
}

func TestGetItems(t *testing.T) {
	tests := []struct {
		name string
		cs   *ConcurrentSlice
		want []string
	}{
		{
			name: "Test with empty slice",
			cs:   &ConcurrentSlice{},
			want: []string{},
		},
		{
			name: "Test with non empty slice",
			cs: &ConcurrentSlice{
				items: []string{
					"scorpions",
					"nightwish",
					"epica",
				},
			},
			want: []string{
				"scorpions",
				"nightwish",
				"epica",
			},
		},
	}

	for _, tt := range tests {
		got := tt.cs.GetItems()

		equal := true
		if len(tt.want) != len(got) {
			equal = false
		}
		for i := range tt.want {
			if tt.want[i] != got[i] {
				equal = false
			}
		}

		if !equal {
			t.Errorf("Test \"%s\" failed for Input: %+v, Expected Output: %+v, Actual Output: %+v", tt.name, tt.cs, tt.want, got)
		}
	}
}
