package quiz

import (
	"encoding/csv"
	"reflect"
	"strings"
	"testing"
)

func Test_ReadCSVFile(t *testing.T) {
	in := `fake question,answer`
	csvReader := csv.NewReader(strings.NewReader(in))
	r, err := ReadCSVFile(csvReader)

	if err != nil {
		t.Error("Failed to read csv data")
	}

	assertEqual(t, r, [][]string{{"fake question", "answer"}})
}

func Test_GetQuestionsFromStrings(t *testing.T) {
	testcases := []struct {
		name          string
		in            string
		expected      []Question
		expectedError error
	}{
		{
			name: "converts strings to array of questions",
			in: `fake question,answer
second question,more answer
last question,last answer`,
			expected: []Question{
				{"fake question", "answer"},
				{"second question", "more answer"},
				{"last question", "last answer"},
			},
			expectedError: nil,
		},
		{
			name: "skips when question or answer is missing",
			in: `fake question,answer
second question,more answer
last question,last answer
lol,`,
			expected: []Question{
				{"fake question", "answer"},
				{"second question", "more answer"},
				{"last question", "last answer"},
			},
			expectedError: InvalidInputError{},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			csvReader := csv.NewReader(strings.NewReader(tc.in))
			r, err := ReadCSVFile(csvReader)

			if err != nil {
				t.Error("failed to read csv data")
			}

			got, err := GetQuestionsFromStrings(r)
			assertIsInstance(t, err, tc.expectedError)
			assertEqual(t, got, tc.expected)
		})
	}
}

func assertEqual(t *testing.T, got, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong contents, got %+v want %+v", got, want)
	}
}

func assertIsInstance(t *testing.T, a, b interface{}) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		t.Errorf("%+v is not an instance of %+v", a, b)
	}
}
