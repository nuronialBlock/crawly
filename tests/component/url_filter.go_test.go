package component

import (
	"testing"

	"github.com/nuronialBlock/crawly/component"
)

func TestURLFilter(t *testing.T) {
	tests := []struct {
		name     string
		urls     []string
		domain   string
		expected []int
	}{
		{
			name:     "Test URL filter works ",
			urls:     []string{"http://example.com/path1", "http://example.com/path2/path3", "http://notexample.com"},
			domain:   "example.com",
			expected: []int{2, 1},
		},
		{
			name:     "Test URL filter",
			urls:     []string{"http://example.com/path1", "http://example.com/path2/path3", "http://notexample.com"},
			domain:   "example.com",
			expected: []int{2, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Logf("test name: %s", test.name)
			gotInternal, gotExternal := component.URLFilter(test.urls, test.domain)
			if len(gotInternal) != test.expected[0] {
				t.Errorf("Expected: %v, but got: %v", test.expected[0], gotInternal)
			}

			if len(gotExternal) != test.expected[1] {
				t.Errorf("Expected: %v, but got: %v", test.expected[1], gotExternal)
			}
		})
	}

	t.Log("Passed test for: ParseLineAndExtract")
}
