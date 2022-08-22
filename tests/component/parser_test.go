package component

import (
	"os"
	"testing"

	"github.com/nuronialBlock/crawly/component"
)

func createMockFile() string {
	filename := "temp_test_file.txt"
	file, _ := os.Create(filename)
	defer file.Close()

	content := []byte(`
		hello\n
		<a href="/path1" >\n
		<a href="/path2" >\n
		<a href="https://this.is.an.absolute/path1">\n
	`)

	file.Write(content)

	return filename
}

func deleteMockFile() {
	filename := "temp_test_file.txt"
	err := os.Remove(filename)
	if err != nil {
		panic(err)
	}
}

func TestParseLineAndExtract(t *testing.T) {
	file := createMockFile()

	tests := []struct {
		name     string
		filename string
		baseURL  string
		expected int
	}{
		{
			name:     "Parse and Link test",
			filename: file,
			baseURL:  "https://example.com",
			expected: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Logf("test name: %s", test.name)
			got := component.ParseLineAndExtract(test.filename, test.baseURL)
			if len(got) != test.expected {
				t.Errorf("Expected: %v, but got: %v", test.expected, got)
			}
		})
	}

	t.Log("Passed test for: ParseLineAndExtract")
	deleteMockFile()
}
