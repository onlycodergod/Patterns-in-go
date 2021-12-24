package main

import (
	"errors"
	"testing"
)

func TestUnpacking(t *testing.T) {
	testCases := []struct {
		str           string
		expectedValue string
		expectedError error
	}{
		{
			str:           "a4bc2d5e",
			expectedValue: "aaaabccddddde",
			expectedError: nil,
		},
		{
			str:           "abcd",
			expectedValue: "abcd",
			expectedError: nil,
		},
		{
			str:           "45",
			expectedValue: "",
			expectedError: errors.New("Wrong string"),
		},
		{
			str:           "",
			expectedValue: "",
			expectedError: nil,
		},
		{
			str:           "qwe\\\\5",
			expectedValue: "qwe\\\\\\\\\\",
			expectedError: nil,
		},
	}

	for _, cases := range testCases {
		res, err := Unpacking(cases.str)
		if res != cases.expectedValue && err != cases.expectedError {
			t.Errorf("Want: Val: %s\n Err:%v\n Have: Val:%s\n Err:%s\n", cases.expectedValue, cases.expectedError, res, err)
		}
	}
}
